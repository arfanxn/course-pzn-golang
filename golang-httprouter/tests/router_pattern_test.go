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

func TestRouterPattern(t *testing.T) {
	router := httprouter.New()
	router.GET("/products/:productID/items/:itemID", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		productID := p.ByName("productID")
		itemID := p.ByName("itemID")
		fmt.Fprintln(w, "Products: "+productID)
		fmt.Fprintln(w, "Items: "+itemID)
	})
	router.GET("/resources/*filepath", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		filepath := p.ByName("filepath")
		fmt.Fprintln(w, "Filepath: "+filepath)
	})

	// productID, itemID := 2, 10
	// request := httptest.NewRequest(
	// 	http.MethodGet,
	// 	"http://localhost:8080/products/"+strconv.Itoa(productID)+"/items/"+strconv.Itoa(itemID),
	// nil)
	request := httptest.NewRequest(
		http.MethodGet,
		"http://localhost:8080/resources"+"/images/users/avatars/default.png",
		nil)

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
