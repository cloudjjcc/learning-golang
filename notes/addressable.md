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

