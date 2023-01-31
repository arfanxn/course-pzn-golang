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

func TestPanicHandler(t *testing.T) {
	router := httprouter.New()

	/* Http Request Error handling */
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, i interface{}) {
		fmt.Fprintln(w, "Something went wrong")
	}

	/* URL Routes */
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		panic("You got an error")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		assert.FailNow(t, "Failed to read response body")
	}
	body := string(bytes)

	fmt.Println(body)

	// assert.Equal(t, body, "Hello World")
}
