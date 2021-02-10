系统监控是 Go 语言运行时的重要组成部分，它会每隔一段时间检查 Go 语言运行时，确保程序没有进入异常状态。



# 监控循环



`runtime.main`函数中会新创建一个m,并在上面执行sysmon函数

```go
func main() {
	...
	if GOARCH != "wasm" {
		systemstack(func() {
			newm(sysmon, nil)
		})
	}
	...
}
```



## 检测死锁

