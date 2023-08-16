package handlers

import (
	"fmt"
	"log"
	"weather-bot/user"
	"weather-bot/weather_way"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

//更新处理函数

// 拿到所有更新
func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.Message == nil {
		log.Printf("Received update without message: %+v", update)
		return // Add this line to return early for empty messages
	}

	if update.Message.Location != nil {
		HandleLocation(bot, update)
	} else if update.Message.Text != "" {
		HandleTextMessage(bot, update)
	}
}

func HandleLocation(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	// 处理用户发送的位置
	location := update.Message.Location
	if location == nil {
		user.SendMessage(bot, update.Message.Chat.ID, "获取位置失败")
		return
	}
	fmt.Println(location, "用户位置已经更新")
	user.SendMessage(bot, update.Message.Chat.ID, "位置已经更新")
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
			user.SendMessage(bot, update.Message.Chat.ID, "获取实况天气失败")
			return
		}
		weatherString := user.FormatStructToString(realtimeinfo)
		user.SendMessage(bot, update.Message.Chat.ID, weatherString)
	case "/minutelyweather":
		minutelyinfo, err := weather_way.GetMinutelyWeatherInfo(location.Longitude, location.Latitude)
		if err != nil {
			user.SendMessage(bot, update.Message.Chat.ID, "获取分钟级降水失败")
			return
		}
		weatherString := user.FormatStructToString(minutelyinfo)
		user.SendMessage(bot, update.Message.Chat.ID, weatherString)

	case "/hourlyweather":
		hourlyinfo, err := weather_way.GetHourlyWeatherInfo(location.Longitude, location.Latitude)
		if err != nil {
			user.SendMessage(bot, update.Message.Chat.ID, "获取逐小时天气失败")
			return
		}
		weatherString := user.FormatStructToString(hourlyinfo)
		user.SendMessage(bot, update.Message.Chat.ID, weatherString)

	case "/dailyweather":
		dailyinfo, err := weather_way.GetDailyWeatherInfo(location.Longitude, location.Latitude)
		if err != nil {
			user.SendMessage(bot, update.Message.Chat.ID, "获取逐日天气失败")
			return
		}
		weatherString := user.FormatStructToString(dailyinfo)
		user.SendMessage(bot, update.Message.Chat.ID, weatherString)

	case "/alert":
		alertinfo, err := weather_way.GetAlertWeatherInfo(location.Longitude, location.Latitude)
		if err != nil {
			user.SendMessage(bot, update.Message.Chat.ID, "获取预警信息失败")
			return
		}
		weatherString := user.FormatStructToString(alertinfo)
		user.SendMessage(bot, update.Message.Chat.ID, weatherString)

	case "/adressupdate":
		// 发送位置请求
		user.SendLocationRequest(bot, update.Message.Chat.ID)
	}
}
