+++
title = 'Go接口'
date = 2024-12-30T17:01:41+08:00
draft = false
categories = ["golang之路"]
tags = ["golang", "study"]
+++

# 接口
## 接口类型定义
接口是程序员定义的一种类型，定义中包含了一组方法，用于规定需要实现哪些方法，示例代码如下：
```go
type Animaler interface {
    Calls() string
    Name() string
    Resp(string) string
}
```
上述代码就定义了一种`Animal`接口类型，只要某种**类型**实现了该接口类型所规定的方法，就可以称其实现（implement）了该接口。

需要注意的有以下几点：
- 接口中声明的函数参数类型可以没有参数名
- 接口名的命名方式建议采用`er`结尾，且最好能够表达其基本意义

## 如何实现接口
上文已经说过，**任何类型**只要实现了接口中定义的所有方法，就可以说实现了这个接口，对于上述的例子，我们可以这样来实现。
```go
type Bird struct {

}
func (b *Bird) Calls() string {
    return "Chirpping!"
}

func (b *Bird) Name() string {
    return "Its name is robin."
}

func (b *Bird) Resp(order string) string {
    if order == "do sth" {
        return "Bird is singing!"
    }
    return "Bird is dead..."
}
```
上述这段代码中`Bird`实现了接口`Animaler`，这与java不同的是，接口实现并不需要直接申明该类型实现了什么接口，而是仅需实现接口规定的方法即可。
## 为什么使用接口
举个例子。对于写任务，会存在向文件写、向控制台写，甚至写的格式也会有一些不同，让我们看一下如果没有接口会有什么问题：
```go
type LoggerWritter struct {
}

func (lw *LoggerWritter) Write(msg string) string {
    return fmt.Sprintf("Write to log file, %s\n", msg)
}

type CmdWritter struct {

}

func (cw *CmdWritter) Write(msg string) string {
    return fmt.Sprintf("Write to command line, %s\n", msg)
}

// 定义一个写任务的场景
func RunLoggerWrite(lw LoggerWritter) {
    lw.Write("death!")
}

func RunCommandWrite(cw CmdWritter) {
    cw.Write("boring!")
}
```
在这个例子中，写方法需要单独的函数来执行，十分麻烦。但是如果将变量类型换成接口，逻辑就变得清晰多了。
```go
type Writer interface {
    Write(string) string
}
func DIYWrite(w Writer, msg string) {
    w.Write(msg)
}

func main() {
    // ...
    DIYWrite(loggerWritter, "death!")
    // ...
}
```
在写这个函数中，其实并不需要关注是有谁来写，只需要知道有个能够执行写任务的参数就行了。就好比插座和插头往往并不是一一对应的，只需要三孔插座对应三线插头即可。接口就承担着这个插座的作用，它说明了需要哪些功能，我们仅需向其传递满足这些功能的变量即可。

同样的，实际生活中也存在的许多这样的例子，我们向商家提供二维码付钱时，这个二维码可以是微信支付二维码，也可以是支付宝支付二维码，二维码发行者是谁并不会影响付钱、也不会存在需要不同设备来扫描这个二维码的问题，这就是接口存在的意义。在实际开发过程中，需要注意积极采用接口这一特性，能够大大提高开发效率，降低代码的耦合程度。

## 接口的注意点
1. 一个类型能够实现多个接口，一个接口也可以被多个类型实现，且实现的方式无需显式说明（即无需使用implement来说明实现了什么接口）
2. Go语言中的接口同样可以通过匿名嵌入的方式实现接口的组合，例如：
    ```go
    type Writer interface {
        Write([]byte) error
    }
    type Reader interface {
        Read([]byte) (int, error)
    }
    type IO interface {
        Writer
        Reader
    }
    ```
    上述例子中，IO接口就是Reader接口和Writer接口的组合
3. 空接口是指没有定义任何方法的接口类型，即`interface{}`或`any`，空接口可以作为函数参数以及map值的类型，这样函数以及map值接收的数据就没有要求了
4. 对于接口类型，可以通过类型断言来判断实际类型，具体方式如下：
    ```go
    package main

    type I interface {
        F()
    }

    type T1 struct {
    }

    func (t T1) F() {
        println("T1.F")
    }

    type T2 struct {
    }

    func (t *T2) F() {
        println("T2.F")
    }

    func checkType(i I) {
        // 类型断言判断类型
        switch i.(type) {
        case T1:
            println("T1")
        case *T2:
            println("T2")
        }
        // 类型断言，转换类型，如果成功，则ok值为true
        v, ok := i.(T1)
        if ok {
            v.F()
        }
    }

    func main() {
        var i I
        i = T1{}
        checkType(i)
        i = &T2{}
        checkType(i)
    }
    ```
5. 接口实现可通过值接收者实现，也可通过指针接收者实现，用什么方式实现的，那么接口能接收的就是那种类型，例如上面这段代码中的
    ```go
    var i I
    i = T1{} // 值接收者实现
    i = &T2{} // 指针接收者实现
    ```