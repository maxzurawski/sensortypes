package handlers

import (
	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/sensortypes/dto"
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
