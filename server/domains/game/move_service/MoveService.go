package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

func CheckCorrectMove(responseGetGame dto.CreateGameResponse, board models.Board, fromReal int, toReal int) bool {

	from := FromRealToVirtualIndex(fromReal)
	to := FromRealToVirtualIndex(toReal)

	game := CreateGameStruct(responseGetGame, board)

	figure := game.GetFigureByIndex(from)

	possibleMoves := (*figure).GetPossibleMoves(&game)

	printMoves(possibleMoves)

	if !CheckMove(possibleMoves, to) {
		return false
	}

	return true
}

func CheckIsItCheck(responseGetGame dto.CreateGameResponse, board models.Board, fromReal int, toReal int) (Game, bool) {
	//cellFrom := boardCells[from]
	//cellTo := boardCells[to]

	from := FromRealToVirtualIndex(fromReal)
	to := FromRealToVirtualIndex(toReal)

	gameAfterMove := CreateGameStruct(responseGetGame, board)

	gameAfterMove.ChangeToAndFrom(to, from)

	figure := gameAfterMove.GetFigureByIndex(to)

	if figure != nil {
		gameAfterMove.ChangeKingGameID(figure)
	}

	if gameAfterMove.CheckIsCheck() {
		return Game{}, false
	}

	gameAfterMove.VirtualBoardToReal()

	return gameAfterMove, true
}

func CheckMove(possibleMoves *TheoryMoves, to int) bool {

	if possibleMoves.Up != nil {
		for _, pm := range possibleMoves.Up {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Right != nil {
		for _, pm := range possibleMoves.Right {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Left != nil {
		for _, pm := range possibleMoves.Left {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.UR != nil {
		for _, pm := range possibleMoves.UR {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.UL != nil {
		for _, pm := range possibleMoves.UL {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.DR != nil {
		for _, pm := range possibleMoves.DR {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.DL != nil {
		for _, pm := range possibleMoves.DL {
			if pm == to {
				return true
			}
		}
	}
	if possibleMoves.Kn != nil {
		for _, pm := range possibleMoves.Kn {
			if pm == to {
				return true
			}
		}
	}
	fmt.Println("Запрашиваемого хода нет в массиве")
	return false
}

func printMoves(possibleMoves *TheoryMoves) {
	for _, v := range possibleMoves.Down {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.Up {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.Left {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.Right {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.DL {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.DR {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.UR {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.UL {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}
	for _, v := range possibleMoves.Kn {
		fmt.Print(IndexToCoordinates(FromVirtualToReal(v)), " ")
	}

}

func FromRealToVirtualIndex(x int) int {
	return x + 24 + 4*(x/8) + 2
}

func FromVirtualToReal(x int) int {
	return x - 26 - 4*((x/12)-2)
}
