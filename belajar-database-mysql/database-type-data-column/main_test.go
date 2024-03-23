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

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func InsertData (db *sql.DB) {

	ctx := context.Background()

	_, err := db.ExecContext(ctx, "INSERT INTO customer(id, name, email, balance, rating, birth_date, married) VALUES('yp2','yp', 'yp@gmail.com', 100000, 5.0, NULL, true)")
	if err != nil{
		panic(err)
	}

	fmt.Println("Success Insert Data")
}


func TestDatabase(t *testing.T){

	db := GetConnection()
	InsertData(db)
	defer db.Close()

	ctx := context.Background()

	rows, err := db.QueryContext(ctx, "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer")
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birth_date sql.NullTime
		var createdAt time.Time
		var married bool
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birth_date, &married, &createdAt )

		if err != nil {
			panic(err)
		}

		fmt.Println("id:",id)
		fmt.Println("name:",name)
		if email.Valid {
			fmt.Println("email:",email.String)
		}
		fmt.Println("balance:",balance)
		fmt.Println("rating:",rating)
		if birth_date.Valid {
			fmt.Println("birth_date:",birth_date.Time)
		}
		fmt.Println("createdAt:",createdAt)
		fmt.Println("married:",married)
	}

}