package dbprovider

import "github.com/maxzurawski/sensortypes/model"

func (mgr *manager) GetByType(input string) (*model.SensorType, error) {
	tempSensor := &model.SensorType{}
	err := mgr.db.Unscoped().Where("type = ?", input).Find(tempSensor).Error
	return tempSensor, err
}
