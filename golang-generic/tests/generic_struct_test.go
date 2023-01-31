package tests

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (this *Data[_]) Print() {
	fmt.Println(this)
}

func (this *Data[T]) SetFirst(val T) {
	this.First = val
}

func TestGenericStruct(t *testing.T) {
	data := new(Data[string])
	data.First = "Muhammad"
	data.Second = "Arfan"

	fmt.Println(data)
}

func TestGenericStructMethod(t *testing.T) {
	data := new(Data[string])
	data.First = "Muhammad"
	data.Second = "Arfan"

	data.SetFirst("Ahmad")

	fmt.Println(data)
}
