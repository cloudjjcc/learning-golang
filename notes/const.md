



在同一个 const group 中，如果常量定义与前一行的定义一致，则可以省略类型和值。编译时，会按照前一行的定义自动补全。

```go
const (
		a, b = "golang", 100
		d, e
		f bool = true
		g
	)
fmt.Println(d, e, g)
//golang 100 true
```





常量分为无类型常量和有类型常量两种，`const N = 100`，属于无类型常量，赋值给其他变量时，如果字面量能够转换为对应类型的变量，则赋值成功，例如，`var x int = N`。但是对于有类型的常量 `const M int32 = 100`，赋值给其他变量时，需要类型匹配才能成功

```go

	const N = 100
	var x int = N

	const M int32 = 100
  var y int = M //compiler err:cannot use M (type int32) as type int in assignment
	fmt.Println(x, y)
```

