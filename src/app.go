package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//heartbeat
	r.HandleFunc("/", healthCheck).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":1337", r))
}

func healthCheck (res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode("Lovely weather init fam?")
}
