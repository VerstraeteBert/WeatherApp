package http

import (
	"WeatherApp/driver"
	"WeatherApp/models"
	"WeatherApp/repository"
	readingRepo "WeatherApp/repository/reading"
	"encoding/json"
	"fmt"
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
	payload, err := h.repo.ListReadings(r.Context())
	if err != nil {
		fmt.Print(err)
		respondWithError(w, http.StatusInternalServerError, "OwO we're working vewwy hawd to fix this sowwy")
	}

	respondWithJSON(w, http.StatusOK, payload)
}

func (h *ReadingHandler) AddReading(w http.ResponseWriter, r *http.Request) {
	reading := models.Reading{}

	json.NewDecoder(r.Body).Decode(&reading)

	if reading.Timestamp == "" {
		// Adds proper timestamp for MySQL (RFC 3339)
		reading.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	}

	insertedId, err := h.repo.AddReading(r.Context(), &reading)
	if err != nil {
		fmt.Print(err)
		respondWithError(w, http.StatusInternalServerError, "OwO we're working vewwy hawd to fix this sowwy")
	}

	respondWithJSON(w, http.StatusCreated, map[string]int64{"id": insertedId})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}