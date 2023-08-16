package weather_way

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MinutelyData struct {
	Status          string    `json:"status"`
	Datasource      string    `json:"datasource"`
	Precipitation2h []float64 `json:"precipitation_2h"` //过去2小时降水量
	Precipitation   []float64 `json:"precipitation"`
	Probability     []float64 `json:"probability"`
	Description     string    `json:"description"`
}

func GetMinutelyWeatherInfo(longitude, latitude float64) (MinutelyData, error) {
	// 根据提供的经度和纬度构建 API 请求 URL
	url := fmt.Sprintf("https://api.caiyunapp.com/v2.6/r4vS4ukc44vWETiL/%v,%v/realtime", longitude, latitude) //经度，纬度

	// 发送 HTTP GET 请求获取天气信息
	resp, err := http.Get(url)
	if err != nil {
		return MinutelyData{}, err
	}
	defer resp.Body.Close()

	// 解码 JSON 响应体并将天气信息存储到 MinutelyData 结构体中
	var info MinutelyData
	err = json.NewDecoder(resp.Body).Decode(&info) //创建一个 JSON 解码器
	if err != nil {
		return MinutelyData{}, err
	}
	return info, nil

}
