package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			data.Store(i, i)
		}()
	}

	group.Wait()

	data.Range(func(key, value any) bool {
		fmt.Println("Key:", key, "Value:", value)
		return true
	})
}
