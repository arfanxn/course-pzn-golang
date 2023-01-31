package main

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"net/http/httptest"
	"testing"

	httprouter "github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

//go:embed resources
var resources embed.FS

func TestServeFile(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources")

	router := httprouter.New()

	// Serve file
	router.ServeFiles("/resources/*filepath", http.FS(directory))

	// URL Routes
	//

	filepath := "documents/README.txt"
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/resources/"+filepath, nil)
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
