package app

import (
	"golang-dependency-injection/middlewares"
	"net/http"
)

func NewServer(middleware *middlewares.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: middleware,
	}
}
