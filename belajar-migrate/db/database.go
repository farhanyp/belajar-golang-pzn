package app

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {

	db, err := sql.Open("postgres", "postgres://postgres:12345678@localhost/belajar_golang_migration_pzn?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db

	// go install -tags ‘postgres’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest

	// Untuk migrate
	// migrate create -ext sql -dir db/migrations create_table_category

	// Untuk migrate up
	// migrate -database "postgres://postgres:12345678@localhost/belajar_golang_migration_pzn?sslmode=disable" -path db/migrations up

	// Untuk migrate down
	// migrate -database "postgres://postgres:12345678@localhost/belajar_golang_migration_pzn?sslmode=disable" -path db/migrations down
}