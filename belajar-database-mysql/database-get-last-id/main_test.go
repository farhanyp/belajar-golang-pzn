package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
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


func TestDatabase(t *testing.T){

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()


	email := "yp@gmail.com"
	comment := "ini yp"

	result, err := db.ExecContext(ctx, "INSERT INTO comments(email, comment) VALUES(? , ?)", email, comment)
	if err != nil{
		panic(err)
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Get id:", insertId)

}