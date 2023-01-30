package main

import (
	"fmt"
	"regexp"
	"sync"
	"testing"
)

func TestDispatcher(t *testing.T) {
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
		req := Request{
			Data:    fmt.Sprintf("(Msg_id: %d) -> Hello", i),
			Handler: func(i interface{}) {
				s, ok := i.(string)
				defer wg.Done()

				if !ok {
					t.Fail()
				}

				ok, err := regexp.Match(`WorkerID\: \d* -\> \(MSG_ID: \d*\) -> [A-Z]*\sWorld`, []byte(s))
				if !ok || err != nil {
					t.Fail()
				}
			},
		}

		dispatcher.MakeRequest(req)
	}

	dispatcher.Stop()

	wg.Wait()
}