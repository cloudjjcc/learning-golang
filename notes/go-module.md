# 环境变量

## GOPATH

$GOPATH/src

存放源代码

$GOPATH/bin

存放可执行文件

## GO111MODULE

此环境变量是go module 的开关

`GO111MODULE` 有三个值：`off`, `on`和`auto（默认值）`。

- `GO111MODULE=off`，go命令行将不会支持module功能，寻找依赖包的方式将会沿用旧版本那种通过vendor目录或者GOPATH模式来查找。

- `GO111MODULE=on`，go命令行会使用modules，而一点也不会去GOPATH目录下查找。

- ```
  GO111MODULE=auto
  ```

  ，默认值，go命令行将会根据当前目录来决定是否启用module功能。这种情况下可以分为两种情形：

  - 当前目录在GOPATH/src之外且该目录包含go.mod文件
  - 当前文件在包含go.mod文件的目录下面。

> 当modules 功能启用时，依赖包的存放位置变更为`$GOPATH/pkg`，允许同一个package多个版本并存，且多个项目可以共享缓存的 module。

可以通过go mod 命令来管理模块



## GOSUMDB



go checksum database ,用于在go拉取模块时的数据校验，其值可以是on获取off



## GOPRIVATE

设置代理



## go env -w 

可以通过`go env -w`  设置系统环境变量，它会在$HOME/.config/go/env 中存储，它不会覆盖系统环境变量





# go.mod文件





# 私有模块访问

GONOPROXY，GONOSUMDB和GOPRIVATE

它们三个的值都是逗号分割的URL

GOPRIVATE的值将作为GONOPROXY和GONOSUMDB的默认值，所以一般情况下我们只需要设置GOPRIVATE的值就行了





# 模块缓存



## 全局缓存

同一个模块版本的数据只缓存一份，所有其他模块共享使用

目前所有模块数据缓存在\$GOPATH/pkg/mod和$GOPATH/pkg/sum下

可以用go clean	-modcache 清理所有已缓存的数据