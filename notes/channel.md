



#  



# 创建

channel只能用make创建



```Go
ch := make(chan int) // ch has type 'chan int'
ch2:= make(chan int,3) // chan with buffer
```







# 可比较

两个相同类型的channel可以使用==运算符比较。如果两个channel引用的是相同的对象，那么比较的结果为真。一个channel也可以和nil进行比较



# 单向通道

### 只写

chan<- int

### 只读

<-chan int

# 支持操作



### 发送



### 接收



### 关闭





