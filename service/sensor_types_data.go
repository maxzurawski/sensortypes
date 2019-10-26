package service

import (
	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/sensortypes/dto"
)

func sensorTypesData() []uint {
	var ids []uint

	typeDTO := dto.SensorTypeDTO{
		Type:        "TEMPERATURE",
		Name:        "Temperature sensor",
		Description: "Temperature description",
		Topic:       "TEMPERATURE_TOPIC",
	}
	sensorType, _ := dbprovider.Mgr.Save(typeDTO)
	ids = append(ids, *sensorType.ID)

	typeDTO = dto.SensorTypeDTO{
		Type:        "MOVEMENT",
		Name:        "Movement sensor",
		Description: "Movement Desc",
		Topic:       "MOVEMENT_TOPIC",
	}
	sensorType, _ = dbprovider.Mgr.Save(typeDTO)
	ids = append(ids, *sensorType.ID)
	return ids
}
