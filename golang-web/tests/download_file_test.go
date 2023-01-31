package tests

import (
	"net/http"
	"testing"
)

func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+filename+"\"")
	http.ServeFile(w, r, "../resources/"+filename)
}

func TestDownloadFile(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/download", DownloadFileHandler)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
