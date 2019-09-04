package dbprovider

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/xdevices/sensortypes/dto"
	"github.com/xdevices/utilities/db"
)

// struct for our suite
type UpdateSuiteTest struct {
	suite.Suite
}

// initialize our suite
func TestUpdateSuiteTest(t *testing.T) {
	suite.Run(t, new(UpdateSuiteTest))
}

// setup function
func (u *UpdateSuiteTest) SetupTest() {
	environmentPreparations()
}

// failure test -> check if ID is not 0
func (u *UpdateSuiteTest) TestUpdate_ID_zero_Failure() {

	// Arrange.
	inputDTO := dto.SensorTypeDTO{
		Name:    "Temperature",
		Version: 1,
		Type:    "TEMPERATURE",
	}

	// Act.

	result, err := Mgr.Update(inputDTO)

	// Assert.

	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), "given sensortype of type [TEMPERATURE] does not have valid id", err.Error())
	assert.Nil(u.T(), result)
}

// failure test -> check if type is not empty string
func (u *UpdateSuiteTest) TestUpdate_Type_empty_Failure() {

	// Arrange.

	inputDTO := dto.SensorTypeDTO{
		Name:    "Temperature",
		Version: 1,
		ID:      1,
	}

	// Act.

	result, err := Mgr.Update(inputDTO)

	// Assert.

	assert.NotNil(u.T(), err)
	assert.Nil(u.T(), result)
	assert.Equal(u.T(), "type of sensor type cannot by empty", err.Error())
}

// failure test -> check if sensor type exists
func (u *UpdateSuiteTest) TestUpdate_SensorType_Not_Exist_Failure() {

	// Arrange.

	inputDTO := dto.SensorTypeDTO{
		Name:    "Temperature",
		Version: 1,
		ID:      1,
		Type:    "TEMPERATURE",
	}

	// Act.

	result, err := Mgr.Update(inputDTO)

	// Assert.

	assert.Nil(u.T(), result)
	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), "sensortype not found", err.Error())
}

// failure test -> check if renaming of type is possible, when type of rename is existing
func (u *UpdateSuiteTest) TestUpdate_Renaming_NotPossible_Failure() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()

	prepareSensorTypes()
	sensorTypeToUpdate, _ := Mgr.GetByType("TEMPERATURE")
	assert.NotNil(u.T(), sensorTypeToUpdate.ID)
	*sensorTypeToUpdate.Type = "MOVEMENT"
	dto := Mgr.ConvertFromEntity(sensorTypeToUpdate)
	assert.Equal(u.T(), "MOVEMENT", dto.Type)

	// Act.

	result, err := Mgr.Update(dto)

	// Assert.

	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), "cannot update sensortype of type [TEMPERATURE] into type [MOVEMENT], because type [MOVEMENT] is already present", err.Error())
	assert.Nil(u.T(), result)
}

// failure test -> optimistic lock check (versions check)
func (u *UpdateSuiteTest) TestUpdate_OptimisticLock_Failure() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()
	sensorTypeToUpdate, _ := Mgr.GetByType("TEMPERATURE")
	*sensorTypeToUpdate.Version = 2
	dto := Mgr.ConvertFromEntity(sensorTypeToUpdate)

	// Act.

	result, err := Mgr.Update(dto)

	// Assert.

	assert.NotNil(u.T(), err)
	assert.Equal(u.T(), "entity versions differs from given input. Please update your state of sensortype and try again", err.Error())
	assert.Nil(u.T(), result)
}

// success test -> update sensorTypeDTO.
func (u *UpdateSuiteTest) TestUpdate_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	prepareSensorTypes()
	sensorTypeToUpdate, _ := Mgr.GetByType("TEMPERATURE")
	*sensorTypeToUpdate.Name = "Dummy Test Update"
	dto := Mgr.ConvertFromEntity(sensorTypeToUpdate)
	oldVersion := dto.Version

	// Act.

	result, err := Mgr.Update(dto)

	// Assert.

	assert.Nil(u.T(), err)
	assert.NotNil(u.T(), result)
	assert.Equal(u.T(), "Dummy Test Update", *result.Name)
	assert.Equal(u.T(), oldVersion+1, *result.Version)
}
