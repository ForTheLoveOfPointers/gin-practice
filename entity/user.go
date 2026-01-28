package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `json:"id" gorm:"uniqueIndex"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
