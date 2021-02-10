# time包

## Timer

Timer表示一个单次的计时

```go
// The Timer type represents a single event.
// When the Timer expires, the current time will be sent on C,
// unless the Timer was created by AfterFunc.
// A Timer must be created with NewTimer or AfterFunc.
type Timer struct {
   C <-chan Time
   r runtimeTimer
}
```

可以通过`time.NewTimer`函数或者`time.AfterFunc`创建一个定时器







## Ticker





# 数据结构

[`runtime.timer`](https://draveness.me/golang/tree/runtime.timer) 是 Go 语言计时器的内部表示，每一个计时器都存储在对应处理器的最小四叉堆中，下面是运行时计时器对应的结构体：

```go
type timer struct {
	pp puintptr

	when     int64
	period   int64
	f        func(interface{}, uintptr)
	arg      interface{}
	seq      uintptr
	nextwhen int64
	status   uint32
}
```

- `when` — 当前计时器被唤醒的时间；
- `period` — 两次被唤醒的间隔；
- `f` — 每当计时器被唤醒时都会调用的函数；
- `arg` — 计时器被唤醒时调用 `f` 传入的参数；
- `nextWhen` — 计时器处于 `timerModifiedXX` 状态时，用于设置 `when` 字段；
- `status` — 计时器的状态；

然而这里的 [`runtime.timer`](https://draveness.me/golang/tree/runtime.timer) 只是计时器运行时的私有结构体，对外暴露的计时器使用 [`time.Timer`](https://draveness.me/golang/tree/time.Timer)结体：

```go
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
```

[`time.Timer`](https://draveness.me/golang/tree/time.Timer) 计时器必须通过 [`time.NewTimer`](https://draveness.me/golang/tree/time.NewTimer)、[`time.AfterFunc`](https://draveness.me/golang/tree/time.AfterFunc) 或者 [`time.After`](https://draveness.me/golang/tree/time.After) 函数创建。 当计时器失效时，订阅计时器 Channel 的 Goroutine 会收到计时器失效的时间。



所有的计时器都以最小四叉堆的形式存储在处理器 [`runtime.p`](https://draveness.me/golang/tree/runtime.p) 中。

处理器 [`runtime.p`](https://draveness.me/golang/tree/runtime.p) 中与计时器相关的有以下字段：

- `timersLock` — 用于保护计时器的互斥锁；
- `timers` — 存储计时器的最小四叉堆；
- `numTimers` — 处理器中的计时器数量；
- `adjustTimers` — 处理器中处于 `timerModifiedEarlier` 状态的计时器数量；
- `deletedTimers` — 处理器中处于 `timerDeleted` 状态的计时器数量；



# 触发计时器

Go 语言会在两个模块触发计时器，运行计时器中保存的函数：

- 调度器调度时会检查处理器中的计时器是否准备就绪；
- 系统监控会检查是否有未执行的到期计时器；





## 调度器

[`runtime.checkTimers`](https://draveness.me/golang/tree/runtime.checkTimers) 是调度器用来运行处理器中计时器的函数，它会在发生以下情况时被调用：

- 调度器调用 [`runtime.schedule`](https://draveness.me/golang/tree/runtime.schedule) 执行调度时；
- 调度器调用 [`runtime.findrunnable`](https://draveness.me/golang/tree/runtime.findrunnable) 获取可执行的 Goroutine 时；
- 调度器调用 [`runtime.findrunnable`](https://draveness.me/golang/tree/runtime.findrunnable) 从其他处理器窃取计时器时；





