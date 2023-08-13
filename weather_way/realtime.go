package weather_way

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

// GetWeatherInfo 获取实时天气信息
func GetRealtimeWeatherInfo(longitude, latitude float64) (RealtimeData, error) {
	// 根据提供的经度和纬度构建 API 请求 URL
	url := fmt.Sprintf("https://api.caiyunapp.com/v2.6/r4vS4ukc44vWETiL/%v,%v/realtime", longitude, latitude) //经度，纬度

	// 发送 HTTP GET 请求获取天气信息
	resp, err := http.Get(url)
	if err != nil {
		return RealtimeData{}, err
	}
	defer resp.Body.Close()

	// 解码 JSON 响应体并将天气信息存储到 RealtimeData 结构体中
	var info RealtimeData
	err = json.NewDecoder(resp.Body).Decode(&info) //创建一个 JSON 解码器
	if err != nil {
		return RealtimeData{}, err
	}
	return info, nil

}
