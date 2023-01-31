package controllers

import (
	"encoding/json"
	"golang-restfullapi/helpers"
	"golang-restfullapi/models/apis"
	"golang-restfullapi/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

/* Implementation */
type CategoryController struct {
	CategoryService services.CategoryServiceContract
}

func NewCategoryController(categoryService services.CategoryServiceContract) ControllerContract {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

func (controller CategoryController) Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.Get(r.Context())

	apiResponse := apis.Response{
		Code:   int16(200),
		Status: "OK",
		Data:   categoriesResponse,
	}

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	helpers.PanicIfError(
		encoder.Encode(apiResponse))
}

func (controller CategoryController) Find(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helpers.PanicIfError(err)

	categoryResponse, err := controller.CategoryService.Find(r.Context(), int32(categoryID))
	helpers.PanicIfError(err)
	apiResponse := apis.Response{
		Code:   int16(200),
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	helpers.PanicIfError(
		encoder.Encode(apiResponse))
}

func (controller CategoryController) Save(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	categoryCreateRequest := apis.CategoryCreateRequest{}
	helpers.PanicIfError(decoder.Decode(&categoryCreateRequest))

	categoryResponse := controller.CategoryService.Save(r.Context(), categoryCreateRequest)
	apiResponse := apis.Response{
		Code:   int16(200),
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	helpers.PanicIfError(
		encoder.Encode(apiResponse))
}

func (controller CategoryController) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)

	categoryUpdateRequest := apis.CategoryUpdateRequest{}
	helpers.PanicIfError(
		decoder.Decode(&categoryUpdateRequest))
	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helpers.PanicIfError(err)
	categoryUpdateRequest.Id = int32(categoryID)

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryUpdateRequest)
	apiResponse := apis.Response{
		Code:   int16(200),
		Status: "OK",
		Data:   categoryResponse,
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	helpers.PanicIfError(
		encoder.Encode(apiResponse))

}

func (controller CategoryController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	categoryID, err := strconv.Atoi(params.ByName("categoryID"))
	helpers.PanicIfError(err)

	controller.CategoryService.Delete(r.Context(), int32(categoryID))
	apiResponse := apis.Response{
		Code:   int16(200),
		Status: "OK",
	}

	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	helpers.PanicIfError(
		encoder.Encode(apiResponse))

}
