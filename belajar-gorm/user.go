package main

import (
	"time"

	"gorm.io/gorm"
)

type Name struct{

	FirstName	string	`gorm:"column:first_name"`
	MiddleName	string	`gorm:"column:middle_name"`
	LastName	string	`gorm:"column:last_name"`

}

type User struct {
	ID        string 	`gorm:"primary_key;column:id"`
	Password  string 	`gorm:"column:password"`
	Name      Name 		`gorm:"embedded"`
	CreatedAt time.Time	`gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time	`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	Wallet		Wallet	`gorm:"foreignKey:user_id;references:id"`
}

// membuat nama table manual
func (u *User)	TableName() string{
	return "users"
}

type UserLog struct{
	ID			int		`gorm:"primary_key;column:id;autoIncrement"`
	UserId  	string 	`gorm:"column:user_id"`
	Action  	string 	`gorm:"column:action"`
	CreatedAt time.Time	`gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time	`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (u *UserLog)	TableName() string{
	return "user_logs"
}

type Todo struct{
	ID			int				`gorm:"primary_key;column:id;autoIncrement"`
	UserId  	string 			`gorm:"column:user_id"`
	Title  		string 			`gorm:"column:title"`
	Description string 			`gorm:"column:description"`
	CreatedAt time.Time			`gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time			`gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt	`gorm:"column:deleted_at"`
}

func (u *Todo)	TableName() string{
	return "todos"
}