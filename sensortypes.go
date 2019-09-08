package main

import (
	"github.com/labstack/echo"
	"github.com/xdevices/sensortypes/config"
	"github.com/xdevices/sensortypes/dbprovider"
	"github.com/xdevices/sensortypes/handlers"
	"github.com/xdevices/sensortypes/service"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.HandleGetAll)
	e.GET("/:id", handlers.HandleGetById)
	e.POST("/new", handlers.HandleSave)
	e.PUT("/:id", handlers.HandleUpdate)
	e.DELETE("/:id", handlers.HandleDelete)
	e.Logger.Fatal(e.Start(config.Config().Address()))
}

func init() {
	dbprovider.InitDbManager()
	service.Init()
	manager := config.EurekaManagerInit()
	manager.SendRegistrationOrFail()
	manager.ScheduleHeartBeat(config.Config().ServiceName(), 10)
}
