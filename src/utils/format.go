package utils

import "fmt"

func FormatMessageWithButton(message, buttonText, buttonURL string) (string, *InlineKeyboardMarkup) {
	if buttonText == "" || buttonURL == "" {
		return message, nil
	}
	markup := &InlineKeyboardMarkup{
		InlineKeyboard: [][]InlineKeyboardButton{
			{
				{Text: buttonText, URL: buttonURL},
			},
		},
	}
	return message, markup
}

func FormatRepo(repoFullName string) string {
	return fmt.Sprintf("[%s](https://github.com/%s)", EscapeMarkdownV2(repoFullName), EscapeMarkdownV2URL(repoFullName))
}

func FormatUser(userLogin string) string {
	return fmt.Sprintf("[%s](https://github.com/%s)", EscapeMarkdownV2(userLogin), EscapeMarkdownV2URL(userLogin))
}
