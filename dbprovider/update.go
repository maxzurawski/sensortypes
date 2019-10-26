package dbprovider

import (
	"errors"
	"fmt"
	"time"

	"github.com/maxzurawski/sensortypes/dto"
	"github.com/maxzurawski/sensortypes/model"
	"github.com/maxzurawski/utilities/stringutils"
)

func (mgr *manager) Update(sensorTypeDTO dto.SensorTypeDTO) (*model.SensorType, error) {

	// if dto.ID is Zero -> then error
	if stringutils.IsZero(sensorTypeDTO.ID) {
		return nil, errors.New(fmt.Sprintf("given sensortype of type [%s] does not have valid id", sensorTypeDTO.Type))
	}

	// if Type is Zero -> then error
	if stringutils.IsZero(sensorTypeDTO.Type) {
		return nil, errors.New(fmt.Sprintf("type of sensor type cannot by empty"))
	}

	// if st with ID not exists -> then error
	output, err := mgr.GetById(sensorTypeDTO.ID)
	if err != nil {
		return nil, err
	}

	// s1.type = 'TEMPERATURE' (s1 = 0 temp sensors)
	// s2.type = 'MOVEMENT' (s2 = 2 mvmt sensors)
	// s1.type = 'MOVEMENT'
	targetType, err := mgr.GetByType(sensorTypeDTO.Type)
	inputType, err := mgr.GetById(sensorTypeDTO.ID)
	if !stringutils.IsZero(targetType.ID) && *targetType.ID != sensorTypeDTO.ID && !stringutils.IsZero(inputType.ID) {
		return nil, errors.New(fmt.Sprintf("cannot update sensortype of type [%s] into type [%s], because type [%s] is already present",
			*inputType.Type, sensorTypeDTO.Type, sensorTypeDTO.Type))
	}

	version := *output.Version
	if version != sensorTypeDTO.Version {
		return nil, errors.New("entity versions differs from given input. Please update your state of sensortype and try again")
	}

	// entity.version + 1
	version = version + 1

	// entity.modifiedAt = currentDate
	now := time.Now()

	// save entity
	output.Version = &version
	output.ModifiedAt = &now
	output.Name = &sensorTypeDTO.Name
	output.Type = &sensorTypeDTO.Type
	output.Description = &sensorTypeDTO.Description
	output.Topic = &sensorTypeDTO.Topic
	err = mgr.GetDb().Save(output).Error
	return output, err
}
