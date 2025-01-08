package telegram

import (
	"fmt"
	gamedto "github.com/IvaCheMih/chess/src/domains/game/dto"
	gameservice "github.com/IvaCheMih/chess/src/domains/game/services/move"
	"github.com/IvaCheMih/chess/src/domains/services/test"
	user "github.com/IvaCheMih/chess/src/domains/user/dto"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"sync"
)

type TelegramService struct {
	bot       *tgbotapi.BotAPI
	appURL    string
	authCache authCache
	figures   map[int]string
	games     gamesCache
}

type authCache struct {
	telegramIdToToken     map[int64]string
	accountIdToTelegramId map[int]int64
	mu                    sync.Mutex
}

func newAuthCache() authCache {
	return authCache{
		telegramIdToToken:     make(map[int64]string),
		accountIdToTelegramId: make(map[int]int64),
		mu:                    sync.Mutex{},
	}
}

func NewTelegramBot(token string, appURL string) (TelegramService, error) {
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
		games:     newGamesCache(),
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
				accountInfo, err := test.TelegramAuth(user.TelegramSignInRequest{
					TelegramId: update.FromChat().ID,
					ChatId:     update.FromChat().ID,
				},
					b.appURL)
				if err != nil {
					b.responseError(update.FromChat().ID, update.Message.MessageID, "Failed to sign-in", err)
				}

				b.authCache.mu.Lock()
				b.authCache.telegramIdToToken[update.FromChat().ID] = accountInfo.Token
				b.authCache.accountIdToTelegramId[accountInfo.AccountId] = update.FromChat().ID
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
			default:
				b.move(update)
			}

			continue
		}
	}
}

func (b *TelegramService) newGame(color bool, update tgbotapi.Update) {
	g, err := test.CreateGame(gamedto.CreateGameBody{
		IsWhite: color,
	},
		b.authCache.telegramIdToToken[update.FromChat().ID],
		b.appURL,
	)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to create game", err)
		return
	}

	board, err := test.GetBoard(b.authCache.telegramIdToToken[update.FromChat().ID], g.GameId, b.appURL)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to get board", err)
		return
	}

	b.addGame(update.FromChat().ID, g.GameId)

	fmt.Println("white: ", g.WhiteUserId)
	fmt.Println("black: ", g.BlackUserId)

	if color && g.BlackUserId != 0 {
		b.addOpponent(update.FromChat().ID, g.BlackUserId)

		opponentTelegramId := b.authCache.accountIdToTelegramId[g.BlackUserId]

		b.addOpponent(opponentTelegramId, g.WhiteUserId)
	}

	if !color && g.WhiteUserId != 0 {
		fmt.Println()
		fmt.Println("!color && g.WhiteUserId != 0")

		b.addOpponent(update.FromChat().ID, g.WhiteUserId)

		opponentTelegramId := b.authCache.accountIdToTelegramId[g.WhiteUserId]

		fmt.Println("opponentTelegramId: ", opponentTelegramId)
		fmt.Println("BlackUserId: ", g.BlackUserId)

		b.addOpponent(opponentTelegramId, g.BlackUserId)

		fmt.Println(b.getGame(opponentTelegramId))
		fmt.Println(b.getGame(update.FromChat().ID))
	}

	boardTemplate := b.makeBoardTemplate(board.BoardCells)

	gameOp, err := b.getGame(update.FromChat().ID)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to get game", err)
		return
	}

	if gameOp.opponentId == 0 {
		b.response(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Game created!", &boardTemplate)
		return
	}

	fmt.Println()
	fmt.Println("chat ids")
	fmt.Println("account: ", update.FromChat().ID)
	fmt.Println("opposite account: ", b.authCache.accountIdToTelegramId[gameOp.opponentId])

	b.response(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Game started", &boardTemplate)
	b.response(b.authCache.accountIdToTelegramId[gameOp.opponentId], 0, "Game started", &boardTemplate)
}

func (b *TelegramService) move(update tgbotapi.Update) {
	indexCell, err := parseData(update.CallbackQuery.Data)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to parse data", err)
	}

	g, err := b.getGame(update.FromChat().ID)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to get game", err)
	}

	from, to := b.addMove(update.FromChat().ID, indexCell)
	if from == nil || to == nil {
		// TODO: new template (to move, new figure)
		return
	}

	if *from == *to {
		b.cleanMoves(update.FromChat().ID)
		return
	}

	fmt.Println("Move: ", *from, *to)
	fmt.Println("Move: ", gameservice.IndexToCoordinates(*from), gameservice.IndexToCoordinates(*to))

	_, err = test.CreateMove(gamedto.DoMoveBody{
		From: gameservice.IndexToCoordinates(*from),
		To:   gameservice.IndexToCoordinates(*to),
		//TODO: new figure
	},
		b.authCache.telegramIdToToken[update.FromChat().ID],
		g.id,
		b.appURL,
	)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to do move", err)
		return
	}

	b.cleanMoves(update.FromChat().ID)

	board, err := test.GetBoard(b.authCache.telegramIdToToken[update.FromChat().ID], g.id, b.appURL)
	if err != nil {
		b.responseError(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "Failed to get board", err)
		return
	}

	boardTemplate := b.makeBoardTemplate(board.BoardCells)

	opponentTelegramId := b.authCache.accountIdToTelegramId[g.opponentId]

	fmt.Println()
	fmt.Println("chat ids")
	fmt.Println("account: ", update.FromChat().ID)
	fmt.Println("opposite account: ", opponentTelegramId)

	b.response(update.FromChat().ID, update.CallbackQuery.Message.MessageID, "New move", &boardTemplate)
	b.response(opponentTelegramId, 0, "New move", &boardTemplate)
}
