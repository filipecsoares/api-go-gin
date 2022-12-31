package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
