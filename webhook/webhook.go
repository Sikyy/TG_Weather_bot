package webhook

import (
	"log"
	"net/http"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SetupWebhook(bot *tgbotapi.BotAPI) {
	//设置 Webhook
	_, err := bot.SetWebhook(tgbotapi.NewWebhookWithCert("weather.siky.me", "fullchain.pem"))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Setting webhook OK")

	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Webhook Info: %+v", info)

	// 如果 Webhook 设置成功，Telegram 服务器会返回一个 OK 的结果
	if info.LastErrorDate != 0 {
		log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}

	// 启动 HTTPS 服务器
	go http.ListenAndServeTLS("127.0.0.1:8443", "fullchain.pem", "privkey.pem", nil)
}

func ShowWebhookInfo(bot *tgbotapi.BotAPI) {
	info, err := bot.GetWebhookInfo()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Webhook Info: %+v", info)
	log.Printf("URL: %s", info.URL)
	log.Printf("Last Error Date: %d", info.LastErrorDate)
	log.Printf("Last Error Message: %s", info.LastErrorMessage)

}
