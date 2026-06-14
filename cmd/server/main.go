package main

import (
	"Obsys/internal/api"
	"Obsys/internal/database"
	"Obsys/internal/metrices"
	"Obsys/internal/monitoring"
	"Obsys/internal/servers"
	"Obsys/internal/service"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Println(envErr)
	}

	password := os.Getenv("PASSWORD")
	log.Println(password)

	db, err := database.NewPostgresDB("postgres://postgres:" + password + "@host.docker.internal:5432/postgres")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connection successful")

	metricService := service.NewService(db)

	metricChan := make(chan metrices.Metric, 1000)
	queryChan := make(chan (service.RequestObj), 1000)

	handler := api.NewHandler(metricChan, queryChan, metricService)
	router := servers.SetUpRouter(handler)

	go service.SaveWorker(metricChan, queryChan, metricService)

	monitoring.RegisterMetrics()
	log.Println("Server running on port 8090")

	errServer := http.ListenAndServe(":8090", router)

	if errServer != nil {
		log.Fatal("Unable to start server ", errServer)
	}
}
