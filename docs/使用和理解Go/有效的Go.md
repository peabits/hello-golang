[⬅️ 上一页: 没有了](#) 🚦 [下一页: 编辑器插件和IDE ➡️](编辑器插件和IDE.md)

[返回上一级: 使用和理解Go ⬆️](../使用和理解Go.md)

# 有效的Go

## 简介

Go是一种新的语言。虽然它借鉴了现有语言的思想，但它具有不寻常的特性，使得有效的Go程序与用其亲属编写的程序在性质上有所不同。将C++或Java程序直接翻译成Go不太可能产生令人满意的结果--Java程序是用Java编写的，不是Go。另一方面，从围棋的角度来考虑问题，可能会产生一个成功的但完全不同的程序。换句话说，要写好Go，了解它的属性和习性是很重要的。了解 Go 编程的既定惯例也很重要，例如命名、格式化、程序构造等，这样你写的程序就能让其他 Go 程序员容易理解。

本文档给出了编写清晰、习惯的 Go 代码的提示。它是对语言规范、Go之旅和如何编写Go代码的补充，您应该首先阅读这些内容。

2022年1月添加的注释：本文档是为2009年发布的Go编写的，此后没有进行过重大更新。虽然它是了解如何使用语言本身的好指南，但由于语言的稳定性，它几乎没有提到库，也没有提到Go生态系统自编写以来的重大变化，如构建系统、测试、模块和多态性。目前还没有计划对其进行更新，因为已经发生了太多的事情，而且有一大批不断增加的文档、博客和书籍对现代Go的使用进行了很好的描述。Effective Go仍然是有用的，但读者应该明白它远不是一个完整的指南。请看第28782期的内容。

### 例子

Go软件包的源码不仅是作为核心库，也是作为如何使用该语言的例子。此外，许多包都包含了可以工作的、独立的可执行的例子，你可以直接从golang.org网站上运行，比如这个例子（如果需要的话，点击 "例子 "一词来打开它）。如果你对如何处理一个问题或如何实现某些东西有疑问，库中的文档、代码和例子可以提供答案、想法和背景。

## 格式

格式问题是最有争议的，但也是影响最小的。人们可以适应不同的格式化风格，但如果他们不需要这样做，那就更好了。如果每个人都遵守相同的风格，那么在这个问题上投入的时间就会更少。问题是如何在没有冗长的规定性风格指南的情况下接近这个乌托邦。

在Go中，我们采取了一种不同寻常的方法，让机器来处理大多数的格式化问题。gofmt程序（也可以用go fmt，它在包级而不是源文件级操作）读取Go程序，并以缩进和垂直对齐的标准风格发出源代码，保留并在必要时重新编排注释。如果你想知道如何处理一些新的布局情况，请运行gofmt；如果答案似乎不对，请重新安排你的程序（或提交关于gofmt的错误），不要绕过它。

举例来说，没有必要花时间把结构的字段上的注释排成一排。Gofmt会帮你做到这一点。给出的声明是

```go
type T struct {
    name string // name of the object
    value int // its value
}
```

gofmt将排成一列。

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

标准包中的所有Go代码都已用gofmt格式化。

一些格式化的细节仍然存在。非常简单。

- 缩进
    
    我们使用制表符来缩进，gofmt默认会发出制表符。只有在必须使用时才使用空格。

- 行的长度
    
    Go没有行长限制。不要担心打好的卡片会溢出来。如果觉得一行太长，就把它包起来，用一个额外的制表符缩进。

- 圆括号

    Go需要的括号比C和Java少：控制结构（if、for、switch）的语法中没有括号。另外，运算符的优先级层次更短，更清晰，所以
    ```go
    x<<8 + y<<16
    ```
    意思是间距所暗示的，与其他语言不同。

## 注释

Go提供了C语言风格的/* */块注释和C++风格的//行注释。行注释是标准的；块注释主要作为包注释出现，但在表达式中或禁用大段代码时很有用。
出现在顶级声明前的注释，没有中间的新行，被认为是对声明本身的记录。这些 "文档注释 "是特定 Go 包或命令的主要文档。关于文档注释的更多信息，请参见 "Go 文档注释"。

## 名称

名称在Go中和其他语言中一样重要。它们甚至有语义上的影响：一个名字在包之外的可见性由其第一个字符是否为大写决定。因此，值得花一点时间来讨论Go程序中的命名规则。

### 包名称

当一个包被导入时，包的名称成为内容的访问器。在

```go
import name
```

后，导入的包就可以谈论bytes.Buffer了。如果每个使用包的人都能用同样的名字来指代它的内容，那是很有帮助的，这意味着包的名字应该是好的：短小、简明、令人回味。按照惯例，包被赋予小写的单字名称；不应该有下划线或混合大写。倾向于简洁的一面，因为每个使用你的包的人都会输入这个名字。而且不要担心先验的碰撞问题。包的名称只是导入时的默认名称；它不需要在所有的源代码中都是唯一的，而且在极少的碰撞情况下，导入的包可以选择一个不同的名称在本地使用。在任何情况下，混淆都是罕见的，因为导入的文件名正好决定了哪个包被使用。

另一个惯例是包的名字是其源目录的基名；src/encoding/base64中的包被导入为 "encoding/base64"，但名字是base64，而不是encoding_base64，也不是encodingBase64。

一个包的导入者会使用这个名字来指代它的内容，所以包中导出的名字可以使用这个事实来避免重复。(不要使用import .符号，它可以简化那些必须在它们所测试的包之外运行的测试，但在其他方面应该避免。) 例如，bufio包中的缓冲阅读器类型被称为Reader，而不是BufReader，因为用户看到它是bufio.Reader，这是一个清晰、简洁的名字。此外，因为导入的实体总是用它们的包名来称呼，所以bufio.Reader与io.Reader并不冲突。同样，制造ring.Ring新实例的函数--这是Go中构造函数的定义--通常会被称为NewRing，但由于Ring是包导出的唯一类型，并且包被称为ring，所以它被称为New，包的客户看到的是ring.New。使用包的结构来帮助你选择好的名字。

另一个简短的例子是once.Do；once.Do(setup)读起来很好，不会因为写once.DoOrWaitUntilDone(setup)而得到改善。长名字不会自动使事情变得更易读。一个有用的文档注释往往比一个超长的名字更有价值。

### 获取器

Go 并不提供自动支持获取器和设置器的功能。自己提供获取器和设置器并无不妥，而且这样做往往是合适的，但在获取器的名称中加入Get既不习惯也没有必要。如果你有一个叫做owner的字段（小写，未导出），getter方法应该叫做Owner（大写，已导出），而不是GetOwner。使用大写的名字进行导出，提供了区分字段和方法的钩子。如果需要的话，一个setter函数可能会被称为SetOwner。这两个名字在实践中都很好读。

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### 接口名称

按照惯例，单一方法的接口是由方法名称加上-er后缀或类似的修饰来构建代理名词的。Reader, Writer, Formatter, CloseNotifier等等。

有许多这样的名字，尊重它们和它们所捕获的函数名称是有成效的。读取、写入、关闭、冲洗、字符串等等都有规范的签名和含义。为了避免混淆，不要给你的方法一个这样的名字，除非它有相同的签名和含义。相反，如果你的类型实现了一个与某个著名类型的方法具有相同含义的方法，那么就给它相同的名称和签名；把你的字符串转换方法称为String而不是ToString。

### 混合帽

最后，Go中的惯例是使用MixedCaps或mixedCaps而不是下划线来写多字名。

## 分号

与C语言一样，Go的形式语法使用分号来结束语句，但与C语言不同的是，这些分号并不出现在源代码中。词典使用一个简单的规则，在扫描时自动插入分号，所以输入的文本基本上没有分号。

这个规则是这样的。如果换行前的最后一个符号是一个标识符（包括int和float64这样的词），一个基本的字面意义，如数字或字符串常数，或其中一个符号

```go
break continue fallthrough return ++ -- ) }
```

词法分析器总是在该标记后插入一个分号。这可以概括为："如果换行出现在一个可以结束语句的标记之后，就插入一个分号"。

分号也可以省略在结尾括号之前，所以一个语句，如

```go
go func() { for { dst <- <-src } }()
```

不需要分号。不言而喻，Go程序只有在for循环条款等地方才有分号，以分隔初始化器、条件和延续元素。如果你在写代码时，也需要用分号来分隔一行中的多个语句。

分号插入规则的一个后果是，你不能把控制结构（if、for、switch或select）的开头括号放在下一行。如果你这样做，分号将被插入到大括号之前，这可能导致不必要的影响。像这样写

```go
if i < f() {
    g()
}
```

不是这样的

```go
if i < f()  // wrong!
{           // wrong!
    g()
}
```

## 控制结构

Go的控制结构与C语言的控制结构有关，但在一些重要方面有所不同。没有do或while循环，只有略微通用的for；switch更加灵活；if和switch接受一个可选的初始化语句，就像for一样；break和continue语句接受一个可选的标签，以确定中断或继续的内容；还有一些新的控制结构，包括一个类型开关和一个多路通信复用器，select。语法也略有不同：没有小括号，主体必须始终以大括号为界。

### If

在Go中，一个简单的if看起来像这样。

```go
if x > 0 {
    return y
}
```

强制性大括号鼓励在多行上写简单的if语句。这样做是很好的风格，特别是当主体包含控制语句时，如返回或中断。

由于if和switch接受初始化语句，所以经常看到用它来设置一个局部变量。

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

在Go库中，你会发现当if语句不流入下一条语句时——即语句体以break、continue、goto或return结尾——不必要的else被省略了。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

这是一个常见情况的例子，代码必须防范一连串的错误情况。如果成功的控制流顺着页面运行，在错误情况出现时消除它们，那么代码就会读得很好。由于错误情况往往以返回语句结束，因此产生的代码不需要else语句。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f)
```

### 重新声明和重新赋值

一个旁证。上一节的最后一个例子演示了:=短声明形式如何工作的一个细节。调用os.Open的声明是这样的

```go
f, err := os.Open(name)
```

这个语句声明了两个变量，f和err。几行之后，对f.Stat的调用是：。

```go
d, err := f.Stat()
```

这看起来就像它声明了d和err。但是请注意，err出现在两个语句中。这种重复是合法的：err在第一条语句中声明，但在第二条语句中只是重新赋值。这意味着对f.Stat的调用使用了上面声明的现有err变量，只是给它一个新的值。

在:=声明中，即使变量v已经被声明过，也可以出现，前提是。

- 该声明与现有的v的声明在同一范围内（如果v已经在外部范围内声明，该声明将创建一个新变量）。
- 初始化中的相应值是可以分配给v的，并且
- 至少有一个其他变量是由该声明创建的。

这个不寻常的属性是纯粹的实用主义，使得在一个长的if-else链中很容易使用一个单一的err值，例如。你会看到它经常被使用。

这里值得注意的是，在Go中，函数参数和返回值的范围与函数主体相同，尽管它们在词汇上出现在包围主体的大括号之外。

### For

Go for 循环与C语言相似，但不相同。它统一了for和while，没有do-while。有三种形式，其中只有一种有分号。

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

简短的声明使其很容易在循环中直接声明索引变量。

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

如果你在一个`array`、`slice`、`string`或`map`上循环，或从一个`channel`上读取，`range`可以管理循环。

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

如果你只需要`range`的第一个项目（键或索引），就放弃第二个。

```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

