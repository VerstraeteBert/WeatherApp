package main

import (
	"github.com/VerstraeteBert/WeatherApp/driver"
	readingHandler "github.com/VerstraeteBert/WeatherApp/handler/http"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	godotenv.Load()

	// Connect to MySQL DB
	connection, err := driver.ConnectSQL(
		os.Getenv("DATABASE_URL"),
	)
	if err != nil {
		log.Fatalf("Couldn't connect to database: %v", err)
	}

	uri, err := url.Parse(os.Getenv("MQTT_SERVER"))
	if err != nil {
		log.Fatalf("Couldn't fetch CloudMQTT URL from env: %v", err)
	}

	// Setting up MQTT Client
	mqttClient, err := driver.ConnectMQTT(
		os.Getenv("MQTT_CLIENT_ID"),
		uri,
	)

	rh := readingHandler.NewReadingHandler(connection)

	// HTTP Routes
	router := chi.NewRouter()
	router.Get("/readings", rh.ListReadings)

	// MQTT Routes
	mqttClient.Subscribe("weatherapp/reading", 2, rh.AddReading)

	// Run server
	log.Printf("Server listening on port %s", os.Getenv("API_PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("API_PORT"), router))
}
