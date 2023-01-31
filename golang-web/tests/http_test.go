package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var url = "http://localhost:8080"

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, (url + "/"), nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsString := string(body)
	_ = bodyAsString

	// assert.True(t, (bodyAsString == "Hello World"))
}
