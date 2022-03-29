package main

import (
	"fmt"
	"go-restfulapi/app"
	"go-restfulapi/controller"
	"go-restfulapi/helper"
	"go-restfulapi/repository"
	"go-restfulapi/service"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	port := 3000

	server := http.Server{
		Addr: "localhost:" + strconv.Itoa(port),
		Handler: router,
	}

	fmt.Println("Server running at port: " + strconv.Itoa(port))
	err := server.ListenAndServe()
	helper.PanicError(err)

}