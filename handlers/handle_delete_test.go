package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"

	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/sensortypes/service"
	"github.com/maxzurawski/utilities/db"

	"github.com/stretchr/testify/suite"
)

// suite struct
type DeleteSuite struct {
	suite.Suite
}

// init suite
func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}

// setup test
func (ds *DeleteSuite) SetupTest() {
	service.ServiceEnvironmentPreparations()
}

// failure check -> :id is going to be "d"
func (ds *DeleteSuite) TestDelete_Failure_on_not_parsable_id() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("d")

	// Act.

	_ = HandleDelete(c)

	// Assert.

	assert.Equal(ds.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(ds.T(), "strconv.Atoi: parsing \"d\": invalid syntax", httpError.Message)
}

// failure check -> sensor type of given :id does not exists
func (ds *DeleteSuite) TestDelete_Failure_SensorType_Not_exists() {

	// Arrange.

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues("99")

	// Act.

	_ = HandleDelete(c)

	// Assert.

	assert.Equal(ds.T(), http.StatusBadRequest, res.Code)
	var httpError echo.HTTPError
	_ = json.Unmarshal(res.Body.Bytes(), &httpError)
	assert.Equal(ds.T(), "sensortype not found", httpError.Message)
}

// successful delete
func (ds *DeleteSuite) TestDelete_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()
	dto, _ := service.Service.GetByType("DUMMY_A")

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/:id", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	c.SetPath("/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprintf("%d", dto.ID))

	// Act.

	_ = HandleDelete(c)

	// Assert.

	assert.Equal(ds.T(), http.StatusOK, res.Code)
	var result bool
	_ = json.Unmarshal(res.Body.Bytes(), &result)
	assert.True(ds.T(), result)
}
