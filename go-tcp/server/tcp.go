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