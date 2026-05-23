package main

import (
	"Obsys/internal/api"
	"Obsys/internal/metrices"
	"Obsys/internal/servers"
	"log"
	"net/http"
)

func main() {
	storage := metrices.NewStorage()
	handler := api.NewHandler(storage)
	router := servers.SetUpRouter(handler)

	log.Println("Server running on port 8090")

	err := http.ListenAndServe(":8090", router)

	if err != nil {
		log.Fatal("Unable to start server ", err)
	}
}
