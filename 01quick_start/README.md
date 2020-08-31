## GOlang的特点

Go 语言除了简洁、优雅之外，还有一些非常强大的特性。包括：
并发与协程
基于消息传递的通信方式
丰富实用的内置数据类型
函数多返回值
defer 机制
反射(reflect)
高性能 HTTP Server

## 实际使用

```
// 要生成 Go 可执行程序，必须建立一个名为 main 的 package，
// 并且在该 package 中必须包含一个名为 main() 的函数。
// GoPath 是 Go 的工作目录，GoRoot 是 Go 的安装目录。
```
使用 go env 命令可以查看 GoPath 和 GoRoot，类似fmt以及math包都在GoRoot目录下，GitHub 上的开源包就要使用GoPath。


## GoPath


GoPath 目录约定有三个子目录
* src：存放源代码。按照 Go 语言约定，go run，go install 等命令默认会在此路径下执行；
* pkg：存放编译时生成的中间文件（ * .a ）；
* bin： 存放编译后生成的可执行文件 （ 在项目内执行 go install，会在 bin 目录下生成一个可执行文件）。
当我们要引用 GitHub 上的开源包时，比如使用 github.com/gomodule/redigo/redis 这个包来用 golang 进行 redis 的操作。可以执行 go get github.com/gomodule/redigo/redis 命令，将会在 GoPath 的 src 目录下生成一个 /github.com/gomodule/redigo 这样的目录结构，后面即可在项目中像引用fmt包一样引用。
