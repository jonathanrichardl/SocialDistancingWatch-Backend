package telegram

import (
	"bytes"
	"capstone/pkg/logger"
	"capstone/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Telegram struct {
	logger *logger.LoggerInstance
}

func NewTelegram(Logger *logger.LoggerInstance) *Telegram {
	return &Telegram{logger: Logger}

}

func (t *Telegram) SendTelegramMessage(wg *sync.WaitGroup, ChannelName int64, Payload string, ImageLink string) {
	reqBody := &models.TelegramOutgoingMessage{
		ChatID: ChannelName,
		Text:   Payload,
	}

	photoRequest := &models.TelegramPhoto{
		ChatID:    ChannelName,
		PhotoLink: ImageLink,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	photoBytes, err := json.Marshal(photoRequest)

	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = http.Post("https://api.telegram.org/bot2030379612:AAEI9HSNeWg8CQYHnqgVl9I5uCW--pt8Ggs/sendMessage",
		"application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = http.Post("https://api.telegram.org/bot2030379612:AAEI9HSNeWg8CQYHnqgVl9I5uCW--pt8Ggs/sendPhoto",
		"application/json", bytes.NewBuffer(photoBytes))
	if err != nil {
		fmt.Println(err.Error())
	}
	wg.Done()
}
