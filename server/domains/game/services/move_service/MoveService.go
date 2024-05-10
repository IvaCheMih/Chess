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

	printMoves(possibleMoves)

	coordinatesToChange := []int{from, to}

	isTrueMove := CheckMove(possibleMoves, &coordinatesToChange)

	return isTrueMove, coordinatesToChange
}

func CheckIsItCheck(responseGetGame dto.CreateGameResponse, board models.Board, from int, to int, indexesToChange []int) (Game, bool) {
	//cellFrom := boardCells[from]
	//cellTo := boardCells[to]

	from = indexesToChange[0]
	to = indexesToChange[1]

	gameAfterMove := CreateGameStruct(responseGetGame, board)

	gameAfterMove.ChangeToAndFrom(to, from)

	if len(indexesToChange) > 2 {
		if indexesToChange[2] != -1 {
			gameAfterMove.ChangeToAndFrom(indexesToChange[2], indexesToChange[3])
		} else {
			fig := gameAfterMove.GetFigureByIndex(indexesToChange[3])
			(*fig).Delete()
		}
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

func CheckMove(possibleMoves *TheoryMoves, coordinatesToChange *[]int) bool {
	crdFrom := IndexToFieldCoordinates((*coordinatesToChange)[0])
	crdTo := IndexToFieldCoordinates((*coordinatesToChange)[1])

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

	if possibleMoves.Castling != nil {
		for _, pm := range possibleMoves.Castling {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				crdRook := []int{}
				GetNewRookCoordinatesIfCastling((*coordinatesToChange)[1], &crdRook)
				*coordinatesToChange = append(*coordinatesToChange, crdRook[0])
				*coordinatesToChange = append(*coordinatesToChange, crdRook[1])
				return true
			}
		}
	}

	if possibleMoves.EnPass != nil {
		for _, pm := range possibleMoves.EnPass {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				*coordinatesToChange = append(*coordinatesToChange, -1)
				*coordinatesToChange = append(*coordinatesToChange, FieldCoordinatesToIndex([]int{crdFrom[0], crdTo[1]}))
				return true
			}
		}
	}

	fmt.Println("Запрашиваемого хода нет в массиве")
	return false
}

func GetNewRookCoordinatesIfCastling(to int, crd *[]int) {
	switch to {
	case 2:
		*crd = append(*crd, 0)
		*crd = append(*crd, 3)
	case 6:
		*crd = append(*crd, 7)
		*crd = append(*crd, 5)
	case 57:
		*crd = append(*crd, 56)
		*crd = append(*crd, 59)
	case 62:

		*crd = append(*crd, 63)
		*crd = append(*crd, 61)
	}
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
