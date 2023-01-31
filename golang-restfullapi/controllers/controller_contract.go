package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

/* Contract */
type ControllerContract interface {
	Get(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Find(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Save(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
