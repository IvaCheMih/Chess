package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

func CheckCorrectMove(responseGetGame dto.CreateGameResponse, board models.Board, from int, to int) bool {

	game := CreateGameStruct(responseGetGame, board)

	figure := game.GetFigureByIndex(from)

	possibleMoves := (*figure).GetPossibleMoves(&game)

	printMoves(possibleMoves)

	if !CheckMove(possibleMoves, to) {
		return false
	}

	return true
}

func CheckIsItCheck(responseGetGame dto.CreateGameResponse, board models.Board, from int, to int) (Game, bool) {
	//cellFrom := boardCells[from]
	//cellTo := boardCells[to]

	gameAfterMove := CreateGameStruct(responseGetGame, board)

	gameAfterMove.ChangeToAndFrom(to, from)

	figure := gameAfterMove.GetFigureByIndex(to)

	if figure != nil {
		gameAfterMove.ChangeKingGameID(figure)
	}

	if gameAfterMove.CheckIsCheck() {
		return Game{}, false
	}

	//gameAfterMove.VirtualBoardToReal()

	return gameAfterMove, true
}

func CheckMove(possibleMoves *TheoryMoves, to int) bool {

	crdTo := IndexToFieldCoordinates(to)

	if possibleMoves.Up != nil {
		for _, pm := range possibleMoves.Up {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.Right != nil {
		for _, pm := range possibleMoves.Right {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.Left != nil {
		for _, pm := range possibleMoves.Left {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.UR != nil {
		for _, pm := range possibleMoves.UR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.UL != nil {
		for _, pm := range possibleMoves.UL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.DR != nil {
		for _, pm := range possibleMoves.DR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.DL != nil {
		for _, pm := range possibleMoves.DL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	if possibleMoves.Kn != nil {
		for _, pm := range possibleMoves.Kn {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true
			}
		}
	}
	fmt.Println("Запрашиваемого хода нет в массиве")
	return false
}

func printMoves(possibleMoves *TheoryMoves) {
	for _, v := range possibleMoves.Down {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.Up {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.Left {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.Right {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.DL {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.DR {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.UR {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.UL {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}
	for _, v := range possibleMoves.Kn {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}

}
