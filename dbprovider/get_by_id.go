package dbprovider

import (
	"errors"

	"github.com/xdevices/sensortypes/model"
)

func (mgr *manager) GetById(id uint) (*model.SensorType, error) {
	output := &model.SensorType{}
	err := mgr.GetDb().Where("id = ?", id).Find(output).Error
	if err != nil {
		return nil, errors.New("sensortype not found")
	}
	return output, err
}
