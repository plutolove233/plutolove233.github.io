+++
title = 'Go指针、函数、结构体、接口以及方法'
date = 2024-11-08T22:39:44+08:00
draft = false
pre = "go-basic"
categories = ["golang之路"]
tags = ["golang", "study"]
+++
## 指针
和C语言一样，Go中也存在指针，一般采用如下方式来获取变量的地址：
```go
a := 12
var ptr *int = &a
```
如果我们想要获取该指针变量指向地址对应的内容时，我们可采用如下方式：
```go
fmt.Println(*ptr)
```
`*`操作符和`&`操作符分别代表了取值操作和取地址操作。

对于引用类型的变量，在Go中需要注意不仅需要声明，而且需要为其分配内存空间，不然就是对一个空指针进行操作了，具体的例子如下：
```go
func main(){
    var a *int
    *a = 100
}
```
为了修改上述代码，我们可以添加如下操作：
```go
func main() {
    var a *int
    a = new(int)
    *a = 100
}
```
`new(T)`函数能够返回传入类型`T`的内存地址，并且内存中对应的值为类型的零值。
## 函数
函数的定义方式如下：
```go
func 函数名(参数) (返回值) {
    函数体
}
```
其中需要说明的是：
- 参数由参数变量和参数变量的类型组成，多个参数之间使用`,`分隔。
- 如果相邻的参数类型相同，例如`a int, b int`可以简写为`a, b int`
- 如果传入参数数量不确定，可用`...T`来表示可变参数。例如`func (a string, x ...int)`
- 固定参数和可变参数搭配时，可变参数得放在固定参数后面。
- 可变参数可以是为一个数组
- 返回值由返回值变量和其变量类型组成，也可以**只写**返回值的类型，多个返回值必须用()包裹，并用,分隔。
- 如果返回值部分是由返回值变量和变量类型组成的话，函数体内部可以直接使用这个返回值，并在最后仅使用`return`来表示返回函数结果
我们定义一个a+b的函数作为示例：
```go
func intSum(a int, b int) res int {
    res = a + b
    return
}

int main() {
    b := intSum(1, 2) // 调用函数
}
```
### 函数类型
在go中，函数作为一等公民，因此可以作为变量进行传输。例如：
```go
func add(x, y int) int {
    return x + y
}

func minus(x, y int) int {
    return x - y
}

func calc(x, y int, op func(int, int)int) int {
    return op(x, y)
}

func main() {
    res := calc(1, 2, add)
    println(res)
}
```
但是在上述实现中，op的函数签名过长，实际应用中并不现实，而且不能突出这个变量的意义，因此我们可以为这个类型定义一个类型：
```go
type IntOperation func(int,int)int
```
可以理解成为`func(int,int)int`这个函数签名定义新的名字，当然，我们也可以为基本类型定义新名字：
```go
type MyInt int
```
通过上述方法，我们就可以简化代码为如下形式：
```go
type IntOperation func(int,int)int

func add(x, y int) int {
    return x + y
}

func minus(x, y int) int {
    return x - y
}

func calc(x, y int, op IntOperation) int {
    return op(x, y)
}

func main() {
    res := calc(1, 2, add)
    println(res)
}
```
### 匿名函数和闭包
函数同样可以作为返回值，在go中不能在函数内部定义函数，因此采用匿名函数来实现，例如：
```go
func main() {
    add := func(x, y int) int {
        fmt.Println(x+y)
    }
    add(1, 2)
    func (x, y int) int {
        fmt.Println(x+y)
    }(1, 2)
}
```
当函数的返回值是一个函数，并且该函数还包含相关引用环境时，我们就称其为闭包，例如：
```go
func addr() func(int) int {
    x := 10
    return func(y int) int {
        x += y
        return x
    }
}

func main() {
    f := addr()
    print(f(10)) // 20
    print(f(20)) // 40
}
```
### defer语句
`defer`语句会将其后面跟随的语句进行延迟处理,在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行，也就是说，**先被defer的语句最后被执行，最后被defer的语句，最先被执行**，例如：
```go
func deferDemo() {
    println("Hello")
    defer println("defer 1")
    defer println("defer 2")
    println("World")
}
/*
Hello
World
defer 2
defer 1
*/
```
其执行原理如下图所示：

