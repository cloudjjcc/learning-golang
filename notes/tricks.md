总结一些go语言小技巧



# 编译时检查接口实现

```go
var _ error = (*MyError)(nil)//断言MyError实现error接口
```

这个技巧配合GOLAND IDE可以很方便在结构提上实现指定接口的所有方法







# []byte转string



```go
*(*string)(unsafe.Pointer(&b.buf))//借助指针转换
```

这个技巧在`strings.Builder.String()`方法中用到，可以不经过内存拷贝将字节切片转为string







