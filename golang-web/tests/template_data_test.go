package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/name.gohtml"))
	t.ExecuteTemplate(w, "name.gohtml", map[string]any{
		"Title": "Hello World",
		"Name":  "Jack Dorsey",
	})
}

func TestTemplateData(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()
 
	TemplateDataHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsStr := string(body)

	fmt.Println(bodyAsStr)
}
