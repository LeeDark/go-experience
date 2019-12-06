package channel_singleton

import (
	"fmt"
	"testing"
	"time"
)

func TestStartMutiexInstance(t *testing.T) {
	singleton := GetMutexInstance()
	singleton2 := GetMutexInstance()

	n := 5000

	for i := 0; i < n; i++ {
		go singleton.AddOne()
		go singleton2.AddOne()
	}

	fmt.Printf("Before loop, current count is %d\n", singleton.GetCount())

	var val int
	for val != n*2 {
		val = singleton.GetCount()
		time.Sleep(10 * time.Millisecond)
	}
}
