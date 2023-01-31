package tests

import (
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func MustValidIDPhoneNumber(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)

	if ok {
		return strings.Contains(value, "62")
	} else {
		return false
	}
}

func MustEqualsIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2() // get second field's value

	if !ok {
		panic("Field not ok")
	}

	val1 := strings.ToUpper(field.Field().String())
	val2 := strings.ToUpper(value.String())

	return val1 == val2
}

func MustValidUserName(level validator.StructLevel) {
	user := level.Current().Interface().(User)

	if len(user.Name) >= 2 {

	} else {
		level.ReportError(user.Name, "Name", "Name", "user-name", "")
	}
}

func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("idphone", MustValidIDPhoneNumber)

	phoneNumber := "62 777 888 999"
	err := validate.Var(phoneNumber, "idphone")

	if err != nil {
		t.Error(err.Error())
	}
}

func TestCustomCrossFieldValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equals_ignore_case", MustEqualsIgnoreCase)

	type Nothing struct {
		One string `validate:"required,field_equals_ignore_case=Two"`
		Two string
	}

	nothing := Nothing{
		One: "NOThing",
		Two: "nOthiNG",
	}

	if err := validate.Struct(nothing); err != nil {
		t.Error(err.Error())
	}
}

func TestCustomStructValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidUserName, User{})

	user := User{
		Id:   "1234AAAA",
		Name: "sule",
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
			"Gaming",
		},
	}

	if err := validate.Struct(user); err != nil {
		t.Error(err.Error())
	}

}
