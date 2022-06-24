package controllers

import (
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
		// TODO: Mensagem de ID deve ser um Integer
	}
	db := database.GetDatabase()
	var establishment models.Estabelecimento
	err = db.First(&establishment, id).Error
	
	if err != nil {
		// TODO: Mensagem de não encontrou nenhum estabelecimento
	}

	return c.String(http.StatusOK, "Bateu "+ idText)
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
