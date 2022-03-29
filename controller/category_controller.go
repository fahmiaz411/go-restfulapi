package controller

import (
	"go-restfulapi/helper"
	"go-restfulapi/model/web"
	"go-restfulapi/model/web/category_web"
	"go-restfulapi/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	Create(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}

 type CategoryControllerImpl struct {
	CategoryService service.CategoryService
 }

 func NewCategoryController(service service.CategoryService)CategoryController{
	return &CategoryControllerImpl{
		CategoryService: service,
	}
 }

 func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	categoryRequest := category_web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(r, &categoryRequest)

	categoryResponse := controller.CategoryService.Create(r.Context(), categoryRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
 }

 
 func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	categoryRequest := category_web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(r, &categoryRequest)

	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicError(err)

	categoryRequest.Id = categoryId

	categoryResponse := controller.CategoryService.Update(r.Context(), categoryRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
 }

 func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicError(err)

	response := controller.CategoryService.Delete(r.Context(), categoryId)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: response,
	}

	helper.WriteToResponseBody(w, webResponse)
 }

 func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	categoryId, err := strconv.Atoi(p.ByName("categoryId"))
	helper.PanicError(err)

	categoryResponse := controller.CategoryService.FindById(r.Context(), categoryId)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
 }

 func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, p httprouter.Params){
	categoryResponses := controller.CategoryService.FindAll(r.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Ok",
		Data: categoryResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
 }

