package darwin

import "strings"

type BoardService struct {
	ScheduledDeparture string    `xml:"std" json:"scheduledDeparture"`
	EstimatedDeparture string    `xml:"etd" json:"estimatedDeparture"`
	ScheduledArrival   string    `xml:"sta" json:"scheduledArrival"`
	EstimatedArrival   string    `xml:"eta" json:"estimatedArrival"`
	Platform           string    `xml:"platform" json:"platform"`
	Operator           string    `xml:"operator" json:"operator"`
	OperatorCode       string    `xml:"operatorCode" json:"operatorCode"`
	ServiceType        string    `xml:"serviceType" json:"serviceType"`
	ServiceID          ServiceID `xml:"serviceID" json:"serviceID"`
	Origin             Location  `xml:"origin>location" json:"origin"`
	Destination        Location  `xml:"destination>location" json:"destination"`
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