如果你只需要`range`的第二个项目（值），使用空白标识符，即下划线，来放弃第一个项目。

```go
sum := 0
for _, value := range array {
    sum += value
}
```

空白标识符有很多用途，在后面的章节中会介绍。

对于字符串，`range`为你做更多的工作，通过解析UTF-8来分解出各个Unicode码位。错误的编码会消耗一个字节并产生替换符文U+FFFD。(名称(与相关的内建类型)符文是Go术语，指单一的Unicode码位。详见语言规范）。这个循环

```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

打印

```
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

最后，Go没有逗号运算符，而且++和--是语句而不是表达式。因此，如果你想在一个for中运行多个变量，你应该使用并行赋值（尽管这排除了++和--）。

```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

### Switch

Go的switch比C的更通用。表达式不需要是常数或甚至是整数，案例从上到下进行评估，直到找到一个匹配的案例，如果switch没有表达式，则切换为真。因此，将if-else-if-else链写成一个switch是可能的，而且是习惯性的。

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

没有自动的落差，但案例可以用逗号分隔的列表呈现。

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

虽然在Go中并不像其他一些类C语言那样常见，但break语句可以用来提前终止一个switch。但有时需要脱离周围的循环，而不是开关，在Go中可以通过在循环上加一个标签并 "断开 "该标签来实现。这个例子展示了这两种用法。

```go
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```

当然，continue语句也接受一个可选的标签，但它只适用于循环。

作为本节的结束，这里有一个使用两个switch语句的字节切片的比较例程。

```go
// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}
```

### Type switch

一个switch也可以用来发现一个接口变量的动态类型。这样的类型切换使用类型断言的语法，括号内有关键字type。如果switc在表达式中声明了一个变量，那么该变量将在每个子句中具有相应的类型。在这种情况下，重用名称也是一种习惯，实际上是在每种情况下声明一个具有相同名称但不同类型的新变量。

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
```

## 函数

### 多个返回值

Go的一个不寻常的特点是，函数和方法可以返回多个值。这种形式可以用来改进C语言程序中的一些笨拙的习惯：带内错误返回，如EOF的-1，以及修改按地址传递的参数。

在C语言中，写错误的信号是一个负数，错误代码隐藏在一个易失性位置。在Go中，Write可以返回一个计数和一个错误。"是的，你写了一些字节，但不是全部，因为你填满了设备"。包os中的文件的Write方法的签名是。

```go
func (file *File) Write(b []byte) (n int, err error)
```

正如文档中所说，当n != len(b)时，它返回写入的字节数和一个非零的错误。这是一种常见的风格，更多的例子见错误处理一节。

类似的方法避免了传递一个指向返回值的指针来模拟引用参数的需要。这里有一个头脑简单的函数，从一个字节片的某个位置抓取一个数字，返回数字和下一个位置。

```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```

你可以用它来扫描一个输入片断b中的数字，像这样。

```go
for i := 0; i < len(b); {
    x, i = nextInt(b, i)
    fmt.Println(x)
}
```

### 命名的结果参数

Go函数的返回或结果 "参数 "可以被命名，并作为常规变量使用，就像传入参数一样。当命名时，它们在函数开始时被初始化为其类型的零值；如果函数执行一个没有参数的返回语句，结果参数的当前值将被用作返回值。

这些名字不是强制性的，但它们可以使代码更短、更清晰：它们是文档。如果我们给nextInt的结果命名，那么哪个返回的int是哪个就很明显了。

```go
func nextInt(b []byte, pos int) (value, nextPos int)
```

因为命名的结果是初始化的，并与一个不加修饰的返回相联系，它们可以简化以及澄清。下面是io.ReadFull的一个版本，很好地使用了它们。

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

### 延迟

## 数据

### 使用new进行分配

Go有两个分配原语，即内置函数new和make。它们做不同的事情，适用于不同的类型，这可能令人困惑，但规则很简单。我们先来谈谈new。这是一个分配内存的内置函数，但与其他一些语言中的同名函数不同，它并不初始化内存，只是将其清零。也就是说，new(T)为一个类型为T的新项目分配了零的存储空间，并返回其地址，即一个类型为*T的值。用Go的术语来说，它返回一个指向新分配的T类型的零值的指针。

由于new返回的内存是清零的，所以在设计数据结构时，安排每种类型的零值无需进一步初始化就可以使用，是很有帮助的。这意味着数据结构的用户可以用new创建一个数据结构并直接开始工作。例如，bytes.Buffer的文档指出，"Buffer的零值是一个准备使用的空缓冲区。" 同样，sync.Mutex也没有一个明确的构造函数或Init方法。相反，sync.Mutex的零值被定义为一个未锁定的mutex。

”零值有用“属性具有传递性。考虑这个类型声明。

```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

SyncedBuffer类型的值也可以在分配或刚刚声明时立即使用。在下一个片段中，p和v都会正确工作，不需要进一步安排。

```go
p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

### 构造函数和复合字面量

有时零值还不够好，需要一个初始化构造函数，就像这个源自包os的例子。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

这里面有很多样板。我们可以使用复合字面量来简化它，复合字面量是一个表达式，每次评估都会创建一个新的实例。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

注意，与C语言不同的是，返回局部变量的地址是完全可以的；与该变量相关的存储在函数返回后仍然存在。事实上，获取一个复合字面量的地址在每次评估时都会分配一个新的实例，所以我们可以把最后两行结合起来。

```go
    return &File{fd, name, nil, 0}
```

一个复合字面量的字段是按顺序排列的，而且必须全部出现。然而，通过将元素明确标记为字段：值对，初始化器可以以任何顺序出现，缺失的字段将作为各自的零值。因此，我们可以说

```go
    return &File{fd: fd, name: name}
```

作为一种限制性的情况，如果一个复合字面完全不包含字段，那么它将为该类型创建一个零值。new(File)和&File{}这两个表达式是等价的。


也可以为数组、切片和映射创建复合字元，字段标签可以是索引或映射的键值。在这些例子中，无论Enone、Eio和Einval的值如何，只要它们是不同的，初始化就会起作用。

```go
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```

### 使用make进行分配

回到分配上。内置函数make(T, args)的目的与new(T)不同。它只创建切片、映射和通道，并返回一个初始化（非归零）的T（而不是*T）类型的值。区别的原因是，这三种类型代表的是对数据结构的引用，在使用前必须进行初始化。例如，一个切片是一个包含指向数据的指针（在一个数组内）、长度和容量的三个项目的描述符，在这些项目被初始化之前，这个切片是空的。对于切片、映射和通道，make初始化内部数据结构，并准备好使用的值。比如说

```go
make([]int, 10, 100)
```

分配一个100个整数的数组，然后创建一个长度为10，容量为100的slice结构，指向数组的前10个元素。(当创建一个slice时，容量可以省略；更多信息请看关于slice的部分)。相反，new([]int)返回一个指向新分配的、归零的分片结构的指针，也就是说，一个指向零分片值的指针。

这些例子说明了`new`和`make`之间的区别。

```go
var p *[]int = new([]int)       // allocates slice structure; *p == nil; rarely useful
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Idiomatic:
v := make([]int, 100)
```

记住，make只适用于map、slices和channel，不返回指针。要获得一个显式指针，需要用new分配，或者显式地获取一个变量的地址。

### 数组

### 切片

### 二维切片

### 映射

### 打印

### 附加

## 初始化

### 常量

### 变量

### 初始化函数

## 方法

### 指针值

正如我们在ByteSize中看到的那样，可以为任何命名的类型（除了指针或接口）定义方法；接收器不一定是一个结构。

在上面关于分片的讨论中，我们写了一个Append函数。我们可以把它定义为片上的方法。要做到这一点，我们首先声明一个命名的类型，我们可以将该方法与之绑定，然后使该方法的接收器成为该类型的值。

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above.
}
```

这仍然需要该方法返回更新的切片。我们可以通过重新定义该方法以获取一个指向ByteSlice的指针作为其接收器来消除这种笨拙，因此该方法可以覆盖调用者的片断。

```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

