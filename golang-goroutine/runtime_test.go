package main

import (
	"fmt"
	"runtime"
	"testing"
)

func TestRuntimeInfo(t *testing.T) {
	fmt.Println("Total CPU :", runtime.NumCPU())
	fmt.Println("Total Thread :", runtime.GOMAXPROCS(-1))
	fmt.Println("Total Goroutine :", runtime.NumGoroutine())
}
