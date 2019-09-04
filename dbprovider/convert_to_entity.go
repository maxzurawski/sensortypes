package dbprovider

import (
	"time"

	"github.com/xdevices/sensortypes/dto"
	"github.com/xdevices/sensortypes/model"
)

func (mgr *manager) ConvertToEntity(input dto.SensorTypeDTO) *model.SensorType {
	now := time.Now()
	entity := &model.SensorType{
		Version:     &input.Version,
		Topic:       &input.Topic,
		Name:        &input.Name,
		Type:        &input.Type,
		Description: &input.Description,
		CreatedAt:   &now,
	}
	return entity
}
