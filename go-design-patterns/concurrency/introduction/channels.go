package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel := make(chan string)
	go func() {
		channel <- "Hello World!"
	}()

	message := <- channel
	fmt.Println(message)

	//
	var wait sync.WaitGroup

	wait.Add(1)
	go func() {
		channel <- "Hello World!"
		println("Finishing goroutine")
		wait.Done()
	}()

	time.Sleep(time.Second)
	message = <- channel
	fmt.Println(message)
	wait.Wait()
}
