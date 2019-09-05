package service

import "github.com/xdevices/sensortypes/dto"

func (s *service) Save(input dto.SensorTypeDTO) (*dto.SensorTypeDTO, error) {
	sensorType, err := s.Mgr.Save(input)
	if err != nil {
		return nil, err
	}
	dto := s.Mgr.ConvertFromEntity(sensorType)
	return &dto, err
}
