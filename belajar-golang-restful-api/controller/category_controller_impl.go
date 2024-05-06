package controller

import (
	"belajar-golang-restful-api/exception"
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/web"
	"belajar-golang-restful-api/repository"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryRepository repository.CategoryRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewCategoryControllerImpl(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate)CategoryController{
	return &CategoryControllerImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		Validate: validate,
	}
}

func(repo *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	categoryCreateRequest := web.CategoryCreateRequest{}

	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	err := repo.Validate.Struct(categoryCreateRequest)
	helper.IfError(err)

	categoryResponse := repo.CategoryRepository.Save(request.Context(), repo.DB, categoryCreateRequest)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(repo *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	CategoryUpdateRequest := web.CategoryUpdateRequest{}

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.IfError(err)

	CategoryUpdateRequest.Id = id

	helper.ReadFromRequestBody(request, &CategoryUpdateRequest)

	err = repo.Validate.Struct(CategoryUpdateRequest)
	helper.IfError(err)

	categoryResponse := repo.CategoryRepository.Update(request.Context(), repo.DB, CategoryUpdateRequest)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(repo *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.IfError(err)

	category, err := repo.CategoryRepository.FindById(request.Context(), repo.DB, id)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	repo.CategoryRepository.Delete(request.Context(), repo.DB, category.Id)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(repo *CategoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.IfError(err)

	category, err := repo.CategoryRepository.FindById(request.Context(), repo.DB, id)
	if err != nil{
		panic(exception.NewNotFoundError(err.Error()))
	}

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: category,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func(repo *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params){

	categories := repo.CategoryRepository.FindAll(request.Context(), repo.DB)

	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categories,
	}

	helper.WriteToResponseBody(writer, webResponse)
}