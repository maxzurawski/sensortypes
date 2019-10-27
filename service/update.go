package service

import "github.com/maxzurawski/sensortypes/dto"

func (s *service) Update(input dto.SensorTypeDTO) (*dto.SensorTypeDTO, error) {
	sensorType, err := s.Mgr.Update(input)
	if err != nil {
		return nil, err
	}
	dto := s.Mgr.ConvertFromEntity(sensorType)
	return &dto, err
}
