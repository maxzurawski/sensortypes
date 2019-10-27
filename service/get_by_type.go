package service

import "github.com/maxzurawski/sensortypes/dto"

func (s *service) GetByType(input string) (*dto.SensorTypeDTO, error) {
	sensorType, err := s.Mgr.GetByType(input)
	if err != nil {
		return nil, err
	}
	dto := s.Mgr.ConvertFromEntity(sensorType)
	return &dto, err
}
