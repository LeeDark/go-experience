package main

import (
	"container/ring"
	"fmt"
	"sync"
	"time"
)

type SafeRing struct {
	mu sync.Mutex
	r  *ring.Ring
}

func (sr *SafeRing) Len() int {
	return sr.r.Len()
}

func (sr *SafeRing) SetValue(value interface{}) {
	sr.mu.Lock()
	sr.r.Value = value
	sr.mu.Unlock()
}

func (sr *SafeRing) Next() *ring.Ring {
	sr.mu.Lock()
	next := sr.r.Next()
	sr.mu.Unlock()
	return next
}

func main() {
	r := ring.New(3)
	//sr := &SafeRing{r: ring.New(3)}
	s := []string{"IP1", "IP2", "IP3"}
	//for i := 0; i < sr.Len(); i++ {
	//	sr.SetValue(s[i])
	//	sr = sr.Next()
	//}
	for i := 0; i < r.Len(); i++ {
		r.Value = s[i]
		r = r.Next()
	}

	r.Do(func(p interface{}) {
		//fmt.Println(p)
		fmt.Println(p.(string))
	})

	// 10 goroutines
	next := func(g int) {
		//defer wg.Done()

		v := r.Value.(string)
		r = r.Next()
		fmt.Printf("goroutine %d, ring value: %s\n", g, v)
	}

	//var wg sync.WaitGroup
	//wg.Add(10)
	for i := 0; i < 10; i++ {
		go next(i)
		time.Sleep(10 * time.Millisecond)
	}
	//wg.Wait()
	time.Sleep(2 * time.Second)
}
