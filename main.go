package main

import (
	"fmt"
	"log"
	"net/http"

	"weather-bot/weather_way"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupWebhook(bot *tgbotapi.BotAPI) {
	// 设置 Webhook
	_, err := bot.SetWebhook(tgbotapi.NewWebhookWithCert("weather.siky.me", "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}

	// 如果 Webhook 设置成功，Telegram 服务器会返回一个 OK 的结果
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	// 启动 HTTP 服务器
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "privkey.pem", nil)
}

// 拿到所有更新
func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Location != nil {
		HandleLocation(bot, update)
	} else if update.Message.Text != "" {
		HandleTextMessage(bot, update)
	}
}

// 在 HandleUpdate 外部定义发送位置请求函数
func SendLocationRequest(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "请点击下面按钮分享您的位置")
	requestLocation := tgbotapi.NewKeyboardButtonLocation("分享位置")
	keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(requestLocation))
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

func HandleLocation(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// 处理用户发送的位置
	location := update.Message.Location
	fmt.Println(location, "用户位置已经更新")
	SendMessage(bot, update.Message.Chat.ID, "位置已经更新")
}

// 将实时天气信息格式化为字符串
func FormatWeatherInfoToString(info weather_way.RealtimeData) string {
	// 根据需要，将 RealtimeData 中的字段信息格式化为字符串
	// 例如：return fmt.Sprintf("实时天气：温度 %.1f°C，湿度 %.1f%%", info.Temperature, info.Humidity)
	return "TODO"
}

func HandleTextMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// 拿到实时位置
	location := update.Message.Location
	// 处理用户发送的文本消息
	text := update.Message.Text

	switch text {

	case "/realweather":
		realtimeinfo, err := weather_way.GetRealtimeWeatherInfo(location.Longitude, location.Latitude)
		if err != nil {
			SendMessage(bot, update.Message.Chat.ID, "获取实况天气失败")
			return
		}
		weatherString := FormatWeatherInfoToString(realtimeinfo)
		SendMessage(bot, update.Message.Chat.ID, weatherString)

	case "/adressupdate":
		// 发送位置请求
		SendLocationRequest(bot, update.Message.Chat.ID)
	}
}

// bot发消息给用户
func SendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}

func main() {

	bot, err := tgbotapi.NewBotAPI("6674995287:AAGK820XcZiKkzYEBK_jRaFMk0eImWX97-c")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	// 设置 Webhook 和 HTTP 服务器
	SetupWebhook(bot)

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		HandleUpdate(bot, update)
	}

}
