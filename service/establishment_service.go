package service

import (
	// "modulo/controllers"
	"modulo/database"
	"modulo/models"
	"modulo/view"

	// "github.com/labstack/echo"
	"gorm.io/gorm"
)

func CamadaService() *gorm.DB {
	db := database.GetDatabase()
	return db
}

func ShowEstablishmentService(id int) (view.ServiceEstablishment, error) {
	db := database.GetDatabase()
	var establishment models.Establishment
	err := db.First(&establishment, id).Error


	if err != nil {
		return view.ServiceEstablishment{}, err
	}

	return view.EstablishmentView(establishment), nil

}

