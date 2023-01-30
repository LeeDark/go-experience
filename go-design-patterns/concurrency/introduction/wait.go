package main

import (
	"fmt"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(1)

	go func() {
		fmt.Println("Hello World")
		wait.Done()
	}()

	wait.Wait()

	goRoutines := 5
	wait.Add(goRoutines)

	for i := 0; i < goRoutines; i++ {
		go func(goRoutineID int) {
			fmt.Printf("ID:%d: Hello goroutines!\n", goRoutineID)
			wait.Done()
		}(i)
	}

	wait.Wait()
}