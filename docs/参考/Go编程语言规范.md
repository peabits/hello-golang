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

```go
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

函数类型表示具有相同参数和结果类型的所有函数的集合。一个未初始化的函数类型的变量的值是nil。

```go
FunctionType   = "func" Signature .
Signature      = Parameters [ Result ] .
Result         = Parameters | Type .
Parameters     = "(" [ ParameterList [ "," ] ] ")" .
ParameterList  = ParameterDecl { "," ParameterDecl } .
ParameterDecl  = [ IdentifierList ] [ "..." ] Type .
```

在一个参数或结果的列表中，名称（IdentifierList）必须全部存在或全部不存在。如果存在，每个名称代表指定类型的一个项目（参数或结果），签名中所有非空白的名称必须是唯一的。如果没有，每个类型代表该类型的一个项目。参数和结果列表总是用括号表示，但如果正好有一个未命名的结果，则可以写成未括号的类型。

在一个函数签名中，最后传入的参数可以有一个以`...`为前缀的类型。有这样一个参数的函数被称为variadic，可以用零个或多个参数来调用该参数。

```go
func()
func(x int) int
func(a, _ int, z float32) bool
func(a, b int, z float32) (bool)
func(prefix string, values ...int)
func(a, b int, z float64, opt ...interface{}) (success bool)
func(int, int, float64) (float64, *[]int)
func(n int) func(p *T)
```

### 接口类型

接口类型定义了类型集合。接口类型的变量可以存储该接口类型集中的任何类型的值。这样的类型被称为实现了该接口。一个未初始化的接口类型的变量的值是nil。

```go
InterfaceType  = "interface" "{" { InterfaceElem ";" } "}" .
InterfaceElem  = MethodElem | TypeElem .
MethodElem     = MethodName Signature .
MethodName     = identifier .
TypeElem       = TypeTerm { "|" TypeTerm } .
TypeTerm       = Type | UnderlyingType .
UnderlyingType = "~" Type .
```

一个接口类型是由一个接口元素的列表指定的。一个接口元素是一个方法或一个类型元素，其中一个类型元素是一个或多个类型项的联合。一个类型项要么是一个单一的类型，要么是一个单一的底层类型。


#### 基本接口

在其最基本的形式中，接口指定了一个（可能是空的）方法列表。这种接口所定义的类型集是实现所有这些方法的类型集，而相应的方法集则完全由接口所指定的方法组成。其类型集可以完全由方法列表来定义的接口被称为基本接口。

```go
// A simple File interface.
interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}
```

每个显式指定方法的名称必须唯一且不能为空。

```go
interface {
	String() string
	String() string  // illegal: String not unique
	_(x int)         // illegal: method must have non-blank name
}
```

可以使用一个以上的类型实现一个接口。例如，如果两个类型S1和S2的方法设置为

```go
func (p T) Read(p []byte) (n int, err error)
func (p T) Write(p []byte) (n int, err error)
func (p T) Close() error
```

(其中T代表S1或S2），那么文件接口就由S1和S2实现，而不管S1和S2可能有什么其他方法或共享什么方法。

作为一个接口的类型集成员的每个类型都实现了该接口。任何给定的类型都可以实现几个不同的接口。例如，所有类型都实现了空接口，它代表了所有（非接口）类型的集合。

```go
interface{}
```

方便起见，预先声明的类型any是空接口的别名。

同样地，考虑这个接口规范，它出现在一个类型声明中，定义了一个叫做Locker的接口。

```go
type Locker interface {
	Lock()
	Unlock()
}
```

如果S1和S2也实现了

```go
func (p T) Lock() { … }
func (p T) Unlock() { … }
```

它们同时实现了Locker接口和File接口。

#### 嵌入接口

在一个稍微一般的形式中，接口T可以使用一个（可能是限定的）接口类型名E作为接口元素。T的类型集是由T的明确声明的方法和T的嵌入接口的类型集定义的类型集的交集。换句话说，T的类型集是实现T的所有明确声明的方法和E的所有方法的所有类型的集合。

```go
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter's methods are Read, Write, and Close.
type ReadWriter interface {
	Reader  // includes methods of Reader in ReadWriter's method set
	Writer  // includes methods of Writer in ReadWriter's method set
}
```

当嵌入接口时，具有相同名称的方法必须具有相同的签名。

```go
type ReadCloser interface {
	Reader   // includes methods of Reader in ReadCloser's method set
	Close()  // illegal: signatures of Reader.Close and Close are different
}
```

#### 通用接口

在最一般的形式下，接口元素也可以是一个任意的类型术语T，或者是一个指定底层类型T的~T形式的术语，或者是一个术语t1|t2|...|tn的联合。与方法规范一起，这些元素能够精确定义一个接口的类型集，如下所示。

- 空接口的类型集是所有非接口类型的集合。
- 一个非空接口的类型集是其接口元素的类型集的交集。
- 一个方法规范的类型集是其方法集包括该方法的所有非接口类型的集合。
- 一个非接口类型术语的类型集是仅由该类型组成的集合。
- 一个形式为~T的术语的类型集是其基本类型为T的所有类型的集合。
- 术语t1|t2|...|tn的类型集是各术语的类型集的联合。

量化 "所有非界面类型的集合 "不仅指手头程序中声明的所有（非界面）类型，而且指所有可能程序中的所有可能类型，因此是无限的。同样地，给定实现某个特定方法的所有非接口类型的集合，这些类型的方法集的交集将正好包含该方法，即使手头的程序中的所有类型总是将该方法与另一个方法配对。

根据结构，一个接口的类型集永远不会包含一个接口类型。

```go
// An interface representing only the type int.
interface {
	int
}

