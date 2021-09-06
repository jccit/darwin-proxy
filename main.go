package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"

	"github.com/joho/godotenv"

	"github.com/jccit/darwin-proxy/internal/darwin"
)

func returnJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response[:]))
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	reqPath := strings.Split(r.URL.Path, "/")

	if len(reqPath) > 2 {
		method := reqPath[1]
		param, _ := url.PathUnescape(reqPath[2])

		switch method {
		case "departures":
			if utf8.RuneCountInString(param) != 3 {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				response := darwin.GetDepartures(strings.ToUpper(param), r)
				returnJSON(w, response)
			}
			return
		}
	}

	http.NotFound(w, r)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	http.HandleFunc("/", requestHandler)
	log.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
