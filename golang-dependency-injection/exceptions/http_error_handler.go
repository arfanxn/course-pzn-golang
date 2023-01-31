package exceptions

import (
	"encoding/json"
	"golang-dependency-injection/helpers"
	"golang-dependency-injection/models/apis"
	"net/http"

	"github.com/go-playground/validator"
)

func HTTPErrorHandler(w http.ResponseWriter, r *http.Request, err any) {

	if notFoundError(w, r, err) {
		return
	}

	if validationErrors(w, r, err) {
		return
	}

	internalServerError(w, r, err)
}

func internalServerError(w http.ResponseWriter, r *http.Request, err any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := apis.Response{
		Code:   http.StatusInternalServerError,
		Status: "Internal Server Error",
		Data:   err,
	}

	encoder := json.NewEncoder(w)
	helpers.PanicIfError(
		encoder.Encode(response))
}

func notFoundError(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		statusCode := http.StatusNotFound
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		response := apis.Response{
			Code:   int16(statusCode),
			Status: "Not Found",
			Data:   exception.Error,
		}

		encoder := json.NewEncoder(w)
		helpers.PanicIfError(
			encoder.Encode(response))
		return true
	} else {
		return false
	}
}

func validationErrors(w http.ResponseWriter, r *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		statusCode := http.StatusBadRequest
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		response := apis.Response{
			Code:   int16(statusCode),
			Status: "Bad Request",
			Data:   exception.Error(),
		}

		encoder := json.NewEncoder(w)
		helpers.PanicIfError(
			encoder.Encode(response))
		return true
	} else {
		return false
	}
}
