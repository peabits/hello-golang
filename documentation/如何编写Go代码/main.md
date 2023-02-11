
# 如何编写Go代码

[TOC]

## 简介

本文演示了在模块中开发一个简单的Go包，并介绍了go工具，这是获取、构建和安装Go模块、包和命令的标准方式。

注意：本文档假设您使用的是`Go 1.13`或更高版本，并且没有设置`GO111MODULE`环境变量。如果你正在寻找本文档的旧版本，即模块之前的版本，它被归档在[这里](./previous.md)。

## 代码组织

Go程序被组织成包。包是同一目录下的源文件的集合，它们被编译在一起。在一个源文件中定义的函数、类型、变量和常量对同一包内的所有其他源文件都是可见的。

一个资源库包含一个或多个模块。一个模块是相关Go包的集合，它们被一起发布。一个Go代码库通常只包含一个模块，位于代码库的根部。那里有一个名为`go.mod`的文件声明了模块路径：模块内所有包的导入路径前缀。该模块包含包含其`go.mod`文件的目录中的包，以及该目录的子目录，直到包含另一个`go.mod`文件的下一个子目录（如果有的话）。

注意，你不需要在构建之前将你的代码发布到一个远程仓库。一个模块可以在本地定义而不属于一个版本库。然而，组织你的代码是一个好习惯，就像你有一天会发布它一样。

每个模块的路径不仅作为其包的导入路径前缀，而且还指出go命令应该在哪里下载它。例如，为了下载`golang.org/x/tools`模块，go命令会查阅[https://golang.org/x/tools](https://golang.org/x/tools)（这里有更多描述）所指示的仓库。
``
导入路径是一个用于导入软件包的字符串。一个包的导入路径是它的模块路径与模块中的子目录相连接。例如，模块`github.com/google/go-cmp`在`cmp/`目录下包含一个包。该包的导入路径是`github.com/google/go-cmp/cmp`。标准库中的包没有模块路径前缀。
``

## 你的第一个程序

要编译和运行一个简单的程序，首先选择一个模块路径（我们将使用`example/user/hello`）并创建一个`go.mod`文件来声明它。

```bash
$ mkdir hello # Alternatively, clone it if it already exists in version control.
$ cd hello
$ go mod init example/user/hello
go: creating new go.mod: module example/user/hello
$ cat go.mod
module example/user/hello

go 1.16
$
```

Go源文件中的第一条语句必须是包名。可执行命令必须始终使用包名。

接下来，在该目录下创建一个名为`hello.go`的文件，包含以下Go代码。

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, world.")
}
```

现在你可以用go工具构建和安装该程序。

```bash
$ go install example/user/hello
$
```

这个命令建立了`hello`命令，产生了一个可执行的二进制文件。然后，它将该二进制文件安装为`$HOME/go/bin/hello`（或者，在Windows下，`%USERPROFILE%\gobinhello.exe`）。

安装目录是由`GOPATH`和`GOBIN`环境变量控制的。如果设置了`GOBIN`，二进制文件将被安装到该目录。如果设置了`GOPATH`，二进制文件将被安装到`GOPATH`列表中第一个目录的`bin`子目录中。否则，二进制文件将被安装到默认的`GOPATH`（`$HOME/go`或`%USERPROFILE%\go`）的`bin`子目录。

