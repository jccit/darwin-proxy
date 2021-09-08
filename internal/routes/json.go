package routes

import (
	"fmt"
	"net/http"
)

func returnJSON(w http.ResponseWriter, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(response[:]))
}
