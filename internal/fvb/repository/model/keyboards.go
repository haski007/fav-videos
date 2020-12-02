package model

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func NewOriginalURLMarkup(url string) tgbotapi.InlineKeyboardMarkup {
	return tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.InlineKeyboardButton{
				Text: "Original",
				URL:  &url,
			}))
}
