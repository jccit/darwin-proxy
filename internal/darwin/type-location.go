package darwin

type Location struct {
	CRS  string `xml:"crs" json:"crs"`
	Name string `xml:"locationName" json:"name"`
}

type CallingPoint struct {
	Location
	ScheduledTime string `xml:"st" json:"scheduledTime"`
	EstimatedTime string `xml:"et" json:"estimatedTime"`
	ActualTime    string `xml:"at" json:"actualTime"`
}
