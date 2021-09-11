package darwin

import (
	"encoding/xml"
	"net/http"
	"strings"

	"github.com/jccit/darwin-proxy/internal/soap"
)

const serviceXML = `<ldb:GetServiceDetailsRequest>||FILTER||</ldb:GetServiceDetailsRequest>`

type ServiceDetailResponse struct {
	Detail ServiceDetail `xml:"Body>GetServiceDetailsResponse>GetServiceDetailsResult"`
}

type ServiceDetail struct {
	Location
	GeneratedAt             string         `xml:"generatedAt"`
	ServiceType             string         `xml:"serviceType"`
	Platform                int            `xml:"platform"`
	Operator                string         `xml:"operator"`
	OperatorCode            string         `xml:"operatorCode"`
	ScheduledDeparture      string         `xml:"std"`
	EstimatedDeparture      string         `xml:"etd"`
	SubsequentCallingPoints []CallingPoint `xml:"subsequentCallingPoints>callingPointList>callingPoint"`
}

func getServiceRequestXML(filter string) string {
	body := strings.Replace(serviceXML, "||FILTER||", filter, 1)
	parts := []string{"<soap:Body>", body, "</soap:Body>", "</soap:Envelope>"}
	combined := strings.Join(parts, "\n")
	return strings.Replace(combined, "\n", "", -1)
}

func GetServiceDetails(id string, r *http.Request) ServiceDetailResponse {
	soapReq := getServiceRequestXML(ServiceSelector(id))
	response := soap.SendDarwinRequest(soapReq, r)

	var parsedResponse ServiceDetailResponse
	xml.Unmarshal(response, &parsedResponse)

	return parsedResponse
}
