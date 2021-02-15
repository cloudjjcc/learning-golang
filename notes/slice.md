# 数据结构

编译期间的切片是 [`cmd/compile/internal/types.Slice`](https://draveness.me/golang/tree/cmd/compile/internal/types.Slice) 类型的，但是在运行时切片可以由如下的 [`reflect.SliceHeader`](https://draveness.me/golang/tree/reflect.SliceHeader) 结构体表示，其中:

- `Data` 是指向数组的指针;
- `Len` 是当前切片的长度；
- `Cap` 是当前切片的容量，即 `Data` 数组的大小：

```go
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}
```

`Data` 是一片连续的内存空间，这片内存空间可以用于存储切片中的全部元素，数组中的元素只是逻辑上的概念，底层存储其实都是连续的，所以我们可以将切片理解成一片连续的内存空间加上长度与容量的标识。

![golang-slice-struct](https://img.draveness.me/2019-02-20-golang-slice-struct.png)



# 创建切片

```go
// 从数组创建切片 等同于切片字面量创建sli:=[]int{1,2,3}
arr := [...]int{1, 2, 3}
sli := arr[:]
// 从已有切片创建，共享底层数组
sli2 := sli[1:]
fmt.Printf("arr:(%p,%T),sli:(%p,%T),sli2:(%p,%T)", &arr, arr, sli, sli,sli2,sli2)
// arr:(0xc00001e240,[3]int),sli:(0xc00001e240,[]int),sli2:(0xc00001e248,[]int)
// 通过make创建
sli3:=make([]int,0,5)
// 通过new创建
sli4:=new([]int)
```





# 访问元素

`len(slice)` 或者 `cap(slice)` 在一些情况下会直接替换成切片的长度或者容量，不需要在运行时获取



# 追加和扩容

append方法可以向切片追加元素，当原切片的cap不够时会发生扩容

```go
sli:=[]int{1,2,3}
fmt.Printf("sli(data_addr:%p，type_addr:%p,cap:%d)\n",sli,&sli,cap(sli))

sli=append(sli,4)// cap<len+1,切片扩容，底层数组变化
fmt.Printf("sli(data_addr:%p，type_addr:%p,cap:%d)\n",sli,&sli,cap(sli))

sli2 := append(sli, 5)// 创建新的切片，sli2和sli共享底层数组
fmt.Printf("sli2:(data_addr:%p，type_addr:%p,cap:%d)\n",sli2,&sli2,cap(sli2))
/*
sli(data_addr:0xc0000d0040，type_addr:0xc0000b6080,cap:3)
sli(data_addr:0xc0000a0030，type_addr:0xc0000b6080,cap:6)
sli2:(data_addr:0xc0000a0030，type_addr:0xc0000b60e0,cap:6)
*/
```

扩容是为切片分配新的内存空间并拷贝原切片中元素的过程

+ 新容量计算

  在分配内存空间之前需要先确定新的切片容量，运行时根据切片的当前容量选择不同的策略进行扩容：

  1. 如果期望容量大于当前容量的两倍就会使用期望容量；
  2. 如果当前切片的长度小于 1024 就会将容量翻倍；
  3. 如果当前切片的长度大于 1024 就会每次增加 25% 的容量，直到新容量大于期望容量；

+ 内存对齐

  需要根据切片中的元素大小对齐内存

+ 内存分配

  调用`runtime.mallocgc` 方法分配内存

+ 拷贝切片

  使用 [`runtime.memmove`](https://draveness.me/golang/tree/runtime.memmove) 将原数组内存中的内容拷贝到新申请的内存中。



# 切片深拷贝



可以利用copy函数进行切片深拷贝

```go
src := []string{"hello", "world", "小明"}
dst := make([]string, 3)
copy(dst, src)
fmt.Printf("src_data_addr:%p,dst_data_addr:%p\n",src,dst)
fmt.Println(dst)
/* 输出
src_data_addr:0xc00007c3c0,dst_data_addr:0xc00007c3f0
[hello world 小明]
*/
```



编译器会使用 [`runtime.slicecopy`](https://draveness.me/golang/tree/runtime.slicecopy) 替换运行期间调用的 `copy`，该函数的实现很简单：

```go
func slicecopy(to, fm slice, width uintptr) int {
	if fm.len == 0 || to.len == 0 {
		return 0
	}
	n := fm.len
	if to.len < n {
		n = to.len
	}
	if width == 0 {
		return n
	}
	...

	size := uintptr(n) * width
	if size == 1 {
		*(*byte)(to.array) = *(*byte)(fm.array)
	} else {
		memmove(to.array, fm.array, size)
	}
	return n
}
```

无论是编译期间拷贝还是运行时拷贝，两种拷贝方式都会通过 [`runtime.memmove`](https://draveness.me/golang/tree/runtime.memmove) 将整块内存的内容拷贝到目标的内存区域中：

![golang-slice-copy](https://img.draveness.me/2019-02-20-golang-slice-copy.png)

**语言切片的拷贝**

相比于依次拷贝元素，[`runtime.memmove`](https://draveness.me/golang/tree/runtime.memmove) 能够提供更好的性能。需要注意的是，整块拷贝内存仍然会占用非常多的资源，在大切片上执行拷贝操作时一定要注意对性能的影响。





# nil切片和空切片

```go
	var sli []int//nil 切片
	fmt.Printf("%v,%p,%p\n", sli, sli, &sli)
	sli2 := []int{}//空切片
	fmt.Printf("%v,%p,%p\n", sli2, sli2, &sli2)
	sli3 := make([]int, 0)//空切片
	fmt.Printf("%v,%p,%p\n", sli3, sli3, &sli3)
/*
[],0x0,0xc00000e120
[],0x1852ed8,0xc00000e140
[],0x1852ed8,0xc00000e1a0
*/
```

nil切片的数组指针为nil

空切片的数组指针为zerobase

