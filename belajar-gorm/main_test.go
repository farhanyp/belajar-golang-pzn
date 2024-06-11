package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DBConnection() *gorm.DB{

	dsn := "host=localhost user=postgres password=12345678 dbname=belajar_gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

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

func TestCreateUser(t *testing.T){

	user := User{
		ID: "1",
		Password: "rahasia",
		Name: Name{
			FirstName: "farhan",
			MiddleName: "yudha",
			LastName: "pratama",
		},
	}

	response := DB.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)

}

func TestCreateBatchUser(t *testing.T){
	var users []User

	for i := 2; i <= 10; i++ {
		users = append(users, User{
			ID: strconv.Itoa(i),
			Name: Name{
				FirstName: "farhan" + strconv.Itoa(i),
				MiddleName: "yudha" + strconv.Itoa(i),
				LastName: "pratama" + strconv.Itoa(i),
			},
		})
	}



	response := DB.Create(&users)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(9), response.RowsAffected)

}

func TestTransactionRollback(t *testing.T){

	err :=DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(
			&User{
				ID: "11", 
				Password: "rahasia", 
				Name: Name{ 
					FirstName: "farhan", 
					MiddleName: "yudha", 
					LastName: "pratama",
					},
				},
			).Error

		if err != nil {
			return err
		}

		err = tx.Create(
			&User{
				ID: "12", 
				Password: "rahasia", 
				Name: Name{ 
					FirstName: "farhan", 
					MiddleName: "yudha", 
					LastName: "pratama",
					},
				},
			).Error

		if err != nil {
			return err
		}

		return nil
	})

	assert.Nil(t, err)

}

func TestTransactionRollbackWithError(t *testing.T){

	err :=DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(
			&User{
				ID: "13", 
				Password: "rahasia", 
				Name: Name{ 
					FirstName: "farhan", 
					MiddleName: "yudha", 
					LastName: "pratama",
					},
				},
			).Error

		if err != nil {
			return err
		}

		err = tx.Create(
			&User{
				ID: "11", 
				Password: "rahasia", 
				Name: Name{ 
					FirstName: "farhan", 
					MiddleName: "yudha", 
					LastName: "pratama",
					},
				},
			).Error

		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)

}

func TestQuerySingleObject(t *testing.T){
	user := User{}
	result := DB.First(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "1", user.ID)

	user = User{}
	result = DB.Last(&user)
	assert.Nil(t, result.Error)
	assert.Equal(t, "9", user.ID)
}

func TestQuerySingleObjectInlineCondition(t *testing.T){
	user := User{}
	result := DB.Take(&user, "id = $1", "2")
	assert.Nil(t, result.Error)
	assert.Equal(t, "2", user.ID)
}

func TestQueryAllObject(t *testing.T){
	var users []User
	result := DB.Find(&users)
	assert.Nil(t, result.Error)
	assert.Equal(t, 12, len(users))
}