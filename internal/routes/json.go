package routes

import (
	"fmt"
	"net/http"
)

func returnJSON(w http.ResponseWriter, response []byte) {
	fmt.Fprintln(w, string(response[:]))
}
