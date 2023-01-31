package tests

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBag(t *testing.T) {
	names := Bag[string]{"Muhammad", "Arfan"}
	PrintBag(names)
}
