package http

import (
	"encoding/json"
	"github.com/VerstraeteBert/WeatherApp/driver"
	"github.com/VerstraeteBert/WeatherApp/models"
	"github.com/VerstraeteBert/WeatherApp/repository"
	readingRepo "github.com/VerstraeteBert/WeatherApp/repository/reading"
	"github.com/eclipse/paho.mqtt.golang"
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

func (h *ReadingHandler) AddReading(client mqtt.Client, message mqtt.Message) {
	// Using anonymous struct for validation
	type readingMessageValidator struct {
		ClientId string `json:"clientId"`
		Reading  struct {
			DegreesCelsius float32 `json:"degreesCelsius"`
		} `json:"reading"`
	}

	rmv := new(readingMessageValidator)
	err := json.Unmarshal(message.Payload(), rmv)
	if err != nil {
		log.Printf("Failed to unmarshal reading message: %v", err)
		return
	}

	log.Print(rmv.Reading.DegreesCelsius)

	// Arbitrary bounds for temperature
	if rmv.Reading.DegreesCelsius > 150 || rmv.Reading.DegreesCelsius < -50 {
		log.Print("Degrees Celcius should be between -50 and 150 inclusive")
		return
	}

	reading := models.Reading{
		// Adds proper timestamp for MySQL (RFC 3339 without timezones)
		Timestamp:      time.Now().Format("2006-01-02 15:04:05"),
		DegreesCelsius: rmv.Reading.DegreesCelsius,
	}

	_, err = h.repo.AddReading(&reading)
	if err != nil {
		log.Printf("Failed to save reading to database: %v", err)
		return
	}
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
