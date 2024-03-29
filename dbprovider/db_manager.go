package dbprovider

import (
	"fmt"

	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	"github.com/maxzurawski/sensortypes/config"
	"github.com/maxzurawski/sensortypes/dto"
	"github.com/maxzurawski/sensortypes/model"
	"github.com/maxzurawski/utilities/db"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Mgr DBManager

type DBManager interface {
	Save(sensorTypeDTO dto.SensorTypeDTO) (*model.SensorType, error)
	Update(sensorTypeDTO dto.SensorTypeDTO) (*model.SensorType, error)
	GetById(id uint) (*model.SensorType, error)
	GetByType(input string) (*model.SensorType, error)
	ConvertToEntity(input dto.SensorTypeDTO) *model.SensorType
	ConvertFromEntity(input *model.SensorType) dto.SensorTypeDTO
	Delete(id uint) (bool, error)
	GetAll() ([]model.SensorType, error)
	GetDb() *gorm.DB

	//cache
	MapToCachedType(input *model.SensorType) dto.CachedTypeDTO
}

type manager struct {
	db *gorm.DB
}

func InitDbManager() {
	dbPath := config.Config().DBPath()

	if path, exists := db.AdjustDBPath(dbPath); !exists {
		dbPath = path
	}

	db2, err := gorm.Open("sqlite3", dbPath)

	if err != nil {
		erroMsg := fmt.Sprintf("failed to init db[%s]", dbPath)
		log.Fatal(erroMsg, err)
	}

	db2.SingularTable(true)
	db2.AutoMigrate(&model.SensorType{})
	Mgr = &manager{db: db2}
}
