package database

import (
	"go-oo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var (
	DB  *gorm.DB
	err error
)

func InitializeDatabase() {
	DB, err := gorm.Open(sqlite.Open("database.db"))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados!")
	}

	DB.AutoMigrate(&models.ContaCorrente{})
	DB.AutoMigrate(&models.ContaPoupanca{})
	DB.AutoMigrate(&models.Customer{})
}
