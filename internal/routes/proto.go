package routes

import (
	"net/http"
)

func returnProto(w http.ResponseWriter, response []byte) {
	w.Write(response)
}
