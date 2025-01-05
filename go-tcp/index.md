# Go TCP编程


通过前面几个章节的学习，对Go语言的基础语法有了基本的了解，现在在这几个章节的基础上，结合`net`库实现TCP服务端和客户端之间的通信。
## TCP服务端
首先我们先来看一下代码
```go
// server/main.go
package main

import (
	"bufio"
	"fmt"
	"net"
)

var id int = 0

func process(id int, c net.Conn) {
	defer c.Close()
	for {
		reader := bufio.NewReader(c)
		var data [128]byte
		n, err := reader.Read(data[:])
		if err != nil {
			fmt.Println("read from connection failed, err=", err)
			break
		}
		recv := string(data[:n])
		fmt.Printf("get data from client %d, data=%s", id, recv)
		c.Write([]byte(recv))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
    // 错误处理
	if err != nil {
		fmt.Println("build listener failed, err=", err)
		return
	}
	defer listener.Close()

	for {
		// 等待客户端与服务端建立连接，否则就阻塞
		conn, err := listener.Accept()
		id += 1
		if err != nil {
			fmt.Println("build connection failed, error=", err)
			continue
		}
		go process(id, conn)
	}
}
```
我们逐条解析一下这段代码。首先看一下变量定义`var id int = 0`，这个变量是为了区分每条服务端与客户端之间的连接。之后我们进入主函数，看一下代码逻辑：
1. 通过`net.Listen`函数，在本地`8080`端口建立监听器，从而能够对客户端的请求作出响应。
2. 之后通过死循环，确保程序能够一直运行。每轮循环中，`listener`会通过`Accept`方法对每次的建立连接的请求创建连接通道，如果没有客户端与服务端连接时，程序会阻塞在这一步。
3. 之后创建一个协程，运行`process`函数，用来实现客户端和服务端这条连接的通信与数据处理。
4. `process`函数中，会一直对连接读取数据，在终端打印客户端传来的数据，并将该数据返回给客户端。其中关于`bufio`库的操作会在之后关于Go常用库一章中详细说明。
## TCP客户端
同样的，我们先给出代码，并对代码进行解释：
```go
// client/main.go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("can't build connection to server, err=", err)
		return
	}
	defer conn.Close()
	readBuffer := bufio.NewReader(os.Stdin)
	for {
		msg, err := readBuffer.ReadString('\n')
		if err != nil {
			fmt.Println("Get msg from console failed, err=", err)
			return
		}
		info := strings.Trim(msg, "\r\n")
		if strings.ToUpper(info) == "Q" { // 按Q退出，断开连接
			return
		}
		_, err = conn.Write([]byte(info))
		if err != nil {
			fmt.Println("send msg to server failed, err=", err)
			continue
		}
		resp := [512]byte{}
		n, err := conn.Read(resp[:])
		if err != nil {
			fmt.Println("get message from server failed, err=", err)
			return
		}
		fmt.Printf("Get resp from server, msg=%s\n", string(resp[:n]))
	}
}
```
对于客户端的代码，就相对简单些。
1. 通过`net.Dial`方法，向本地`8080`端口请求建立连接。
2. 通过`bufio.NewReader`方法，从控制台读取输入数据。
3. 之后通过死循环，获取要发送的数据信息，并通过`conn.Write`方法，向连接发送数据。
4. 之后再通过`conn.Read`方法，获取服务端返回结果，并输出到控制台上。
