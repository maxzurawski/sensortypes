package dbprovider

import (
	"github.com/maxzurawski/sensortypes/dto"
	"github.com/maxzurawski/sensortypes/model"
)

func (mgr *manager) ConvertFromEntity(input *model.SensorType) dto.SensorTypeDTO {
	result := dto.SensorTypeDTO{}
	result.ID = *input.ID
	result.Name = *input.Name
	result.Type = *input.Type
	result.Topic = *input.Topic
	result.Version = *input.Version
	result.Description = *input.Description
	return result
}
