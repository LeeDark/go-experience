package main

import (
	"strconv"
	"sync"
	"time"
)

func main() {
	bufferSize := 20
	dispatcherTimeout := 60 * time.Second
	var dispatcher Dispatcher = NewDispatcher(bufferSize, dispatcherTimeout)

	workers := 10
	for i := 0; i < workers; i++ {
		w := NewWorker(i)
		dispatcher.LaunchWorker(w)
	}

	requests := 100

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		time.Sleep(100 * time.Millisecond)
		task := NewTask(strconv.Itoa(i), &wg)
		dispatcher.PushTask(task)
	}

	dispatcher.Stop()
	wg.Wait()
}
