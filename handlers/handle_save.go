package handlers

import (
	"net/http"

	"github.com/maxzurawski/sensortypes/publishers"

	"github.com/maxzurawski/sensortypes/service"

	"github.com/labstack/echo"
	"github.com/maxzurawski/sensortypes/dto"
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
	publishers.TypesConfigChangePublisher().PublishSaveChange("", afterSaveDTO)

	return c.JSON(http.StatusCreated, afterSaveDTO)
}
