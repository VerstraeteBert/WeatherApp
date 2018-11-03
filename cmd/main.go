package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var router *chi.Mux

func main() {
	godotenv.Load()
	apiPort := os.Getenv("API_PORT")

	router = chi.NewRouter()

	router.Get("/", checkPulse)
	router.Get("/readings", listReadings)

	log.Fatal(http.ListenAndServe(":" + apiPort, router))
}

func listReadings(w http.ResponseWriter, r *http.Request) {
	//TODO implement
}

func checkPulse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprint("Isn't the weather lovely, fam?")))
}

