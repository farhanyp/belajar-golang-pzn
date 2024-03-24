package repository

import (
	"context"
	"driver-pattern-repository/entity"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"testing"
	"time"
)

func GetConnection() *sql.DB{

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestCommentInsert(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	CommentRepository := NewCommentRepository(db)

	comment := entity.Comment{
		Email: "tes1@gmail.com",
		Comment: "Ini Tes 1",
	}

	result, err := CommentRepository.Insert(context.Background(), comment)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	CommentRepository := NewCommentRepository(db)


	result, err := CommentRepository.FindById(context.Background(), 35)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {

	db := GetConnection()
	defer db.Close()

	CommentRepository := NewCommentRepository(db)


	comments, err := CommentRepository.FindAll(context.Background())
	if err != nil{
		fmt.Println(err)
	}

	for _, comment := range comments {
		
		fmt.Println(comment)
	}
}