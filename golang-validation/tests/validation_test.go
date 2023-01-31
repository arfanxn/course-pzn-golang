package tests

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate *validator.Validate = validator.New()

	if validate == nil {
		t.Error("Validator should not be nil")
	}
}

func TestValidationField(t *testing.T) {
	var validate *validator.Validate = validator.New()

	var username string = " "
	err := validate.Var(username, "required")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestValidationTwoField(t *testing.T) {
	var validate *validator.Validate = validator.New()

	password, confirmPassword := "password", "password"

	err := validate.VarWithValue(password, confirmPassword, "required,alphanum,eqfield,min=8,max=16")

	if err != nil {
		t.Error(err.Error())
	}
}
