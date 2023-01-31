package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAcync(group *sync.WaitGroup) {
	group.Add(1)
	defer group.Done()
	fmt.Println("Hello world")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAcync(group)
	}

	group.Wait()
	fmt.Println("Done")
}
