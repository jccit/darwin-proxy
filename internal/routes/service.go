package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/jccit/darwin-proxy/internal/darwin"
)

func Service(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := darwin.ServiceID(ps.ByName("id"))
	format := r.URL.Query().Get("format")

	if format == "proto" {
		response := darwin.GetServiceDetailsAsProto(id.URLDecode(), r)
		returnProto(w, response)
	} else {
		response := darwin.GetServiceDetailsAsJSON(id.URLDecode(), r)
		returnJSON(w, response)
	}
}
