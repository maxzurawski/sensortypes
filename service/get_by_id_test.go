package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
)

// suite struct
type GetByIdSuiteTest struct {
	suite.Suite
}

// initialize suite
func TestGetByIdSuiteTest(t *testing.T) {
	suite.Run(t, new(GetByIdSuiteTest))
}

// setup test
func (g *GetByIdSuiteTest) SetupTest() {
	environmentPreparations()
}

// failure check
func (g *GetByIdSuiteTest) TestGetById_Failure() {

	// Arrange & Act.

	result, err := Service.GetById(uint(999))

	// Assert.

	assert.NotNil(g.T(), err)
	assert.Equal(g.T(), "sensortype not found", err.Error())
	assert.Nil(g.T(), result)
}

// success check
func (g *GetByIdSuiteTest) TestGetById_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	sensorTypesData()
	sensorType, _ := dbprovider.Mgr.GetByType("MOVEMENT")

	// Act.

	result, err := Service.GetById(*sensorType.ID)

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), "Movement sensor", result.Name)
}
