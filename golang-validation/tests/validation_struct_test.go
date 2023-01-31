package tests

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username        string `json:"username" validate:"required,email|alphanum"`
	Password        string `json:"password" validate:"required,alphanum,min=8,max=16"`
	ConfirmPassword string `json:"password" validate:"eqfield=Password"`
}

func TestValidationStruct(t *testing.T) {
	var validate *validator.Validate = validator.New()

	loginRequest := LoginRequest{
		Username:        "arfanxn",
		Password:        "password",
		ConfirmPassword: "password",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		t.Error(err.Error())
	}
}

func TestValidationStructWithErrDetail(t *testing.T) {
	var validate *validator.Validate = validator.New()

	loginRequest := LoginRequest{
		Username: "arfanxn",
		Password: "@*#$",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		for _, fieldError := range validationErrors {
			fmt.Println(
				"Error: " + fieldError.Field() +
					" on tag " + fieldError.Tag() +
					" with error message: " + fieldError.Error(),
			)
		}

		//
	}
}

func TestValidationTwoStructField(t *testing.T) {
	var validate *validator.Validate = validator.New()

	loginRequest := LoginRequest{
		Username:        "arfanxn",
		Password:        "password",
		ConfirmPassword: "password",
	}

	err := validate.Struct(loginRequest)

	if err != nil {
		t.Error(err.Error())
	}
}
