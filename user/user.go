package user

import (
	"weather-bot/weather_way"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// 在 HandleUpdate 外部定义发送位置请求函数
func SendLocationRequest(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "请点击下面按钮分享您的位置")
	requestLocation := tgbotapi.NewKeyboardButtonLocation("分享位置")
	keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(requestLocation))
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

// 将实时天气信息格式化为字符串
func FormatWeatherInfoToString(info weather_way.RealtimeData) string {
	// 根据需要，将 RealtimeData 中的字段信息格式化为字符串
	// 例如：return fmt.Sprintf("实时天气：温度 %.1f°C，湿度 %.1f%%", info.Temperature, info.Humidity)
	return "TODO"
}

// bot发消息给用户
func SendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}
