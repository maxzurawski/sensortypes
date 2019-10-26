package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/sensortypes/dto"
	"github.com/maxzurawski/utilities/db"

	"github.com/labstack/echo"

	"github.com/maxzurawski/sensortypes/service"

	"github.com/stretchr/testify/suite"
)

// suite struct
type UpdateSuite struct {
	suite.Suite
}

// init suite
func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}

// setup test
func (us *UpdateSuite) SetupTest() {
	service.ServiceEnvironmentPreparations()
}

// failure check -> :id is "c"
func (us *UpdateSuite) TestUpdate_Failure_ID_not_numeric() {

	// Arrange.

	inputDTO := dto.SensorTypeDTO{
		ID:    1,
		Name:  "Test",
		Topic: "Test",
		Type:  "Test",
	}
	bytes, _ := json.Marshal(inputDTO)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:id", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("c")

	// Act.

	_ = HandleUpdate(c)

	// Assert.

	assert.Equal(us.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(us.T(), "strconv.Atoi: parsing \"c\": invalid syntax", httpError.Message)
}

// failure check -> body of the put request is simply a message "this is test update msg"
func (us *UpdateSuite) TestUpdate_Body_Failure() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:id", strings.NewReader("improper payload"))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Act.

	_ = HandleUpdate(c)

	// Assert.

	assert.Equal(us.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(us.T(), "code=400, message=Syntax error: offset=1, error=invalid character 'i' looking for beginning of value", httpError.Message)
}

// failure check -> :id = 1, sensortype with id = 10
func (us *UpdateSuite) TestUpdate_Failure_ID_not_match() {

	// Arrange.

	inputDTO := dto.SensorTypeDTO{
		ID:    10,
		Name:  "Test",
		Topic: "Test",
		Type:  "Test",
	}
	bytes, _ := json.Marshal(inputDTO)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:id", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	// Act.

	_ = HandleUpdate(c)

	// Assert.

	assert.Equal(us.T(), http.StatusConflict, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(us.T(), "given sensortype cannot be updated under id: 1", httpError.Message)
}

// failure check ->
// 1. create sensortype with type: "UPDATE_SENSOR"
// 2. get it's ID and version
// 3. try to update sensortype by changing it's topic (but with version = 10)
func (us *UpdateSuite) TestUpdate_Failure_versions_not_match() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	inputDTO := dto.SensorTypeDTO{
		Name:  "Test",
		Topic: "Test",
		Type:  "UPDATE_SENSOR",
	}

	// save it
	typeDTO, _ := service.Service.Save(inputDTO)
	assert.Equal(us.T(), "UPDATE_SENSOR", typeDTO.Type)
	assert.NotNil(us.T(), typeDTO.ID)

	typeDTO.Version = 10
	typeDTO.Topic = "Test 2"

	bytes, _ := json.Marshal(typeDTO)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:id", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", typeDTO.ID))

	// Act.

	_ = HandleUpdate(c)

	// Assert.

	assert.Equal(us.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(us.T(), "entity versions differs from given input. Please update your state of sensortype and try again", httpError.Message)
}

// success check -> update sensortype, we change it's topic -> check if version is + 1 to original
func (us *UpdateSuite) TestUpdate_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	inputDTO := dto.SensorTypeDTO{
		Name:  "Test",
		Topic: "Test",
		Type:  "UPDATE_SENSOR",
	}

	// save it
	typeDTO, _ := service.Service.Save(inputDTO)
	assert.Equal(us.T(), "UPDATE_SENSOR", typeDTO.Type)
	assert.NotNil(us.T(), typeDTO.ID)

	typeDTO.Topic = "Test 2"

	bytes, _ := json.Marshal(typeDTO)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/:id", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", typeDTO.ID))

	// Act.

	_ = HandleUpdate(c)

	// Assert.

	assert.Equal(us.T(), http.StatusAccepted, res.Code)
	var resultDTO dto.SensorTypeDTO
	_ = json.Unmarshal(res.Body.Bytes(), &resultDTO)
	assert.Equal(us.T(), typeDTO.Version+1, resultDTO.Version)
	assert.Equal(us.T(), "Test 2", resultDTO.Topic)
}
