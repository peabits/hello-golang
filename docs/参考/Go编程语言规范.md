[⬅️ 上一页: 命令文档](命令文档.md) 🚦 [下一页: Go模块参考 ➡️](Go模块参考.md)

[返回上一级: 参考 ⬆️](../参考.md)

# Go编程语言规范

## 简介

这是Go编程语言的参考手册。没有泛型的Go1.18之前的版本可以在[这里](https://go.dev/doc/go1.17_spec.html "https://go.dev/doc/go1.17_spec.html")找到。更多信息和其他文件，请参见[golang.org](golang.org "golang.org")。

Go是一种通用语言，设计时考虑到了系统编程。它是强类型和垃圾收集的，并且明确支持并发编程。程序是由包构成的，包的属性允许有效管理依赖关系。

语法紧凑且易于解析，便于自动工具的分析，如集成开发环境。

## 符号

语法是使用Extended Backus-Naur Form（EBNF）的变体来指定的。

```
Syntax      = { Production } .
Production  = production_name "=" [ Expression ] "." .
Expression  = Term { "|" Term } .
Term        = Factor { Factor } .
Factor      = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions是由术语和以下运算符构成的表达式，其优先级越来越高。

```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

小写的production名称用于识别词性（终端）标记。非终端则使用CamelCase。词汇标记用双引号""或反引号``括起来。

`a...b`的形式表示从a到b的一组字符作为备选。在规范的其他地方也使用水平省略号...来非正式地表示各种枚举或没有进一步指定的代码片断。字符...（相对于三个字符...而言）不是Go语言的标记。

## 源代码表示法

源代码是以UTF-8编码的Unicode文本。该文本没有被规范化，因此，一个重音代码点与由重音和字母组合而成的相同字符不同；它们被视为两个代码点。为了简单起见，本文将使用不合格的术语字符来指代源文本中的Unicode码位。

每个代码点都是不同的；例如，大写字母和小写字母是不同的字符。

实施限制。为了与其他工具兼容，编译器可能不允许源文本中出现NUL字符（U+0000）。

实施限制。为了与其他工具兼容，编译器可以忽略UTF-8编码的字节顺序标记（U+FEFF），如果它是源文本中的第一个Unicode代码点。在源文本中的其他地方可以不允许有字节序标记。

### 字符

以下术语用于表示特定的Unicode字符类别。

```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
```

在[Unicode标准8.0](https://www.unicode.org/versions/Unicode8.0.0/ "https://www.unicode.org/versions/Unicode8.0.0/")中，第4.5节 "一般类别 "定义了一组字符类别。Go将字母类别Lu、Ll、Lt、Lm或Lo中的所有字符视为Unicode字母，而将数字类别Nd中的字符视为Unicode数字。

### 字母和数字

下划线字符_（U+005F）被认为是一个小写字母。

```
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```

## 词汇元素

### 注释

注释作为程序文档。有两种形式：

1. 行注释以字符序列`//`开始，在行的末尾停止。
2. 一般注释以字符序列`/*`开始，并以随后的第一个字符序列`*/`停止。

注释不能从[符文](#)或[字符串](#)字面量开始，也不能从注释开始。一个不包含换行符的一般注释就像一个空格。任何其他的注释就像一个换行符。

### 标记

标记构成了Go语言的词汇表。有四类：标识符、关键词、运算符和标点符号以及字词。由空格（U+0020）、水平制表符（U+0009）、回车符（U+000D）和换行符（U+000A）组成的白色空间被忽略，除非它将原本合并为一个标记的标记分开。此外，换行或文件结束可能会触发插入[分号](#分号)。在将输入分成标记时，下一个标记是构成有效标记的最长的字符序列。

### 分号

正式的语法在一些产品中使用分号“;”作为终止符。Go程序可以通过以下两条规则省略大部分的分号：

1. 当输入被分成标记时，分号会立即在某行的最后一个标记后自动插入标记流，如果这个标记是：
    - 一个[标识符](#标识符)
    - 一个[整数](#整数字面量)、[浮点数](#浮点字面量)、[虚数](#虚数字面量)、符文或字符串字面量
    - 关键字之一：break、continue、fallthrough或return
    - 操作符和标点符号之一：++, --, ), ], 或 }。
2. 为了让复杂的语句占用一行，在结尾的"）"或"}"之前可以省略分号。

为了反映习惯性使用，本文件中的代码示例使用这些规则省略分号。

### 标识符

标识符命名程序实体，如变量和类型。一个标识符是一个或多个字母和数字的序列。标识符中的第一个字符必须是一个字母。

```
identifier = letter { letter | unicode_digit } .
letter     = unicode_letter | "_" .
```

```
a
_x9
ThisVariableIsExported
αβ
```

一些标识符是[预先声明](#)的。

### 关键词

以下关键词是保留的，不得作为标识符使用。

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

### 操作符和标点符号

以下字符序列代表[运算符](#)（包括[赋值运算符](#)）和标点符号。

```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

### 整数字面量

整数字面量是一串数字，代表一个[整数常量](#)。一个可选的前缀设置一个非十进制的基数。二进制为0b或0B，八进制为0、0o或0O，十六进制为0x或0X。单一的0被认为是十进制的0。在十六进制字段中，字母a到f和A到F代表数值10到15。

为了便于阅读，下划线字符_可以出现在基数前缀之后或连续的数字之间；这种下划线不会改变字面的值。

```
int_lit        = decimal_lit | binary_lit | octal_lit | hex_lit .
decimal_lit    = "0" | ( "1" … "9" ) [ [ "_" ] decimal_digits ] .
binary_lit     = "0" ( "b" | "B" ) [ "_" ] binary_digits .
octal_lit      = "0" [ "o" | "O" ] [ "_" ] octal_digits .
hex_lit        = "0" ( "x" | "X" ) [ "_" ] hex_digits .

decimal_digits = decimal_digit { [ "_" ] decimal_digit } .
binary_digits  = binary_digit { [ "_" ] binary_digit } .
octal_digits   = octal_digit { [ "_" ] octal_digit } .
hex_digits     = hex_digit { [ "_" ] hex_digit } .
```

```
42
4_2
0600
0_600
0o600
0O600       // second character is capital letter 'O'
0xBadFace
0xBad_Face
0x_67_7a_2f_cc_40_c6
170141183460469231731687303715884105727
170_141183_460469_231731_687303_715884_105727

_42         // an identifier, not an integer literal
42_         // invalid: _ must separate successive digits
4__2        // invalid: only one _ at a time
0_xBadFace  // invalid: _ must separate successive digits
```

### 浮点数字面量

浮点数字面量是浮点常量的十进制或十六进制表示。

十进制浮点数字面量由整数部分（十进制数字）、小数点、小数部分（十进制数字）和指数部分（e或E后面有可选的符号和十进制数字）组成。整数部分或分数部分中的一个可以省略；小数点或指数部分中的一个可以省略。一个指数值将尾数（整数和小数部分）按10exp进行缩放。

十六进制浮点数字面量由0x或0X前缀、整数部分（十六进制数字）、小数点、小数部分（十六进制数字）和指数部分（p或P，后面是可选的符号和十进制数字）组成。整数部分或小数部分中的一个可以省略；弧度点也可以省略，但指数部分是必须的。(这个语法与IEEE 754-2008 §5.12.3中给出的语法一致。)指数值exp将尾数(整数和小数部分)按2exp缩放。

为了便于阅读，下划线字符_可以出现在基数前缀之后或连续的数字之间；这种下划线不会改变字面值。

```
float_lit         = decimal_float_lit | hex_float_lit .

decimal_float_lit = decimal_digits "." [ decimal_digits ] [ decimal_exponent ] |
                    decimal_digits decimal_exponent |
                    "." decimal_digits [ decimal_exponent ] .
decimal_exponent  = ( "e" | "E" ) [ "+" | "-" ] decimal_digits .

hex_float_lit     = "0" ( "x" | "X" ) hex_mantissa hex_exponent .
hex_mantissa      = [ "_" ] hex_digits "." [ hex_digits ] |
                    [ "_" ] hex_digits |
                    "." hex_digits .
hex_exponent      = ( "p" | "P" ) [ "+" | "-" ] decimal_digits .
```

```
0.
72.40
072.40       // == 72.40
2.71828
1.e+0
6.67428e-11
1E6
.25
.12345E+5
1_5.         // == 15.0
0.15e+0_2    // == 15.0

0x1p-2       // == 0.25
0x2.p10      // == 2048.0
0x1.Fp+0     // == 1.9375
0X.8p-0      // == 0.5
0X_1FFFP-16  // == 0.1249847412109375
0x15e-2      // == 0x15e - 2 (integer subtraction)

0x.p1        // invalid: mantissa has no digits
1p-2         // invalid: p exponent requires hexadecimal mantissa
0x1.5e-2     // invalid: hexadecimal mantissa requires p exponent
1_.5         // invalid: _ must separate successive digits
1._5         // invalid: _ must separate successive digits
1.5_e1       // invalid: _ must separate successive digits
1.5e_1       // invalid: _ must separate successive digits
1.5e1_       // invalid: _ must separate successive digits
```

### 虚数字面量

### 符文字面量

### 字符串字面量

## 常量

有布尔常量、符文常量、整数常量、浮点数常量、复数常量和字符串常量。符文、整数、浮点数和复数常量统称为数字常量。

常量值由符文、整数、浮点数、虚数或字符串字头、表示常量的标识符、常量表达式、结果为常量的转换或一些内置函数的结果值表示，如unsafe.Sizeof应用于某些值，cap或len应用于某些表达式，real和imag应用于复数常量，复数应用于数字常量。布尔真值由预先声明的常数true和false表示。预先声明的标识符iota表示一个整数常数。

一般来说，复数常量是常数表达的一种形式，在该节中讨论。

数字常量代表任意精度的精确值，不会溢出。因此，没有表示IEEE-754负零、无穷大和非数字值的常数。

常量可以是类型化的或非类型化的。字面常量，true，false，iota，以及某些只包含未定型常量操作数的常量表达式是未定型的。

常量可以通过常量声明或转换明确地给出类型，或者在变量声明或赋值语句中使用时隐含地给出类型，或者作为表达式的操作数。如果常量值不能被表示为相应类型的值，那就是一个错误。如果该类型是一个类型参数，常量将被转换为类型参数的非常量值。

一个没有类型的常量有一个默认的类型，在需要类型值的情况下，常量被隐含地转换为这个类型，例如，在一个短变量声明中，如i := 0，没有明确的类型。非类型常量的默认类型分别是bool, rune, int, float64, complex128或string，这取决于它是一个布尔常量、rune、整数、浮点、复合或字符串常量。

实施限制：尽管数字常量在语言中具有任意的精度，但编译器可以使用有限精度的内部表示法来实现它们。这就是说，每个实现都必须：

- 表示整数常量，至少有256位。
- 用至少256位的尾数和至少16位的有符号二进制指数来表示浮点数常量，包括复数的各个部分。
- 如果不能精确表示一个整数常量，请给出一个错误。
- 如果由于溢出而无法表示一个浮点数常量或复数常量，请给出一个错误。
- 如果由于精度的限制，无法表示浮点数常量或复数常量，请将其四舍五入到最接近的可表示常量。

这些要求既适用于字面常量，也适用于常量表达式的评估结果。

## 变量

变量是一个存储位置，用于保存一个值。允许的值的集合是由变量的类型决定的。

变量声明或对于函数参数和结果，函数声明或函数字面的签名为一个命名的变量保留存储空间。调用内置函数new或获取复合字面的地址，在运行时为一个变量分配存储空间。这样的匿名变量是通过一个（可能是隐含的）指针指示来引用的。

数组、切片和结构类型的结构化变量具有可以单独寻址的元素和字段。每个这样的元素都像一个变量。

变量的静态类型（或仅仅是类型）是在其声明中给出的类型，在新调用或复合字面中提供的类型，或结构化变量的元素的类型。接口类型的变量也有一个独特的动态类型，它是运行时分配给变量的值的（非接口）类型（除非该值是预先声明的标识符nil，它没有类型）。动态类型在执行过程中可能会发生变化，但是存储在接口变量中的值总是可以被分配到变量的静态类型中。

```
var x interface{}  // x is nil and has static type interface{}
var v *T           // v has value nil, static type *T
x = 42             // x has value 42 and dynamic type int
x = v              // x has value (*T)(nil) and dynamic type *T
```

变量的值是通过在表达式中引用该变量来检索的；它是最近分配给该变量的值。如果一个变量还没有被赋值，它的值就是其类型的零值。

## 类型

一个类型决定了一组值，以及针对这些值的操作和方法。一个类型可以用一个类型名来表示（如果有的话），如果这个类型是泛型，它后面必须有类型参数。一个类型也可以用一个类型名来指定，它由现有的类型组成一个类型。

```
Type      = TypeName [ TypeArgs ] | TypeLit | "(" Type ")" .
TypeName  = identifier | QualifiedIdent .
TypeArgs  = "[" TypeList [ "," ] "]" .
TypeList  = Type { "," Type } .
TypeLit   = ArrayType | StructType | PointerType | FunctionType | InterfaceType | liceType | MapType | ChannelType .
```

该语言预先声明了某些类型的名称。其他类型是通过类型声明或类型参数列表引入的。复合类型——数组、结构体、指针、函数、接口、切片、映射和通道类型——它们可以用类型字面来构造。

预先声明的类型、定义的类型和类型参数被称为命名类型。如果在别名声明中给出的类型是一个命名的类型，那么一个别名就表示一个命名的类型。

### 布尔类型

布尔类型表示由预先声明的常量`true`和`false`表示的布尔真值的集合。预先声明的布尔类型是`bool`；它是一个定义的类型。

### 数值类型

一个整数、浮点或复数类型分别代表了整数、浮点或复数的数值集合。它们被统称为数值类型。预先声明的独立于体系结构的数字类型是：

```
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
```

n位整数的值是n位宽，并使用二进制补码算法表示。

还有一组预先声明的整数类型，具有特定于实现的大小：

```
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

为了避免可移植性问题，所有的数字类型都是定义的类型，因此是不同的，除了byte（uint8的别名）和rune（int32的别名）。当不同的数字类型在表达式或赋值中混合时，需要进行明确的转换。例如，int32和int不是同一类型，尽管它们在一个特定的架构上可能具有相同的大小。

### 字符串类型

一个字符串类型代表了字符串值的集合。一个字符串值是一个（可能为空）的字节序列。字节的数量被称为字符串的长度，并且永远不会是负数。字符串是不可改变的：一旦创建，就不可能改变字符串的内容。预先声明的字符串类型是string；它是一种定义的类型。

一个字符串s的长度可以通过内置函数len来发现。如果字符串是一个常数，那么长度就是一个编译时常数。一个字符串的字节可以通过整数索引0到len(s)-1来访问。取这样一个元素的地址是非法的；如果s[i]是一个字符串的第i个字节，那么&s[i]是无效的。

### 数组类型

一个数组是由单一类型的元素组成的编号序列，称为元素类型。元素的数量被称为数组的长度，并且永远不会是负数。

```
ArrayType   = "[" ArrayLength "]" ElementType .
ArrayLength = Expression .
ElementType = Type .
```

长度是数组类型的一部分；它必须评估为一个非负常数，可由int类型的值表示。数组a的长度可以用内置函数len来发现。元素可以通过整数索引0到len(a)-1来寻址。数组类型总是一维的，但可以组成多维的类型。

```
[32]byte
[2*N] struct { x, y int32 }
[1000]*float64
[3][5]int
[2][2][2]float64  // same as [2]([2]([2]float64))
```

如果一个数组类型T只是数组或结构类型，那么该数组类型T不能直接或间接地有一个T类型的元素，或包含T的类型作为一个组件。

```
// invalid array types
type (
	T1 [10]T1                 // element type of T1 is T1
	T2 [10]struct{ f T2 }     // T2 contains T2 as component of a struct
	T3 [10]T4                 // T3 contains T3 as component of a struct in T4
	T4 struct{ f T3 }         // T4 contains T4 as component of array T3 in a struct
)

// valid array types
type (
	T5 [10]*T5                // T5 contains T5 as component of a pointer
	T6 [10]func() T6          // T6 contains T6 as component of a function type
	T7 [10]struct{ f []T7 }   // T7 contains T7 as component of a slice in a struct
)
```

### 切片类型

一个切片是一个底层数组的连续段的描述符，并提供对该数组元素的编号序列的访问。一个切片类型表示其元素类型的所有数组切片的集合。元素的数量被称为切片的长度，并且永远不会是负数。一个未初始化的切片的值是nil。

```
SliceType = "[" "]" ElementType .
```

切片s的长度可以通过内置函数len发现；与数组不同，它可能在执行过程中发生变化。元素可以通过整数索引0到len(s)-1来寻址。一个给定元素的分片索引可能小于底层数组中同一元素的索引。

一个切片，一旦初始化，总是与容纳其元素的底层数组相关。因此，一个切片与它的数组以及同一数组的其他切片共享存储；相反，不同的数组总是代表不同的存储。

一个切片的底层数组可以延伸到该切片的末端。容量是对这一范围的衡量：它是切片的长度和切片之外的数组的长度之和；一个长度达到该容量的切片可以通过从原始切片中切出一个新的切片来创建。一个切片的容量a可以通过内置函数cap(a)来发现。

对于一个给定的元素类型T，一个新的、初始化的切片值可以使用内置函数make来创建，该函数接收一个切片类型和指定长度和可选容量的参数。用make创建的切片总是分配一个新的、隐藏的数组，返回的切片值指向该数组。也就是说，执行

```
make([]T, length, capacity)
```

产生的切片与分配数组并进行切片相同，所以这两个表达式是等价的：

```
make([]int, 50, 100)
new([100]int)[0:50]
```

像数组一样，切片总是一维的，但可以组成更高维的对象。对于数组的数组，内部数组在结构上总是相同的长度；但是对于切片的切片（或切片的数组），内部长度可以动态变化。此外，内部切片必须被单独初始化。

### 结构体类型

结构体是一个命名元素的序列，称为字段，每个字段都有一个名称和一个类型。字段名可以明确指定（IdentifierList）或隐式指定（EmbeddedField）。在一个结构体中，非空白字段名必须是唯一的。

```
StructType    = "struct" "{" { FieldDecl ";" } "}" .
FieldDecl     = (IdentifierList Type | EmbeddedField) [ Tag ] .
EmbeddedField = [ "*" ] TypeName [ TypeArgs ] .
Tag           = string_lit .
```

```
// An empty struct.
struct {}

// A struct with 6 fields.
struct {
	x, y int
	u float32
	_ float32  // padding
	A *[]int
	F func()
}
```

一个声明了类型但没有明确字段名的字段被称为嵌入字段。一个内嵌字段必须被指定为一个类型名T或一个指向非界面类型名*T的指针，T本身不能是一个指针类型。未限定的类型名作为字段名。

```
// A struct with four embedded fields of types T1, *T2, P.T3 and *P.T4
struct {
	T1        // field name is T1
	*T2       // field name is T2
	P.T3      // field name is T3
	*P.T4     // field name is T4
	x, y int  // field names are x and y
}
```

下面的声明是非法的，因为字段名在结构类型中必须是唯一的。

```
struct {
	T     // conflicts with embedded field *T and *P.T
	*T    // conflicts with embedded field T and *P.T
	*P.T  // conflicts with embedded field T and *T
}
```

如果x.f是一个合法的选择器，表示该字段或方法f，那么结构x中的一个嵌入式字段或方法f被称为晋升。

提升的字段与结构的普通字段一样，只是它们不能作为结构的复合字段名使用。

给定一个结构类型S和一个命名的类型T，推广的方法会被包含在结构的方法集中，具体如下：

- 如果S包含一个内嵌字段T，S和*S的方法集都包括具有接收器T的推广方法。
- 如果S包含一个嵌入式字段*T，S和*S的方法集都包括带有接收者T或*T的推广方法。

一个字段声明后面可以有一个可选的字符串字面标签，它成为相应字段声明中所有字段的属性。一个空的标签字符串等同于一个没有的标签。标签通过反射接口可见，并参与结构体的类型识别，但其他方面则被忽略。

```
struct {
	x, y float64 ""  // an empty tag string is like an absent tag
	name string  "any string is permitted as a tag"
	_    [4]byte "ceci n'est pas un champ de structure"
}

// A struct corresponding to a TimeStamp protocol buffer.
// The tag strings define the protocol buffer field numbers;
// they follow the convention outlined by the reflect package.
struct {
	microsec  uint64 `protobuf:"1"`
	serverIP6 uint64 `protobuf:"2"`
}
```

如果一个结构体类型T只是数组或结构体类型，那么该结构体类型T不能直接或间接地包含T类型的字段或包含T的类型作为一个组件。

```
// invalid struct types
type (
	T1 struct{ T1 }            // T1 contains a field of T1
	T2 struct{ f [10]T2 }      // T2 contains T2 as component of an array
	T3 struct{ T4 }            // T3 contains T3 as component of an array in struct T4
	T4 struct{ f [10]T3 }      // T4 contains T4 as component of struct T3 in an array
)

// valid struct types
type (
	T5 struct{ f *T5 }         // T5 contains T5 as component of a pointer
	T6 struct{ f func() T6 }   // T6 contains T6 as component of a function type
	T7 struct{ f [10][]T7 }    // T7 contains T7 as component of a slice in an array
)
```

### 指针类型

指针类型表示给定类型的变量的所有指针的集合，称为指针的基本类型。一个未初始化的指针的值是nil。

```
PointerType = "*" BaseType .
BaseType    = Type .
```

```
*Point
*[4]int
```

### 函数类型

### 接口类型

### 映射类型

映射是一个无序的元素组，由一种类型的元素（称为元素类型）组成，由另一种类型的唯一键集（称为键类型）进行索引。一个未初始化的地图的值是nil。

```
MapType     = "map" "[" KeyType "]" ElementType .
KeyType     = Type .
```

比较运算符 == 和 != 必须为键类型的操作数完全定义；因此键类型不能是函数、映射或切片。如果键类型是一个接口类型，这些比较运算符必须为动态键值定义；失败将导致运行时的恐慌。

```
map[string]int
map[*T]struct{ x, y float64 }
map[string]interface{}
```

映射的元素数量被称为其长度。对于一个映射m来说，它可以用内置函数len来发现，并且在执行过程中可能会改变。在执行过程中可以用赋值添加元素，用索引表达式检索元素；可以用内置函数delete删除元素。

一个新的、空的映射值是用内置的函数make制作的，它以映射类型和一个可选的容量提示作为参数。

```
make(map[string]int)
make(map[string]int, 100)
```

初始容量并不约束其大小：映射的增长是为了适应其中存储的项目数量，但nil映射除外。一个nil映射等同于一个空映射，只是不能添加任何元素。

### 通道类型

## 类型和值的属性

### 底层类型

### 核心类型

### 类型特征

两种类型要么相同，要么不同。

一个被命名的类型总是与任何其他类型不同。否则，如果两个类型的底层类型字在结构上是等同的；也就是说，它们具有相同的字面结构，并且相应的组件具有相同的类型。详细来说：

- 如果两个数组类型有相同的元素类型和相同的数组长度，则它们是相同的。
- 如果两个切片类型有相同的元素类型，那么它们就是相同的。
- 如果两个结构体类型有相同的字段序列，并且相应的字段有相同的名称、相同的类型和相同的标签，那么它们就是相同的。来自不同包的非输出字段名总是不同的。
- 如果两个指针类型有相同的基本类型，那么它们就是相同的。
- 如果两个函数类型有相同数量的参数和结果值，相应的参数和结果类型是相同的，并且两个函数都是可变的，或者都不是。参数和结果名称不需要匹配。
- 如果两个接口类型定义了相同的类型集，那么它们就是相同的。
- 如果两个映射类型有相同的键和元素类型，它们就是相同的。
- 如果两个通道类型有相同的元素类型和相同的方向，那么它们是相同的。
- 如果两个实例化的类型的定义类型和所有类型参数都是相同的，那么它们就是相同的。

### 可分配性

### 可表示性

### 方法集

## 块

一个块是一个可能是空的声明和语句的序列，在匹配的大括号内。

```
Block = "{" StatementList "}" .
StatementList = { Statement ";" } .
```

除了源代码中的显性区块外，还有隐性区块。

1. universe块包含了所有的Go源文本。
2. 每个包都有一个包块，包含该包的所有Go源文本。
3. 每个文件都有一个文件块，包含该文件的所有Go源文本。
4. 每个“if”、“for”和“switch”语句都被认为是在自己的隐含块中。
5. “switch”或“select”语句中的每个子句都是一个隐式块。

块嵌套并影响范围。

## 声明和范围

### 标签作用域

### 空白标识符

### 预先声明的标识符

### 输出的标识符

### 标识符的唯一性

### 常数声明

### 符号

### 类型声明

### 类型参数声明

### 变量声明

### 短期变量声明

### 函数声明

### 方法声明

## 表达式

### 操作数

### 资格标识符

### 复合字词

### 函数字词

### 初级表达式

### 选择器

### 方法表达式

### 方法值

### 索引表达式

### 分片表达式

### 类型断言

### 调用

### 将参数传递给......参数

### 实例化

### 类型推理

### 操作员

### 算术运算符

### 比较运算符

### 逻辑运算符

### 地址运算符

### 接收运算符

### 转换

### 常数表达式

### 评估的顺序

## 语句

### 终止语句

### 空语句

### 有标签的语句

### 表达式语句

### 发送语句

### IncDec语句

### 赋值语句

### 条件语句

### 选择语句

### 赋值语句

### Go语句

### 选择语句

### 返回语句

### 中断语句

### 继续语句

### Goto语句

### 突破语句

### 延迟语句

## 内置功能

### 关闭

### 长度和容量

### 分配

### 制作切片、地图和通道

### 附加到和复制切片

### 删除地图元素

### 操纵复数

### 处理恐慌

### 引导

## 包

### 源文件组织

### 包条款

### 导入声明

### 一个实例包

## 程序的初始化和执行

### 零值

### 包的初始化

### 程序执行

## 错误

## 运行时惊慌失措


## 系统考虑

### 包装不安全

### 尺寸和对齐保证

[返回顶部 🔝](#Go编程语言规范)