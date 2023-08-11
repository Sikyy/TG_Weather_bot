package weather

type WeatherInfo struct {
	Status     string     `json:"status"`
	APIVersion string     `json:"api_version"`
	APIStatus  string     `json:"api_status"`
	Lang       string     `json:"lang"`
	Unit       string     `json:"unit"`
	Tzshift    int        `json:"tzshift"`
	Timezone   string     `json:"timezone"`
	ServerTime int        `json:"server_time"`
	Location   []int      `json:"location"`
	Result     ResultData `json:"result"`
}

// ResultData 天气信息
type ResultData struct {
	Alert            AlertData    `json:"alert"`
	Realtime         RealtimeData `json:"realtime"`
	Minutely         MinutelyData `json:"minutely"`
	Hourly           HourlyData   `json:"hourly"`
	Daily            DailyData    `json:"daily"`
	Primary          int          `json:"primary"`
	ForecastKeyPoint string       `json:"forecast_keypoint"`
}
