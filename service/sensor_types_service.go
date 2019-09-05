package service

import (
	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/sensortypes/dto"
)

type SensorTypesService interface {
	GetAll() ([]dto.SensorTypeDTO, error)
	GetById(id uint) (*dto.SensorTypeDTO, error)
	GetByType(input string) (*dto.SensorTypeDTO, error)
	// Save
	// Update
	// Delete
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
