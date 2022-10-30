package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	//done := At(time.Now().Add(time.Minute), func() {
	//	fmt.Println("Hello world")
	//})
	//<-done // wait for fn to be done

	Every(time.Second, func() {
		fmt.Println("Hello world")
	})

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}