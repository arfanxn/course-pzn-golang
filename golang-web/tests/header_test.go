package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeaderHandler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")

	fmt.Fprintln(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, (url + "/"), nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeaderHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsString := string(body)

	fmt.Println(bodyAsString)
}
