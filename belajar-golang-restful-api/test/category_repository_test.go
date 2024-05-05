package test

import (
	"belajar-golang-restful-api/helper"
	"belajar-golang-restful-api/model/domain"
	"belajar-golang-restful-api/repository"
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func TestCategoryRepositorySaveSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()

	// ClearDatabase(db)
    
    category := domain.Category{Name: "Tes Category1 "}
    repo.Save(context.Background(), db, category)
}

func TestCategoryRepositoryUpdateSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()

    
    category := domain.Category{Id: 3, Name: "Tes Update Category"}
    savedCategory := repo.Update(context.Background(), db, category)

	assert.NotEmpty(t, savedCategory)
}

func TestCategoryRepositoryDeleteSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
	
	repo.Delete(context.Background(), db, 3)
}

func TestCategoryRepositoryFindByIdSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
	
    result, _ := repo.FindById(context.Background(), db, 4)

	fmt.Println(result)

	assert.NotEmpty(t, result)
}

func TestCategoryRepositoryFindByIdNotFound(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
	
    _, err := repo.FindById(context.Background(), db, 5)
	
	fmt.Println(err)

	assert.NotEmpty(t, err)
}

func TestCategoryRepositoryFindAllSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
	
    result := repo.FindAll(context.Background(), db)

	fmt.Println(result)

	assert.NotEmpty(t, result)
}

func SetupDatabase() *sql.DB{

	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.IfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func ClearDatabase(db *sql.DB){
	db.Exec("delete from category")
}


