+++
title = 'go语言小技巧--编译约束'
date = 2023-12-04T19:13:11+08:00
draft = false
tags=['golang', 'tips']
categories=['golang小技巧']
+++

### 简介

我们能在go代码中添加一段注释，用来表示一种编译约束

```go
//go:build [tags]
```

通过这个并且配合`go build -tags="..."`的方式，就能实现带有约束的编译。

举个例子：

我们编写一下代码：主程序申明如下

```go
//main.go
package main

var configArr []string

func main(){
    for _, v := range configArr {
        print(v, " ")
    }
}
```

同时，我们会编写另外两个用于初始化的代码，分别为dev和test

```go
//go:build dev

package main

func init() {
    configArr = append(configArr, "dev")
}
```

```go
//go:build test

package main

func init() {
    configArr = append(configArr, "test")
}
```

我们可以看到，在两个同一包下，我们定义了两个不同的初始化功能，分别为`configArr`添加相应的内容。

当我们采用不同方式编译代码时，会看到如下情况：

![](./result_1.png)

能够看到编译时携带的tag不同，结果也就不同。

我们可以很直观的想到可以通过这种方法，实现不同模式下的编译

### 环境设置

当然我们也可以通过添加`//go:build linux`的方式对本段代码进行编译环境的限制。

例如我们对上述代码的`go:build`注释进行修改，可以看到如下效果：

```go
//go:build linux

package main

func init() {
    configArr = append(configArr, "dev")
}
```

```go
//go:build windows

package main

func init() {
    configArr = append(configArr, "test")
}
```

这时，我们直接进行代码编译，能够看到如下效果：

![](./result_2.png)

可以看到，我们只将test部分的代码编译，而dev的代码并没有编译

### 格式

举个例子：`// go:build linux && amd64 && ！cgo || darwin`，通过这个我们就能发现，可以通过逻辑运算符来表明编译约束。
