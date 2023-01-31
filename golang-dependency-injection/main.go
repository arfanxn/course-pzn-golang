package main

import (
	"golang-dependency-injection/helpers"
)

func main() {
	server := InitializeServer()

	err := server.ListenAndServe()
	helpers.PanicIfError(err)
}
