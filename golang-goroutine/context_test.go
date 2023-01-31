package main

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB1 := context.WithValue(contextA, "b1", "B1")
	contextB2 := context.WithValue(contextA, "b2", "B2")

	// Child of context B1
	contextC1 := context.WithValue(contextB1, "c1", "C1")
	contextC2 := context.WithValue(contextB1, "c2", "C2")

	// Child of context B2
	contextC3 := context.WithValue(contextB2, "c3", "C3")

	fmt.Println(contextA)

	fmt.Println(contextB1)
	fmt.Println(contextB2)

	fmt.Println(contextC1)
	fmt.Println(contextC2)
	fmt.Println(contextC3)

	fmt.Println(contextC3.Value("c3")) // get self context value
	fmt.Println(contextC3.Value("b2")) // get parent context value
	fmt.Println(contextC3.Value("b1")) // get value from another parent of context (this will result nil coz the child does not have a relationship to another parent of context)
	fmt.Println(contextA.Value("b1"))  // get value from child context (this will result nil coz the child does not have accesible from parent context)

}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Start - Total Goroutine:", runtime.NumGoroutine())

	ctx := context.Background()
	ctxCounter, ctxCancelCounter := context.WithCancel(ctx)
	_ = ctxCancelCounter

	destination := CreateCounter(ctxCounter)
	fmt.Println("Middle - Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			ctxCancelCounter()
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Finish - Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Start - Total Goroutine:", runtime.NumGoroutine())

	ctx := context.Background()
	ctxCounter, ctxCancelCounter := context.WithTimeout(ctx, 5*time.Second) // will be cancelled after 5 seconds
	_ = ctxCancelCounter

	destination := CreateCounter(ctxCounter)
	fmt.Println("Middle - Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			ctxCancelCounter()
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Finish - Total Goroutine:", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Start - Total Goroutine:", runtime.NumGoroutine())

	ctx := context.Background()
	ctxCounter, ctxCancelCounter := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	_ = ctxCancelCounter

	destination := CreateCounter(ctxCounter)
	fmt.Println("Middle - Total Goroutine:", runtime.NumGoroutine())

	for n := range destination {
		fmt.Println("Counter:", n)
		if n == 10 {
			ctxCancelCounter()
			break
		}
	}

	time.Sleep(2 * time.Second)

	fmt.Println("Finish - Total Goroutine:", runtime.NumGoroutine())
}
