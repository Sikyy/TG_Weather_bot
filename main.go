package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

func main() {
	//加载配置
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	certFile := viper.GetString("cert_file")
	keyFile := viper.GetString("key_file")

	bot, err := tgbotapi.NewBotAPI("6674995287:AAGK820XcZiKkzYEBK_jRaFMk0eImWX97-c")
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	wh, _ := tgbotapi.NewWebhookWithCert("https://weather.siky.me:8443/"+bot.Token, tgbotapi.FilePath(certFile))

	_, err = bot.Request(wh)
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
	go http.ListenAndServeTLS("0.0.0.0:8443", certFile, keyFile, nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