// An interface representing all types with underlying type int.
interface {
	~int
}

// An interface representing all types with underlying type int that implement the String method.
interface {
	~int
	String() string
}

// An interface representing an empty type set: there is no type that is both an int and a string.
interface {
	int
	string
}
```

在一个形式为~T的术语中，T的底层类型必须是它自己，而且T不能是一个接口。

```go
type MyInt int

interface {
	~[]byte  // the underlying type of []byte is itself
	~MyInt   // illegal: the underlying type of MyInt is not MyInt
	~error   // illegal: error is an interface
}
```

联合元素表示类型集合的联合：

```go
// The Float interface represents all floating-point types
// (including any named types whose underlying types are
// either float32 or float64).
type Float interface {
	~float32 | ~float64
}
```

形式为T或~T的术语中的类型T不能是类型参数，所有非界面术语的类型集必须是成对不相交的（类型集的成对交集必须为空）。给定一个类型参数P。

```go
interface {
	P                // illegal: P is a type parameter
	int | ~P         // illegal: P is a type parameter
	~int | MyInt     // illegal: the type sets for ~int and MyInt are not disjoint (~int includes MyInt)
	float32 | Float  // overlapping type sets but Float is an interface
}
```

实现限制：一个联合（有一个以上的术语）不能包含预先声明的标识符可比性或指定方法的接口，或者嵌入可比性或指定方法的接口。

不是基本的接口只能作为类型约束使用，或者作为其他接口的元素作为约束使用。它们不能成为值或变量的类型，或其他非接口类型的组成部分。

```go
var x Float                     // illegal: Float is not a basic interface

var x interface{} = Float(nil)  // illegal

type Floatish struct {
	f Float                 // illegal
}
```

一个接口类型T不能直接或间接地嵌入一个是、包含或嵌入T的类型元素。

```go
// illegal: Bad may not embed itself
type Bad interface {
	Bad
}

// illegal: Bad1 may not embed itself using Bad2
type Bad1 interface {
	Bad2
}
type Bad2 interface {
	Bad1
}

// illegal: Bad3 may not embed a union containing Bad3
type Bad3 interface {
	~int | ~string | Bad3
}

