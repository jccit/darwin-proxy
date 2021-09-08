package routes

import (
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"

	"github.com/jccit/darwin-proxy/internal/darwin"
)

func Departures(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	crs := ps.ByName("crs")
	response := darwin.GetDepartures(strings.ToUpper(crs), r)
	returnJSON(w, response)
}
