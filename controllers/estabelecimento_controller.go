package controllers

import (
	"fmt"
	"modulo/database"
	"modulo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ShowEstablishment(c echo.Context) error {

	idText := c.Param("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"erro": "O ID deve ser inteiro"})
	}

	db := database.GetDatabase()
	var establishment models.Estabelecimento
	err = db.First(&establishment, id).Error

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"erro": "Não existe nenhum estabelecimento com este ID",
		})
	}

	return c.JSON(http.StatusOK, establishment)
}

func ShowEstablishments(c echo.Context) error {

	db := database.GetDatabase()

	var establishments []models.Estabelecimento

	err := db.Find(&establishments).Error

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, establishments)
}

func CreateEstablishment(c echo.Context) error {
	var establishment models.Estabelecimento
	err := c.Bind(&establishment)
	//Validar dados vazios ou nulos

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O corpo passado não é um estabelecimento",
		})
	}

	db := database.GetDatabase()
	err = db.Create(&establishment).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível criar o estabelecimento",
		})
	}

	return c.JSON(http.StatusCreated, establishment)
}

func UpdateEstablishment(c echo.Context) error {
	var updatedEstablishment models.Estabelecimento
	err := c.Bind(&updatedEstablishment)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O corpo passado não é um estabelecimento",
		})
	}

	var optionalEstablishment models.Estabelecimento
	db := database.GetDatabase()
	err = db.First(&optionalEstablishment, updatedEstablishment.ID).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": fmt.Sprintf("O estabelecimento com ID %d não existe", updatedEstablishment.ID),
		})
	}

	err = db.Save(&updatedEstablishment).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível atualizar o estabelecimento",
		})
	}

	return c.JSON(http.StatusOK, updatedEstablishment)
}

func DeleteEstablishment(c echo.Context) error {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O ID deve ser inteiro",
		})
	}

	db := database.GetDatabase()
	err = db.First(&models.Estabelecimento{}, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não existe nenhum estabelecimento com este ID",
		})
	}

	err = db.Delete(&models.Estabelecimento{}, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível deletar o estabelecimento",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
