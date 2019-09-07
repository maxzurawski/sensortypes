package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	dto2 "github.com/xdevices/sensortypes/dto"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/sensortypes/service"
	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleGetByIdSuite struct {
	suite.Suite
}

// init suite
func TestHandleGetByIdSuite(t *testing.T) {
	suite.Run(t, new(HandleGetByIdSuite))
}

// setup test
func (h *HandleGetByIdSuite) SetupTest() {
	service.ServiceEnvironmentPreparations()
}

// failure check -> :id = abc
func (h *HandleGetByIdSuite) TestGetById_Failure_Id_not_parsable_to_int() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("abc")

	// Act.

	_ = HandleGetById(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(h.T(), "strconv.Atoi: parsing \"abc\": invalid syntax", httpError.Message)
}

// failure check -> :id = 0
func (h *HandleGetByIdSuite) TestGetById_Failure_Id_Is_Zero() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("0")

	// Act.

	_ = HandleGetById(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(h.T(), "id cannot be zero", httpError.Message)
}

// success check -> no content (204)
func (h *HandleGetByIdSuite) TestGetById_Success_NoContent() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Act.

	_ = HandleGetById(c)

	// Assert.

	assert.Equal(h.T(), http.StatusNoContent, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(h.T(), "sensortype not found", httpError.Message)
}

// success check -> with content (200)
func (h *HandleGetByIdSuite) TestGetById_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	id := prepareSensorTypeWithId()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", id))

	// Act.

	_ = HandleGetById(c)

	// Assert.

	assert.Equal(h.T(), http.StatusOK, res.Code)
	var stResult dto2.SensorTypeDTO
	_ = json.Unmarshal(res.Body.Bytes(), &stResult)
	assert.Equal(h.T(), "Dummy Sensor A", stResult.Name)
	assert.Equal(h.T(), "DUMMY_A", stResult.Type)
	assert.Equal(h.T(), "DUMMY_A_TOPIC", stResult.Topic)
}
