package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var startTemplate = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("New game", "new"),
	),
)

var newGameTemplate = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("white", "white"),
		tgbotapi.NewInlineKeyboardButtonData("black", "black"),
	),
)

var endGameTemplate = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("give up", "giveUp"),
		tgbotapi.NewInlineKeyboardButtonData("draw", "draw"),
	),
)
