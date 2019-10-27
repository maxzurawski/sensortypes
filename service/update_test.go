package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/utilities/db"
	"github.com/stretchr/testify/suite"
)

// suite struct
type UpdateSuite struct {
	suite.Suite
}

// initialize our suite
func TestUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateSuite))
}

// provide setup test function
func (us *UpdateSuite) SetupTest() {
	ServiceEnvironmentPreparations()
}

// failure check -> version is different then version stored in database
func (us *UpdateSuite) TestUpdate_Failure() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	sensorTypesData()
	dtoToUpdate, _ := Service.GetByType("TEMPERATURE")
	dtoToUpdate.Version = 11

	// Act.

	result, err := Service.Update(*dtoToUpdate)

	// Assert.

	assert.Nil(us.T(), result)
	assert.NotNil(us.T(), err)
	assert.Equal(us.T(), "entity versions differs from given input. Please update your state of sensortype and try again", err.Error())

}

// success check -> update sensor type
func (us *UpdateSuite) TestUpdate_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	sensorTypesData()
	dtoToUpdate, _ := Service.GetByType("TEMPERATURE")
	assert.Equal(us.T(), "Temperature sensor", dtoToUpdate.Name)

	dtoToUpdate.Name = "Dummy Temperature sensor"

	// Act.

	result, err := Service.Update(*dtoToUpdate)

	// Assert.

	assert.Nil(us.T(), err)
	assert.Equal(us.T(), "Dummy Temperature sensor", result.Name)
	assert.Equal(us.T(), dtoToUpdate.Version+1, result.Version)
}
