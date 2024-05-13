package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

func CheckCorrectMove(responseGetGame dto.CreateGameResponse, board models.Board, from int, to int) (bool, []int) {

	game := CreateGameStruct(responseGetGame, board)

	figure := game.GetFigureByIndex(from)

	if !game.IsItYourFigure(figure) {
		return false, []int{}
	}

	possibleMoves := (*figure).GetPossibleMoves(&game)

	coordinatesToChange := []int{from, to}

	return CheckMove(possibleMoves, coordinatesToChange)
}

func CheckIsItCheck(responseGetGame dto.CreateGameResponse, board models.Board, indexesToChange []int) (Game, bool) {
	from := indexesToChange[0]
	to := indexesToChange[1]

	gameAfterMove := CreateGameStruct(responseGetGame, board)

	gameAfterMove.ChangeToAndFrom(to, from)

	figure := gameAfterMove.GetFigureByIndex(to)

	if figure != nil {
		gameAfterMove.ChangeKingGameID(figure)
	}

	if gameAfterMove.CheckIsCheck() {
		return Game{}, false
	}

	gameAfterMove.ChangeCastlingFlag(figure)

	gameAfterMove.ChangeLastPawnMove(figure, from, to)

	return gameAfterMove, true
}

func CheckMove(possibleMoves *TheoryMoves, coordinatesToChange []int) (bool, []int) {
	crdFrom := IndexToFieldCoordinates((coordinatesToChange)[0])
	crdTo := IndexToFieldCoordinates((coordinatesToChange)[1])

	if possibleMoves.Up != nil {
		for _, pm := range possibleMoves.Up {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Right != nil {
		for _, pm := range possibleMoves.Right {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Left != nil {
		for _, pm := range possibleMoves.Left {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.UR != nil {
		for _, pm := range possibleMoves.UR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.UL != nil {
		for _, pm := range possibleMoves.UL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.DR != nil {
		for _, pm := range possibleMoves.DR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.DL != nil {
		for _, pm := range possibleMoves.DL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Kn != nil {
		for _, pm := range possibleMoves.Kn {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}

	if possibleMoves.Castling != nil {
		for _, pm := range possibleMoves.Castling {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				crdRook := GetNewRookCoordinatesIfCastling((coordinatesToChange)[1])
				coordinatesToChange = append(coordinatesToChange, crdRook[0])
				coordinatesToChange = append(coordinatesToChange, crdRook[1])
				return true, coordinatesToChange
			}
		}
	}

	if possibleMoves.EnPass != nil {
		for _, pm := range possibleMoves.EnPass {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {

				coordinatesToChange = append(coordinatesToChange, -1)
				coordinatesToChange = append(coordinatesToChange, FieldCoordinatesToIndex([]int{crdTo[0], crdFrom[1]}))

				return true, coordinatesToChange
			}
		}
	}

	fmt.Println("Запрашиваемого хода нет в массиве")
	return false, []int{}
}

func GetNewRookCoordinatesIfCastling(to int) []int {
	crd := []int{}

	switch to {
	case 2:
		crd = append(crd, 0)
		crd = append(crd, 3)
	case 6:
		crd = append(crd, 7)
		crd = append(crd, 5)
	case 57:
		crd = append(crd, 56)
		crd = append(crd, 59)
	case 62:

		crd = append(crd, 63)
		crd = append(crd, 61)
	}

	return crd
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

	for _, v := range possibleMoves.Castling {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}

	for _, v := range possibleMoves.EnPass {
		fmt.Print(IndexToCoordinates(FieldCoordinatesToIndex(v)), " ")
	}

}
