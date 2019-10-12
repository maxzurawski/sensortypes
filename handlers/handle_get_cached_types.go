package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/xdevices/sensortypes/service"
	"github.com/xdevices/utilities/resterror"
)

func HandleGetCachedTypes(c echo.Context) error {

	dtos, e := service.Service.GetCachedSensorTypes()

	if e != nil {
		return c.JSON(http.StatusInternalServerError, resterror.New(e.Error()))
	}

	if len(dtos) == 0 {
		return c.JSON(http.StatusNoContent, dtos)
	}

	return c.JSON(http.StatusOK, dtos)
}
