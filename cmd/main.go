package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"time"
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
	//TODO implement
}

func checkPulse(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprint("Isn't the weather lovely, fam?")))
}