// illegal: Bad4 may not embed an array containing Bad4 as element type
type Bad4 interface {
	[10]Bad4
}
```

#### 实现接口

类型T实现了接口I，如果

- T不是一个接口，并且是I的类型集的一个元素；
- T是一个接口，并且T的类型集是I的类型集的一个子集。

如果T实现了一个接口，那么一个T类型的值就实现了该接口。

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

## 内置函数

内置函数是[预先声明]()的。它们像其他函数一样被调用，但其中一些函数接受一个类型而不是表达式作为第一个参数。

内置函数没有标准的Go类型，所以它们只能出现在[调用表达式]()中；它们不能作为函数值使用。

### 关闭（Close）

对于一个核心类型为通道的参数ch，内置函数close会记录该通道上不再有任何值被发送。如果ch是一个只接收的通道，这就是一个错误。发送或关闭一个关闭的通道会引起运行时的恐慌。关闭nil通道也会引起运行时的恐慌。在调用close后，在任何先前发送的值被接收后，接收操作将返回通道类型的零值而不阻塞。多值接收操作会返回一个接收值以及通道是否被关闭的指示。

### 长度和容量

内置函数len和cap接受各种类型的参数，并返回一个int类型的结果。该实现保证结果总是适合于一个int。

```go
Call      Argument type    Result

len(s)    string type      string length in bytes
          [n]T, *[n]T      array length (== n)
          []T              slice length
          map[K]T          map length (number of defined keys)
          chan T           number of elements queued in channel buffer
          type parameter   see below

cap(s)    [n]T, *[n]T      array length (== n)
          []T              slice capacity
          chan T           channel buffer capacity
          type parameter   see below
```

如果参数类型是一个类型参数P，调用len(e)（或cap(e)）必须对P的类型集中的每个类型有效。其结果是参数的长度（或容量），其类型与P被实例化的类型参数相对应。

分片的容量是指底层数组中分配到的元素的数量。在任何时候，以下关系都是成立的：

```go
0 <= len(s) <= cap(s)
```

一个空的切片、映射或通道的长度为0。一个空的切片或通道的容量为0。

如果s是一个字符串常数，表达式len(s)就是常数。如果s的类型是一个数组或指向数组的指针，并且表达式s不包含通道接收或（非常量）函数调用，那么表达式len(s)和cap(s)是常量；在这种情况下，s不会被评估。否则，len和cap的调用不是常量，s被评估。

```go
const (
	c1 = imag(2i)                    // imag(2i) = 2.0 is a constant
	c2 = len([10]float64{2})         // [10]float64{2} contains no function calls
	c3 = len([10]float64{c1})        // [10]float64{c1} contains no function calls
	c4 = len([10]float64{imag(2i)})  // imag(2i) is a constant and no function call is issued
	c5 = len([10]float64{imag(z)})   // invalid: imag(z) is a (non-constant) function call
)
var z complex128
```

### 分配

内置函数new接收一个类型T，在运行时为该类型的变量分配存储空间，并返回一个指向它的*T类型的值。该变量被初始化，如初始值一节中所述。

```go
new(T)
```

例如

```go
type S struct { a int; b float64 }
new(S)
```

为一个S类型的变量分配存储空间，对其进行初始化（a=0，b=0.0），并返回一个包含该位置的地址的*S类型的值。

### 生成切片、映射和通道

内置函数make接收一个类型T，后面可以选择一个特定类型的表达式列表。T的核心类型必须是一个片断、地图或通道。它返回一个类型为T（不是*T）的值。存储器被初始化，如初始值一节中所述。

```go
Call             Core type    Result

make(T, n)       slice        slice of type T with length n and capacity n
make(T, n, m)    slice        slice of type T with length n and capacity m

make(T)          map          map of type T
make(T, n)       map          map of type T with initial space for approximately n elements

