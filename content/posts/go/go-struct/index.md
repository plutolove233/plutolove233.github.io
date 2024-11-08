+++
title = 'Go指针、函数、结构体、接口以及方法'
date = 2024-11-08T22:39:44+08:00
draft = true
categories = ["golang之路"]
tags = ["golang", "study"]
+++
# 指针

# 函数
函数的定义方式如下：
```go
func 函数名(参数) (返回值) {
    函数体
}
```

# 结构体
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

# 推荐阅读：
1. [Go结构体的内存布局](https://www.liwenzhou.com/posts/Go/struct-memory-layout/)