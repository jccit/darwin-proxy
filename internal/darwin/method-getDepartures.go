package darwin

import (
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/jccit/darwin-proxy/internal/soap"
)

const departuresXML = `<ldb:GetDepartureBoardRequest>||FILTER||</ldb:GetDepartureBoardRequest>`

type DepartureBoardResponse struct {
	Board DepartureBoard `xml:"Body>GetDepartureBoardResponse>GetStationBoardResult"`
}

type DepartureBoard struct {
	GeneratedAt       string         `xml:"generatedAt"`
	Location          string         `xml:"locationName"`
	CRS               string         `xml:"crs"`
	PlatformAvailable bool           `xml:"platformAvailable"`
	Services          []BoardService `xml:"trainServices>service"`
}

func getDeparturesRequestXML(filter string) string {
	body := strings.Replace(departuresXML, "||FILTER||", filter, 1)
	parts := []string{"<soap:Body>", body, "</soap:Body>", "</soap:Envelope>"}
	combined := strings.Join(parts, "\n")
	return strings.Replace(combined, "\n", "", -1)
}

func GetDepartures(crs string, r *http.Request) DepartureBoardResponse {
	soapReq := getDeparturesRequestXML(CRSSelector(crs))
	response := soap.SendDarwinRequest(soapReq, r)

	var parsedResponse DepartureBoardResponse
	xml.Unmarshal(response, &parsedResponse)

	return parsedResponse
}
