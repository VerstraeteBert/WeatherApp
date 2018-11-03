package models

type Reading struct {
	ID int `json:"id"`
	Timestamp string `json:"timestamp"`
	DegreesCelcius float32 `json:"degreesCelcius"`
}
