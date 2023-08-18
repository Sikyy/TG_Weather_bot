package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("6674995287:AAGK820XcZiKkzYEBK_jRaFMk0eImWX97-c")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	whConfig := tgbotapi.NewWebhookWithCert("https://weather.siky.me:8443/"+bot.Token, "fullchain.pem")

	_, err = bot.SetWebhook(whConfig)
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

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "fullchain.pem", "privkey.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}

}
