package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {

	ticker := time.NewTicker(1 * time.Second)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		fmt.Println("Ticker stopped after 5 seconds")
	}()

	for t := range ticker.C {
		fmt.Println(t)
	}

	

}
