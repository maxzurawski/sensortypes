package dbprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maxzurawski/utilities/db"
	"github.com/stretchr/testify/suite"
)

// define struct for tests
type GetByTypeSuiteTest struct {
	suite.Suite
}

// func to initialize TestSuiteContext
func TestGetByTypeSuiteTest(t *testing.T) {
	suite.Run(t, new(GetByTypeSuiteTest))
}

// func to setup tests
func (g *GetByTypeSuiteTest) SetupTest() {
	EnvironmentPreparations()
}

// unit test for success
func (g *GetByTypeSuiteTest) TestGetByType_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()

	// Act.

	result, err := Mgr.GetByType("TEMPERATURE")

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result.ID)
	assert.Equal(g.T(), "temperature", *result.Name)
	assert.Equal(g.T(), "Simple temperature sensor type", *result.Description)
}

// unit test for failure
func (g *GetByTypeSuiteTest) TestGetByType_Failure() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()

	// Act.

	result, err := Mgr.GetByType("SMOKE_SENSOR")

	// Assert.

	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "record not found", err.Error())
	assert.Nil(g.T(), result.ID)
}
