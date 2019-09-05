package service

import "github.com/xdevices/sensortypes/dto"

func (s *service) GetByType(input string) (*dto.SensorTypeDTO, error) {
	sensorType, err := s.Mgr.GetByType(input)
	dto := s.Mgr.ConvertFromEntity(sensorType)
	return &dto, err
}
