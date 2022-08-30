package service

import (
	"fmt"
	"modulo/database"
	"modulo/models"
	"modulo/view"
)

func ShowStoreService(id int) (view.ServiceStore, error) {
	db := database.GetDatabase()
	var store models.Store
	err := db.First(&store, id).Error

	if err != nil {
		return view.ServiceStore{}, err
	}

	return view.StoreView(store), nil

}

func ShowStoresService() ([]models.Store, error) {
	var stores []models.Store

	db := database.GetDatabase()

	err := db.Find(&stores).Error

	return stores, err
}

func CreateStoreService(store models.Store) error {

	db := database.GetDatabase()

	err := db.Create(&store).Error

	return err

}

func UpdateStoreService(updatedStore models.Store) (models.Store, string) {

	db := database.GetDatabase()

	var optionalStore models.Store
	err := db.First(&optionalStore, updatedStore.ID).Error

	errText := ""

	if err != nil {
		errText = fmt.Sprintf("A loja com ID %d não existe", updatedStore.ID)
	}

	err = db.Save(&updatedStore).Error

	if err != nil {
		errText = "Não foi possível atualizar a loja"
	}

	return optionalStore, errText

}

func DeleteStoreService(id int) string {
	db := database.GetDatabase()

	err := db.First(&models.Store{}, id).Error

	if err != nil {
		return "Não existe nenhuma loja com este ID"
	}

	err = db.Delete(&models.Store{}, id).Error

	if err != nil {
		return "Não foi possível deletar a loja"
	}

	return ""
}