事实上，我们可以做得更好。如果我们修改我们的函数，使它看起来像一个标准的Write方法，像这样。

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```

那么*ByteSlice类型就满足了标准的io.Writer接口，这很方便。例如，我们可以打印成一个。

```go
    var b ByteSlice
    fmt.Fprintf(&b, "This hour has %d days\n", 7)
```

我们传递一个ByteSlice的地址，因为只有*ByteSlice满足io.Writer的要求。关于接收器的指针与值的规则是，值方法可以在指针和值上调用，但指针方法只能在指针上调用。

这条规则的产生是因为指针方法可以修改接收者；在值上调用它们会导致方法接收到值的副本，所以任何修改都会被丢弃。因此，该语言不允许这种错误。不过，有一个方便的例外。当值是可寻址的，语言通过自动插入地址操作符来处理在值上调用指针方法的常见情况。在我们的例子中，变量b是可寻址的，所以我们可以只用b.Write来调用它的Write方法。编译器将为我们把它改写成(&b).Write。

顺便说一下，在一个字节片上使用Write的想法是实现bytes.Buffer的核心。


## 接口和其他类型

### 接口

### 转换

### 接口转换和类型断言

### 通用性

### 接口和方法

## 空白标识符

### 多重分配中的空白标识符

### 未使用的导入和变量

### 进口的副作用

### 接口检查

## 嵌入

## 并发

### 通过通信共享

并发编程是一个很大的话题，这里只有一些针对Go的亮点。

许多环境中的并发编程由于需要实现对共享变量的正确访问而变得困难。Go鼓励采用一种不同的方法，即在通道上传递共享值，事实上，不同的执行线程从不主动共享。在任何时候，只有一个goroutine可以访问该值。通过设计，数据竞赛不会发生。为了鼓励这种思维方式，我们将其简化为一个口号。

    不要通过共享内存来通信；相反，通过通信来共享内存。

这种方法可能会走得太远。例如，通过在整数变量周围放置一个mutex来进行引用计数可能是最好的。但作为一种高层次的方法，使用通道来控制访问使得编写清晰、正确的程序更加容易。

思考这个模型的一种方式是考虑一个典型的单线程程序在一个CPU上运行。它不需要同步原语。现在运行另一个这样的实例；它也不需要同步。现在让这两个人进行通信；如果通信是同步器，那么仍然不需要其他同步。例如，Unix流水线就非常符合这种模式。虽然Go的并发方法源于Hoare的通信顺序进程（CSP），但它也可以被看作是Unix管道的类型安全的概括。

### 协程（Goroutines）

它们被称为goroutines，因为现有的术语（线程、协程、进程等等）都包含着不准确的含义。一个goroutine有一个简单的模型：它是一个在同一地址空间内与其他goroutine同时执行的函数。它是轻量级的，除了分配堆栈空间外，几乎不需要花费什么。而堆栈开始时很小，所以它们很便宜，并根据需要通过分配（和释放）堆存储来增长。

Goroutines被复用到多个操作系统线程上，因此如果一个线程发生阻塞，例如在等待I/O时，其他线程继续运行。它们的设计隐藏了许多线程创建和管理的复杂性。

在一个函数或方法的调用前加上go关键字，在一个新的goroutine中运行该调用。当调用完成后，goroutine会悄悄地退出。(其效果类似于Unix shell的&符号，用于在后台运行一个命令）。

```go
go list.Sort()  // run list.Sort concurrently; don't wait for it.
```

一个函数的字面量在一个goroutine的调用中可以很方便。

```go
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

