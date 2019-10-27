package service

import "github.com/maxzurawski/sensortypes/dto"

func (s *service) GetAll() ([]dto.SensorTypeDTO, error) {
	types, err := s.Mgr.GetAll()
	var result []dto.SensorTypeDTO
	for _, entity := range types {
		result = append(result, s.Mgr.ConvertFromEntity(&entity))
	}
	return result, err
}
