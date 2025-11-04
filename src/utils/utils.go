package utils

import "strings"

// EscapeMarkdownV2 escapes characters for Telegram's MarkdownV2 format.
func EscapeMarkdownV2(text string) string {
	replacer := strings.NewReplacer(
		"_", "\\_",
		"*", "\\*",
		"[", "\\[",
		"]", "\\]",
		"(", "\\(",
		")", "\\)",
		"~", "\\~",
		"`", "\\`",
		">", "\\>",
		"#", "\\#",
		"+", "\\+",
		"-", "\\-",
		"=", "\\=",
		"|", "\\|",
		"{", "\\{",
		"}", "\\}",
		".", "\\.",
		"!", "\\!",
	)
	return replacer.Replace(text)
}

// EscapeMarkdownV2URL escapes characters for the URL part of a MarkdownV2 link.
func EscapeMarkdownV2URL(text string) string {
	replacer := strings.NewReplacer(
		"(", "\\(",
		")", "\\)",
	)
	return replacer.Replace(text)
}
