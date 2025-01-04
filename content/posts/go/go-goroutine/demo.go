package main

import (
	"fmt"
	"time"
)

func work(ch chan int) {
	data, ok := <- ch
	if ok {
		fmt.Println(data)
		return
	}
	fmt.Println("channel is empty")
}

func main() {
	ch := make(chan int)
	go work(ch)
	ch <- 1
	time.Sleep(1 * time.Second)
}