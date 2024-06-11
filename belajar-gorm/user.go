package main

import "time"

type User struct {
	ID        string 	`gorm:"primary_key;column:id"`
	Password  string 	`gorm:"column:password"`
	Name      string 	`gorm:"column:name"`
	CreatedAt time.Time	`gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time	`gorm:"column:update_at;autoCreateTime;autoUpdateTime"`
}

// membuat nama table manual
func (u *User)	TableName() string{
	return "users"
}