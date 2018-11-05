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
		log.Print(err)
		respondWithError(w, http.StatusInternalServerError, "OwO we're working vewwy hawd to fix this sowwy")
		return
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (h *ReadingHandler) AddReading(w http.ResponseWriter, r *http.Request) {
	//TODO Validation?
	reading := models.Reading{}

	json.NewDecoder(r.Body).Decode(&reading)

	if reading.Timestamp == "" {
		// Adds proper timestamp for MySQL (RFC 3339)
		reading.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	}

	insertedId, err := h.repo.AddReading(&reading)
	if err != nil {
		log.Print(err)
		respondWithError(w, http.StatusInternalServerError, "OwO we're working vewwy hawd to fix this sowwy")
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
	respondWithJSON(w, code, map[string]string{"message": msg})
}
