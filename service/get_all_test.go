package service

import (
	"testing"

	"github.com/maxzurawski/sensortypes/dbprovider"

	"github.com/maxzurawski/utilities/db"
	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/suite"
)

// suites struct
type ServiceGetAllSuiteTest struct {
	suite.Suite
}

// initialize suites
func TestServiceGetAllSuiteTest(t *testing.T) {
	suite.Run(t, new(ServiceGetAllSuiteTest))
}

// setup func
func (s *ServiceGetAllSuiteTest) SetupTest() {
	ServiceEnvironmentPreparations() // DRY - DON't REPEAT YOURSELF
}

// failure test
func (s *ServiceGetAllSuiteTest) TestGetAll_Failure() {

	// Arrange & Act.

	result, err := Service.GetAll()

	// Assert.

	assert.Nil(s.T(), err)
	assert.Nil(s.T(), result)
}

// success test
func (s *ServiceGetAllSuiteTest) TestGetAll_Success() {

	// Arrange.

	cleaner := db.DeleteCreatedEntities(dbprovider.Mgr.GetDb())
	defer cleaner()
	sensorTypesData()

	// Act.

	result, err := Service.GetAll()

	// Assert.

	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), result)
	assert.Equal(s.T(), 2, len(result))
}
