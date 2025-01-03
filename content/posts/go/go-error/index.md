+++
title = 'Go Error'
date = 2025-01-03T19:50:55+08:00
draft = false
categories = ["golang之路"]
tags = ["golang", "study"]
+++
# Go error接口和错误处理
## error接口与错误处理
go与java不同的是，并不会采用`try-catch`方式来处理代码中出现的错误，而是通过`if err != nil`的方式来处理。其中error接口类型定义如下：
```go
type error interface {
    Error() string
}
```
`Error()`函数返回的结果是对错误内容的描述。由于`error`是接口类型，默认零值是`nil`，以此来判断是否发生错误，例如：
```go
file, err := os.Open("./a.txt")
if err != nil {
    fmt.Println("Open file failed, error=", err.Error())
    return
}
```
## 自定义错误
### errors库
可以通过errors库中的New函数，即可创建一个自定义的错误：
```go
import "errors"

func query(id int) (int, error) {
    if id < 1 {
        return 0, errors.New("invalid id")
    }
    return 1, nil
}

func main() {
    if res, err := query(0); err != nil {
        fmt.Println("get infomation failed, err=", err)
        return
    }
    fmt.Printf("get %d information", res)
}
```
### fmt.Errorf()结构化定义错误
当表达的错误字符串需要结构化表达时，可采用该方法实现，例如：
```go
fmt.Errorf("query database failed, error=%w", err)
```
这其中设计到错误链技术，可阅读这篇[博客](https://tonybai.com/2023/05/14/a-guide-of-using-go-error-chain/)，对该知识获取更加详细的理解，在一般开发过程中，并不需要考虑到这些。
### 自定义类型，实现接口
这种做法很好理解，就是自定义一个结构体，来表示特定的错误（这种方式更加类似java的做法），例如：
```go
type OpError struct {
    Op string
}
func (e *OpError) Error() string {
    return fmt.Sprintf("you don't have auth to exec %s", e.Op)
}
```