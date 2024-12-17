package main

import (
	"log"
	"login/config"
	"login/routes"
	"net/http"
)

func main() {
	config.ConnectDB()
	routes.RegisterRoutes()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	log.Println("Server started in port 8080")
}
