当我们在 Go 语言中使用 `select` 控制结构时，会遇到两个有趣的现象：

1. `select` 能在 Channel 上进行非阻塞的收发操作；
2. `select` 在遇到多个 Channel 同时响应时，会随机执行一种情况；



### 非阻塞的收发

在通常情况下，`select` 语句会阻塞当前 Goroutine 并等待多个 Channel 中的一个达到可以收发的状态。但是如果 `select` 控制结构中包含 `default` 语句，那么这个 `select` 语句在执行时会遇到以下两种情况：

1. 当存在可以收发的 Channel 时，直接处理该 Channel 对应的 `case`；
2. 当不存在可以收发的 Channel 时，执行 `default` 中的语句；





### 直接阻塞



首先介绍的是最简单的情况，也就是当 `select` 结构中不包含任何 `case`。我们截取 [`cmd/compile/internal/gc.walkselectcases`](https://draveness.me/golang/tree/cmd/compile/internal/gc.walkselectcases) 函数的前几行代码：

```go
func walkselectcases(cases *Nodes) []*Node {
	n := cases.Len()

	if n == 0 {
		return []*Node{mkcall("block", nil, nil)}
	}
	...
}
```

这段代码很简单并且容易理解，它直接将类似 `select {}` 的语句转换成调用 [`runtime.block`](https://draveness.me/golang/tree/runtime.block) 函数：

```go
func block() {
	gopark(nil, nil, waitReasonSelectNoCases, traceEvGoStop, 1)
}
```

[`runtime.block`](https://draveness.me/golang/tree/runtime.block) 的实现非常简单，它会调用 [`runtime.gopark`](https://draveness.me/golang/tree/runtime.gopark) 让出当前 Goroutine 对处理器的使用权并传入等待原因 `waitReasonSelectNoCases`。

简单总结一下，空的 `select` 语句会直接阻塞当前 Goroutine，导致 Goroutine 进入无法被唤醒的永久休眠状态。





### 单一管道 

如果当前的 `select` 条件只包含一个 `case`，那么编译器会将 `select` 改写成 `if` 条件语句。下面对比了改写前后的代码：

```go
// 改写前
select {
case v, ok <-ch: // case ch <- v
    ...    
}

// 改写后
if ch == nil {
    block()
}
v, ok := <-ch // case ch <- v
...
```



#### 发送 

首先是 Channel 的发送过程，当 `case` 中表达式的类型是 `OSEND` 时，编译器会使用条件语句和 [`runtime.selectnbsend`](https://draveness.me/golang/tree/runtime.selectnbsend) 函数改写代码：

```go
select {
case ch <- i:
    ...
default:
    ...
}

if selectnbsend(ch, i) {
    ...
} else {
    ...
}
```

这段代码中最重要的就是 [`runtime.selectnbsend`](https://draveness.me/golang/tree/runtime.selectnbsend)，它为我们提供了向 Channel 非阻塞地发送数据的能力。我们在 Channel 一节介绍了向 Channel 发送数据的 [`runtime.chansend`](https://draveness.me/golang/tree/runtime.chansend) 函数包含一个 `block` 参数，该参数会决定这一次的发送是不是阻塞的：

```go
func selectnbsend(c *hchan, elem unsafe.Pointer) (selected bool) {
	return chansend(c, elem, false, getcallerpc())
}
```

由于我们向 [`runtime.chansend`](https://draveness.me/golang/tree/runtime.chansend) 函数传入了非阻塞，所以在不存在接收方或者缓冲区空间不足时，当前 Goroutine 都不会阻塞而是会直接返回。



#### 接收

由于从 Channel 中接收数据可能会返回一个或者两个值，所以接收数据的情况会比发送稍显复杂，不过改写的套路是差不多的：

```go
// 改写前
select {
case v <- ch: // case v, ok <- ch:
    ......
default:
    ......
}

// 改写后
if selectnbrecv(&v, ch) { // if selectnbrecv2(&v, &ok, ch) {
    ...
} else {
    ...
}
```

返回值数量不同会导致使用函数的不同，两个用于非阻塞接收消息的函数 [`runtime.selectnbrecv`](https://draveness.me/golang/tree/runtime.selectnbrecv)和 [`runtime.selectnbrecv2`](https://draveness.me/golang/tree/runtime.selectnbrecv2) 只是对 [`runtime.chanrecv`](https://draveness.me/golang/tree/runtime.chanrecv) 返回值的处理稍有不同：

```go
func selectnbrecv(elem unsafe.Pointer, c *hchan) (selected bool) {
	selected, _ = chanrecv(c, elem, false)
	return
}

func selectnbrecv2(elem unsafe.Pointer, received *bool, c *hchan) (selected bool) {
	selected, *received = chanrecv(c, elem, false)
	return
}
```

因为接收方不需要，所以 [`runtime.selectnbrecv`](https://draveness.me/golang/tree/runtime.selectnbrecv) 会直接忽略返回的布尔值，而 [`runtime.selectnbrecv2`](https://draveness.me/golang/tree/runtime.selectnbrecv2) 会将布尔值回传给调用方。与 [`runtime.chansend`](https://draveness.me/golang/tree/runtime.chansend) 一样，[`runtime.chanrecv`](https://draveness.me/golang/tree/runtime.chanrecv) 也提供了一个 `block` 参数用于控制这次接收是否阻塞。



