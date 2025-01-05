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
		if strings.ToUpper(info) == "Q" {
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