make(T)          channel      unbuffered channel of type T
make(T, n)       channel      buffered channel of type T, buffer size n
```

每个大小参数n和m必须是整数类型，有一个只包含整数类型的类型集，或者是一个没有类型的常数。一个常数的大小参数必须是非负的，并且可以用int类型的值来表示；如果它是一个未定型的常数，它的类型是int。如果n和m都被提供并且是常数，那么n必须不大于m。对于切片和通道，如果n在运行时是负数或者大于m，就会发生运行时恐慌。

```go
s := make([]int, 10, 100)       // slice with len(s) == 10, cap(s) == 100
s := make([]int, 1e3)           // slice with len(s) == cap(s) == 1000
s := make([]int, 1<<63)         // illegal: len(s) is not representable by a value of type int
s := make([]int, 10, 0)         // illegal: len(s) > cap(s)
c := make(chan int, 10)         // channel with a buffer size of 10
m := make(map[string]int, 100)  // map with initial space for approximately 100 elements
```

调用map类型和大小提示n的make将创建一个具有初始空间的map，以容纳n个map元素。准确的行为是依赖于实现的。

### 追加和复制切片

内置的函数append和copy可以帮助进行常见的分片操作。对于这两个函数，其结果与参数所引用的内存是否重叠无关。

变量函数append将零个或多个值x追加到一个片断s，并返回与s相同类型的结果片断。值x被传递给一个类型为...E的参数，各自的参数传递规则适用。作为一个特例，如果s的核心类型是[]字节，append也接受第二个参数，其核心类型是bytestring，后面是.... 这种形式附加了字节片或字符串的字节。

```go
append(s S, x ...E) S  // core type of S is []E
```

如果s的容量不足以容纳额外的值，append会分配一个新的、足够大的底层数组，以容纳现有的分片元素和额外的值。否则，append会重新使用底层数组。

```go
s0 := []int{0, 0}
s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

var t []interface{}
t = append(t, 42, 3.1415, "foo")   //                             t == []interface{}{42, 3.1415, "foo"}

var b []byte
b = append(b, "bar"...)            // append string contents      b == []byte{'b', 'a', 'r' }
```

函数copy将片状元素从源src复制到目的dst，并返回复制的元素数量。两个参数的核心类型必须是具有相同元素类型的片断。复制的元素数是len(src)和len(dst)的最小值。作为一种特殊情况，如果目的地的核心类型是[]字节，copy也接受一个核心类型为bytestring的源参数。这种形式将字节片或字符串中的字节复制到字节片中。

```go
copy(dst, src []T) int
copy(dst []byte, src string) int
```

例子：

```go
var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
var s = make([]int, 6)
var b = make([]byte, 5)
n1 := copy(s, a[0:])            // n1 == 6, s == []int{0, 1, 2, 3, 4, 5}
n2 := copy(s, s[2:])            // n2 == 4, s == []int{2, 3, 4, 5, 4, 5}
n3 := copy(b, "Hello, World!")  // n3 == 5, b == []byte("Hello")
```

### 删除映射元素

内置函数delete从地图m中删除键值为k的元素，值k必须可以分配给m的键类型。

```go
delete(m, k)  // remove element m[k] from map m
```

如果m的类型是一个类型参数，该类型集的所有类型必须是地图，而且它们必须都有相同的键类型。

如果映射m是nil或者元素m[k]不存在，那么delete是一个w无操作（no-op）。

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

当通过声明或调用new为一个变量分配存储空间时，或者通过复合字面或调用make创建一个新的值时，如果没有提供明确的初始化，那么这个变量或值将被赋予一个默认值。这种变量或值的每个元素都被设置为其类型的零值：布尔类型为false，数字类型为0，字符串为""，指针、函数、接口、片断、通道和地图为nil。这种初始化是递归进行的，因此，举例来说，如果没有指定值，一个结构数组的每个元素都将被归零。

这两个简单的声明是等价的：

```go
var i int
var i int = 0
```

在

```go
type T struct { i int; f float64; next *T }
t := new(T)
```

之后，以下情况成立：

```go
t.i == 0
t.f == 0.0
t.next == nil
```

同样的情况也会发生在

```go
var t T
```

之后。

### 包的初始化

在一个包内，包级变量的初始化是逐步进行的，每一步都会选择在声明顺序中最早的、与未初始化的变量没有依赖关系的变量。

更确切地说，如果一个包级变量还没有被初始化，并且没有初始化表达式，或者它的初始化表达式与未初始化的变量没有依赖关系，那么这个包级变量就被认为可以被初始化。初始化的过程是重复初始化在声明顺序中最早的、准备好初始化的下一个包级变量，直到没有准备好初始化的变量。

如果在这个过程结束时，仍有任何变量未被初始化，那么这些变量就是一个或多个初始化循环的一部分，程序是无效的。

变量声明左侧的多个变量由右侧的单个（多值）表达式初始化，并一起初始化。如果左手边的任何一个变量被初始化，所有这些变量都在同一步骤中被初始化。

```go
var x = a
var a, b = f() // a and b are initialized together, before x is initialized
```

为了实现包的初始化，空白变量与声明中的任何其他变量一样被处理。

在多个文件中声明的变量的声明顺序是由文件呈现给编译器的顺序决定的。在第一个文件中声明的变量要在第二个文件中声明的任何变量之前声明，以此类推。

依赖性分析不依赖于变量的实际值，只依赖于源文件中对它们的词法引用，并进行转述分析。例如，如果一个变量x的初始化表达式引用了一个函数，该函数的主体引用了变量y，那么x就依赖于y。

- 对一个变量或函数的引用是一个表示该变量或函数的标识符。
- 对方法m的引用是一个方法值或者形式为t.m的方法表达式，其中t的（静态）类型不是一个接口类型，并且方法m在t的方法集中，由此产生的函数值t.m是否被调用并不重要。
- 如果x的初始化表达式或主体（对于函数和方法）包含对y的引用或对依赖于y的函数或方法的引用，那么一个变量、函数或方法x就依赖于一个变量y。

例如，给定声明

```go
var (
	a = c + b  // == 9
	b = f()    // == 4
	c = f()    // == 5
	d = 3      // == 5 after initialization has finished
)

