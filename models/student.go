package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Id    int    `json:"id"`
	Name  string `json:"name" validate:"nonzero"`
	Email string `json:"email" validate:"min=3,max=80"`
}

func (s *Student) Validate() error {
	if err := validator.Validate(s); err != nil {
		return err
	}
	return nil
}
