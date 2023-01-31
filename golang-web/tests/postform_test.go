package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func PostFormHandler(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	name := string(request.PostForm.Get("name"))
	fmt.Fprintln(writer, "Name: "+name)
}

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("name=Arfan")
	request := httptest.NewRequest(http.MethodPost, (url + "/"), requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	PostFormHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsString := string(body)

	fmt.Println(bodyAsString)
}
