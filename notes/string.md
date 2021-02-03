# rune

+ `rune`用来表示`Unicode`码点，实际上是`int32`的别名

```go
// rune is an alias for int32 and is equivalent to int32 in all ways. It is
// used, by convention, to distinguish character values from integer values.
type rune = int32
```

+ 可以通过单引号声明`rune`字面量

```go
ru:='哈'
fmt.Printf("%T,%c\n",ru,ru)
// int32,哈
```

  

# string 字面量

+ 双引号（std）

+ 倒引号(raw)

```go
strA:="hello"
strB:=`
		小明
`
fmt.Print(strA,strB)
/*输出
hello
                小明
*/
```



# string的遍历

+ 可以通过`for-range` 的方式遍历字符串,得到的i是字节索引，v是`rune`类型

```go
for i, v := range "hi,我是小明" {
		fmt.Printf("i:%d,v:%c\n", i, v)
}
/* 输出
i:0,v:h,type v:int32
i:1,v:i,type v:int32
i:2,v:,,type v:int32
i:3,v:我,type v:int32
i:6,v:是,type v:int32
i:9,v:小,type v:int32
i:12,v:明,type v:int32
*/
```



# string 的长度

+ 可以通过内置函数`Len`获取string的长度得到的是字节长度。

+ 如果想要得到`rune`长度，可以使用`utf8.RuneCountInString(str)`
+ 可以直接将字符串转化为`[]byte`或`[]rune`

+ 不能对`string`使用`Cap`函数

```go
	demoStr := "hi,我是小明"
	lenStr := len(demoStr)
	lenBytes := len([]byte(demoStr))
	countRunes := utf8.RuneCountInString(demoStr)
	lenRunes := len([]rune(demoStr))
	fmt.Printf("lenStr:%d,lenBytes:%d,countRunes:%d,lenRunes:%d",lenStr,lenBytes,countRunes,lenRunes)
// lenStr:15,lenBytes:15,countRunes:7,lenRunes:7
```



# string 的拼接

+ 直接使用+号连接

  最终会调用 [`runtime.concatstrings`](https://draveness.me/golang/tree/runtime.concatstrings)
  如果同一个表达式使用多次+ 号，会一次计算出最终长度并将所有字符串拷贝到新分配的内存空间

  ```go
  str1:="hello"+"world"
  str2:= "hello" + `world`
  c := 'a' + 'b' // c 的类型是int32
  fmt.Printf("str1:%s\nstr2:%s\nc:%v",str1,str2,c)
  /* 输出
  str1:helloworld
  str2:helloworld
  c:195
  */
  ```

  

+ 使用fmt.Sprintf()函数

  支持格式化输出，内部使用[]byte,效率较低

  适合需要格式化不同类型数据的场景

+ 使用strings.Join

  如果待拼接的字符串已经在一个切片中则优先使用此方法，内部使用了string.Builder

  ```go
  strA := "hello"
  strB := "world"
  strC := strings.Join([]string{strA, strB}, "")
  ```

+ 使用bytes.Buffer

  ```go
  func BenchmarkConcatWithBuffer(b *testing.B) {
  	strA := "hello"
  	buffer := bytes.Buffer{}
  	buffer.Grow(b.N*5) // 提前分配内存
  	for i := 0; i < b.N; i++ {
  		buffer.WriteString(strA)
  	}
  }
  // BenchmarkConcatWithBuffer-12    	160003474	         6.47 ns/op
  ```

+ 使用string.Builder

  在需要多次连接字符串时的最佳选择

  ```go
  func BenchmarkConcatWithBuilder(b *testing.B) {
  	strA := "hello"
  	builder := strings.Builder{}
  	builder.Grow(b.N*5)
  	for i := 0; i < b.N; i++ {
  		builder.WriteString(strA)
  	}
  }
  // BenchmarkConcatWithBuilder-12    	317321160	         3.44 ns/op
  ```

+ strings.Builder 中的String()方法通过使用unsafe.Pointer 避免了[]byte转化为string时内存分配，故而性能更高

  ```go
  // String returns the accumulated string.
  func (b *Builder) String() string {
  	return *(*string)(unsafe.Pointer(&b.buf))
  }
  ```

  

![string-concat-and-copy](https://img.draveness.me/2019-12-31-15777265631620-string-concat-and-copy.png)

**字符串的拼接和拷贝**



# string 运行时结构

+ 字符串在 Go 语言中的接口其实非常简单，每一个字符串在运行时都会使用如下的 [`reflect.StringHeader`](https://draveness.me/golang/tree/reflect.StringHeader) 表示，其中包含指向字节数组的指针和数组的大小

```go
// StringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
type StringHeader struct {
	Data uintptr
	Len  int
}
```



# 类型转换

![string-bytes-conversion](https://img.draveness.me/2019-12-31-15777265631625-string-bytes-conversion.png)

**字符串和字节数组的转换**

字符串和 `[]byte` 中的内容虽然一样，但是字符串的内容是只读的，我们不能通过下标或者其他形式改变其中的数据，而 `[]byte` 中的内容是可以读写的。不过无论从哪种类型转换到另一种都需要拷贝数据，而内存拷贝的性能损耗会随着字符串和 `[]byte` 长度的增长而增长。

