package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"

	"github.com/jccit/darwin-proxy/internal/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := httprouter.New()
	router.GET("/departures/:crs", routes.Departures)

	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
