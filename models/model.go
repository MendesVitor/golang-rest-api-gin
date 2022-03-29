package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name"`
	Cpf  string `json:"cpf"`
	Rg   string `json:"Rg"`
}
