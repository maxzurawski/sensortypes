package dbprovider

import (
	"testing"

	"github.com/maxzurawski/utilities/db"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

// suite struct
type GetAllSuiteTest struct {
	suite.Suite
}

// initialize suite
func TestGetAllSuteTest(t *testing.T) {
	suite.Run(t, new(GetAllSuiteTest))
}

// setup func
func (g *GetAllSuiteTest) SetupTest() {
	EnvironmentPreparations()
}

// failure check -> no sensortypes are present
func (g *GetAllSuiteTest) TestGetAll_Failure() {

	// Arrange & Act.

	result, err := Mgr.GetAll()

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), 0, len(result))
}

// success check -> use prepareSensorTypes and find all of them
func (g *GetAllSuiteTest) TestGetAll_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()

	// Act.

	result, err := Mgr.GetAll()

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), 2, len(result))
}
