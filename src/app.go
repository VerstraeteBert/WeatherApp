package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	//heartbeat
	r.HandleFunc("/", handleHeartbeat).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":1337", r))
}

func handleHeartbeat (res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Lovely weather init fam?")
}
