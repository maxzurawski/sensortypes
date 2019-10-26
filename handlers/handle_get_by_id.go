package handlers

import (
	"net/http"
	"strconv"

	"github.com/maxzurawski/utilities/stringutils"

	"github.com/labstack/echo"
	"github.com/maxzurawski/sensortypes/service"
)

func HandleGetById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, err.Error()))
	}
	if stringutils.IsZero(id) {
		return c.JSON(http.StatusBadRequest, echo.NewHTTPError(http.StatusBadRequest, "id cannot be zero"))
	}
	dto, err := service.Service.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusNoContent, echo.NewHTTPError(http.StatusNoContent, err.Error()))
	}
	return c.JSON(http.StatusOK, dto)
}
