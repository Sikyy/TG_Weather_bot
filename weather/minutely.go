package weather

type MinutelyData struct {
	Status          string    `json:"status"`
	Datasource      string    `json:"datasource"`
	Precipitation2h []float64 `json:"precipitation_2h"` //过去2小时降水量
	Precipitation   []float64 `json:"precipitation"`
	Probability     []float64 `json:"probability"`
	Description     string    `json:"description"`
}
