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

func TestRouterNotFound(t *testing.T) {
	router := httprouter.New()

	/* Http Request Not Found Handling */
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "URL Not Found")
	})

	/* URL Routes */
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "Hello World")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/not-found/11717171717271y481unfjabfjabf", nil)
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
