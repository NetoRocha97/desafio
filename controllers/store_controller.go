package controllers

import (
	"modulo/models"
	"modulo/service"
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

	store, err := service.ShowStoreService(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "Não foi possível encontrar a loja",
		})
	}

	return c.JSON(http.StatusOK, store)
}

func ShowStores(c echo.Context) error {
	stores, err := service.ShowStoresService()

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

	err = service.CreateStoreService(store)

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

	service.UpdateStoreService(updatedStore)

	optionalStore, errText := service.UpdateStoreService(updatedStore)
	
	if errText != "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": errText,
		})
	}

	return c.JSON(http.StatusOK, optionalStore)
}

func DeleteStore(c echo.Context) error {
	idText := c.Param("id")
	id, err := strconv.Atoi(idText)

	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": "O ID deve ser inteiro",
		})
	}

	errText := service.DeleteStoreService(id)

	if errText != "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"erro": errText,
		})
	}

	return c.NoContent(http.StatusNoContent)
}
