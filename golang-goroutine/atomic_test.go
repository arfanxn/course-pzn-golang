package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	var group sync.WaitGroup = sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {

			defer group.Done()
			group.Add(1)

			for j := 0; j < 100; j++ {
				atomic.AddInt64(&counter, 1)
			}

		}()
	}
	group.Wait()

	fmt.Println("Counter :", counter)
}
