package service

import "github.com/maxzurawski/sensortypes/dto"

func (s *service) GetById(id uint) (*dto.SensorTypeDTO, error) {
	sensorType, err := s.Mgr.GetById(id)
	if err != nil {
		return nil, err
	}
	dto := s.Mgr.ConvertFromEntity(sensorType)
	return &dto, err
}
