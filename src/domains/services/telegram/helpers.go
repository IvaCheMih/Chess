package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func (b *TelegramService) responseError(chatId int64, messageID int, text string, err error) {
	log.Println(err)
	msg := tgbotapi.NewMessage(chatId, text)
	msg.ReplyToMessageID = messageID

	_, err = b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}

func (b *TelegramService) response(chatId int64, messageID int, text string, template *tgbotapi.InlineKeyboardMarkup) {
	msg := tgbotapi.NewMessage(chatId, text)
	if messageID != 0 {
		msg.ReplyToMessageID = messageID
	}

	if template != nil {
		msg.ReplyMarkup = *template
	}

	_, err := b.bot.Send(msg)
	if err != nil {
		log.Println(err)
	}
}
