package main

import (
	"fmt"
	"github.com/VerstraeteBert/WeatherApp/driver"
	readingHandler "github.com/VerstraeteBert/WeatherApp/handler/http"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	connection, err := driver.ConnectSQL(
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		fmt.Printf("Couldn't connect to database: %v", err)
		os.Exit(-1)
	}

	router := chi.NewRouter()

	rh := readingHandler.NewReadingHandler(connection)

	xd := 2
	log.Println(xd)
	router.Get("/readings", rh.ListReadings)
	router.Post("/readings", rh.AddReading)

	fmt.Printf("Server listening on port %s", os.Getenv("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), router))
}
