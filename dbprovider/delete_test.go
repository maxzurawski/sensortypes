package dbprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
)

// define suite struct for delete
type DeleteSuiteTest struct {
	suite.Suite
}

// initialize our suite object
func TestDeleteSuiteTest(t *testing.T) {
	suite.Run(t, new(DeleteSuiteTest))
}

// setup environment
func (d *DeleteSuiteTest) SetupTest() {
	EnvironmentPreparations()
}

// failure check -> try to delete not existing sensor type
func (d *DeleteSuiteTest) TestDelete_Failure() {

	// Arrange & Act.

	result, err := Mgr.Delete(uint(9876))

	// Assert.

	assert.NotNil(d.T(), err)
	assert.Equal(d.T(), "sensortype not found", err.Error())
	assert.False(d.T(), result)
}

// success check -> delete existing sensor type
func (d *DeleteSuiteTest) TestDelete_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()
	sensorType, _ := Mgr.GetByType("TEMPERATURE")
	assert.NotNil(d.T(), sensorType.ID)

	// Act.

	result, err := Mgr.Delete(*sensorType.ID)

	// Assert.

	assert.Nil(d.T(), err)
	assert.True(d.T(), result)
}
