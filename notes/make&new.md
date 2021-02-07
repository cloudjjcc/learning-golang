### make

`make` 的作用是初始化内置的数据结构，也就是我们在前面提到的切片、哈希表和 Channel[2](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-make-and-new/#fn:2)



我们在代码中往往都会使用如下所示的语句初始化这三类基本类型，这三个语句分别返回了不同类型的数据结构：

```go
slice := make([]int, 0, 100)
hash := make(map[int]bool, 10)
ch := make(chan int, 5)
```

1. `slice` 是一个包含 `data`、`cap` 和 `len` 的结构体 [`reflect.SliceHeader`](https://draveness.me/golang/tree/reflect.SliceHeader)；
2. `hash` 是一个指向 [`runtime.hmap`](https://draveness.me/golang/tree/runtime.hmap) 结构体的指针；
3. `ch` 是一个指向 [`runtime.hchan`](https://draveness.me/golang/tree/runtime.hchan) 结构体的指针；



### new

编译器会在中间代码生成阶段通过以下两个函数处理该关键字：

1. [`cmd/compile/internal/gc.callnew`](https://draveness.me/golang/tree/cmd/compile/internal/gc.callnew) 会将关键字转换成 `ONEWOBJ` 类型的节点[2](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-make-and-new/#fn:2)；

2. `cmd/compile/internal/gc.state.expr`

    

   会根据申请空间的大小分两种情况处理：

   1. 如果申请的空间为 0，就会返回一个表示空指针的 `zerobase` 变量；
   2. 在遇到其他情况时会将关键字转换成 [`runtime.newobject`](https://draveness.me/golang/tree/runtime.newobject) 函数：

[`runtime.newobject`](https://draveness.me/golang/tree/runtime.newobject) 函数会获取传入类型占用空间的大小，调用 [`runtime.mallocgc`](https://draveness.me/golang/tree/runtime.mallocgc) 在堆上申请一片内存空间并返回指向这片内存空间的指针：

```go
func newobject(typ *_type) unsafe.Pointer {
	return mallocgc(typ.size, typ, true)
}
```