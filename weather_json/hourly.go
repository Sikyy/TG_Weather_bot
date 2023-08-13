package weather_json

type HourlyData struct {
	Status               string                `json:"status"`
	Description          string                `json:"description"`
	Precipitation        []Precipitation       `json:"precipitation"`
	Temperature          []Temperature         `json:"temperature"`
	Apparent_temperature []ApparentTemperature `json:"apparent_temperature"`
	Wind                 []Wind                `json:"wind"`
	Humidity             []Humidity            `json:"humidity"`
	Cloudrate            []Cloudrate           `json:"cloudrate"`
	Skycon               []Skycon              `json:"skycon"`
	Pressure             []Pressure            `json:"pressure"`
	Visibility           []Visibility          `json:"visibility"`
	Dswrf                []Dswrf               `json:"dswrf"`
	Air_quality          AirQualityData        `json:"air_quality"`
}

type Precipitation struct {
	Datetime    string  `json:"datetime"`
	Value       float64 `json:"value"`
	Probability int     `json:"probability"`
}

type Temperature struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type ApparentTemperature struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type Wind struct {
	Datetime  string  `json:"datetime"`
	Speed     float64 `json:"speed"`
	Direction int     `json:"direction"`
}

type Humidity struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type Cloudrate struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type Skycon struct {
	Datetime string `json:"datetime"`
	Value    string `json:"value"`
}

type Pressure struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type Visibility struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type Dswrf struct {
	Datetime string  `json:"datetime"`
	Value    float64 `json:"value"`
}

type AirQualityData struct {
	Aqi_2 AqiData_2 `json:"aqi"`
	Pm25  Pm25Data  `json:"pm25"`
}

type AqiData_2 struct {
	Detatime string    `json:"detatime"`
	Value    ValueData `json:"value"`
}

type Pm25Data struct {
	Detatime string  `json:"detatime"`
	Value    float64 `json:"value"`
}

type ValueData struct {
	Chn float64 `json:"chn"`
	Usa float64 `json:"usa"`
}
