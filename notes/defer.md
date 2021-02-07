

### 基本语法

defer 用在函数中，后面跟着函数调用表达式，会在函数返回时执行。

```go
func TestDefer(t *testing.T) {
	defer fmt.Println("bey bey")
	fmt.Println("do something ...")
}
/*
do something ...
bey bey
*/
```





### 作用域

`defer` 传入的函数不是在退出代码块的作用域时执行的，它只会在当前函数和方法返回之前被调用。



### 预计算参数

调用 `defer` 关键字会立刻拷贝函数中引用的外部参数

```go
func TestDefer2(t *testing.T) {
	start:=time.Now()
	defer fmt.Println(time.Since(start))// 会立刻执行time.Since,导致统计到的时间不是函数花费的时间
	time.Sleep(1*time.Second)
}
//183ns
```







## 数据结构 

在介绍 `defer` 函数的执行过程与实现原理之前，我们首先来了解一下 `defer` 关键字在 Go 语言源代码中对应的数据结构：

```go
type _defer struct {
	siz       int32
	started   bool
	openDefer bool
	sp        uintptr
	pc        uintptr
	fn        *funcval
	_panic    *_panic
	link      *_defer
}
```

[`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 结构体是延迟调用链表上的一个元素，所有的结构体都会通过 `link` 字段串联成链表。

![golang-defer-link](https://img.draveness.me/2020-01-19-15794017184603-golang-defer-link.png)

**图 5-10 延迟调用链表**

我们简单介绍一下 [`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 结构体中的几个字段：

- `siz` 是参数和结果的内存大小；
- `sp` 和 `pc` 分别代表栈指针和调用方的程序计数器；
- `fn` 是 `defer` 关键字中传入的函数；
- `_panic` 是触发延迟调用的结构体，可能为空；
- `openDefer` 表示当前 `defer` 是否经过开放编码的优化；

除了上述的这些字段之外，[`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 中还包含一些垃圾回收机制使用的字段，这里为了减少理解的成本就都省去了。



`runtime._defer` 结构体的分配：

+ 堆分配

+ 栈分配（1.13）

+ 开放编码（1.14）

- [`runtime.deferproc`](https://draveness.me/golang/tree/runtime.deferproc) 负责创建新的延迟调用；
- [`runtime.deferreturn`](https://draveness.me/golang/tree/runtime.deferreturn) 负责在函数调用结束时执行所有的延迟调用；



### 创建延迟调用

[`runtime.deferproc`](https://draveness.me/golang/tree/runtime.deferproc) 会为 `defer` 创建一个新的 [`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 结构体、设置它的函数指针 `fn`、程序计数器 `pc` 和栈指针 `sp` 并将相关的参数拷贝到相邻的内存空间中：

```go
func deferproc(siz int32, fn *funcval) {
	sp := getcallersp()
	argp := uintptr(unsafe.Pointer(&fn)) + unsafe.Sizeof(fn)
	callerpc := getcallerpc()

	d := newdefer(siz)
	if d._panic != nil {
		throw("deferproc: d.panic != nil after newdefer")
	}
	d.fn = fn
	d.pc = callerpc
	d.sp = sp
	switch siz {
	case 0:
	case sys.PtrSize:
		*(*uintptr)(deferArgs(d)) = *(*uintptr)(unsafe.Pointer(argp))
	default:
		memmove(deferArgs(d), unsafe.Pointer(argp), uintptr(siz))
	}

	return0()
}
```

最后调用的 [`runtime.return0`](https://draveness.me/golang/tree/runtime.return0) 是唯一一个不会触发延迟调用的函数，它可以避免递归 [`runtime.deferreturn`](https://draveness.me/golang/tree/runtime.deferreturn) 的递归调用。

[`runtime.deferproc`](https://draveness.me/golang/tree/runtime.deferproc) 中 [`runtime.newdefer`](https://draveness.me/golang/tree/runtime.newdefer) 的作用是想尽办法获得 [`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 结构体，这里包含三种路径：

1. 从调度器的延迟调用缓存池 `sched.deferpool` 中取出结构体并将该结构体追加到当前 Goroutine 的缓存池中；
2. 从 Goroutine 的延迟调用缓存池 `pp.deferpool` 中取出结构体；
3. 通过 [`runtime.mallocgc`](https://draveness.me/golang/tree/runtime.mallocgc) 在堆上创建一个新的结构体；



无论使用哪种方式，只要获取到 [`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 结构体，它都会被追加到所在 Goroutine `_defer`链表的最前面。

![golang-new-defer](https://img.draveness.me/2020-01-19-15794017184614-golang-new-defer.png)

**追加新的延迟调用**

`defer` 关键字的插入顺序是从后向前的，而 `defer` 关键字执行是从前向后的，这也是为什么后调用的 `defer` 会优先执行。





### 执行延迟调用 

[`runtime.deferreturn`](https://draveness.me/golang/tree/runtime.deferreturn) 会从 Goroutine 的 `_defer` 链表中取出最前面的 [`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 并调用 [`runtime.jmpdefer`](https://draveness.me/golang/tree/runtime.jmpdefer) 传入需要执行的函数和参数：

```go
func deferreturn(arg0 uintptr) {
	gp := getg()
	d := gp._defer
	if d == nil {
		return
	}
	sp := getcallersp()
	...

	switch d.siz {
	case 0:
	case sys.PtrSize:
		*(*uintptr)(unsafe.Pointer(&arg0)) = *(*uintptr)(deferArgs(d))
	default:
		memmove(unsafe.Pointer(&arg0), deferArgs(d), uintptr(d.siz))
	}
	fn := d.fn
	gp._defer = d.link
	freedefer(d)
	jmpdefer(fn, uintptr(unsafe.Pointer(&arg0)))
}
```

[`runtime.jmpdefer`](https://draveness.me/golang/tree/runtime.jmpdefer) 是一个用汇编语言实现的运行时函数，它的主要工作是跳转到 `defer` 所在的代码段并在执行结束之后跳转回 [`runtime.deferreturn`](https://draveness.me/golang/tree/runtime.deferreturn)。



## 栈上分配 

在默认情况下，我们可以看到 Go 语言中 [`runtime._defer`](https://draveness.me/golang/tree/runtime._defer) 结构体都会在堆上分配，如果我们能够将部分结构体分配到栈上就可以节约内存分配带来的额外开销。

Go 语言团队在 1.13 中对 `defer` 关键字进行了优化，当该关键字在函数体中最多执行一次时，编译期间的 [`cmd/compile/internal/gc.state.call`](https://draveness.me/golang/tree/cmd/compile/internal/gc.state.call) 会将结构体分配到栈上并调用 [`runtime.deferprocStack`](https://draveness.me/golang/tree/runtime.deferprocStack)：





## 开放编码

Go 语言在 1.14 中通过开发编码（Open Coded）实现 `defer` 关键字，该设计使用代码内联优化 `defer` 关键的额外开销并引入函数数据 `funcdata` 管理 `panic` 的调用[3](https://draveness.me/golang/docs/part2-foundation/ch05-keyword/golang-defer/#fn:3)，该优化可以将 `defer` 的调用开销从 1.13 版本的 ~35ns 降低至 ~6ns 左右：

然而开放编码作为一种优化 `defer` 关键字的方法，它不是在所有的场景下都会开启的，开发编码只会在满足以下的条件时启用：

1. 函数的 `defer` 数量少于或者等于 8 个；
2. 函数的 `defer` 关键字不能在循环中执行；
3. 函数的 `return` 语句与 `defer` 语句的乘积小于或者等于 15 个；