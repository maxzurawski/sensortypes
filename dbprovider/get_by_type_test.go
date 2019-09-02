package dbprovider

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
)

// get by type suite
type GetByTypeSuiteTest struct {
	suite.Suite
}

// prepare setup func (aka before func) -> just like previously in get by id test challenge
func (g *GetByTypeSuiteTest) SetupTest() {
	_ = os.Setenv("SERVICE_NAME", "sensortypes")
	_ = os.Setenv("HTTP_PORT", "8101")
	_ = os.Setenv("EUREKA_SERVICE", "http://xdevicesdev.home:8761")
	_ = os.Setenv("DB_PATH", "/Users/l0cke/.databases/xdevices/test/sensortypes.db")
	InitDbManager()
}

// prepare 2 tests: success
func (g *GetByTypeSuiteTest) TestGetByType_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()

	// Act.

	sensorType, err := Mgr.GetByType("TEMPERATURE")

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), sensorType.ID)
	assert.Equal(g.T(), "temperature", *sensorType.Name)
	assert.Equal(g.T(), "Simple temperature sensor type", *sensorType.Description)
	assert.Equal(g.T(), "temperature_topic", *sensorType.Topic)

}

// & failure -> use previously created test data
func (g *GetByTypeSuiteTest) TestGetByType_Failure() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()

	// Act.

	sensorType, err := Mgr.GetByType("HUMIDITY")

	// Assert.

	assert.Nil(g.T(), sensorType.ID)
	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "record not found", err.Error())
}

func TestGetByTypeSuiteTest(t *testing.T) {
	suite.Run(t, new(GetByTypeSuiteTest))
}
