package game

import (
	"errors"
	"github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"gorm.io/gorm"
)

func checkCorrectRequestSideUser(userId int, game models.Game) error {
	if !game.IsActive() {
		return errors.New("game is not active")
	}

	if game.Side && userId != game.WhiteUserId {
		return errors.New("its not your move now")
	}

	if !game.Side && userId != game.BlackUserId {
		return errors.New("its not your move now")
	}
	return nil
}

func indexToCoordinates(index int) string {
	y := int('8') - (index / 8)
	x := (index % 8) + int('A')

	return string(byte(x)) + string(byte(y))
}

func coordinatesToIndex(coordinates string) int {
	x := int(coordinates[0]) - int('A')
	y := int('8') - int(coordinates[1])

	return (y * 8) + x
}

func checkCellOnBoardByIndex(index int) bool {
	coordinates := indexToCoordinates(index)
	if coordinates[0] >= byte('A') && coordinates[0] <= byte('H') {
		if coordinates[1] >= byte('1') && coordinates[1] <= byte('8') {
			return true
		}
	}
	return false
}

func checkCorrectRequest(move dto.DoMoveBody) bool {
	from, to := coordinatesToIndex(move.From), coordinatesToIndex(move.To)

	if !checkCellOnBoardByIndex(from) || !checkCellOnBoardByIndex(to) {
		return false
	}

	return true
}

func updateBoardAfterMove(tx *gorm.DB, g *GamesService, board models.Board, newFigureId int, indexesToChange []int) error {
	var err error
	from := indexesToChange[0]
	to := indexesToChange[1]

	if board.Cells[to] != nil {
		err = g.boardRepo.Delete(tx, board.Cells[to].Id)
		if err != nil {
			return err
		}
	}

	if newFigureId != 0 {
		err = g.boardRepo.UpdateNewFigure(tx, board.Cells[from].Id, to, newFigureId)
	} else {
		err = g.boardRepo.Update(tx, board.Cells[from].Id, to)
	}

	if err != nil {
		return err
	}

	if len(indexesToChange) > 2 {
		if indexesToChange[2] == -1 {
			err = g.boardRepo.Delete(tx, board.Cells[indexesToChange[3]].Id)
		} else {
			err = g.boardRepo.Update(tx, board.Cells[indexesToChange[2]].Id, indexesToChange[3])
		}
	}

	return err
}

func fromModelsToDtoCreateGame(response models.Game) dto.CreateGameResponse {
	return dto.CreateGameResponse{
		GameId:             response.Id,
		IsCheckWhite:       response.IsCheckWhite,
		IsCheckBlack:       response.IsCheckBlack,
		Status:             response.Status.ToString(),
		EndReason:          response.EndReason.ToString(),
		WhiteKingCastling:  response.WhiteKingCastling,
		BlackKingCastling:  response.BlackKingCastling,
		WhiteRookACastling: response.WhiteRookACastling,
		WhiteRookHCastling: response.WhiteRookHCastling,
		BlackRookACastling: response.BlackRookACastling,
		BlackRookHCastling: response.BlackRookHCastling,
		BlackUserId:        response.BlackUserId,
		WhiteUserId:        response.WhiteUserId,
		LastPawnMove:       response.LastPawnMove,
		Side:               response.Side,
	}

}

func newGame() models.Game {
	return models.Game{
		EndReason: models.NotEndgame,
		Status:    models.Created,
		Side:      true,
	}
}
