# 如何编写Go代码（使用GOPATH）

## 简介

如果你是Go的新手，请参考最近的[如何编写Go代码](main.md)。

本文演示了一个简单的 Go 包的开发，并介绍了 go 工具，这是获取、构建和安装 Go 包和命令的标准方法。

go工具要求你以特定的方式组织你的代码。请仔细阅读本文档。它解释了最简单的方法来启动和运行你的Go安装。

类似的解释可以以截屏的形式提供。

## 代码组织

### 概述

- Go程序员通常将所有的Go代码保存在一个工作区。
- 一个工作区包含许多版本控制库（例如由Git管理）。
- 每个仓库包含一个或多个包。
- 每个包由一个或多个Go源文件组成，放在一个目录中。
- 一个包的目录路径决定了它的导入路径。

请注意，这与其他编程环境不同，在这些环境中，每个项目都有一个单独的工作空间，而且工作空间与版本控制库紧密相连。

### 工作空间

一个工作区是一个目录层次结构，其根部有两个目录。

- src 包含 Go 源文件，
- bin 包含可执行命令。

go工具会构建并安装二进制文件到bin目录中。

src子目录通常包含多个版本控制库（如Git或Mercurial），跟踪一个或多个源代码包的开发。

为了让你了解工作区的实际情况，这里有一个例子。

```
bin/
    hello                          # command executable
    outyet                         # command executable
src/
    golang.org/x/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```

上面的树显示了一个包含两个资源库（example和image）的工作空间。示例库包含两个命令（hello和outyet）和一个库（stringutil）。图像资源库包含bmp包和其他几个包。

一个典型的工作区包含许多源码库，其中有许多包和命令。大多数Go程序员将他们所有的Go源代码和依赖关系放在一个工作区中。

请注意，不应该使用符号链接将文件或目录链接到你的工作区。

命令和库是由不同种类的源码包构建的。我们将在后面讨论这种区别。

### GOPATH环境变量

GOPATH环境变量指定了你的工作空间的位置。它默认为你的主目录内的一个名为go的目录，因此在Unix中为\$HOME/go，在Plan 9中为\$home/go，在Windows中为%USERPROFILE%\go（通常为C:\Users\YourName\go）。

如果你想在一个不同的位置工作，你需要将GOPATH设置为该目录的路径。(另一个常见的设置是设置GOPATH=$HOME。)注意，GOPATH不能与您的 Go 安装路径相同。

命令 go env GOPATH 打印当前有效的 GOPATH；如果环境变量未设置，则打印默认位置。

为方便起见，请将工作区的bin子目录添加到您的PATH中。

```bash
$ export PATH=$PATH:$(go env GOPATH)/bin
```

为了简洁起见，本文件其余部分的脚本使用\$GOPATH而不是\$(go env GOPATH)。如果你没有设置GOPATH，要想让脚本按照写好的内容运行，你可以在这些命令中用$HOME/go代替，否则运行。

```bash
$ export GOPATH=$(go env GOPATH)
```

要了解更多关于GOPATH环境变量的信息，请参见 "go help gopath"。

### 导入路径

导入路径是一个字符串，用于唯一地识别一个包。一个包的导入路径与它在工作区或远程资源库中的位置相对应（解释如下）。

来自标准库的包被赋予简短的导入路径，如 "fmt "和 "net/http"。对于你自己的包，你必须选择一个不太可能与未来添加到标准库或其他外部库相冲突的基本路径。

如果你把你的代码保存在某个源码库中，那么你应该使用该源码库的根作为你的基本路径。例如，如果你有一个GitHub账户，地址是github.com/user，这应该是你的基本路径。

注意，你不需要在构建代码之前将其发布到远程仓库。这只是一个好习惯，把你的代码组织起来，就像你有一天会发布它一样。在实践中，你可以选择任何任意的路径名称，只要它对标准库和更大的Go生态系统来说是唯一的。

我们将使用github.com/user作为我们的基本路径。在你的工作空间里创建一个目录，用来保存源代码。

```bash
$ mkdir -p $GOPATH/src/github.com/user
```

### 你的第一个程序

要编译和运行一个简单的程序，首先选择一个包的路径（我们将使用github.com/user/hello），并在你的工作空间内创建一个相应的包目录。

```bash
$ mkdir $GOPATH/src/github.com/user/hello
```

接下来，在该目录下创建一个名为hello.go的文件，包含以下Go代码。

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world.")
}
```

现在你可以用go工具构建和安装该程序。

```bash
$ go install github.com/user/hello
```

注意，你可以在你系统的任何地方运行这个命令。go工具通过在GOPATH指定的工作空间内寻找github.com/user/hello包来找到源代码。

如果你在软件包目录下运行go install，你也可以省略软件包路径。

```bash
$ cd $GOPATH/src/github.com/user/hello
$ go install
```

该命令构建了hello命令，产生了一个可执行的二进制文件。然后，它将该二进制文件安装到工作区的bin目录下，作为hello（或者，在Windows下，hello.exe）。在我们的例子中，这将是\$GOPATH/bin/hello，也就是\$HOME/go/bin/hello。

go工具只有在发生错误时才会打印输出，所以如果这些命令没有产生输出，它们就已经成功执行了。

现在你可以在命令行中输入程序的完整路径来运行该程序。

```bash
$ $GOPATH/bin/hello
Hello, world.
```

或者，由于你已经将$GOPATH/bin添加到你的PATH中，只需输入二进制名称即可。

```bash
$ hello
Hello, world.
```

如果你使用的是源码控制系统，现在是初始化存储库、添加文件并提交第一个修改的好时机。同样，这一步是可选的：你不需要使用源码控制来编写Go代码。

```bash
$ cd $GOPATH/src/github.com/user/hello
$ git init
Initialized empty Git repository in /home/user/go/src/github.com/user/hello/.git/
$ git add hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 7 insertion(+)
 create mode 100644 hello.go