func f() int {
	d++
	return d
}
```

注意，初始化表达式中子表达式的顺序是不相关的：在这个例子中，a = c + b和a = b + c的初始化顺序是一样的。

依赖性分析是按包进行的；只考虑引用当前包中声明的变量、函数和（非接口）方法。如果变量之间存在其他隐藏的数据依赖关系，那么这些变量之间的初始化顺序是不指定的。

例如，给定声明

```go
var x = I(T{}).ab()   // x has an undetected, hidden dependency on a and b
var _ = sideEffect()  // unrelated to x, a, or b
var a = b
var b = 42

type I interface      { ab() []int }
type T struct{}
func (T) ab() []int   { return []int{a, b} }
```

变量a将在b之后被初始化，但x是在b之前，b和a之间，还是在a之后被初始化，因此也没有指定调用sideEffect()的时刻（在x被初始化之前或之后）。

变量也可以使用在包块中声明的名为init的函数进行初始化，没有参数，也没有结果参数。

```go
func init() { … }
```

每个包可以定义多个这样的函数，甚至在一个源文件中也可以。在包块中，init标识符只能用于声明init函数，但标识符本身并没有被声明。因此，init函数不能在程序中的任何地方被引用。

一个没有导入的包的初始化方法是给它所有的包级变量分配初始值，然后按照它们在源文件中出现的顺序调用所有的init函数，可能是在多个文件中出现的，就像提交给编译器的那样。如果一个包有导入，在初始化包本身之前，导入的包会被初始化。如果多个包导入了一个包，导入的包将只被初始化一次。从结构上看，导入包可以保证不存在循环的初始化依赖。

包的初始化--变量的初始化和init函数的调用--发生在一个goroutine中，按顺序，一次一个包。一个init函数可以启动其他goroutine，它们可以与初始化代码同时运行。然而，初始化总是对init函数进行排序：在前一个函数返回之前，它不会调用下一个函数。

为了保证初始化行为的可重复性，我们鼓励构建系统以词法文件名的顺序向编译器提交属于同一软件包的多个文件。

### 程序执行

一个完整的程序是通过将一个叫做main包的未导入的包与它所导入的所有包连接起来，过渡性地创建的。主包必须有包名main，并声明一个不需要参数也不需要返回值的函数main。

```go
func main() { ... }
```

程序的执行从初始化main包开始，然后调用函数main。当该函数的调用返回时，程序就退出了。它不会等待其他（非main）goroutine完成。

## 错误

## 运行时惊慌失措


## 系统考虑

### 包装不安全

### 尺寸和对齐保证

[返回顶部 🔝](#Go编程语言规范)