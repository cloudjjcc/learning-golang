# 经典循环

一个常见的 for 循环代码会被 [`cmd/compile/internal/gc.state.stmt`](https://draveness.me/golang/tree/cmd/compile/internal/gc.state.stmt) 转换成下面的控制结构，该结构中包含了 4 个不同的块，这些代码块之间的连接表示汇编语言中的跳转关系，与我们理解的 for 循环控制结构没有太多的差别。

![golang-for-loop-ssa](https://img.draveness.me/2020-01-17-15792766877627-golang-for-loop-ssa.png)

**Go 语言循环生成的 SSA 代码**

[机器码生成](https://draveness.me/golang/docs/part1-prerequisite/ch02-compile/golang-machinecode/)阶段会将这些代码块转换成机器码，以及指定 CPU 架构上运行的机器语言，就是我们在前面编译得到的汇编指令。



# 范围循环



编译器会在编译期间将所有 for-range 循环变成经典循环。从编译器的视角来看，就是将 `ORANGE` 类型的节点转换成 `OFOR` 节点:

![Golang-For-Range-Loop](https://img.draveness.me/2020-01-17-15792766926441-Golang-For-Range-Loop.png)

** 范围循环、普通循环和 SSA**





## arrayClear 优化

优化 Go 语言遍历数组或者切片并删除全部元素的逻辑：

```go
// 原代码
for i := range a {
	a[i] = zero
}

// 优化后
if len(a) != 0 {
	hp = &a[0]
	hn = len(a)*sizeof(elem(a))
	memclrNoHeapPointers(hp, hn)
	i = len(a) - 1
}
```



### 切片或数组

对于所有的 range 循环，Go 语言都会在编译期将原切片或者数组赋值给一个新变量 `ha`，在赋值的过程中就发生了拷贝，而我们又通过 `len` 关键字预先获取了切片的长度，所以在循环中追加新的元素也不会改变循环执行的次数



而遇到这种同时遍历索引和元素的 range 循环时，Go 语言会额外创建一个新的 `v2` 变量存储切片中的元素，**循环中使用的这个变量 v2 会在每一次迭代被重新赋值而覆盖，赋值时也会触发拷贝**

因此当我们想要访问数组中元素所在的地址时，不应该直接获取 range 返回的变量地址 `&v2`，而应该使用 `&a[index]` 这种形式。



### 哈希表 

在遍历哈希表时，编译器会使用 [`runtime.mapiterinit`](https://draveness.me/golang/tree/runtime.mapiterinit) 和 [`runtime.mapiternext`](https://draveness.me/golang/tree/runtime.mapiternext) 两个运行时函数重写原始的 for-range 循环：

```go
ha := a
hit := hiter(n.Type)
th := hit.Type
mapiterinit(typename(t), ha, &hit)
for ; hit.key != nil; mapiternext(&hit) {
    key := *hit.key
    val := *hit.val
}
```



![golang-range-map-and-buckets](https://img.draveness.me/2020-01-17-15792766877646-golang-range-map-and-buckets.png)

**哈希表的遍历过程**

简单总结一下哈希表遍历的顺序，首先会选出一个绿色的正常桶开始遍历，随后遍历所有黄色的溢出桶，最后依次按照索引顺序遍历哈希表中其他的桶，直到所有的桶都被遍历完成。





### 字符串



遍历字符串的过程与数组、切片和哈希表非常相似，只是在遍历时会获取字符串中索引对应的字节并将字节转换成 `rune`。我们在遍历字符串时拿到的值都是 `rune` 类型的变量，`for i, r := range s {}` 的结构都会被转换成如下所示的形式：

```go
ha := s
for hv1 := 0; hv1 < len(ha); {
    hv1t := hv1
    hv2 := rune(ha[hv1])
    if hv2 < utf8.RuneSelf {
        hv1++
    } else {
        hv2, hv1 = decoderune(ha, hv1)
    }
    v1, v2 = hv1t, hv2
}
```



### 通道

使用 range 遍历 Channel 也是比较常见的做法，一个形如 `for v := range ch {}` 的语句最终会被转换成如下的格式：

```go
ha := a
hv1, hb := <-ha
for ; hb != false; hv1, hb = <-ha {
    v1 := hv1
    hv1 = nil
    ...
}
```

