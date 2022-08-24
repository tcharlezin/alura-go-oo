package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Nome      string
	Sobrenome string
}
