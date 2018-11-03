package main

import (
	"WeatherApp/driver"
	readingHandler "WeatherApp/handler/http"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)


func main() {
	godotenv.Load()

	connection, err := driver.ConnectSQL (
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		)
	if err != nil {
		fmt.Printf("Couldn't connect to database: %v", err)
		os.Exit(-1)
	}

	router := chi.NewRouter()

	rh := readingHandler.NewReadingHandler(connection)

	router.Get("/readings", rh.ListReadings)

	fmt.Printf("Server listening on port %s", os.Getenv("API_PORT"))
	log.Fatal(http.ListenAndServe(":" + os.Getenv("API_PORT"), router))

}

