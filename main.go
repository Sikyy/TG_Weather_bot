package main

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//加载配置
	viper.SetConfigFile("config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	botToken := viper.GetString("bot_token")
	webhookURL := viper.GetString("webhook_url")
	certFilePath := viper.GetString("cert_file_path")
	keyFilePath := viper.GetString("key_file_path")

	//设置日志
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync() // 释放资源

	//初始化bot
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		logger.Error("Error initializing bot", zap.Error(err))
		return
	}

	//开启调试
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	//设置webhook
	webhook, _ := tgbotapi.NewWebhookWithCert(webhookURL+bot.Token, tgbotapi.FilePath(certFilePath))

	_, err = bot.Request(webhook)
	if err != nil {
		logger.Error("Error setting up webhook", zap.Error(err))
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		logger.Error("Error setting up webhook", zap.Error(err))
	}

	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", certFilePath, keyFilePath, nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
