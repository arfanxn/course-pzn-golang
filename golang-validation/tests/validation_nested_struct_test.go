package tests

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

type Address struct {
	City    string `json:"city" validate:"required"`
	Country string `json:"country" validate:"required"`
}

type User struct {
	Id        string    `json:"id" validate:"required"`
	Name      string    `json:"name" validate:"required"`
	Addresses []Address `json:"addresses" validate:"dive,required"`
	Hobbies   []string  `json:"hobbies" validate:"dive,required,min=1"`
}

func TestValidationNestedStruct(t *testing.T) {
	validate := validator.New()

	user := User{
		Id:   "1234AAAA",
		Name: "Muhammad Arfan",
		Addresses: []Address{
			{
				City:    "Jakarta",
				Country: "Indonesia",
			},
			{
				City:    "Los Angeles",
				Country: "US",
			},
		},
		Hobbies: []string{
			" ",
		},
	}

	err := validate.Struct(user)

	if err != nil {
		t.Error(err.Error())
	}
}
