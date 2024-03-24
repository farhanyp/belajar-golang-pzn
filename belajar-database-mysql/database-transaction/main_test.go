package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
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


	tx, err := db.Begin()
	if err != nil{
		panic(err)
	}



	for i := 0; i < 10; i++ {
		email := "yp" + strconv.Itoa(i) + "@gmail.com" 
		comment := "ini yp" + strconv.Itoa(i)
		script :=  "INSERT INTO comments(email, comment) VALUES(?,?)"

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		insertId, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Get id:", insertId)
	}

	// err = tx.Rollback()	
	// if err != nil {
	// 	panic(err)
	// }

	err = tx.Commit()	
	if err != nil {
		panic(err)
	}


}