### 栈调整

intel 或 AT&T 汇编提供了 push 和 pop 指令族，plan9 中没有 push 和 pop，栈的调整是通过对硬件 SP 寄存器进行运算来实现的，例如:

```go
SUBQ $0x18, SP // 对 SP 做减法，为函数分配函数栈帧
...               // 省略无用代码
ADDQ $0x18, SP // 对 SP 做加法，清除函数栈帧
```

通用的指令和 IA64 平台差不多，下面分节详述。



### 数据搬运

常数在 plan9 汇编用 \$num 表示，可以为负数，默认情况下为十进制。可以用 $0x123 的形式来表示十六进制数。

```go
MOVB $1, DI      // 1 byte
MOVW $0x10, BX   // 2 bytes
MOVD $1, DX      // 4 bytes
MOVQ $-10, AX     // 8 bytes
```

可以看到，搬运的长度是由 MOV 的后缀决定的，这一点与 intel 汇编稍有不同，看看类似的 IA64 汇编:

```asm
mov rax, 0x1   // 8 bytes
mov eax, 0x100 // 4 bytes
mov ax, 0x22   // 2 bytes
mov ah, 0x33   // 1 byte
mov al, 0x44   // 1 byte
```

plan9 的汇编的操作数的方向是和 intel 汇编相反的，与 AT&T 类似。

```go
MOVQ $0x10, AX ===== mov rax, 0x10
       |    |------------|      |
       |------------------------|
```



### 常见计算指令

```go
ADDQ  AX, BX   // BX += AX
SUBQ  AX, BX   // BX -= AX
IMULQ AX, BX   // BX *= AX
```

类似数据搬运指令，同样可以通过修改指令的后缀来对应不同长度的操作数。例如 ADDQ/ADDW/ADDL/ADDB。





### 条件跳转/无条件跳转

```go
// 无条件跳转
JMP addr   // 跳转到地址，地址可为代码中的地址，不过实际上手写不会出现这种东西
JMP label  // 跳转到标签，可以跳转到同一函数内的标签位置
JMP 2(PC)  // 以当前指令为基础，向前/后跳转 x 行
JMP -2(PC) // 同上

// 有条件跳转
JNZ target // 如果 zero flag 被 set 过，则跳转
```



### 通用寄存器

plan9 中使用寄存器不需要带 r 或 e 的前缀，例如 rax，只要写 AX 即可:

```go
MOVQ $101, AX = mov rax, 101
```

下面是通用通用寄存器的名字在 IA64 和 plan9 中的对应关系:

| IA64  | RAX  | RBX  | RCX  | RDX  | RDI  | RSI  | RBP  | RSP  | R8   | R9   | R10  | R11  | R12  | R13  | R14  | RIP  |
| :---- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| Plan9 | AX   | BX   | CX   | DX   | DI   | SI   | BP   | SP   | R8   | R9   | R10  | R11  | R12  | R13  | R14  | PC   |