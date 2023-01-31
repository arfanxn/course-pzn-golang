package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserContract interface {
	GetName() string
}

type User struct {
	Name string
}

func (this *User) GetName() string {
	return this.Name
}

func GetName[T UserContract](object T) string {
	return object.GetName()
}

func TestGenericInterfaceInheritance(t *testing.T) {
	expected := "arfanxn"
	actual := GetName(&User{Name: expected})
	assert.Equal(t, expected, actual)
}
