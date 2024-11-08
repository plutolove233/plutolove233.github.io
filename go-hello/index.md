# Hello Go!


在完成了上个章节的Go编译器安装后，本章节将介绍Go的基础语法，并通过几个例子进行展示。
# Hello World
学习每个语言的第一步永远是学会如何在控制台打印😊，现在让我们看一下go如何实现`Hello World`：
```go
// main.go
package main

import "fmt"

func main(){
    fmt.Println("Hello World!")
}
```
> 在本文件的目录下，在控制台中执行`go build main.go && main`或`go run main.go`观察控制台输出的结果。

我们逐行分析说明：
- 第1行使用`package main`申明这是一个主包，我们需要在主包下面定义`func main()`函数，用于声明程序入口。
- 第3行使用`import`来导入外部包`"fmt"`，该包通常用于实现格式化的I/O操作，这个与C的`<cstdio>`相似。
- 第5行使用`func main()`来申明编译入口以及程序执行的开始位置。
- 第6行使用`fmt.Println("Hello World!")`向控制台输出`Hello World!`

通过上述例子，我们介绍下go语言一个程序所需要的三个部分：

1. 包名。每个go程序都需要申明该程序所属的包，在本章节中，我们所有代码的包名均为`main`，而对于go的包管理方式，我们将在后续的章节中进行介绍。
2. 导入外部包。想要实现一段代码，不可避免需要调用已经封装好的程序，因此通过`import`来导入一些必要的外部包，例如本例子中的`fmt`，以及一些常用的`bytes`，`net`等，这些均是已经封装好的标准库。这些标准库的位置在`${GOROOT}/src`中
3. 函数。go语言函数申明格式为`func functionName(args)`并用花括号包裹函数体。需要注意的是，不像C++可以将花括号换行，go语言中花括号的换行是不允许换行的，并使用tab进行代码缩进。
接下来我们将介绍go的基本语法
# GO基本语法
## 变量和常量
### 变量
变量定义如下方式如下：
```go
var a int
var a int = 10
var a = 10
a := 10
```
上述为4种不同表达形式的变量申明方式，其中a为变量名，`int`为变量类型，后两种方式中，go能够自动推理出各个变量的类型。

go的变量名可以用26个英文字母，`_`以及数字组合而成

> 数字不能作为变量的起始字符

此外go中存在一些关键字，其具备特殊含义，包含一下25个，这些单词不能作为变量名：
```go
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```
此外还存在一些保留字，这些同样不建议作为变量名：
```go
    Constants:    true  false  iota  nil

        Types:    int  int8  int16  int32  int64  
                  uint  uint8  uint16  uint32  uint64  uintptr
                  float32  float64  complex128  complex64
                  bool  byte  rune  string  error

    Functions:   make  len  cap  new  append  copy  close  delete
                 complex  real  imag
                 panic  recover

```
如果需要定义多个变量可以如此声明：
```go
var (
    a int = 19
    b bool = true
)
或
var (
    a = 10
    b = false
)
```
### 常量定义
与变量定义类似，可以如此定义
```go
const a int = 10
const a = 10 // notice the type of a is 'untyped int'
const (
    a = 10
    b string = "12a"
)
```
需要注意的是，常量的批量定义中，如果没有确定常量的类型，go并不会直接推到出常量的类型，而是`untyped ...`，关于这个问题可以查看[官方文档](https://go.dev/blog/constants)，对于`untyped`常量会导致的问题，可以查看[疑惑: Go const 导致程序结果错乱?](https://mp.weixin.qq.com/s/9OngnAzyKfL7Y2nAYtaxhA)
### 多重赋值
这个特点有点像js中的多重赋值，我们可以这么做
```go
    x, y, z := 1, 2, 3
```
这时，`x`的值为1，`y`的值就为2，……
你可以随处使用该方法来实现多重赋值，例如可以用在常量定义中
### 匿名变量
go中会使用`_`来表示一个仅用于接收，但不能使用的变量，常用于当我们想要忽略多重赋值中的某个值时。例如：
```go
    x, _, z := 1, 2, 3
```
### iota
`iota`是`go`语言的常量计数器，只能在常量的表达式中使用。

`iota`在`const`关键字出现时将被重置为0。`const`中每新增一行常量声明将使`iota`计数一次。使用iota能简化定义，在定义枚举时很有用。例如：
```go
const (
    a = iota // a = 0
    b        // b = 1
    c        // c = 2
)
或
const (
    _ = iota
    KB = 1<<(10*iota)
    MB
	GB
)
```
可以发现，后续的常量推导，均是依据离这条语句最近的iota表达来推导的。
# 代码风格
在go的代码中，变量命名的风格通常具有一个约定俗成的规则，这里我将总结一下我遇到的一些规则
- 对于常量名，一般全部大写，并用`_`隔开，例如：`MAX_CONNECTION`。
- 对于变量名以及函数名，一般采用驼峰风格，例如`isHashed`。如果需要将这个函数或变量设为public，则将首字母也大写。
- 当然变量的定义最好能够直观表现其意义，最好别用拼音、无意义的一个字符串来定义一个变量。

