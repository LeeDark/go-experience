package main

import (
	"fmt"
	"log"
	"sync"
)

type RequestHandler func(interface{})

type Request struct {
	Data interface{}
	Handler RequestHandler
}

func NewStringRequest(s string, id int, wg *sync.WaitGroup) Request {
	myRequest := Request{
		Data:    s,
		Handler: func(i interface{}) {
			defer wg.Done()
			res, ok := i.(string)
			if !ok {
				log.Fatal("Invalid casting to string")
			}
			fmt.Println(res)
		},
	}

	return myRequest
}

func main() {
	bufferSize := 100
	var dispatcher Dispatcher = NewDispatcher(bufferSize)

	workers := 3
	for i := 0; i < workers; i++ {
		var w WorkerLauncher = &PreffixSuffixWorker{
			id:      i,
			prefixS: fmt.Sprintf("WorkerID: %d -> ", i),
			suffixS: " World",
		}

		dispatcher.LaunchWorker(w)
	}

	requests := 10

	var wg sync.WaitGroup
	wg.Add(requests)

	for i := 0; i < requests; i++ {
		req := NewStringRequest(fmt.Sprintf("(Msg_id: %d) -> Hello", i), i, &wg)
		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}