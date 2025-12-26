# Fmt包介绍

## 用途
`fmt`包用于实现读写操作，以及格式化字符串表达，具体使用方式如下：
### 输出
可通过`Printf`，`Print`以及`Println`函数实现向控制台输出指定内容。
首先介绍下最常用的`Printf`函数，包含两个参数，分别代表格式以及输出的内容。举个例子：
```go
package main

import "fmt"

func main() {
    var a int = 3
    fmt.Printf("this is a example, a = %d\n", a)
}
```
熟悉`C`语言的，能够发现这个部分的代码和`stdio.h`库中的`printf`函数相似。使用此函数主要需要了解不同占位符以及转义所代表的意义，因此本部分将主要介绍此类内容。我们用`%d`来表示这个位置数据的类型是整数类型；`\n`是转义符，表示回车。不同类型与其占位符的对应关系如下：
```
%d => int   %f => float   %t => bool    %s => string    %p => 引用类型(地址)
```
对于一些不确定的值，可以用`%v`来代替，例如结构体、字典、切片以及数组，均可以使用`%v`来输出内容。

如果想要输出变量的类型，可使用`%T`来实现，例如：
```go
import "fmt"

func main() {
	var arr = []int{1, 2, 3}
	fmt.Printf("%T\n", arr)
}
```
其输出结果是`[]int`

转义符以及其代表的含义可见如下内容：
<table align="center">
    <tbody align="center" valign="center">
        <tr>
            <td bgcolor="gray"><b><font color="white">转义符</font></b></td>
            <td bgcolor="gray"><b><font color="white">含义</font></b></td>
        </tr>
        <tr>
            <td>\n</td>
            <td>换行符</td>
        </tr>
        <tr>
            <td>\r</td>
            <td>回车符</td>
        </tr>
        <tr>
            <td>\t</td>
            <td>制表符</td>
        </tr>
        <tr>
            <td>\'</td>
            <td>单引号</td>
        </tr>
        <tr>
            <td>\"</td>
            <td>双引号</td>
        </tr>
        <tr>
            <td>\\</td>
            <td>反斜杠</td>
        </tr>
    </tbody>
</table>

### 格式化占位符
#### 通用占位符
```
%v  => 值的默认格式表示
%+v => 类似%v，但输出结构体时会添加字段名
%#v => 值的Go语法表示
%T  => 打印值的类型
```
#### 数值类型占位符
**整数**
```
%b => 二进制    %c => 该值对应的unicode码值
%d => 十进制    %x => 表示为十六进制，使用a-f
%o => 八进制    %X => 表示为十六进制，使用A-F
```
**浮点数**
```
%b => 二进制    %e/%E => 科学计数法
%f => 有小数部分但无指数部分
```
**字符串和[]byte**
```
%s => 直接输出字符串或者[]byte
%q => 该值对应的双引号括起来的go语法字符串字面值
%x => 每个字节用两字符十六进制数表示，适用a-f
%x => 每个字节用两字符十六进制数表示，适用A-F
```

### 输入
可通过`Scan`，`Scanf`，`Scanln`

