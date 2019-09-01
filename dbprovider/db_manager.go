package dbprovider

import (
	"fmt"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/jinzhu/gorm"
	"github.com/xdevices/sensortypes/config"
	"github.com/xdevices/sensortypes/dto"
	"github.com/xdevices/sensortypes/model"
	"github.com/xdevices/utilities/db"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var Mgr DBManager

type DBManager interface {
	Save(sensorTypeDTO dto.SensorTypeDTO) (*model.SensorType, error)
	Update(sensorTypeDTO dto.SensorTypeDTO) (*model.SensorType, error)
	GetById(id uint) (*model.SensorType, error)
	GetByType(input string) (*model.SensorType, error)
	//Delete(id uint) (bool, error)
	//GetAll() ([]model.SensorType, error)
	//GetById(id uint) (*model.SensorType, error)
	GetDb() *gorm.DB
}

type manager struct {
	db *gorm.DB
}

func convertToEntity(input dto.SensorTypeDTO) *model.SensorType {
	now := time.Now()
	entity := &model.SensorType{
		Version:     &input.Version,
		Topic:       &input.Topic,
		Name:        &input.Name,
		Type:        &input.Type,
		Description: &input.Description,
		CreatedAt:   &now,
	}
	return entity
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
