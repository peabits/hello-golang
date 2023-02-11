
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

你可以使用`go env`命令，为未来的go命令可移植地设置环境变量的默认值。

```bash
$ go env -w GOBIN=/somewhere/else/bin
$
```

要取消之前由`go env -w`设置的变量，请使用`go env -u`。

```bash
$ go env -u GOBIN
$
```

像`go install`这样的命令在包含当前工作目录的模块的上下文中应用。如果工作目录不在`example/user/hello`模块内，`go install`可能会失败。

为了方便起见，go命令接受相对于工作目录的路径，如果没有给出其他路径，则默认为当前工作目录下的软件包。因此，在我们的工作目录下，以下命令都是等价的。

```bash
$ go install example/user/hello

$ go install .

$ go install
```

接下来，让我们运行该程序，以确保它的工作。为了方便起见，我们将把安装目录添加到我们的`PATH`中，使运行二进制文件变得容易。

```bash
# Windows users should consult https://github.com/golang/go/wiki/SettingGOPATH
# for setting %PATH%.
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
$ hello
Hello, world.
$
```

如果你使用的是源码控制系统，现在是初始化存储库、添加文件并提交第一个修改的好时机。同样，这一步是可选的：你不需要使用源码控制来编写Go代码。

```bash
$ git init
Initialized empty Git repository in /home/user/hello/.git/
$ git add go.mod hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 go.mod hello.go
$
```

go命令通过请求相应的`HTTPS URL`并读取嵌入`HTML`响应中的元数据来定位包含给定模块路径的资源库（见`go help importpath`）。许多托管服务已经为包含Go代码的仓库提供了元数据，所以让你的模块供他人使用的最简单方法通常是使其模块路径与仓库的`URL`一致。

### 从你的模块中导入包

让我们写一个`morestrings`包并从`hello`程序中使用它。首先，为该包创建一个名为`$HOME/hello/morestrings`的目录，然后在该目录下创建一个名为`reverse.go`的文件，内容如下。

```go
// Package morestrings implements additional functions to manipulate UTF-8
// encoded strings, beyond what is provided in the standard "strings" package.
package morestrings

// ReverseRunes returns its argument string reversed rune-wise left to right.
func ReverseRunes(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

因为我们的`ReverseRunes`函数以大写字母开头，所以它被导出，可以在其他导入我们的`morestrings`包的软件包中使用。

让我们用`go build`来测试一下该软件包的编译情况。

```bash
$ cd $HOME/hello/morestrings
$ go build
$
```

这不会产生一个输出文件。相反，它将编译后的包保存在本地的构建缓存中。

在确认了`morestrings`包的构建之后，让我们在`hello`程序中使用它。要做到这一点，请修改你原来的`$HOME/hello/hello.go`以使用`morestrings`包。

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
}
```

安装`hello`程序。

```bash
$ go install example/user/hello
```

运行新版本的程序，你应该看到一个新的、反转的信息。

```bash
$ hello
Hello, Go!
```

### 从远程模块导入软件包

一个导入路径可以描述如何使用Git或Mercurial这样的修订控制系统来获取软件包的源代码。go工具使用这个属性来自动从远程仓库获取软件包。例如，要在你的程序中使用`github.com/google/go-cmp/cmp`。

```go
package main

import (
    "fmt"

    "example/user/hello/morestrings"
    "github.com/google/go-cmp/cmp"
)

func main() {
    fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
```

现在你有了对外部模块的依赖，你需要下载该模块并在你的`go.mod`文件中记录其版本。`go mod tidy`命令为导入的软件包添加缺少的模块需求，并删除不再使用的模块的需求。

```bash
$ go mod tidy
go: finding module for package github.com/google/go-cmp/cmp
go: found github.com/google/go-cmp/cmp in github.com/google/go-cmp v0.5.4
$ go install example/user/hello
$ hello
Hello, Go!
  string(
-     "Hello World",
+     "Hello Go",
  )
$ cat go.mod
module example/user/hello

go 1.16

require github.com/google/go-cmp v0.5.4
$
```

模块的依赖性被自动下载到`GOPATH`环境变量所指示的目录下的`pkg/mod`子目录中。一个给定版本的模块的下载内容在所有需要该版本的模块之间共享，所以go命令将这些文件和目录标记为只读。要删除所有下载的模块，你可以在`go clean`中传递`-modcache`标志。

```bash
$ go clean -modcache
$
```

## 测试

Go有一个由`go test`命令和测试包组成的轻量级测试框架。

你可以通过创建一个名称以`_test.go`结尾的文件来编写测试，该文件包含名为`TestXXX`的函数，其签名为`func（t *testing.T）`。测试框架运行每个这样的函数；如果该函数调用一个失败函数，如`t.Error`或`t.Fail`，则认为测试失败。
通过创建包含以下Go代码的`$HOME/hello/morestrings/reverse_test.go`文件，向`morestrings`包添加一个测试。

```go
package morestrings

import "testing"

func TestReverseRunes(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := ReverseRunes(c.in)
        if got != c.want {
            t.Errorf("ReverseRunes(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
```

然后用`go test`运行该测试。

```bash
$ cd $HOME/hello/morestrings
$ go test
PASS
ok  	example/user/hello/morestrings 0.165s
$
```

运行`go help test`，更多细节见测试包文档。

## 下一步工作

订阅 `golang-announce` 邮件列表，以便在 Go 的新稳定版本发布时获得通知。

请参阅Effective Go，了解有关编写清晰、简洁的Go代码的技巧。

参加 Go 之旅来学习这门语言。

访问文档页面，了解有关 Go 语言及其库和工具的一系列深度文章。

## 获得帮助

要获得实时帮助，请向社区管理的 gophers Slack 服务器中有用的 gophers 询问（在此获取邀请）。

用于讨论 Go 语言的官方邮件列表是 Go Nuts。

使用 Go 问题跟踪器报告错误。