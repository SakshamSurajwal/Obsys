package main

import (
	"Obsys/internal/api"
	"Obsys/internal/metrices"
	"Obsys/internal/servers"
	"Obsys/internal/service"
	"log"
	"net/http"
)

func main() {
	storage := metrices.NewStorage()
	metricService := service.NewService(storage)

	metricChan := make(chan metrices.Metric, 1000)
	queryChan := make(chan (chan []metrices.Metric), 1000)

	handler := api.NewHandler(metricChan, metricService)
	router := servers.SetUpRouter(handler)

	go service.SaveWorker(metricChan, queryChan, metricService)

	log.Println("Server running on port 8090")

	err := http.ListenAndServe(":8090", router)

	if err != nil {
		log.Fatal("Unable to start server ", err)
	}
}
