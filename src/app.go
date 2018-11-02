package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load("../.env")

	r := mux.NewRouter()

	//heartbeat
	r.HandleFunc("/", healthCheck).Methods(http.MethodGet)

	apiPort := os.Getenv("API_PORT")

	log.Fatal(http.ListenAndServe(apiPort, r))
}

func healthCheck (res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("Lovely weather init fam?")
}
