package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

func CheckCorrectMove(responseGetGame dto.CreateGameResponse, board models.Board, requestFromTo dto.DoMoveRequest) bool {
	fmt.Println(300)
	game := CreateGameStruct(responseGetGame, board)
	fmt.Println(301)

	figure := game.GetFigureByIndex(game.CoordinatesToIndex(requestFromTo.From))
	fmt.Println(302)

	possibleMoves := (*figure).GetPossibleMoves(&game)
	fmt.Println(303)

	if !CheckMove(possibleMoves, game.CoordinatesToIndex(requestFromTo.To)) {
		return false
	}
	fmt.Println(304)

	return true
}

func CheckIsItCheck(responseGetGame dto.CreateGameResponse, board models.Board, from int, to int) (Game, bool) {
	//cellFrom := boardCells[from]
	//cellTo := boardCells[to]

	fmt.Println(400)

	gameAfterMove := CreateGameStruct(responseGetGame, board)

	for i, fig := range gameAfterMove.Figures {
		if fig != nil {
			fmt.Println(" ", i, ":", (*fig).GetType())
		} else {
			fmt.Println(" ", i, ": нет ")
		}
	}

	fmt.Println(401)

	gameAfterMove.ChangeToAndFrom(to, from)

	fmt.Println(402)
	figure := gameAfterMove.GetFigureByIndex(to)
	fmt.Println(4022)
	if figure != nil {
		gameAfterMove.ChangeKingGameID(figure)
	}

	fmt.Println(403)

	if gameAfterMove.CheckIsCheck() {
		return Game{}, false
	}
	fmt.Println(404)
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
