package models

type Reading struct {
	ID             int64   `json:"id"`
	Timestamp      string  `json:"timestamp"`
	DegreesCelsius float32 `json:"degreesCelsius"`
}
