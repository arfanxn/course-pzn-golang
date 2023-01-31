package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMethodNotAllowed(t *testing.T) {
	router := httprouter.New()

	/* Http Request Method Not Allowed Handling */
	router.MethodNotAllowed = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Method Not Allowed")
	})

	/* URL Routes */
	router.POST("/login", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "You have been authenticated")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/login", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		assert.FailNow(t, "Failed to read response body")
	}
	body := string(bytes)

	fmt.Println(body)
}
