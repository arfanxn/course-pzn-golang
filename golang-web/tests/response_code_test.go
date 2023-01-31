package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCodeHandler(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(writer, "Name is empty")
	} else {
		fmt.Fprintln(writer, "Name: "+name)
	}

}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, (url + "?name="), nil)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	ResponseCodeHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsString := string(body)

	fmt.Println(bodyAsString)
	fmt.Println("Response Code:", response.StatusCode)
}
