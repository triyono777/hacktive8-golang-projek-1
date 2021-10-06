package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Todo struct {
	GormModel
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func (todo *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(todo)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}

func (todo *Todo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(todo)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
