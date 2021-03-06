# 堆内存分配

线程缓存分配（Thread-Caching Malloc，TCMalloc）是用于分配内存的机制

Go 语言的内存分配器就借鉴了 TCMalloc 的设计实现高速的内存分配，它的核心理念是使用多级缓存将对象根据大小分类，并按照类别实施不同的分配策略。



当我们使用new关键字，实际上调用了运行时函数newobject,最终调用mallocgc方法分配内存

```go
// implementation of new builtin
// compiler (both frontend and SSA backend) knows the signature
// of this function
func newobject(typ *_type) unsafe.Pointer {
   return mallocgc(typ.size, typ, true)
}
```



运行时分配器都会引入线程缓存（Thread Cache）、中心缓存（Central Cache）和页堆（Page Heap）三个组件分级管理内存：

![multi-level-cache](https://img.draveness.me/2020-02-29-15829868066457-multi-level-cache.png)

**图 7-6 多级缓存内存分配**



Go 语言的内存分配器包含内存管理单元、线程缓存、中心缓存和页堆几个重要组件，本节将介绍这几种最重要组件对应的数据结构 [`runtime.mspan`](https://draveness.me/golang/tree/runtime.mspan)、[`runtime.mcache`](https://draveness.me/golang/tree/runtime.mcache)、[`runtime.mcentral`](https://draveness.me/golang/tree/runtime.mcentral) 和 [`runtime.mheap`](https://draveness.me/golang/tree/runtime.mheap)，我们会详细介绍它们在内存分配器中的作用以及实现。

![go-memory-layout](https://img.draveness.me/2020-02-29-15829868066479-go-memory-layout.png)

**图 7-10 Go 程序的内存布局**





## 微对象

tinyAllocator



## 小对象





## 大对象

Mheap



# 垃圾回收



# 栈内存

栈区的内存一般由编译器自动分配和释放

Go 语言的汇编代码包含 BP 和 SP 两个栈寄存器，它们分别存储了栈的基址指针和栈顶的地址，栈内存与函数调用的关系非常紧密，我们在函数调用一节中曾经介绍过栈区，BP 和 SP 之间的内存就是当前函数的调用栈。

![stack-registers](https://img.draveness.me/2020-03-23-15849514795843-stack-registers.png)

**图 7-43 栈寄存器与内存**



## 逃逸分析



Go 语言的编译器使用逃逸分析决定哪些变量应该在栈上分配，哪些变量应该在堆上分配，其中包括使用 `new`、`make` 和字面量等方法隐式分配的内存，Go 语言的逃逸分析遵循以下两个不变性：

1. 指向栈对象的指针不能存在于堆中；
2. 指向栈对象的指针不能在栈对象回收后存活；



