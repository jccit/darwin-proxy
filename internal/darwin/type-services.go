package darwin

import "strings"

type BoardService struct {
	ScheduledDeparture string    `xml:"std"`
	EstimatedDeparture string    `xml:"etd"`
	ScheduledArrival   string    `xml:"sta"`
	EstimatedArrival   string    `xml:"eta"`
	Platform           string    `xml:"platform"`
	Operator           string    `xml:"operator"`
	OperatorCode       string    `xml:"operatorCode"`
	ServiceType        string    `xml:"serviceType"`
	ServiceID          ServiceID `xml:"serviceID"`
	Origin             Location  `xml:"origin>location"`
	Destination        Location  `xml:"destination>location"`
}

type ServiceID string

func (s ServiceID) URLEncode() string {
	out := string(s)
	out = strings.ReplaceAll(out, "/", "_")
	out = strings.ReplaceAll(out, "+", "-")

	return out
}

func (s ServiceID) URLDecode() string {
	out := string(s)
	out = strings.ReplaceAll(out, "_", "/")
	out = strings.ReplaceAll(out, "-", "+")

	return out
}

func (s ServiceID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + s.URLEncode() + `"`), nil
}
