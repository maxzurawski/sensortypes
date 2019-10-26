package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/utilities/db"
)

// suite struct
type DeleteSuite struct {
	suite.Suite
}

// initialize suite
func TestDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteSuite))
}

// setup test function
func (ds *DeleteSuite) SetupTest() {
	ServiceEnvironmentPreparations()
}

// failure delete -> try to delete sensor type with ID 999
func (ds *DeleteSuite) TestDelete_Failure() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	sensorTypesData()

	// Act.

	result, err := Service.Delete(uint(999))

	// Assert.

	assert.NotNil(ds.T(), err)
	assert.Equal(ds.T(), "sensortype not found", err.Error())
	assert.False(ds.T(), result)
}

// success check -> ID from MOVEMENT sensor
func (ds *DeleteSuite) TestDelete_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	sensorTypesData()
	sensorTypeToDelete, _ := Service.GetByType("MOVEMENT")
	dtos, _ := Service.GetAll()
	assert.Equal(ds.T(), 2, len(dtos))

	// Act.

	result, err := Service.Delete(sensorTypeToDelete.ID)

	// Assert.

	assert.Nil(ds.T(), err)
	assert.True(ds.T(), result)
	dtos, _ = Service.GetAll()
	assert.Equal(ds.T(), 1, len(dtos))
}
