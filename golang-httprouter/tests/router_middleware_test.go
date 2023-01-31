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

type LogMiddleware struct {
	http.Handler
}

func (middlware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Before")
	middlware.Handler.ServeHTTP(w, r)
	fmt.Fprintln(w, "After")
}

func TestMiddleware(t *testing.T) {
	router := httprouter.New()

	/* URL Routes */
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "Hello World")
	})

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	middleware := new(LogMiddleware)
	middleware.Handler = router
	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		assert.FailNow(t, "Failed to read response body")
	}
	body := string(bytes)

	fmt.Println(body)
}
