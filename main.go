package main

import (
	"log"

	"weather-bot/handlers"
	"weather-bot/webhook"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("6674995287:AAGK820XcZiKkzYEBK_jRaFMk0eImWX97-c")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	// 设置 Webhook 和 HTTP 服务器
	webhook.SetupWebhook(bot)
	webhook.ShowWebhookInfo(bot)

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		handlers.HandleUpdate(bot, update)
	}

}
