package tests

import (
	"fmt"
	"testing"
)

type GetterSetterContract[T any] interface {
	GetValue() T
	SetValue(value T)
}

/* Individual function */

func SetAndGetValue[T any](object GetterSetterContract[T], valueToSet T) T {
	object.SetValue(valueToSet)
	return object.GetValue()
}

/* Implementation of getter setter contract  */

type Parcel[T any] struct {
	Value T
}

func (this *Parcel[T]) SetValue(value T) {
	this.Value = value
}

func (this *Parcel[T]) GetValue() T {
	return this.Value
}

/* Testing methods */

func TestGenericInterfae(t *testing.T) {
	parcel := new(Parcel[string])
	SetAndGetValue[string](parcel, "Jack Dorsey")

	fmt.Println(parcel)
}
