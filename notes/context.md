上下文 [`context.Context`](https://draveness.me/golang/tree/context.Context) Go 语言中用来设置截止日期、同步信号，传递请求相关值的结构体。


```go
type Context interface {
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}
```





## 默认上下文 [#](https://draveness.me/golang/docs/part3-runtime/ch06-concurrency/golang-context/#612-默认上下文)

[`context`](https://github.com/golang/go/tree/master/src/context) 包中最常用的方法还是 [`context.Background`](https://draveness.me/golang/tree/context.Background)、[`context.TODO`](https://draveness.me/golang/tree/context.TODO)，这两个方法都会返回预先初始化好的私有变量 `background` 和 `todo`，它们会在同一个 Go 程序中被复用：

```go
func Background() Context {
	return background
}

func TODO() Context {
	return todo
}
```

这两个私有变量都是通过 `new(emptyCtx)` 语句初始化的，它们是指向私有结构体 [`context.emptyCtx`](https://draveness.me/golang/tree/context.emptyCtx) 的指针，这是最简单、最常用的上下文类型：

```go
type emptyCtx int

func (*emptyCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*emptyCtx) Done() <-chan struct{} {
	return nil
}

func (*emptyCtx) Err() error {
	return nil
}

func (*emptyCtx) Value(key interface{}) interface{} {
	return nil
}
```

从上述代码中，我们不能够发现 [`context.emptyCtx`](https://draveness.me/golang/tree/context.emptyCtx) 通过空方法实现了 [`context.Context`](https://draveness.me/golang/tree/context.Context) 接口中的所有方法，它没有任何功能。