在Go中，函数的字面量是封闭的：实现确保函数所引用的变量只要是有效的，就会一直存在。

这些例子不是很实用，因为这些函数没有办法发出完成的信号。为此，我们需要通道。

### 通道（Channels）

像映射一样，通道是用make分配的，产生的值作为对一个底层数据结构的引用。如果提供了一个可选的整数参数，它将设置通道的缓冲区大小。对于无缓冲区或同步通道来说，默认值是0。

```go
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

无缓冲通道将通信（数值交换）与同步相结合，保证两个计算（goroutine）处于已知状态。

有很多使用通道的习语。这里有一个可以让我们开始。在上一节中，我们在后台启动了一个排序。一个通道可以让启动的goroutine等待排序完成。

```go
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

接收者总是阻塞，直到有数据要接收。如果信道没有缓冲区，发送方就会阻塞，直到接收方收到该值。如果信道有缓冲区，发送方只阻断到值被复制到缓冲区；如果缓冲区满了，这意味着要等到某个接收方检索到一个值。

缓冲通道可以像信号灯一样使用，比如说用来限制吞吐量。在这个例子中，传入的请求被传递给handle，handle向通道发送一个值，处理请求，然后从通道接收一个值，为下一个消费者准备 "信号"。通道缓冲区的容量限制了同时处理的调用的数量。

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```

一旦MaxOutstanding处理程序在执行过程中，任何更多的处理程序都会阻断向充满的通道缓冲区发送，直到现有的一个处理程序完成并从缓冲区接收。

不过，这种设计有一个问题。Serve为每个传入的请求创建一个新的goroutine，即使在任何时候只有MaxOutstanding的请求可以运行。因此，如果请求来的太快，程序就会消耗无限的资源。我们可以通过改变Serve对goroutine的创建进行把关来解决这一缺陷。这里有一个明显的解决方案，但要注意它有一个bug，我们随后会修复。

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // Buggy; see explanation below.
            <-sem
        }()
    }
}
```

错误在于，在Go的for循环中，循环变量在每次迭代中都被重复使用，所以req变量在所有goroutines中都是共享的。这不是我们想要的。我们需要确保req在每个goroutine中是唯一的。这里有一个方法，将req的值作为参数传递给goroutine中的closure。

func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
}

将这个版本与之前的版本进行比较，可以看到闭包的声明和运行方式的不同。另一个解决方案是直接创建一个同名的新变量，如本例：

```go
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // Create new instance of req for the goroutine.
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

写起来可能看起来很奇怪

```go
req := req
```

但在Go中这样做是合法的，而且是习惯性的。你会得到一个同名的新版本的变量，故意在本地影射循环变量，但对每个goroutine是唯一的。

回到编写服务器的一般问题上，另一种能很好地管理资源的方法是启动固定数量的处理goroutines，都从请求通道中读取。goroutines的数量限制了同时调用处理的数量。这个Serve函数也接受一个通道，它将被告知退出；在启动goroutines后，它阻止从该通道接收。

```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // Start handlers
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // Wait to be told to exit.
}
```

### 通道的通道

### 并行化

### 泄漏的缓冲区

## 错误

### 恐慌

### 恢复

## 一个网络服务器


[返回顶部 🔝](#有效的Go)