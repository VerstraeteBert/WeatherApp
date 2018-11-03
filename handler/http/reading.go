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
		//TODO responds with err
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