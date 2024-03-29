package service

import (
	"testing"

	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/utilities/db"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

// create struct for getbytype suite
type SuiteGetByType struct {
	suite.Suite
}

// initialize getbytype suite
func TestSuiteGetByType(t *testing.T) {
	suite.Run(t, new(SuiteGetByType))
}

// setup tests
func (g *SuiteGetByType) SetupTest() {
	ServiceEnvironmentPreparations()
}

// failure test
func (g *SuiteGetByType) TestGetByType_Failure() {

	// Arrange & Act.

	result, err := Service.GetByType("TEMPERATURE")

	// Assert.

	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "record not found", err.Error())
	assert.Nil(g.T(), result)
}

// success test
func (g *SuiteGetByType) TestGetByType_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	sensorTypesData()

	// Act.

	result, err := Service.GetByType("TEMPERATURE")

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), "TEMPERATURE_TOPIC", result.Topic)
}
