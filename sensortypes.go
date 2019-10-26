package main

import (
	"github.com/labstack/echo"
	"github.com/maxzurawski/sensortypes/config"
	"github.com/maxzurawski/sensortypes/dbprovider"
	"github.com/maxzurawski/sensortypes/handlers"
	"github.com/maxzurawski/sensortypes/publishers"
	"github.com/maxzurawski/sensortypes/service"
)

func main() {
	e := echo.New()
	e.GET("/", handlers.HandleGetAll)
	e.GET("/:id", handlers.HandleGetById)
	e.POST("/new", handlers.HandleSave)
	e.PUT("/:id", handlers.HandleUpdate)
	e.DELETE("/:id", handlers.HandleDelete)
	e.GET("/cachetypes/", handlers.HandleGetCachedTypes)
	e.Logger.Fatal(e.Start(config.Config().Address()))
}

func init() {
	dbprovider.InitDbManager()
	service.Init()
	manager := config.EurekaManagerInit()
	manager.SendRegistrationOrFail()
	manager.ScheduleHeartBeat(config.Config().ServiceName(), 10)

	publishers.Init()
}
