package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/xdevices/sensortypes/publishers"

	"github.com/xdevices/sensortypes/service"

	"github.com/xdevices/sensortypes/dto"

	"github.com/labstack/echo"
)

// PUT /:id body: dto.SensorTypeDTO
func HandleUpdate(c echo.Context) error {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	var sensorTypeDTO dto.SensorTypeDTO
	if err = c.Bind(&sensorTypeDTO); err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}

	if sensorTypeDTO.ID != uint(id) {
		return c.JSON(http.StatusConflict, echo.NewHTTPError(http.StatusConflict, fmt.Sprintf("given sensortype cannot be updated under id: %d", id)))
	}

	oldType, _ := service.Service.GetByType(sensorTypeDTO.Type)
	resultDTO, err := service.Service.Update(sensorTypeDTO)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	publishers.TypesConfigChangePublisher().PublishUpdateChange(oldType, resultDTO)

	return c.JSON(http.StatusAccepted, resultDTO)
}
