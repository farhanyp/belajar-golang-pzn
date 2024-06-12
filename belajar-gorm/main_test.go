package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func TestQueryCondition(t *testing.T){
	var users []User
	err := DB.Where("first_name like ?", "%farhan%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestOrOperator(t *testing.T) {
	var users []User
	err := DB.Where("first_name like ?", "%farhan%").Or("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 12, len(users))
}

func TestNotOperator(t *testing.T) {
	var users []User
	err := DB.Not("first_name like ?", "%User%").Where("password = ?", "rahasia").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestSelectFields(t *testing.T) {
	var users []User
	err := DB.Select("id", "first_name").Find(&users).Error
	assert.Nil(t, err)

	for _, user := range users {
		assert.NotNil(t, user.ID)
		assert.NotEqual(t, "", user.Name.FirstName)
	}

	assert.Equal(t, 12, len(users))
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "farhan",
			LastName:  "", // tidak bisa, karena dianggap default value
		},
		Password: "rahasia",
	}

	var users []User
	err := DB.Where(userCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestMapCondition(t *testing.T) {
	mapCondition := map[string]interface{}{
		"middle_name": "yudha",
		"last_name":   "pratama",
	}

	var users []User
	err := DB.Where(mapCondition).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 3, len(users))
}

func TestOrderLimitOffset(t *testing.T) {
	var users []User
	err := DB.Order("id asc, first_name desc").Limit(5).Offset(5).Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 5, len(users))
}

type UserResponse struct{

	ID 			string
	FirstName	string
	LastName	string

}

func TestQueryNonModel(t *testing.T){
	var users []UserResponse
	err := DB.Model(&User{}).Select("id", "first_name", "first_name").Find(&users).Error
	assert.Nil(t, err)
	assert.Equal(t, 12, len(users))

}

func TestUpdate(t *testing.T) {
	user := User{}
	err := DB.Take(&user, "id = ?", "1").Error
	assert.Nil(t, err)

	user.Name.FirstName = "farhan"
	user.Name.MiddleName = ""
	user.Name.LastName = "pratama"
	user.Password = "rahasia123"

	err = DB.Save(&user).Error
	assert.Nil(t, err)
}

func TestUpdateSelectedColumns(t *testing.T) {
	err := DB.Model(&User{}).Where("id = ?", "1").Updates(map[string]interface{}{
		"middle_name": "",
		"last_name":   "Morro",
	}).Error
	assert.Nil(t, err)

	err = DB.Model(&User{}).Where("id = ?", "1").Update("password", "diubahlagi").Error
	assert.Nil(t, err)

	err = DB.Where("id = ?", "1").Updates(User{
		Name: Name{
			FirstName: "Farhan",
			LastName:  "Yp",
		},
	}).Error
	assert.Nil(t, err)
}

func TestAutoIncrement(t *testing.T){
	for i := 0; i < 10; i++ {
		userLog := UserLog{
			UserId: "1",
			Action: "Test Action",
		}

		err := DB.Create(&userLog).Error
		assert.Nil(t, err)

		assert.NotEqual(t, 0, userLog.ID)
		fmt.Println(userLog.ID)
		
	}
}

func TestDelete(t *testing.T){

	var user User
	err := DB.Take(&user, "id = $1", "12").Error
	assert.Nil(t, err)

	err = DB.Delete(&user).Error
	assert.Nil(t, err)

	err = DB.Delete(&user, "id = $1", "11").Error
	assert.Nil(t, err)

	err = DB.Where("id = $1", "10").Delete(&user).Error
	assert.Nil(t, err)

}

func TestSoftDelete(t *testing.T) {
	todo := Todo{
		UserId:      "1",
		Title:       "Todo 1",
		Description: "Description 1",
	}
	err := DB.Create(&todo).Error
	assert.Nil(t, err)

	err = DB.Delete(&todo).Error
	assert.Nil(t, err)
	assert.NotNil(t, todo.DeletedAt)

	var todos []Todo
	err = DB.Find(&todos).Error
	assert.Nil(t, err)
	assert.Equal(t, 0, len(todos))
}

func TestUnscoped(t *testing.T) {
	var todo Todo
	err := DB.Unscoped().First(&todo, "id = $1", 1).Error
	assert.Nil(t, err)
	fmt.Println(todo)

	err = DB.Unscoped().Delete(&todo).Error
	assert.Nil(t, err)

	var todos []Todo
	err = DB.Unscoped().Find(&todos).Error
	assert.Nil(t, err)
}

func TestLock(t *testing.T) {
	err := DB.Transaction(func(tx *gorm.DB) error {
		var user User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&user, "id = ?", "1").Error
		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Morro"
		err = tx.Save(&user).Error
		return err
	})
	assert.Nil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 1000000,
	}

	err := DB.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestRetrieveRelation(t *testing.T) {
	var user User
	err := DB.Model(&User{}).Preload("Wallet").Take(&user, "id = $1", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestRetrieveRelationJoin(t *testing.T) {
	var user User
	err := DB.Model(&User{}).Joins("Wallet").Take(&user, "users.id = $1", "1").Error
	assert.Nil(t, err)

	assert.Equal(t, "1", user.ID)
	assert.Equal(t, "1", user.Wallet.ID)
}

func TestUserAndAddresses(t *testing.T) {
	user := User{
		ID:       "2",
		Password: "rahasia",
		Name: Name{
			FirstName: "User 50",
		},
		Wallet: Wallet{
			ID:      "2",
			UserId:  "2",
			Balance: 1000000,
		},
		Addresses: []Address{
			{
				UserId:  "2",
				Address: "Jalan A",
			},
			{
				UserId:  "2",
				Address: "Jalan B",
			},
		},
	}

	err := DB.Save(&user).Error
	assert.Nil(t, err)
}

func TestPreloadJoinOneToMany(t *testing.T) {
	var users []User
	err := DB.Model(&User{}).Preload("Addresses").Joins("Wallet").Find(&users).Error
	assert.Nil(t, err)
}

func TestTakePreloadJoinOneToMany(t *testing.T) {
	var user User
	err := DB.Model(&User{}).Preload("Addresses").Joins("Wallet").
		Take(&user, "users.id = ?", "2").Error
	assert.Nil(t, err)
}