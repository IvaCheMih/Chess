package move_service

import "github.com/IvaCheMih/chess/server/domains/game/dto"

type Game struct {
	N             int
	WhiteClientId *int
	BlackClientId *int
	Figures       []*Figure
	IsCheckWhite  IsCheck
	IsCheckBlack  IsCheck
}

type IsCheck struct {
	IsItCheck  bool
	KingGameID int
}

func CreateGameStruct(game dto.ResponseGetGame, cells []dto.BoardCell, lastMove dto.Move) Game {

	return Game{
		N:             8,
		WhiteClientId: &game.WhiteUserId,
		BlackClientId: &game.BlackUserId,
		Figures:       CreateDefaultField(cells),
		IsCheckWhite:  IsCheck{lastMove.IsCheckWhite, lastMove.WhiteKingCell},
		IsCheckBlack:  IsCheck{lastMove.IsCheckBlack, lastMove.BlackKingCell},
	}
}
