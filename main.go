package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"github.com/jccit/darwin-proxy/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Failed to load .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.GET("/departures/:crs", routes.Departures)
	router.GET("/service/:id", routes.Service)

	log.Println("Server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
