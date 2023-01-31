package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayoutHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"../templates/body.gohtml", "../templates/header.gohtml",
	))
	t.ExecuteTemplate(w, "body.gohtml", map[string]any{
		"Title": "Hello World",
		"Name":  "Jack Dorsey",
	})

}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateLayoutHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsStr := string(body)

	fmt.Println(bodyAsStr)
}
