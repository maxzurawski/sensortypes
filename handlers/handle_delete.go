package handlers

import (
	"net/http"
	"strconv"

	"github.com/xdevices/sensortypes/service"

	"github.com/labstack/echo"
)

// NOTE: DELETE /:id
func HandleDelete(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	result, err := service.Service.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, result)
}