```

将代码推送到远程仓库是留给读者的一个练习。

### 你的第一个库

让我们写一个库并在hello程序中使用它。

同样，第一步是选择一个包路径（我们将使用github.com/user/stringutil）并创建包目录。

```bash
$ mkdir $GOPATH/src/github.com/user/stringutil
```

接下来，在该目录下创建一个名为reverse.go的文件，内容如下。

```go
// Package stringutil contains utility functions for working with strings.
package stringutil

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
```

现在，用go build来测试软件包的编译情况。

```bash
$ go build github.com/user/stringutil
```

或者，如果你在软件包的源目录下工作，只需。

```bash
$ go build
```

这不会产生一个输出文件。相反，它将编译后的包保存在本地构建缓存中。

在确认stringutil包构建完成后，修改你原来的hello.go（它在$GOPATH/src/github.com/user/hello中）以使用它。

```go
package main

import (
	"fmt"

	"github.com/user/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
}
```

安装Hello程序。

```bash
$ go install github.com/user/hello
```

运行新版本的程序，你应该看到一个新的、反转的信息。

```bash
$ hello
Hello, Go!
```

完成上述步骤后，你的工作区应该是这样的。完成上述步骤后，你的工作区应该是这样的。

```
bin/
    hello                 # command executable
src/
    github.com/user/
        hello/
            hello.go      # command source
        stringutil/
            reverse.go    # package source
```

### 软件包名称

Go源代码文件中的第一条语句必须是

```go
package name
```

其中 name 是软件包的默认导入名称。(一个包中的所有文件必须使用相同的名称）。

Go的惯例是，包的名称是导入路径的最后一个元素：以 "crypto/rot13 "导入的包应被命名为rot13。

可执行命令必须始终使用包名main。

没有要求包名在链接到一个二进制文件的所有包中是唯一的，只要求导入路径（它们的完整文件名）是唯一的。

参见Effective Go以了解更多关于Go的命名规则。

## 测试

Go有一个由go test命令和测试包组成的轻量级测试框架。

你通过创建一个名称以_test.go结尾的文件来编写测试，该文件包含名为TestXXX的函数，其签名为func（t *testing.T）。测试框架运行每个这样的函数；如果该函数调用一个失败函数，如t.Error或t.Fail，则认为测试失败。

通过创建包含以下Go代码的$GOPATH/src/github.com/user/stringutil/reverse_test.go文件，向stringutil包添加一个测试。

```go
package stringutil

import "testing"

func TestReverse(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, c := range cases {
		got := Reverse(c.in)
		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
```

然后用go test运行该测试。

```bash
$ go test github.com/user/stringutil
ok  	github.com/user/stringutil 0.165s
```

像往常一样，如果你是在软件包目录下运行go工具，你可以省略软件包路径。

```bash
$ go test
ok  	github.com/user/stringutil 0.165s
```

运行go help test，更多细节见测试包文档。

## 远程软件包

一个导入路径可以描述如何使用Git或Mercurial这样的修订控制系统来获取软件包的源代码。go工具使用这个属性来自动从远程仓库获取软件包。例如，本文档中描述的例子也保存在GitHub golang.org/x/example的Git仓库中。如果你在包的导入路径中包含了仓库的URL，go get将自动获取、构建和安装它。

```bash
$ go get golang.org/x/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!
```

如果指定的软件包不存在于一个工作区，go get将把它放在GOPATH指定的第一个工作区中。(如果该软件包已经存在，go get将跳过远程获取，其行为与go install相同。)

在发出上述go get命令后，工作区的目录树现在应该是这样的。

```
bin/
    hello                           # command executable
src/
    golang.org/x/example/
	.git/                       # Git repository metadata
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
    github.com/user/
        hello/
            hello.go                # command source
        stringutil/
            reverse.go              # package source
            reverse_test.go         # test source
```

GitHub上的hello命令依赖于同一仓库中的stringutil包。hello.go文件中的导入使用了相同的导入路径约定，所以go get命令也能定位并安装依赖的包。

```go
import "golang.org/x/example/stringutil"
```

这个惯例是让你的 Go 包供他人使用的最简单方法。Pkg.go.dev 和 Go Wiki 提供了外部 Go 项目的列表。

关于使用go工具的远程仓库的更多信息，请参见go help importpath。

## 下一步是什么

订阅 golang-announce 邮件列表，以便在 Go 的新稳定版本发布时获得通知。

请参阅Effective Go，了解有关编写清晰、简洁的Go代码的技巧。

参加 Go 之旅以了解该语言本身。

访问文档页面，了解有关 Go 语言及其库和工具的一系列深度文章。

## 获得帮助

要想获得实时帮助，请向Libera.Chat IRC服务器上的#go-nuts提问。

用于讨论 Go 语言的官方邮件列表是 Go Nuts。

使用Go问题跟踪器报告错误。
