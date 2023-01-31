package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouterParams(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Products: "+p.ByName("id"))
	})

	productId := 2
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/"+strconv.Itoa(productId), nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		assert.FailNow(t, "Failed to read response body")
	}
	body := string(bytes)
	assert.Equal(t, body, "Products: "+strconv.Itoa(productId))
}
