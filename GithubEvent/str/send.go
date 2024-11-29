package str

import (
	"fmt"
	"github-webhook/GithubEvent/config"
	"io"
	"log"
	"net/http"
	"strings"
)

func sendToTelegram(chatID, message string) {
	if message == "" || chatID == "" {
		return
	}

	if config.BotToken == "" {
		log.Println("Telegram bot token is not set")
		return
	}

	telegramURL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.BotToken)
	payload := fmt.Sprintf(`{"chat_id":"%s", "text":"%s", "parse_mode":"HTML", "disable_web_page_preview": true}`, chatID, message)
	req, err := http.NewRequest("POST", telegramURL, strings.NewReader(payload))
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request to Telegram:", err)
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Println("Error response from Telegram:", resp.Status)
	} else {
		log.Println("Message sent to Telegram")
	}
}
