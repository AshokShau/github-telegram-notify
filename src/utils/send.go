package utils

import (
	"errors"
	"fmt"
	"github-webhook/src/config"
	"io"
	"log"
	"net/http"
	"strings"
)

func SendToTelegram(chatID, message string) error {
	if message == "" || chatID == "" {
		return errors.New("message or chatID is empty")
	}

	if config.BotToken == "" {
		log.Println("Telegram bot token is not set")
		return errors.New("telegram bot token is not set")
	}

	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.BotToken)
	payload := fmt.Sprintf(`{"chat_id":"%s", "text":"%s", "parse_mode":"HTML", "disable_web_page_preview": true}`, chatID, message)
	req, err := http.NewRequest("POST", telegramURL, strings.NewReader(payload))
	if err != nil {
		log.Println("Error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to Telegram:", err)
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Println("Error response from Telegram:", resp.Status)
		return err
	} else {
		log.Println("Message sent to Telegram")
		return nil
	}
}
