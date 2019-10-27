package service

import "github.com/maxzurawski/sensortypes/dto"

func (s *service) GetCachedSensorTypes() ([]dto.CachedTypeDTO, error) {

	types, err := s.Mgr.GetAll()

	if err != nil {
		return nil, err
	}

	var dtos []dto.CachedTypeDTO
	for _, item := range types {
		cachedType := s.Mgr.MapToCachedType(&item)
		dtos = append(dtos, cachedType)
	}

	return dtos, nil
}
