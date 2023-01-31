package middlewares

import (
	"encoding/json"
	"golang-dependency-injection/helpers"
	"golang-dependency-injection/models/apis"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "SECRET" == r.Header.Get("X-Auth-Token") {
		middleware.Handler.ServeHTTP(w, r)
	} else {
		statusCode := http.StatusUnauthorized
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)

		response := apis.Response{
			Code:   int16(statusCode),
			Status: "Unauthorized",
		}

		encoder := json.NewEncoder(w)
		helpers.PanicIfError(
			encoder.Encode(response))
	}
}
