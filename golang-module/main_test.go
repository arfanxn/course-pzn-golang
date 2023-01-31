package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	fmt.Println("Before starting tests")

	m.Run()

	fmt.Println("After finished tests")
}

func TestSayHello(t *testing.T) {
	actual := SayHello()
	expected := "Hello"
	require.Equal(t, expected, actual)
	// assert.Equal(t, expected, actual)
	fmt.Println("Test succeeded")
}
func TestSubTest(t *testing.T) {
	t.Run("Arfan", func(t *testing.T) {
		actual := SayHello()
		expected := "Hello"
		require.Equal(t, expected, actual)
	})
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello()
 	}
}
