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

## 变量

## 类型

### 布尔类型

### 数值型

### 字符串类型

### 结构体类型

### 切片类型

### 结构类型

### 指针类型

### 函数类型

### 接口类型

### 映射类型

### 通道类型

类型和值的属性
底层类型
核心类型
类型特征
可分配性
可表示性
方法集

块状物

声明和范围
标签作用域
空白标识符
预先声明的标识符
输出的标识符
标识符的唯一性
常数声明
符号
类型声明
类型参数声明
变量声明
短期变量声明
函数声明
方法声明

表达式
操作数
资格标识符
复合字词
函数字词
初级表达式
选择器
方法表达式
方法值
索引表达式
分片表达式
类型断言
调用
将参数传递给......参数
实例化
类型推理
操作员
算术运算符
比较运算符
逻辑运算符
地址运算符
接收运算符
转换
常数表达式
评估的顺序

语句
终止语句
空语句
有标签的语句
表达式语句
发送语句
IncDec 语句
赋值语句
如果语句
开关语句
赋值语句
Go语句
选择语句
返回语句
中断语句
继续语句
Goto语句
突破语句
延迟语句

内置功能
关闭
长度和容量
分配
制作片断、地图和通道
附加到和复制片断
删除地图元素
操纵复数
处理恐慌
引导

包裹
源文件组织
包条款
导入声明
一个实例包

程序的初始化和执行
零值
包的初始化
程序执行

错误

运行时惊慌失措

系统考虑
包装不安全
尺寸和对齐保证

[返回顶部 🔝](#Go编程语言规范)