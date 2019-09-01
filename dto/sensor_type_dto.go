package dto

type SensorTypeDTO struct {
	ID          uint   `json:"id"`
	Version     uint   `json:"version"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
}
