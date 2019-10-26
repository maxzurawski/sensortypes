package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/xdevices/sensortypes/service"
)

func HandleGetAll(c echo.Context) error {
	dtos, err := service.Service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if len(dtos) == 0 {
		return c.JSON(http.StatusNoContent, dtos)
	}
	return c.JSON(http.StatusOK, dtos)
}
