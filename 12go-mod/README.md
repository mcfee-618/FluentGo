## go mod

module是一个相关Go包的集合，它是源代码更替和版本控制的单元。模块由源文件形成的go.mod文件的根目录定义，包含go.mod文件的目录也被称为模块根。modules取代旧的的基于GOPATH方法来指定在工程中使用哪些源文件或导入包。模块路径是导入包的路径前缀，go.mod文件定义模块路径，并且列出了在项目构建过程中使用的特定版本。

* GOPATH的缺陷

  * 不支持多个版本：当有多个项目时，不同项目对于依赖库的版本需求不一致时，无法在一个GOPATH下面放置不同版本的依赖项。典型的例子：当有多项目时候，A项目依赖C 1.0.0，B项目依赖C 2.0.0，由于没有依赖项版本的概念，C 1.0.0和C 2.0.0无法同时在GOPATH下共存，解决办法是分别为A项目和B项目设置GOPATH，将不同版本的C源代码放在两个GOPATH中，彼此独立.

  * 依赖项列表无法数据化：在Go Modules之前，没有任何语义化的数据可以知道当前项目的所有依赖项，需要手动找出所有依赖。对项目而言，需要将所有的依赖项全部放入源代码控制中。如果剔除某个依赖，需要在源代码中手工确认某个依赖是否剔除。

## 使用Go Modules

Go Modules是语义化版本管理的依赖项的包管理工具；它解决了GOPATH存在的缺陷，最重要的是，它是Go官方出品。

* go mod命令

```
    go mod init：生成 go.mod 文件
    go mod download : 下载 go.mod 文件中指明的所有依赖
    go mod tidy: 整理现有的依赖,运行下面命令可以移出所有代码中不需要的包：
    go mod graph: 查看现有的依赖结构
    go mod edit : 编辑 go.mod 文件
    go mod vendor : 导出项目所有的依赖到vendor目录
    go mod verify: 校验一个模块是否被篡改过
    go mod why: 查看为什么需要依赖某模块
```

*  go mod环境变量

Go语言提供了 GO111MODULE这个环境变量来作为 Go modules 的开关，其允许设置以下参数：

auto：只要项目包含了 go.mod 文件的话启用 Go modules，目前在 Go1.11 至 Go1.14 中仍然是默认值。
on：启用 Go modules，推荐设置，将会是未来版本中的默认值。
off：禁用 Go modules，不推荐设置。

```
    go env -w GO111MODULE=on 开启Go Modules
```

* 使用步骤

  * go mod init xxx 初始化Go modules 生成go.mod文件
  * go get xxx库 下载源码包并执行go install，会更新go.mod文件
  
  ```
  module github.com/aceld/modules_test 模块路径
  go 1.14 当前Go版本
  require github.com/aceld/zinx v0.0.0-20200221135252-8a8954e75100 // indirect  
  require当前项目依赖的一个特定的必须版本
  // indirect表示该模块为间接依赖
  ```

  * go.sum文件:拉取模块依赖后，会发现多出了一个 go.sum 文件，其详细罗列了当前项目直接或间接依赖的所有模块版本，并写明了那些模块版本的 SHA-256 哈希值，h1 hash 是 Go modules 将目标模块版本的 zip 文件开包后，针对所有包内文件依次进行 hash，然后再把它们的 hash 结果按照固定格式和算法组成总的 hash 值。


  ## go mod替换版本


  go mod edit -replace=old[@v]=new[@v]

  https://www.jianshu.com/p/defde13823f6
  https://www.cnblogs.com/sunsky303/p/12150575.html