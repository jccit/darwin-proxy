package darwin

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/jccit/darwin-proxy/internal/soap"
	"github.com/jccit/darwin-proxy/pb"
	"google.golang.org/protobuf/proto"
)

const departuresXML = `<ldb:GetDepartureBoardRequest>||FILTER||</ldb:GetDepartureBoardRequest>`

type DepartureBoardResponse struct {
	Board DepartureBoard `xml:"Body>GetDepartureBoardResponse>GetStationBoardResult"`
}

type DepartureBoard struct {
	GeneratedAt       string         `xml:"generatedAt" json:"generatedAt"`
	Location          string         `xml:"locationName" json:"location"`
	CRS               string         `xml:"crs" json:"crs"`
	PlatformAvailable bool           `xml:"platformAvailable" json:"platformAvailable"`
	Services          []BoardService `xml:"trainServices>service" json:"services"`
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

func GetDeparturesAsJSON(crs string, r *http.Request) []byte {
	response := GetDepartures(crs, r)
	json, _ := json.Marshal(response)
	return json
}

func GetDeparturesAsProto(crs string, r *http.Request) []byte {
	response := GetDepartures(crs, r)
	jsonOut, _ := json.Marshal(response.Board)

	var protoOut pb.DepartureBoard
	json.Unmarshal(jsonOut, &protoOut)
	encodedProto, _ := proto.Marshal(&protoOut)

	return encodedProto
}
