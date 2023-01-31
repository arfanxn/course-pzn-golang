package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestHandlerServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Home page")
	})

	mux.HandleFunc("/about", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "About page")
	})

	mux.HandleFunc("/info", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.RequestURI)
		fmt.Fprintln(writer, request.Method)
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
