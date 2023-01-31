package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type User struct {
	Name string
}

func (user User) Introduce(strangerName string) string {
	return "Hello " + strangerName + " my name is " + user.Name
}

func TemplateFuncHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNC").Parse(`{{.Introduce "Alexander"}}`))
	t.ExecuteTemplate(w, "FUNC", User{
		Name: "Jack Dorsey",
	})
}

func TestTemplateFunc(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsStr := string(body)

	fmt.Println(bodyAsStr)
}

func TemplateFuncGlobalHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.New("FUNC").Parse(`Name lenght is {{len .Name}}`))
	t.ExecuteTemplate(w, "FUNC", User{
		Name: "Jack Dorsey",
	})
}

func TestTemplateFuncGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFuncGlobalHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyAsStr := string(body)

	fmt.Println(bodyAsStr)
}
