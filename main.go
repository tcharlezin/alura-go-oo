package main

import (
	"go-oo/database"
	"go-oo/features"
)

func main() {
	database.InitializeDatabase()
	features.CreateContaCorrente()
}
