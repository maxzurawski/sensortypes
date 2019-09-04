package service

import "github.com/xdevices/sensortypes/dbprovider"

func environmentPreparations() {
	dbprovider.EnvironmentPreparations()
	Init()
}
