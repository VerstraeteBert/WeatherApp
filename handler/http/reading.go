package http

import (
	"encoding/json"
	"github.com/VerstraeteBert/WeatherApp/driver"
	"github.com/VerstraeteBert/WeatherApp/models"
	"github.com/VerstraeteBert/WeatherApp/repository"
	readingRepo "github.com/VerstraeteBert/WeatherApp/repository/reading"
	"log"
	"net/http"
	"time"
)

func NewReadingHandler(db *driver.DB) *ReadingHandler {
	return &ReadingHandler{
		repo: readingRepo.NewSQLReadingRepo(db.SQL),
	}
}

type ReadingHandler struct {
	repo repository.ReadingRepo
}

func (h *ReadingHandler) ListReadings(w http.ResponseWriter, r *http.Request) {
	payload, err := h.repo.ListReadings()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create request")
		return
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (h *ReadingHandler) AddReading(w http.ResponseWriter, r *http.Request) {
	// Using anonymous struct for validation
	type readingValidator struct {
		DegreesCelcius float32 `json:"degreesCelcius"`
	}

	rv := new(readingValidator)

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	err := d.Decode(&rv)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	// Arbitrary bounds for temperature
	if rv.DegreesCelcius > 150 || rv.DegreesCelcius < -50 {
		respondWithError(w, http.StatusBadRequest, "DegreesCelcius needs to be between -50 and 150 inclusive")
		return
	}

	reading := models.Reading{
		// Adds proper timestamp for MySQL (RFC 3339 without timezones)
		Timestamp:      time.Now().Format("2006-01-02 15:04:05"),
		DegreesCelcius: rv.DegreesCelcius,
	}

	insertedId, err := h.repo.AddReading(&reading)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]int64{"id": insertedId})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	log.Print(msg)
	respondWithJSON(w, code, map[string]string{"message": msg})
}
