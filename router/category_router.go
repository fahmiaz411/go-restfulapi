package router

import (
	"database/sql"
	"go-restfulapi/controller"
	"go-restfulapi/helper"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func StartCategoryRouter(r *httprouter.Router, db *sql.DB, v *validator.Validate) {
	route := helper.RouterWithPrefix{
		Router: r,
		Prefix: "/api",
	}

	c := controller.NewCategoryController(db, v)

	route.GET("/categories", c.FindAll)
	route.GET("/categories/:categoryId", c.FindById)
	
	route.POST("/categories", c.Create)
	route.PUT("/categories/:categoryId", c.Update)
	route.DELETE("/categories/:categoryId", c.Delete)
}