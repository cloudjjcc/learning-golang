# 创建Map

map类型可以写为 map[K]V

其中K对应的key必须是支持==比较运算符的数据类型，所以map可以通过测试key是否相等来判断是否已经存在。

虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做key类型则是一个坏的想法,可能出现的NaN和任何浮点数都不相等

+ make 函数创建

+ map字面量

```go
	m:=map[string]int{}
	var m1 map[string]int
	m2:=make(map[string]int)
	fmt.Printf("m:%p,m1:%p,m2:%p\n",m,m1,m2)
//m:0xc00007c3c0,m1:0x0,m2:0xc00007c3f0
```

  

# 元素操作

+ 查找
+ 删除
+ len()
+ for-range
+ map之间不能进行相等比较；可以和nil进行比较

```go
	m := map[string]int{}
	m["math"] = 90
	m["math"]++
//addr:=&m["math"]  compile error: cannot take address of map element
	fmt.Println(len(m))
// 2
	for k, v := range m { // 注意：map的迭代顺序是不确定的
		fmt.Printf("%s:%d\n",k,v)
	}
/*输出
english:80
math:91
*/
	delete(m,"math")
	if v, ok := m["math"]; ok {
		fmt.Printf("math score is:%d",v)// not output
	}
```



map上的大部分操作，包括查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似。但是向一个nil值的map存入元素将导致一个panic异常

```go
  var m1 map[string]int //m1==nil
	m1["math"]=90 //panic: assignment to entry in nil map
```



# Hash

哈希函数应该能够将不同键映射到不同的索引上

解决hash碰撞：

+ 开放地址法

  [开放寻址法](https://en.wikipedia.org/wiki/Open_addressing)[2](https://draveness.me/golang/docs/part2-foundation/ch03-datastructure/golang-hashmap/#fn:2)是一种在哈希表中解决哈希碰撞的方法，这种方法的核心思想是**依次探测和比较数组中的元素以判断目标键值对是否存在于哈希表中**，如果我们使用开放寻址法来实现哈希表，那么实现哈希表底层的数据结构就是数组

  开放寻址法中对性能影响最大的是**装载因子**，它是数组中元素的数量与数组大小的比值。随着装载因子的增加，线性探测的平均用时就会逐渐增加，这会影响哈希表的读写性能。

+ 链地址法

  实现拉链法一般会使用数组加上链表，不过一些编程语言会在拉链法的哈希中引入红黑树以优化性能

  计算哈希、定位桶和遍历链表三个过程是哈希表读写操作的主要开销，使用拉链法实现的哈希也有装载因子这一概念：

  装载因子:=元素数量÷桶数量



# 初始化原理



# 插入原理



# 扩容



# 删除原理



# 遍历原理