package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

func CheckCorrectMove(responseGetGame dto.ResponseGetGame, boardCells []models.BoardCell, requestFromTo dto.RequestDoMove) bool {
	game := CreateGameStruct(responseGetGame, boardCells)

	figure := game.GetFigureByIndex(game.CoordinatesToIndex(requestFromTo.From))

	possibleMoves := (*figure).GetPossibleMoves(&game)

	if !CheckMove(possibleMoves, game.CoordinatesToIndex(requestFromTo.To)) {
		return false
	}
	return true
}

func CheckIsItCheck(responseGetGame dto.ResponseGetGame, boardCells []models.BoardCell, from int, to int) (Game, bool) {
	//cellFrom := boardCells[from]
	//cellTo := boardCells[to]

	boardCells[to].FigureId = boardCells[from].FigureId

	gameAfterMove := CreateGameStruct(responseGetGame, boardCells)

	figure := gameAfterMove.GetFigureByIndex(to)
	(*figure).ChangeGameIndex(to)

	gameAfterMove.ChangeKingGameID(figure)

	if gameAfterMove.CheckIsCheck() {
		return Game{}, false
	}
	return gameAfterMove, true
}

func CheckMove(possibleMoves *TheoryMoves, to int) bool {

	if possibleMoves.Up != nil {
		for _, pm := range possibleMoves.Up {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Right != nil {
		for _, pm := range possibleMoves.Right {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Left != nil {
		for _, pm := range possibleMoves.Left {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.UR != nil {
		for _, pm := range possibleMoves.UR {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.UL != nil {
		for _, pm := range possibleMoves.UL {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.DR != nil {
		for _, pm := range possibleMoves.DR {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.DL != nil {
		for _, pm := range possibleMoves.DL {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Kn != nil {
		for _, pm := range possibleMoves.Kn {
			fmt.Print(" ", pm)
			if pm == to {
				return true
			}
		}
	}
	fmt.Println("Запрашиваемого хода нет в массиве")
	return false
}
