package routes

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/jccit/darwin-proxy/internal/darwin"
)

func Departures(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	crs := ps.ByName("crs")
	format := r.URL.Query().Get("format")

	if format == "proto" {
		response := darwin.GetDeparturesAsProto(strings.ToUpper(crs), r)
		returnProto(w, response)
	} else {
		response := darwin.GetDeparturesAsJSON(strings.ToUpper(crs), r)
		returnJSON(w, response)
	}
}
