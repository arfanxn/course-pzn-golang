package app

import (
	"fmt"
	"golang-restfullapi/controllers"
	"golang-restfullapi/exceptions"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func NewCategoryRouter(categoryController controllers.ControllerContract) *httprouter.Router {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(w, "Hello World")
	})
	router.GET("/api/categories", categoryController.Get)
	router.GET("/api/categories/:categoryID", categoryController.Find)
	router.POST("/api/categories", categoryController.Save)
	router.PUT("/api/categories/:categoryID", categoryController.Update)
	router.DELETE("/api/categories/:categoryID", categoryController.Delete)

	router.PanicHandler = exceptions.HTTPErrorHandler

	return router
}
