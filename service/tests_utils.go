package service

import "github.com/maxzurawski/sensortypes/dbprovider"

func ServiceEnvironmentPreparations() {
	dbprovider.EnvironmentPreparations()
	Init()
}
