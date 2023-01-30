package main

import (
	"fmt"
	"strings"
	"sync"
)

func main() {
	toUpperSync("Hello Callbacks!", func(v string) {
		fmt.Printf("Sync Callback: %s\n", v)
	})

	var wait sync.WaitGroup
	wait.Add(1)

	toUpperAsync("Hello Callbacks!", func(v string) {
		fmt.Printf("Async Callback: %s\n", v)
		wait.Done()
	})

	println("Waiting async response...")

	wait.Wait()
}

func toUpperSync(word string, f func(string)) {
	f(strings.ToUpper(word))
}

func toUpperAsync(word string, f func(string)) {
	go func() {
		f(strings.ToUpper(word))
	}()
}