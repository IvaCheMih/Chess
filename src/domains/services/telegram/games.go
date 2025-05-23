package telegram

import (
	"errors"
	"sync"
)

type game struct {
	id         int
	from       *int
	to         *int
	opponentId int
}

type gamesCache struct {
	accountIdToGameId map[int64]*game
	mu                sync.Mutex
}

func newGamesCache() gamesCache {
	return gamesCache{
		accountIdToGameId: make(map[int64]*game),
		mu:                sync.Mutex{},
	}
}

func (b *TelegramService) addGame(telegramId int64, gameId int) {
	b.games.mu.Lock()
	defer b.games.mu.Unlock()

	if _, ok := b.games.accountIdToGameId[telegramId]; !ok {
		b.games.accountIdToGameId[telegramId] = &game{
			id: gameId,
		}
	}
}

func (b *TelegramService) addOpponent(telegramId int64, opponentId int) {
	b.games.mu.Lock()
	defer b.games.mu.Unlock()

	if _, ok := b.games.accountIdToGameId[telegramId]; ok {
		id := b.games.accountIdToGameId[telegramId].id

		b.games.accountIdToGameId[telegramId] = &game{
			id:         id,
			opponentId: opponentId,
		}
	}
}

func (b *TelegramService) getGame(telegramId int64) (game, error) {
	b.games.mu.Lock()
	defer b.games.mu.Unlock()

	if g, ok := b.games.accountIdToGameId[telegramId]; ok {
		return *g, nil
	}

	return game{}, errors.New("game not exists")
}

// TODO: endgame
//func (b *TelegramService) removeGame(telegramId int64) {
//	b.games.mu.Lock()
//	defer b.games.mu.Unlock()
//
//	if _, ok := b.games.accountIdToGameId[telegramId]; ok {
//		delete(b.games.accountIdToGameId, telegramId)
//	}
//}

// TODO: add new figure
func (b *TelegramService) addMove(telegramId int64, index int) (*int, *int) {
	b.games.mu.Lock()
	defer b.games.mu.Unlock()

	if _, ok := b.games.accountIdToGameId[telegramId]; ok {
		gameId := b.games.accountIdToGameId[telegramId].id
		opponentId := b.games.accountIdToGameId[telegramId].opponentId

		if b.games.accountIdToGameId[telegramId].from == nil {
			b.games.accountIdToGameId[telegramId] = &game{
				id:         gameId,
				from:       &index,
				opponentId: opponentId,
			}
		} else {
			from := *b.games.accountIdToGameId[telegramId].from
			b.games.accountIdToGameId[telegramId] = &game{
				id:         gameId,
				from:       &from,
				to:         &index,
				opponentId: opponentId,
			}

			return &from, &index
		}
	}

	return nil, nil
}

func (b *TelegramService) cleanMoves(telegramId int64) {
	b.games.mu.Lock()
	defer b.games.mu.Unlock()

	if _, ok := b.games.accountIdToGameId[telegramId]; ok {
		gameId := b.games.accountIdToGameId[telegramId].id
		opponentId := b.games.accountIdToGameId[telegramId].opponentId

		b.games.accountIdToGameId[telegramId] = &game{
			id:         gameId,
			opponentId: opponentId,
		}
	}
}
