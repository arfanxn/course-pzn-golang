package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func CookieHandler(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "Auth"
	cookie.Value = "1234asdf124afas"
	cookie.Path = "/"

	http.SetCookie(writer, cookie)

	fmt.Println(request.Cookies())
}

func TestCookie(t *testing.T) {

	cookie := new(http.Cookie)
	cookie.Name = "Browser"
	cookie.Value = "Firefox"
	cookie.Path = "/"

	request := httptest.NewRequest(http.MethodPost, (url), nil)
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	CookieHandler(recorder, request)

	response := recorder.Result()
	cookies := response.Cookies()

	fmt.Println(cookies)
}
