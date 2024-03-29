package config

import (
	"fmt"
	"os"

	"github.com/maxzurawski/utilities/config"
	"github.com/maxzurawski/utilities/rabbit"
)

type sensorTypesConfig struct {
	config.Manager
	dbPath string
	rabbit.RabbitMQManager
}

var instance *sensorTypesConfig

func Config() *sensorTypesConfig {
	if instance == nil {
		instance = new(sensorTypesConfig)
		instance.Init()
		instance.sensorTypesConfigInit()
	}
	return instance
}

func (c *sensorTypesConfig) sensorTypesConfigInit() {
	if dbPath, err := os.LookupEnv("DB_PATH"); !err {
		panic(fmt.Sprintf("set DB_PATH and try again"))
	} else {
		c.dbPath = dbPath
	}

	if c.ConnectToRabbit() {
		c.RabbitMQManager.InitConnection(c.RabbitURL())
	}
}

func (c *sensorTypesConfig) DBPath() string {
	return c.dbPath
}
