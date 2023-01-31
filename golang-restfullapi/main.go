package main

import (
	"golang-restfullapi/app"
	"golang-restfullapi/controllers"
	"golang-restfullapi/helpers"
	"golang-restfullapi/middlewares"
	"golang-restfullapi/repositories"
	"golang-restfullapi/services"
	"net/http"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repositories.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controllers.NewCategoryController(categoryService)

	router := app.NewCategoryRouter(categoryController)
	middleware := middlewares.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware,
	}

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
