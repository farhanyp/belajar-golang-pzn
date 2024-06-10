package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnection() *gorm.DB{

	dsn := "host=localhost user=postgres password=12345678 dbname=belajar_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db

}

var DB = DBConnection()

func TestDBConnection(t *testing.T){

	assert.NotNil(t, DB)

}