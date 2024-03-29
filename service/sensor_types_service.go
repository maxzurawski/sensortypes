package service

import (
	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/sensortypes/dto"
)

type SensorTypesService interface {
	GetAll() ([]dto.SensorTypeDTO, error)
	GetById(id uint) (*dto.SensorTypeDTO, error)
	GetByType(input string) (*dto.SensorTypeDTO, error)
	Save(input dto.SensorTypeDTO) (*dto.SensorTypeDTO, error)
	Update(input dto.SensorTypeDTO) (*dto.SensorTypeDTO, error)
	Delete(id uint) (bool, error)

	// cache
	GetCachedSensorTypes() ([]dto.CachedTypeDTO, error)
}

var Service SensorTypesService

type service struct {
	Mgr dbprovider.DBManager
}

func Init() {
	s := service{}
	s.Mgr = dbprovider.Mgr
	Service = &s
}
