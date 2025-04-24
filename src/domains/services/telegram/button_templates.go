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

func addEndgameButton(template tgbotapi.InlineKeyboardMarkup) tgbotapi.InlineKeyboardMarkup {
	endgame := tgbotapi.NewInlineKeyboardButtonData("End game", "endgame")
	rowEndGame := tgbotapi.NewInlineKeyboardRow(endgame)

	template.InlineKeyboard = append(template.InlineKeyboard, rowEndGame)
	return template
}

func addCancelButtonButton(template tgbotapi.InlineKeyboardMarkup) tgbotapi.InlineKeyboardMarkup {
	cancel := tgbotapi.NewInlineKeyboardButtonData("Cancel game", "cancel")
	rowCancelGame := tgbotapi.NewInlineKeyboardRow(cancel)

	template.InlineKeyboard = append(template.InlineKeyboard, rowCancelGame)
	return template
}
