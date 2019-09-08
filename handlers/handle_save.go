package handlers

import (
	"net/http"

	"github.com/xdevices/sensortypes/service"

	"github.com/labstack/echo"
	"github.com/xdevices/sensortypes/dto"
)

func HandleSave(c echo.Context) error {
	var sensorTypeDTO dto.SensorTypeDTO
	if err := c.Bind(&sensorTypeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	afterSaveDTO, err := service.Service.Save(sensorTypeDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	return c.JSON(http.StatusCreated, afterSaveDTO)
}
