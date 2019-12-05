package main

import "time"

func main() {
	go helloWorld()

	go func() {
		println("Hello Anonymous!")
	}()

	go func(msg string) {
		println("Hello " + msg + "!")
	}("Anonymous with Data")

	messagePrinter := func(msg string) {
		println(msg)
	}

	go messagePrinter("Hello Function as Var!")

	time.Sleep(time.Second)
}

func helloWorld() {
	println("Hello World!")
}
