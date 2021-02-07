下列情况`x`是不可以寻址的，你不能使用`&x`取得指针：

- 字符串中的字节:
- map对象中的元素
- 接口对象的动态值(通过type assertions获得)
- 常数
- literal值(非composite literal)
- package 级别的函数
- 方法method (用作函数值)
- 中间值(intermediate value):
  - 函数调用
  - 显式类型转换
  - 各种类型的操作 （除了指针引用pointer dereference操作
    - channel receive operations
    - sub-string operations
    - sub-slice operations
    - 加减乘除等运算符



golang中不能寻址的可以总结为：不可变的，临时结果和不安全的。只要符合其中任何一个条件，它就是不可以寻址的。





- 常数为什么不可以寻址?： 如果可以寻址的话，我们可以通过指针修改常数的值，破坏了常数的定义。
- map的元素为什么不可以寻址？:两个原因，如果对象不存在，则返回零值，零值是不可变对象，所以不能寻址，如果对象存在，因为Go中map实现中元素的地址是变化的，这意味着寻址的结果是无意义的。
- 为什么slice不管是否可寻址，它的元素读是可以寻址的？:因为slice底层实现了一个数组，它是可以寻址的。
- 为什么字符串中的字符/字节又不能寻址呢：因为字符串是不可变的。

规范中还有几处提到了 `addressable`:

- 调用一个receiver为指针类型的方法时，使用一个addressable的值将自动获取这个值的指针
- `++`、`--`语句的操作对象必须是addressable或者是map的index操作
- 赋值语句`=`的左边对象必须是addressable,或者是map的index操作，或者是`_`
- 上条同样使用`for ... range`语句



