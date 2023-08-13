package weather_json

type DailyData struct {
	Status              string              `json:"status"`
	Astro               []AstroData         `json:"astro"`
	Precipitation08h20h []PrecipitationData `json:"precipitation_08h_20h"`
	Precipitation20h32h []PrecipitationData `json:"precipitation_20h_32h"`
	Precipitation       []PrecipitationData `json:"precipitation"`
	Temperature         []TemperatureData   `json:"temperature"`
	Temperature08h20h   []TemperatureData   `json:"temperature_08h_20h"`
	Temperature20h32h   []TemperatureData   `json:"temperature_20h_32h"`
	Wind                []WindData_1        `json:"wind"`
	Wind08h20h          []WindData_1        `json:"wind_08h_20h"`
	Wind20h32h          []WindData_1        `json:"wind_20h_32h"`
	Humidity            []OneData           `json:"humidity"`
	Cloudrate           []OneData           `json:"cloudrate"`
	Pressure            []OneData           `json:"pressure"`
	Visibility          []OneData           `json:"visibility"`
	Dswrf               []OneData           `json:"dswrf"`
	AirQuality          AirQualityData_1    `json:"air_quality"`
	Skycon              []SkyconData        `json:"skycon"`
	Skycon08h20h        []SkyconData        `json:"skycon_08h_20h"`
	Skycon20h32h        []SkyconData        `json:"skycon_20h_32h"`
	LifeIndex           LifeIndexData       `json:"life_index"`
}

type AstroData struct {
	Date   string     `json:"date"`
	Sun    SunData    `json:"sun"`
	Sunset SunsetData `json:"moon"`
}

type SunData struct {
	Time string `json:"time"`
}

type SunsetData struct {
	Time string `json:"time"`
}

type PrecipitationData struct {
	Date        string  `json:"date"`
	Max         float64 `json:"max"`
	Min         float64 `json:"min"`
	Avg         float64 `json:"avg"`
	Probability float64 `json:"probability"`
}

type TemperatureData struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}

type WindData_1 struct {
	Date string  `json:"date"`
	Max  MaxData `json:"max"`
	Min  MinData `json:"min"`
	Avg  AvgData `json:"avg"`
}

type MaxData struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"direction"`
}

type MinData struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"direction"`
}

type AvgData struct {
	Speed     float64 `json:"speed"`
	Direction float64 `json:"direction"`
}

type OneData struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}

type AirQualityData_1 struct {
	Aqi  []AqiData_1  `json:"aqi"`
	Pm25 []Pm25Data_1 `json:"pm25"`
}

type AqiData_1 struct {
	Date string    `json:"date"`
	Max  MaxData_1 `json:"max"`
	Min  MinData_1 `json:"min"`
	Avg  AvgData_1 `json:"avg"`
}

type MaxData_1 struct {
	Chn float64 `json:"chn"`
	Usa float64 `json:"usa"`
}

type MinData_1 struct {
	Chn float64 `json:"chn"`
	Usa float64 `json:"usa"`
}

type AvgData_1 struct {
	Chn float64 `json:"chn"`
	Usa float64 `json:"usa"`
}

type Pm25Data_1 struct {
	Date string  `json:"date"`
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Avg  float64 `json:"avg"`
}

type SkyconData struct {
	Data  string `json:"data"`
	Value string `json:"value"`
}

type LifeIndexData struct {
	Ultraviolet []LifeIndexData_1 `json:"ultraviolet"`
	CarWashing  []LifeIndexData_1 `json:"carWashing"`
	Dressing    []LifeIndexData_1 `json:"dressing"`
	ColdRisk    []LifeIndexData_1 `json:"coldRisk"`
}

type LifeIndexData_1 struct {
	Date  string `json:"date"`
	Index string `json:"index"`
	Desc  string `json:"desc"`
}
