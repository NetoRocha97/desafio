package service

import (
	"fmt"
	"modulo/database"
	"modulo/models"
	"modulo/view"

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

func ShowEstablishmentsService() ([]models.Establishment, error) {
	var establishments []models.Establishment

	db := database.GetDatabase()

	err := db.Find(&establishments).Error

	return establishments, err
}

func CreateEstablishmentService(establishment models.Establishment) error {

	db := database.GetDatabase()

	err := db.Create(&establishment).Error

	return err

}

func UpdateEstablishmentService(updatedEstablishment models.Establishment) (models.Establishment, string) {

	db := database.GetDatabase()

	var optionalEstablishment models.Establishment
	err := db.First(&optionalEstablishment, updatedEstablishment.ID).Error

	errText := ""

	if err != nil {
		errText = fmt.Sprintf("O estabelecimento com ID %d não existe", updatedEstablishment.ID)
	}

	err = db.Save(&updatedEstablishment).Error

	if err != nil {
		errText = "Não foi possível atualizar o estabelecimento"
	}

	return optionalEstablishment, errText

}

func DeleteEstablishmentService(id int) string {
	db := database.GetDatabase()

	err := db.First(&models.Establishment{}, id).Error

	if err != nil {
		return "Não existe nenhum estabelecimento com este ID"
	}

	err = db.Delete(&models.Establishment{}, id).Error

	if err != nil {
		return "Não foi possível deletar o estabelecimento"
	}

	return ""
}
