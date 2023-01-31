package tests

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
	"text/template"
)

func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("../templates/upload_file.gohtml"))
	t.ExecuteTemplate(w, "upload_file.gohtml", nil)
}

func UploadFileReceiverHandler(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	fileDestination, err := os.Create("../resources/" + fileHeader.Filename)
	if err != nil {
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil {
		panic(err)
	}
	name := r.PostFormValue("name")

	fmt.Fprintln(w, "Name:", name)
	fmt.Fprintln(w, "Filename:", fileHeader.Filename)
}

func TestUploadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadFileHandler)
	mux.HandleFunc("/upload", UploadFileReceiverHandler)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("../resources")))) 

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
