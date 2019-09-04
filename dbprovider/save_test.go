package dbprovider

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xdevices/sensortypes/dto"
	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
)

type SaveSuiteTest struct {
	suite.Suite
}

func (s *SaveSuiteTest) SetupTest() {
	environmentPreparations()
}

func (s *SaveSuiteTest) TestSave_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()

	typeDTO := dto.SensorTypeDTO{
		Type:        "TEMPERATURE",
		Name:        "Temperature sensor",
		Description: "Temperature Desc",
		Topic:       "temperature_topic",
	}

	// Act.

	result, err := Mgr.Save(typeDTO)

	// Assert.

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), result.ID)
}

// Test failure on ID in sensortype present

func (s *SaveSuiteTest) TestSave_Failure_Id() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()

	typeDTO := dto.SensorTypeDTO{
		ID:          1,
		Type:        "TEMPERATURE",
		Name:        "Temperature sensor",
		Description: "Temperature Desc",
		Topic:       "temperature_topic",
	}

	// Act.

	_, err := Mgr.Save(typeDTO)

	// Assert.

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), "given type has already an id. cannot be saved. [id: 1]", err.Error())
}

// Test failure on type attribute is empty
func (s *SaveSuiteTest) TestSave_Failure_Type_Empty() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()

	typeDTO := dto.SensorTypeDTO{
		Name:        "Temperature sensor",
		Description: "Temperature Desc",
		Topic:       "temperature_topic",
	}

	// Act.

	_, err := Mgr.Save(typeDTO)

	// Assert.

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), "type of sensor type cannot be empty", err.Error())
}

// Test failure where type is already existing
func (s *SaveSuiteTest) TestSave_Failure_Type_Exists() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	id := prepareTemperatureType()
	typeDTO := dto.SensorTypeDTO{
		Type:        "TEMPERATURE",
		Name:        "Temperature sensor",
		Description: "Temperature Desc",
		Topic:       "temperature_topic",
	}

	// Act.

	_, err := Mgr.Save(typeDTO)

	// Assert.

	assert.NotNil(s.T(), err)
	assert.Equal(s.T(), fmt.Sprintf("there is already sensor type of the type: [TEMPERATURE]. id of the sensor type is: [%d]", *id), err.Error())
}

func prepareTemperatureType() *uint {
	typeDTO := dto.SensorTypeDTO{
		Type: "TEMPERATURE",
		Name: "Temperature sensor",
	}

	sensorType, _ := Mgr.Save(typeDTO)
	return sensorType.ID
}

func TestSaveSuiteTest(t *testing.T) {
	suite.Run(t, new(SaveSuiteTest))
}
