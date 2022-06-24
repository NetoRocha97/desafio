package database

import (
	"modulo/models"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.Estabelecimento{})
	db.AutoMigrate(models.Loja{})
}