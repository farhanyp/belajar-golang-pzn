package main

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)


func TestDatabase(t *testing.T){

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/godb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}