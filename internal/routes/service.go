package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/jccit/darwin-proxy/internal/darwin"
)

func Service(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	response := darwin.GetServiceDetails(id, r)
	returnJSON(w, response)
}
