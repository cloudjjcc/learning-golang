NGINX日志统计UV


```shell
awk '{print $1}' access.log | sort | uniq -c | wc -l
```







## Linux shell 可执行的命令有三种：

### 内建命令

- 集成在shell解释器中
- 改变shell本省的属性设置
- io等
- 在执行命令时没有进程的创建与销毁





### shell函数



### 外部命令

- 调用fork函数
- 查找外部命令的位置
- 已找到的程序替代shell执行
- 子进程执行完毕后父进程接着执行下一条命令



使用source执行shell脚本时，不会创建子进程







