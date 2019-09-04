package dbprovider

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xdevices/utilities/db"

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
	_ = os.Setenv("SERVICE_NAME", "sensortypes")
	_ = os.Setenv("HTTP_PORT", "8101")
	_ = os.Setenv("EUREKA_SERVICE", "http://xdevicesdev.home:8761")
	_ = os.Setenv("DB_PATH", "/Users/l0cke/.databases/xdevices/test/sensortypes.db")
	InitDbManager()
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
