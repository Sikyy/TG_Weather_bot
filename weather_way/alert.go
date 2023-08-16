package weather_way

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

func GetAlertWeatherInfo(longitude, latitude float64) (AlertData, error) {
	// 根据提供的经度和纬度构建 API 请求 URL
	url := fmt.Sprintf("https://api.caiyunapp.com/v2.6/r4vS4ukc44vWETiL/%v,%v/realtime", longitude, latitude) //经度，纬度

	// 发送 HTTP GET 请求获取天气信息
	resp, err := http.Get(url)
	if err != nil {
		return AlertData{}, err
	}
	defer resp.Body.Close()

	// 解码 JSON 响应体并将天气信息存储到 AlertData 结构体中
	var info AlertData
	err = json.NewDecoder(resp.Body).Decode(&info) //创建一个 JSON 解码器
	if err != nil {
		return AlertData{}, err
	}
	return info, nil

}
