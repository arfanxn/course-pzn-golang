//go:build wireinject
// +build wireinject

package main

import (
	"golang-dependency-injection/app"
	"golang-dependency-injection/controllers"
	"golang-dependency-injection/middlewares"
	"golang-dependency-injection/repositories"
	"golang-dependency-injection/services"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	repositories.NewCategoryRepository,
	services.NewCategoryService,
	controllers.NewCategoryController,
)

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewCategoryRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middlewares.NewAuthMiddleware,
		app.NewServer,
	)
	return &http.Server{}
}
