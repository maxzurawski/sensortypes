package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/sensortypes/dto"
	"github.com/maxzurawski/utilities/db"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"

	"github.com/maxzurawski/sensortypes/service"

	"github.com/stretchr/testify/suite"
)

// suite struct
type HandleSaveSuite struct {
	suite.Suite
}

// init suite
func TestHandleSaveSuite(t *testing.T) {
	suite.Run(t, new(HandleSaveSuite))
}

// setup test
func (h *HandleSaveSuite) SetupTest() {
	service.ServiceEnvironmentPreparations()
}

// failure test -> content is not parsable to sensortype dto
func (h *HandleSaveSuite) TestSave_Failure_Content_Not_Valid() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/new", strings.NewReader("This is dummy message"))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// Act.

	_ = HandleSave(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(h.T(), "code=400, message=Syntax error: offset=1, error=invalid character 'T' looking for beginning of value", httpError.Message)

}

// failure test -> :id of sensortype is set in body request
func (h *HandleSaveSuite) TestSave_Failure_Id_Should_Not_be_set() {

	// Arrange.

	inputDTO := dto.SensorTypeDTO{
		ID:    9,
		Name:  "Dummy A",
		Topic: "DUMMY_A",
		Type:  "DUMMY_A",
	}

	bytes, _ := json.Marshal(inputDTO)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/new", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// Act.

	_ = HandleSave(c)

	// Assert.

	assert.Equal(h.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(h.T(), "given type has already an id. cannot be saved. [id: 9]", httpError.Message)
}

// success test -> status: create(201)
func (h *HandleSaveSuite) TestSave_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	inputDTO := dto.SensorTypeDTO{
		Name:  "Success Save test sensor",
		Topic: "SUCCESS_A",
		Type:  "SUCCESS_A",
	}

	bytes, _ := json.Marshal(inputDTO)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/new", strings.NewReader(string(bytes)))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// Act.

	_ = HandleSave(c)

	// Assert.

	assert.Equal(h.T(), http.StatusCreated, res.Code)
	var resultDTO dto.SensorTypeDTO
	_ = json.Unmarshal(res.Body.Bytes(), &resultDTO)
	assert.NotNil(h.T(), resultDTO.ID)

}
