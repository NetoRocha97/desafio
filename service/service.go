package service

import (
	// "modulo/controllers"
	"modulo/database"

	"gorm.io/gorm"
)

func CamadaService() *gorm.DB{
	db := database.GetDatabase()
	return db
}
