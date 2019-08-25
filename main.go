package main

import (
	"line-boi/servicemanagement"
	"line-boi/servicemanagement/delivery/http"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/line/line-bot-sdk-go/linebot"
)

func main() {
	startService()
}

func connectLineBot() *linebot.Client {
	bot, err := linebot.New(
		os.Getenv("57fdbf7e00f0a73c912b9980059f14ae"),
		os.Getenv("yT35cNNeUMINaUceIHQZ4Ekf/w9PdmrmHUk7vG9NQwZdVM949YcAvKNrWeASyosY9qblfzfLFyKWNodhlDymCDU91oNtZvSNxZms+PtH7S+z/ffz39kRMyTWsJixwd8AjVkIy3TwedCUb+vARz9YvQdB04t89/1O/w1cDnyilFU="),
	)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func startService() {
	e := echo.New()
	bankCoreInfo := servicemanagement.NewBankCoreServiceInfo()
	http.NewServiceHTTPHandler(e, connectLineBot(), bankCoreInfo)
	e.Logger.Fatal(e.Start(":6000"))
}
