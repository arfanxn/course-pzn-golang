package tests

import (
	"fmt"
	"testing"
)

type Number interface {
	~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Age int

func Max[T Number](val1, val2 T) T {
	switch true {
	case val1 > val2:
		return val1
	default:
		return val2
	}
}

func Min[T interface {
	~int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}](val1, val2 T) T {
	switch true {
	case val1 < val2:
		return val1
	default:
		return val2
	}
}

func TestGenericTypeSets(t *testing.T) {
	max := Max(Age(4444), Age(7777))
	fmt.Println(max)
}

func TestGenericInlineTypeSets(t *testing.T) {
	max := Min(Age(4444), Age(7777))
	fmt.Println(max)
}
