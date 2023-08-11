package weather

// AlertData 预警信息
type AlertData struct {
	Status  string    `json:"status"`
	Content []Content `json:"content"`
	Adcodes []Adcodes `json:"adcodes"`
}

// Content 预警信息
type Content struct {
	Province       string    `json:"province"`
	Status         string    `json:"status"`
	Code           string    `json:"code"`
	Description    string    `json:"description"`
	RegionId       string    `json:"regionId"`
	County         string    `json:"county"`
	Pubtimestamp   int       `json:"pubtimestamp"`
	Latlon         []float64 `json:"latlon"`
	City           string    `json:"city"`
	AlertId        string    `json:"alertId"`
	Title          string    `json:"title"`
	Adcode         string    `json:"adcode"`
	Source         string    `json:"source"`
	Location       string    `json:"location"`
	Request_status string    `json:"request_status"`
}

// Adcodes 预警信息
type Adcodes struct {
	Adcode int    `json:"adcode"`
	Name   string `json:"name"`
}
