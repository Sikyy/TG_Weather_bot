package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// WeatherInfo 结构体定义
type WeatherInfo struct {
	// 定义你的结构体字段
}

func main() {
	bot, err := tgbotapi.NewBotAPI("YOUR_BOT_TOKEN")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	// 设置 Webhook 和 HTTP 服务器
	setupWebhook(bot)

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		handleUpdate(bot, update)
	}
}

func setupWebhook(bot *tgbotapi.BotAPI) {
	// 设置 Webhook
	_, err := bot.SetWebhook(tgbotapi.NewWebhookWithCert("YOUR_WEBHOOK_URL", "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	// 启动 HTTP 服务器
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "privkey.pem", nil)
}

func handleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Location != nil {
		handleLocation(bot, update)
	} else if update.Message.Text != "" {
		handleTextMessage(bot, update)
	}
}

func handleLocation(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// 处理用户发送的位置
	location := update.Message.Location
	fmt.Println(location, "用户位置已经更新")

	// 获取天气信息并发送回复
	weatherInfo, err := GetWeatherInfo()
	if err != nil {
		response := "获取天气信息失败"
		sendMessage(bot, update.Message.Chat.ID, response)
		return
	}

	// 根据天气信息发送回复
	response := getWeatherResponse(weatherInfo)
	sendMessage(bot, update.Message.Chat.ID, response)
}

func handleTextMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// 处理用户发送的文本消息
	text := update.Message.Text

	if text == "今天的天气怎样" {
		weatherInfo, err := GetWeatherInfo()
		if err != nil {
			response := "获取天气信息失败"
			sendMessage(bot, update.Message.Chat.ID, response)
			return
		}

		response := getWeatherResponse(weatherInfo)
		sendMessage(bot, update.Message.Chat.ID, response)
	} else {
		// 回复用户发送的文本
		sendMessage(bot, update.Message.Chat.ID, text)
	}
}

func sendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func GetWeatherInfo() (WeatherInfo, error) {
	url := "https://api.caiyunapp.com/v2.6/YOUR_API_KEY/101.6656,39.2072/realtime"

	resp, err := http.Get(url)
	if err != nil {
		return WeatherInfo{}, err
	}
	defer resp.Body.Close()

	var info WeatherInfo
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return WeatherInfo{}, err
	}

	return info, nil
}

func getWeatherResponse(weatherInfo WeatherInfo) string {
	// 根据天气信息生成回复内容
	// 注意：这里只是一个示例，根据实际情况进行生成
	return "天气预报：" + "TODO"
}
