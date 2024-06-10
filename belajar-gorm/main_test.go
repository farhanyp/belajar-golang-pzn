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

func TestExecuteSQL(t *testing.T){

	err := DB.Exec("insert into sample(id,name) values ($1,$2)", "1","yp1").Error
	assert.Nil(t, err)

	err = DB.Exec("insert into sample(id,name) values ($1,$2)", "2","yp2").Error
	assert.Nil(t, err)

	err = DB.Exec("insert into sample(id,name) values ($1,$2)", "3","yp3").Error
	assert.Nil(t, err)

}

type Sample struct{
	Id		string
	Name	string
}

func TestRawSQL(t *testing.T){

	var sample Sample
	err := DB.Raw("SELECT id, name FROM sample WHERE id = $1", "1").Scan(&sample).Error
	assert.Nil(t, err)
	assert.Equal(t, "yp1", sample.Name)

	var samples []Sample
	err = DB.Raw("SELECT id, name FROM sample").Scan(&samples).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(samples))

}

func TestScanRow(t *testing.T){

	rows, err := DB.Raw("SELECT id, name FROM sample").Rows()
	assert.Nil(t, err)
	defer rows.Close()

	var samples []Sample
	for rows.Next(){

		err := DB.ScanRows(rows, &samples)
		assert.Nil(t, err)

	}
	assert.Equal(t, 3, len(samples))

}