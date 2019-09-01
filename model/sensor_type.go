package model

import "time"

type SensorType struct {
	ID          *uint      `gorm:"primary_key"`
	Version     *uint      `gorm:"column:version"`
	Name        *string    `gorm:"column:name"`
	Type        *string    `gorm:"column:type;unique;type:varchar(255)"`
	Topic       *string    `gorm:"column:topic;unique;type:varchar(255)"`
	Description *string    `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	ModifiedAt  *time.Time `gorm:"column:modified_at"`
}
