package dbprovider

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xdevices/utilities/db"

	"github.com/stretchr/testify/suite"
	"github.com/xdevices/sensortypes/dto"
)

type GetByIdTestSuite struct {
	suite.Suite
}

func TestGetByIdSuiteTest(t *testing.T) {
	suite.Run(t, new(GetByIdTestSuite))
}

func (s *GetByIdTestSuite) SetupTest() {
	_ = os.Setenv("SERVICE_NAME", "sensortypes")
	_ = os.Setenv("HTTP_PORT", "8101")
	_ = os.Setenv("EUREKA_SERVICE", "http://xdevicesdev.home:8761")
	_ = os.Setenv("DB_PATH", "/Users/l0cke/.databases/xdevices/test/sensortypes.db")
	InitDbManager()
}

// test get by id with success
func (g *GetByIdTestSuite) TestGetById_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()
	ids := prepareSensorTypes()

	// Act.

	result, err := Mgr.GetById(ids[0])

	// Assert.

	assert.Nil(g.T(), err)
	assert.NotNil(g.T(), result)
	assert.Equal(g.T(), *result.ID, ids[0])
}

// test get by id failure
func (g *GetByIdTestSuite) TestGetById_Failure() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(Mgr.GetDb())
	defer cleaner()

	// Act.

	result, err := Mgr.GetById(uint(0))

	// Assert.

	assert.NotNil(g.T(), err)
	assert.Nil(g.T(), result)
	assert.Equal(g.T(), "sensortype not found", err.Error())
}

func prepareSensorTypes() []uint {

	var ids []uint

	typeDTO := dto.SensorTypeDTO{
		Type:        "TEMPERATURE",
		Name:        "temperature",
		Topic:       "temperature_topic",
		Description: "Simple temperature sensor type",
	}

	sensorType, _ := Mgr.Save(typeDTO)
	ids = append(ids, *sensorType.ID)

	typeDTO = dto.SensorTypeDTO{
		Type:  "MOVEMENT",
		Name:  "Movement sensor",
		Topic: "Movement_topic",
	}

	sensorType, _ = Mgr.Save(typeDTO)
	ids = append(ids, *sensorType.ID)
	return ids

}
