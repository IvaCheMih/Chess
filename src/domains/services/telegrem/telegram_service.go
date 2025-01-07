package telegrem

import (
	"fmt"
	gamedto "github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/services/test"
	user "github.com/IvaCheMih/chess/src/domains/user/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
)

type TelegramService struct {
	bot       *tgbotapi.BotAPI
	appURL    string
	authCache AuthCache
	figures   map[int]string
}

type AuthCache struct {
	accountIdToToken map[int64]string
	mu               sync.Mutex
}

func newAuthCache() AuthCache {
	return AuthCache{
		accountIdToToken: make(map[int64]string),
		mu:               sync.Mutex{},
	}
}

func NewTelegramBot(token string, appURL string) (TelegramService, error) {
	// token = "7758105509:AAGIrigjvjV4-kKMd3XkP7uZH5ABudcxFYA"
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return TelegramService{}, err
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	return TelegramService{
		bot:       bot,
		appURL:    appURL,
		authCache: newAuthCache(),
		figures:   newFigures(),
	}, nil
}

func (b *TelegramService) StartBot() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			switch update.Message.Command() {
			case "start":
				token, err := test.TelegramAuth(user.TelegramSignInRequest{
					TelegramId: update.FromChat().ID,
					ChatId:     update.FromChat().ID,
				},
					b.appURL)
				if err != nil {
					b.responseError(update.FromChat().ID, update.Message.MessageID, "Failed to sign-in", err)
				}

				fmt.Println()
				fmt.Println(update.FromChat().ID)
				fmt.Println(token.Token)

				b.authCache.mu.Lock()
				b.authCache.accountIdToToken[update.FromChat().ID] = token.Token
				b.authCache.mu.Unlock()

				b.response(update.FromChat().ID, update.Message.MessageID, "Lets start!", &startTemplate)
				continue
			}
		}

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "new":
				b.response(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Choose color: ", &newGameTemplate)
			case "white":
				b.newGame(true, update)
			case "black":
				b.newGame(false, update)
			}
			continue
		}
	}
}

func (b *TelegramService) newGame(color bool, update tgbotapi.Update) {
	fmt.Println()
	fmt.Println(update.FromChat().ID)
	fmt.Println(b.authCache.accountIdToToken[update.FromChat().ID])
	game, err := test.CreateGame(gamedto.CreateGameBody{
		IsWhite: color,
	},
		b.authCache.accountIdToToken[update.FromChat().ID],
		b.appURL,
	)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to create game", err)
		return
	}

	board, err := test.GetBoard(b.authCache.accountIdToToken[update.FromChat().ID], game.GameId, b.appURL)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to get board", err)
		return
	}

	boardTemplate := b.makeBoardTemplate(board)

	b.response(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Game created!", &boardTemplate)
}
