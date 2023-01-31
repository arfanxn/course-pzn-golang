package main

import (
	"embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFileHandler(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "../resources/index.css")
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func ServeFileEmbedHandler(writer http.ResponseWriter, request *http.Request) {
	file, _ := resources.Open("resources/index.css")
	fmt.Fprintln(writer, file)
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbedHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
