package models

type Reading struct {
	ID int64 `json:"id"`
	Timestamp string `json:"timestamp"`
	DegreesCelcius float32 `json:"degreesCelcius"`
}
