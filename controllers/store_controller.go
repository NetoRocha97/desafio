package controllers

import (
	"fmt"
	"modulo/database"
	"modulo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ShowStore(c echo.Context) error {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O Id deve ser inteiro",
		})
	}

	db := database.GetDatabase()
	var store models.Store
	err = db.First(&store, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível encontrar a loja",
		})
	}

	return c.JSON(http.StatusOK, store)
}

func ShowStores(c echo.Context) error {
	var stores []models.Store
	db := database.GetDatabase()
	err := db.Find(&stores).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível listar as lojas: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, stores)
}

func CreateStore(c echo.Context) error {
	var store models.Store
	err := c.Bind(&store)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O corpo passado não é uma loja",
		})
	}

	db := database.GetDatabase()
	err = db.Create(&store).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível criar a loja: " + err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, store)
}

func UpdateStore(c echo.Context) error {
	var updatedStore models.Store
	err := c.Bind(&updatedStore)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O corpo passado não é uma loja",
		})
	}

	var optionalStore models.Store
	db := database.GetDatabase()
	err = db.First(&optionalStore, updatedStore.ID).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": fmt.Sprintf("A loja com ID %d não existe", updatedStore.ID),
		})
	}

	err = db.Save(&updatedStore).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível atualizar o estabelecimento",
		})
	}

	return c.JSON(http.StatusOK, updatedStore)
}

func DeleteStore(c echo.Context) error {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O ID deve ser inteiro",
		})
	}

	db := database.GetDatabase()
	err = db.First(&models.Store{}, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não existe nenhuma loja com este ID",
		})
	}

	err = db.Delete(&models.Store{}, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível deleter a loja",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
