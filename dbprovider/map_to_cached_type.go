package dbprovider

import (
	"github.com/maxzurawski/sensortypes/dto"
	"github.com/maxzurawski/sensortypes/model"
)

func (mgr *manager) MapToCachedType(input *model.SensorType) dto.CachedTypeDTO {

	if input == nil {
		return dto.CachedTypeDTO{}
	}

	return dto.CachedTypeDTO{
		Type:  *input.Type,
		Topic: *input.Topic,
	}
}
