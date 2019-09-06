package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/sensortypes/dto"
	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
)

// suite struct
type SaveSuite struct {
	suite.Suite
}

// initialize suite
func TestSaveSuite(t *testing.T) {
	suite.Run(t, new(SaveSuite))
}

// setup tests
func (ss *SaveSuite) SetupTest() {
	ServiceEnvironmentPreparations()
}

// failure check -> input.ID is 9
func (ss *SaveSuite) TestSave_Failure() {

	// Arrange.

	input := dto.SensorTypeDTO{
		ID:    9,
		Name:  "Dummy",
		Type:  "DUMMY",
		Topic: "DUMMY_TOPIC",
	}

	// Act.

	result, err := Service.Save(input)

	// Assert.

	assert.NotNil(ss.T(), err)
	assert.Equal(ss.T(), "given type has already an id. cannot be saved. [id: 9]", err.Error())
	assert.Nil(ss.T(), result)
}

// success check -> ID, Version, Name
func (ss *SaveSuite) TestSave_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	input := dto.SensorTypeDTO{
		Name:  "Dummy",
		Type:  "DUMMY",
		Topic: "DUMMY_TOPIC",
	}

	// Act.

	result, err := Service.Save(input)

	// Assert.

	assert.Nil(ss.T(), err)
	assert.NotNil(ss.T(), result)
	assert.NotNil(ss.T(), result.ID)
	assert.NotNil(ss.T(), result.Version)
}
