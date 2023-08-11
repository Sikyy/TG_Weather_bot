package weather

type RealtimeData struct {
	Status        string     `json:"status"`
	Temperature   float64    `json:"temperature"`
	Humidity      float64    `json:"humidity"`
	Cloudrate     float64    `json:"cloudrate"`
	Skycon        string     `json:"skycon"`
	Visibility    float64    `json:"visibility"`
	Dswrf         float64    `json:"dswrf"`
	Wind          WindData   `json:"wind"`
	Pressure      float64    `json:"pressure"`
	Apparent_temp float64    `json:"apparent_temperature"`
	Precipitation PrecipData `json:"precipitation"`
	AirQuality    AirData    `json:"air_quality"`
	LifeIndex     LifeData   `json:"life_index"`
}

type WindData struct {
	Speed  float64 `json:"speed"`
	Direct int     `json:"direct"`
}

type PrecipData struct {
	Local   LocalData   `json:"local"`
	Nearest NearestData `json:"nearest"`
}

type LocalData struct {
	Status     string `json:"status"`
	Datasource string `json:"datasource"`
	Intensity  int    `json:"intensity"`
}

type NearestData struct {
	Status    string `json:"status"`
	Distance  int    `json:"distance"`
	Intensity int    `json:"intensity"`
}

type AirData struct {
	Pm25        float64         `json:"pm25"`
	Pm10        float64         `json:"pm10"`
	O3          float64         `json:"o3"`
	So2         float64         `json:"so2"`
	No2         float64         `json:"no2"`
	Co          float64         `json:"co"`
	Aqi         AqiData         `json:"aqi"`
	Description DescriptionData `json:"description"`
}

type AqiData struct {
	Chn float64 `json:"chn"`
	Usa float64 `json:"usa"`
}

type DescriptionData struct {
	Chn string `json:"chn"`
	Usa string `json:"usa"`
}

type LifeData struct {
	Ultraviolet UltravioletData `json:"ultraviolet"`
	Comfort     ComfortData     `json:"comfort"`
}

type UltravioletData struct {
	Index int    `json:"index"`
	Desc  string `json:"desc"`
}

type ComfortData struct {
	Index int    `json:"index"`
	Desc  string `json:"desc"`
}
