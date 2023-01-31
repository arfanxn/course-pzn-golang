package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "URL Queries:")
	fmt.Fprintln(w, r.URL.Query())

	name := r.URL.Query().Get("name")
	fmt.Fprintln(w, "Hello "+name)
}

func TestSayHelloByQuery(t *testing.T) {
	request := httptest.NewRequest(
		http.MethodGet, "http://localhost:8080?name=Arfan", nil)
	recorder := httptest.NewRecorder()

	SayHelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsString := string(body)

	fmt.Println(bodyAsString)

}
