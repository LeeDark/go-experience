package main

import (
	"time"
)

func main() {
	channel := make(chan string, 1)

	go func(ch chan<- string) {
		ch <- "Hello World!"
		println("Finishing goroutine")
	}(channel)

	time.Sleep(time.Second)

	//message := <-channel
	//fmt.Println(message)

	receivingCh(channel)
}

func receivingCh(ch <- chan string) {
	msg := <-ch
	println(msg)

	// Error: invalid operation: ch <- "hello" (send to receive-only type <-chan string)
	//ch <- "hello"
}
