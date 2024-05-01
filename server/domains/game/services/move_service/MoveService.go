package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

func CheckCorrectMove(responseGetGame dto.CreateGameResponse, board models.Board, from int, to int) (bool, bool) {

	game := CreateGameStruct(responseGetGame, board)

	figure := game.GetFigureByIndex(from)

	if !game.IsItYourFigure(figure) {
		return false, false
	}

	possibleMoves := (*figure).GetPossibleMoves(&game)

	printMoves(possibleMoves)

	return CheckMove(possibleMoves, to)
}

func CheckIsItCheck(responseGetGame dto.CreateGameResponse, board models.Board, from int, to int, isCastling bool) (Game, bool) {
	//cellFrom := boardCells[from]
	//cellTo := boardCells[to]

	gameAfterMove := CreateGameStruct(responseGetGame, board)

	gameAfterMove.ChangeToAndFrom(to, from)

	if isCastling {
		gameAfterMove.ChangeRookIfCastling(to)
	}

	figure := gameAfterMove.GetFigureByIndex(to)

	if figure != nil {
		gameAfterMove.ChangeKingGameID(figure)
	}

	if gameAfterMove.CheckIsCheck() {
		return Game{}, false
	}

	gameAfterMove.ChangeCastlingFlag(figure)

	return gameAfterMove, true
}

func CheckMove(possibleMoves *TheoryMoves, to int) (bool, bool) {

	crdTo := IndexToFieldCoordinates(to)

	if possibleMoves.Up != nil {
		for _, pm := range possibleMoves.Up {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.Right != nil {
		for _, pm := range possibleMoves.Right {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.Left != nil {
		for _, pm := range possibleMoves.Left {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.UR != nil {
		for _, pm := range possibleMoves.UR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.UL != nil {
		for _, pm := range possibleMoves.UL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.DR != nil {
		for _, pm := range possibleMoves.DR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.DL != nil {
		for _, pm := range possibleMoves.DL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}
	if possibleMoves.Kn != nil {
		for _, pm := range possibleMoves.Kn {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, false
			}
		}
	}

	if possibleMoves.Castling != nil {
		for _, pm := range possibleMoves.Castling {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, true
			}
		}
	}

	fmt.Println("Запрашиваемого хода нет в массиве")
	return false, false
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
