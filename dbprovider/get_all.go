package dbprovider

import "github.com/xdevices/sensortypes/model"

func (mgr *manager) GetAll() ([]model.SensorType, error) {
	sensorTypes := []model.SensorType{}
	err := mgr.GetDb().Unscoped().Find(&sensorTypes).Error
	return sensorTypes, err
}
