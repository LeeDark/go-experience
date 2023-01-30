package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(1)

	toLowerAsync("Hello Callbacks!", func(v string) {
		toLowerAsync(fmt.Sprintf("Callback: %s\n", v), func(v string) {
			fmt.Printf("Callback within %s", v)
			wait.Done()
		})
	})

	println("Waiting async response...")

	wait.Wait()
}

func toLowerAsync(word string, f func(string)) {
	go func() {
		f(strings.ToLower(word))
	}()
}