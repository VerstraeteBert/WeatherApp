package repository

import (
	"WeatherApp/models"
	"context"
)

type ReadingRepo interface {
	ListReadings(ctx context.Context) ([]*models.Reading, error)
}