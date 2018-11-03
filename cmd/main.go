package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var router *chi.Mux
var db *sql.DB

func init() {
	godotenv.Load()

	router = chi.NewRouter()

	dbSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	db, err = sql.Open("mysql", dbSource)
	if err != nil {
		fmt.Printf("Couldn't establish database connection: %s", err)
		os.Exit(-1)
	}
}

func main() {
	router.Get("/", checkPulse)
	router.Get("/readings", listReadings)

	apiPort := os.Getenv("API_PORT")

	log.Fatal(http.ListenAndServe(":" + apiPort, router))
}

func listReadings(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("select * from readings")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	payload := make([]*Reading, 0)
	for rows.Next() {
		data := new(Reading)
		err := rows.Scan(
			&data.ID,
			&data.Timestamp,
			&data.DegreesCelcius,
		)
		if err != nil {
			panic(err)
		}
		payload = append(payload, data)
	}
	respondwithJSON(w, 200, payload)
}

func checkPulse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprint("Isn't the weather lovely, fam?")))
}

type Reading struct {
	ID int `json:"id"`
	Timestamp string `json:"timestamp"`
	DegreesCelcius float32 `json:"degreesCelcius"`
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
    response, _ := json.Marshal(payload)
    fmt.Println(payload)
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(code)
    w.Write(response)
}