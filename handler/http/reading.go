package http

import (
	"WeatherApp/driver"
	"WeatherApp/repository"
	readingRepo "WeatherApp/repository/reading"
	"encoding/json"
	"fmt"
	"net/http"
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