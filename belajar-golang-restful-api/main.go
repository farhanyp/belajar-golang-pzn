package main

import (
	"belajar-golang-restful-api/app"
	"belajar-golang-restful-api/controller"
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/repository"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryController := controller.NewCategoryControllerImpl(categoryRepository, db, validate)

	router := httprouter.New()
	router.GET("/api/v1/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Hello Api")
	})

	router.GET("/api/v1/categories", categoryController.FindAll)
	router.GET("/api/v1/categories/:categoryId", categoryController.FindById)
	router.POST("/api/v1/categories", categoryController.Create)
	router.PUT("/api/v1/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/v1/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr: "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}