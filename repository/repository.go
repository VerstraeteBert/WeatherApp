package repository

import (
	"WeatherApp/models"
)

type ReadingRepo interface {
	ListReadings() ([]*models.Reading, error)
	AddReading(reading *models.Reading) (int64, error)
}
