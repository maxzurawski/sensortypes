package service

import "github.com/xdevices/sensortypes/dbprovider"

func ServiceEnvironmentPreparations() {
	dbprovider.EnvironmentPreparations()
	Init()
}
