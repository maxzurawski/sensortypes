package handlers

import (
	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/sensortypes/dto"
)

func prepareSensorTypes() {

	sensorType := dto.SensorTypeDTO{
		Name:  "Dummy Sensor A",
		Type:  "DUMMY_A",
		Topic: "DUMMY_A_TOPIC",
	}
	_, _ = dbprovider.Mgr.Save(sensorType)

	sensorType = dto.SensorTypeDTO{
		Name:  "Dummy Sensor B",
		Type:  "DUMMY_B",
		Topic: "DUMMY_B_TOPIC",
	}

	_, _ = dbprovider.Mgr.Save(sensorType)

}

func prepareSensorTypeWithId() uint {

	sensorType := dto.SensorTypeDTO{
		Name:  "Dummy Sensor A",
		Type:  "DUMMY_A",
		Topic: "DUMMY_A_TOPIC",
	}
	entity, _ := dbprovider.Mgr.Save(sensorType)
	return *entity.ID
}
