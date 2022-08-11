package controllers

import (
	"fmt"
	"modulo/models"
	"modulo/service"
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

	establishment, err := service.ShowEstablishmentService(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"erro": "Não existe nenhum estabelecimento com este ID",
		})
	}

	return c.JSON(http.StatusOK, establishment)
}

func ShowEstablishments(c echo.Context) error {

	service := service.CamadaService()

	var establishments []models.Establishment

	err := service.Find(&establishments).Error

	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, establishments)
}

func CreateEstablishment(c echo.Context) error {
	var establishment models.Establishment
	err := c.Bind(&establishment)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O corpo passado não é um estabelecimento",
		})
	}

	service := service.CamadaService()
	err = service.Create(&establishment).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível criar o estabelecimento",
		})
	}

	return c.JSON(http.StatusCreated, establishment)
}

func UpdateEstablishment(c echo.Context) error {
	var updatedEstablishment models.Establishment
	err := c.Bind(&updatedEstablishment)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O corpo passado não é um estabelecimento",
		})
	}

	var optionalEstablishment models.Establishment
	service := service.CamadaService()
	err = service.First(&optionalEstablishment, updatedEstablishment.ID).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": fmt.Sprintf("O estabelecimento com ID %d não existe", updatedEstablishment.ID),
		})
	}

	err = service.Save(&updatedEstablishment).Error

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

	service := service.CamadaService()
	err = service.First(&models.Establishment{}, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não existe nenhum estabelecimento com este ID",
		})
	}

	err = service.Delete(&models.Establishment{}, id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível deletar o estabelecimento",
		})
	}

	return c.NoContent(http.StatusNoContent)
}

func ShowStoresByEstablishment(c echo.Context) error {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O ID deve ser inteiro",
		})
	}

	service := service.CamadaService()
	var stores []models.Store

	err = service.Find(&stores).Where("establishment_id = ?", id).Error

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível acessar o banco de dados",
		})
	}

	return c.JSON(http.StatusOK, stores)
}
