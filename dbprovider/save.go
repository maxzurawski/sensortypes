package dbprovider

import (
	"errors"
	"fmt"

	"github.com/maxzurawski/sensortypes/dto"
	"github.com/maxzurawski/sensortypes/model"
	"github.com/maxzurawski/utilities/stringutils"
)

func (mgr *manager) Save(sensorTypeDTO dto.SensorTypeDTO) (*model.SensorType, error) {

	// check if dto has ID (negative)
	if !stringutils.IsZero(sensorTypeDTO.ID) {
		return nil, errors.New(fmt.Sprintf("given type has already an id. cannot be saved. [id: %d]", sensorTypeDTO.ID))
	}

	// check if dto has a type (positive)
	if stringutils.IsZero(sensorTypeDTO.Type) {
		return nil, errors.New(fmt.Sprintf("type of sensor type cannot be empty"))
	}

	// check if type is not already registered
	tempSensor, err := mgr.GetByType(sensorTypeDTO.Type)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if !stringutils.IsZero(tempSensor.ID) {
		return nil, errors.New(fmt.Sprintf("there is already sensor type of the type: [%s]. id of the sensor type is: [%d]", sensorTypeDTO.Type, *tempSensor.ID))
	}

	// convert dto -> entity (dto -> model)
	entity := mgr.ConvertToEntity(sensorTypeDTO)

	// save our object
	err = mgr.db.Save(entity).Error
	return entity, err

}
