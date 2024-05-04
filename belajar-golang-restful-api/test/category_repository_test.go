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
    
    tx, err := db.Begin()
	helper.IfError(err)

	defer func(){
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			helper.IfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			helper.IfError(errorCommit)
		}
	}()
    
    category := domain.Category{Name: "Test Category"}
    savedCategory := repo.Save(context.Background(), tx, category)

	assert.NotEmpty(t, savedCategory)
}

func TestCategoryRepositoryUpdateSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
    
    tx, err := db.Begin()
	helper.IfError(err)

	defer func(){
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			helper.IfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			helper.IfError(errorCommit)
		}
	}()
    
    category := domain.Category{Id: 3, Name: "Test Category edit"}
    updateCategory := repo.Update(context.Background(), tx, category)

	assert.NotEmpty(t, updateCategory)
}

func TestCategoryRepositoryDeleteSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
    
    tx, err := db.Begin()
	helper.IfError(err)

	defer func(){
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			helper.IfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			helper.IfError(errorCommit)
		}
	}()
	
	repo.Delete(context.Background(), tx, 3)
}

func TestCategoryRepositoryFindByIdSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
    
    tx, err := db.Begin()
	helper.IfError(err)

	defer func(){
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			helper.IfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			helper.IfError(errorCommit)
		}
	}()
	
    result, _ := repo.FindById(context.Background(), tx, 4)

	fmt.Println(result)

	assert.NotEmpty(t, result)
}

func TestCategoryRepositoryFindByIdNotFound(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
    
    tx, err := db.Begin()
	helper.IfError(err)

	defer func(){
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			helper.IfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			helper.IfError(errorCommit)
		}
	}()
	
    _, err = repo.FindById(context.Background(), tx, 5)
	
	fmt.Println(err)

	assert.NotEmpty(t, err)
}

func TestCategoryRepositoryFindAllSuccess(t *testing.T) {
    repo := repository.NewCategoryRepositoryImpl()
    db := SetupDatabase()
    defer db.Close()
    
    tx, err := db.Begin()
	helper.IfError(err)

	defer func(){
		err := recover()
		if err != nil {
			errorRollback := tx.Rollback()
			helper.IfError(errorRollback)
			panic(err)
		} else {
			errorCommit := tx.Commit()
			helper.IfError(errorCommit)
		}
	}()
	
    result := repo.FindAll(context.Background(), tx)

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


