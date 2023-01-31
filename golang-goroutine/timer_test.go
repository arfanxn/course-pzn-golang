package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	returnedTime := <-timer.C
	fmt.Println("Afte r timer got event :")
	fmt.Println(returnedTime)
}

func TestTimerAfter(t *testing.T) {
	fmt.Println(time.Now())
	time := <-time.After(5 * time.Second)
	fmt.Println("After timer got event :")
	fmt.Println(time)
}

func TestTimerAfterFunction(t *testing.T) {
	var seconds time.Duration = 5 * time.Second
	group := sync.WaitGroup{}

	group.Add(1)
	time.AfterFunc(seconds, func() {
		fmt.Println("Runs after", seconds)
		group.Done()
	})

	group.Wait()
}
