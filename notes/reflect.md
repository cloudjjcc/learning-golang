[`reflect`](https://golang.org/pkg/reflect/) 实现了运行时的反射能力，能够让程序操作不同类型的对象[1](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/#fn:1)。反射包中有两对非常重要的函数和类型，两个函数分别是：

- [`reflect.TypeOf`](https://draveness.me/golang/tree/reflect.TypeOf) 能获取类型信息；
- [`reflect.ValueOf`](https://draveness.me/golang/tree/reflect.ValueOf) 能获取数据的运行时表示；



# 三大原则

运行时反射是程序在运行期间检查其自身结构的一种方式。反射带来的灵活性是一把双刃剑，反射作为一种元编程方式可以减少重复代码[2](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/#fn:2)，但是过量的使用反射会使我们的程序逻辑变得难以理解并且运行缓慢。我们在这一节中会介绍 Go 语言反射的三大法则[3](https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-reflect/#fn:3)，其中包括：

1. 从 `interface{}` 变量可以反射出反射对象；
2. 从反射对象可以获取 `interface{}` 变量；
3. 要修改反射对象，其值必须可设置；