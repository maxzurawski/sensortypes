package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/xdevices/sensortypes/dto"

	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"

	"github.com/xdevices/sensortypes/service"

	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleGetAllSuite struct {
	suite.Suite
}

// init suite struct
func TestHandleGetAllSuite(t *testing.T) {
	suite.Run(t, new(HandleGetAllSuite))
}

// setup test
func (hga *HandleGetAllSuite) SetupTest() {
	service.ServiceEnvironmentPreparations()
}

// checking no content
func (hga *HandleGetAllSuite) TestGetAll_NoContent() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Act.

	_ = HandleGetAll(c)

	// Assert.

	assert.Equal(hga.T(), http.StatusNoContent, rec.Code)
}

// checking content
func (hga *HandleGetAllSuite) TestGetAll_WithContent() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	prepareSensorTypes()

	// Act.

	_ = HandleGetAll(c)

	// Assert.

	assert.Equal(hga.T(), http.StatusOK, res.Code)

	var sensorTypes []dto.SensorTypeDTO
	_ = json.NewDecoder(res.Body).Decode(&sensorTypes)
	assert.Equal(hga.T(), 2, len(sensorTypes))

	sensorTypesMap := make(map[string]dto.SensorTypeDTO)

	for _, item := range sensorTypes {
		sensorTypesMap[item.Type] = item
	}

	if value, ok := sensorTypesMap["DUMMY_A"]; ok {
		assert.Equal(hga.T(), "Dummy Sensor A", value.Name)
	}

	if value, ok := sensorTypesMap["DUMMY_B"]; ok {
		assert.Equal(hga.T(), "Dummy Sensor B", value.Name)
	}
}
