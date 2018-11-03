package repository

import (
	"WeatherApp/models"
	"context"
)

type ReadingRepo interface {
	ListReadings(ctx context.Context) ([]*models.Reading, error)
	AddReading(ctx context.Context, reading *models.Reading) (*models.Reading, error)
}