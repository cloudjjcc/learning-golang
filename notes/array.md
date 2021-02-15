# 初始化

```go
arr1 := [3]int{1, 2, 3}
arr2 := [...]int{1, 2, 3} // 编译器推导数组大小

var arr3 [3]int
arr3[0]=1
arr3[1]=2
arr3[2]=3
```

+ 数组初始化后大小无法改变

+ 编译期间的数组类型是由上述的 [`cmd/compile/internal/types.NewArray`](https://draveness.me/golang/tree/cmd/compile/internal/types.NewArray) 函数生成的，该类型包含两个字段，分别是元素类型 `Elem` 和数组的大小 `Bound`，这两个字段共同构成了数组类型，而当前数组是否应该在堆栈中初始化也在编译期就确定了

```go
func NewArray(elem *Type, bound int64) *Type {
	if bound < 0 {
		Fatalf("NewArray: invalid bound %v", bound)
	}
	t := New(TARRAY)
	t.Extra = &Array{Elem: elem, Bound: bound}
	t.SetNotInHeap(elem.NotInHeap())
	return t
}
```



- 对于一个由字面量组成的数组，根据数组元素数量的不同，编译器会在负责初始化字面量的 [`cmd/compile/internal/gc.anylit`](https://draveness.me/golang/tree/cmd/compile/internal/gc.anylit) 函数中做两种不同的优化：

  - 当元素数量小于或者等于 4 个时，会直接将数组中的元素放置在栈上；

  - 当元素数量大于 4 个时，会将数组中的元素放置到静态区并在运行时取出；



# 访问和赋值

```go
a:=arr[0]
// 元素赋值
arr[0]=3
// 作为函数参数时，会将数组整个进行拷贝
f(arr)// func f([3]int)
```

+ 编译时检查

  Go 语言中可以在编译期间的静态类型检查判断数组越界，[`cmd/compile/internal/gc.typecheck1`](https://draveness.me/golang/tree/cmd/compile/internal/gc.typecheck1) 会验证访问数组的索引

  1. 访问数组的索引是非整数时，报错 “non-integer array index %v”；
  2. 访问数组的索引是负数时，报错 “invalid array index %v (index must be non-negative)"；
  3. 访问数组的索引越界时，报错 “invalid array index %v (out of bounds for %d-element array)"；

+ 运行时检查

  Go 语言为数组的访问操作生成了判断数组上限的指令 `IsInBounds` 以及当条件不满足时触发程序崩溃的 `PanicBounds` 指令

对数组的访问和赋值需要同时依赖编译器和运行时，它的大多数操作在[编译期间](https://draveness.me/golang/docs/part1-prerequisite/ch02-compile/golang-compile-intro/)都会转换成直接读写内存，在中间代码生成期间，编译器还会插入运行时方法 [`runtime.panicIndex`](https://draveness.me/golang/tree/runtime.panicIndex)调用防止发生越界错误。





# 什么时候用数组

TODO

