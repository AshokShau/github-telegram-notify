package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github-webhook/src/config"
	"io"
	"log"
	"net/http"
)

type InlineKeyboardButton struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type sendMessagePayload struct {
	ChatID                string                `json:"chat_id"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func SendToTelegram(chatID, message string, markup *InlineKeyboardMarkup) error {
	if config.BotToken == "" {
		log.Println("Telegram bot token is not set")
		return errors.New("telegram bot token is not set")
	}

	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.BotToken)

	payload := sendMessagePayload{
		ChatID:                chatID,
		Text:                  message,
		ParseMode:             "MarkdownV2",
		DisableWebPagePreview: true,
		ReplyMarkup:           markup,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling payload:", err)
		return err
	}

	req, err := http.NewRequest("POST", telegramURL, bytes.NewBuffer(payloadBytes))
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
		body, _ := io.ReadAll(resp.Body)
		log.Printf("Error response from Telegram: %s - %s", resp.Status, string(body))
		return fmt.Errorf("telegram API error: %s", resp.Status)
	}

	log.Println("Message sent to Telegram")
	return nil
}