![](./defer.png)

由于`defer`语句的这一特性，我们可以在defer后面执行例如资源清理、错误处理、解锁等操作。

## 结构体
go中结构体的定义方式如下：
```go
type StructName struct {
    a1 int
    b1 string
    ...
}
```
其中：
- StructName结构体名在同一个包内仅能有一个，不能重复
- a1,b1称为字段名，在每个结构体内唯一，不能重复
例如：
```go
type person struct {
    name string
    age int
    sex int
}
```
当我们声明了一个结构体后，如何实例化：
```go
var p person
或
p := new(person)
或
p := &person{}
```
后面两种得到的是结构体的指针，但是对于结构体而言，并不需要像`C++`中需要`(*p).name="John"`来修改属性，只需要`p.name=John`就行。

既然变量有匿名变量，结构体也会有匿名结构体：
```go
var user = struct {name string, age int}{
    "yizhigopher",
    12,
}
```
上面这段代码不仅展示了匿名结构体的用法，还展示了结构体初始化的一种方式，这种方式通过直接按着结构体属性的顺序赋值即可（需要注意的是，这种方法必须对所有属性都初始化）。

接下来展示其他初始化的方法：
```go
var p person
p.name = "yizhigopher"
// 或
p := person{
    name: "yizhigopher",
    age: 22,
}
// 或
p := &person{
    name: "yizhigopher"
}
```
在go中没有构造函数的说法，需要自己实现构造函数：
```go
func newPerson() *person {
    return &peroson{
        name: "yizhigopher",
        age: 12,
    }
}
```
Go中的结构体继承通过匿名结构体来实现，具体如下：
```go
//Animal 动物
type Animal struct {
	name string
}
//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}
```
其中关于匿名结构体需要注意的几点有：
- 当访问结构体成员时，会先在结构体中查找该字段，找不到再去嵌套的匿名字段中查找。
- 嵌套结构体内部可能存在相同的字段名。在这种情况下为了避免歧义需要通过指定具体的内嵌结构体字段名。

除此之外，结构体还有一个究极大杀器，称为结构体标签（Tag）。`Tag`是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 `Tag`在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
```go
`key1:"value1" key2:"value2"`
```
需要注意的是：
1. 结构体`tag`由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。同一个结构体字段可以设置多个键值对tag，不同的键值对之间使用空格分隔。
2. 为结构体编写`Tag`时，必须**严格遵守键值对的规则**=，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如**不要**在key和value之间添加空格。

我们举个例子来展示`tag`的能力：
```go
//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "yizhigopher",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"男"}
}
```
## 方法
go语言中，方法与OOP中的成员方法类似，是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的`this`或者`self`，具体定义的格式如下：
```go
func (接受者变量 接受者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```
可以发现，方法和函数之间在表达形式上的区别，就是在函数名前面增加了一段用于申明接受者类型的语句。需要注意的是：
- 接收者中的参数变量名在命名时，建议使用接收者类型名称首字母的小写，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
- 接收者类型和参数类似，可以是指针类型和非指针类型。
- 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。举个例子，我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。
```go
type MyInt int

func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}
```
当接受者类型是指针等引用类型时，方法内对于接受者的修改是直接作用于接受者的。但是如果接受者类型是指类型，那么是作用于接受者copy后的对象。可以通过下属例子体现其中的区别：
```go
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

func (p Person) SetAge2(newAge int8) {
	p.age = newAge
}
```
我们应当在以下这些情况，使用指针类型接受者：
1. 需要修改接收者中的值
2. 接收者是拷贝代价比较大的大对象
3. 保证**一致性**，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
## 推荐阅读：
1. [Go结构体的内存布局](https://www.liwenzhou.com/posts/Go/struct-memory-layout/)
2. [通过八个demo搞明白Go语言defer的五大特性
](https://learnku.com/articles/76556)
3. [Go Tag标签详解
](https://learnku.com/articles/78000)