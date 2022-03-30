package main

import (
	"database/sql"
	"fmt"
	"go-restfulapi/app"
	"go-restfulapi/helper"
	"go-restfulapi/router"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func StartRoutes(db *sql.DB) http.Handler {
	
	v := validator.New()
	r := httprouter.New()
	
	router.StartCategoryRouter(r, db, v)

	return r
}

func ValidateDatabase(db *sql.DB){
	
	q_bytes, err := ioutil.ReadFile("sql/db.sql")
	helper.PanicError(err)

	queries := strings.Split(string(q_bytes), ";")
	for _, v := range queries {
		if v != ""{
			_, err := db.Exec(v)
			helper.PanicError(err)
		} 
	}

}

func StartServer(host string, port int) error {
	db := app.NewDB() 
    ValidateDatabase(db)
	handler := StartRoutes(db)

	server := http.Server{
		Addr: host + ":" + strconv.Itoa(port),
		Handler: handler,
	}

	fmt.Println("Server running at port: " + strconv.Itoa(port))
	err := server.ListenAndServe()

	return err
}

func main() {

    err := StartServer("localhost", 3000)
    helper.PanicError(err)

}