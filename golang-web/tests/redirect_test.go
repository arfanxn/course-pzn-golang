package tests

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "RedirectTo")
}

func RedirectFrom(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "RedirectFrom")
	http.Redirect(w, r, "/to", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RedirectFrom)
	mux.HandleFunc("/to", RedirectTo)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
