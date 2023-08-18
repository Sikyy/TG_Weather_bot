package user

import (
	"fmt"
	"reflect"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// 在 HandleUpdate 外部定义发送位置请求函数
func SendLocationRequest(bot *tgbotapi.BotAPI, chatID int64) {
	msg := tgbotapi.NewMessage(chatID, "请点击下面按钮分享您的位置")
	requestLocation := tgbotapi.NewKeyboardButtonLocation("分享位置")
	keyboard := tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(requestLocation))
	msg.ReplyMarkup = keyboard
	bot.Send(msg)
}

// 将天气信息格式化为字符串
func FormatStructToString(data interface{}) string {
	result := ""
	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Struct {
		t := v.Type()
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			fieldValue := v.Field(i).Interface()

			result += fmt.Sprintf("%s: %v\n", field.Name, fieldValue)
		}
	}

	return result
}

// bot发消息给用户
func SendMessage(bot *tgbotapi.BotAPI, chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	bot.Send(msg)
}
