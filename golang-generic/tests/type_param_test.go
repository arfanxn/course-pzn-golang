package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GenericSingleParamater[T any](val T) T {
	fmt.Println(val)
	return val
}

func GenericIsSame[T comparable](val1, val2 T) bool {
	return val1 == val2
}

func GenericMultipleParamater[T1 any, T2 any](val1 T1, val2 T2) {
	fmt.Println(val1)
	fmt.Println(val2)
}

func TestGenericSingleParamater(t *testing.T) {
	expected := "arfanxn"
	actual := GenericSingleParamater(expected)
	assert.Equal(t, expected, actual)
}

func TestGenericMultipleParamater(t *testing.T) {
	GenericMultipleParamater("arfanxn", 1234)
}

func TestGenericIsSame(t *testing.T) {
	isSame := GenericIsSame(22, 22)
	assert.True(t, isSame)
}
