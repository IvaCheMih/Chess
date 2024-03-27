package main

import (
	"fmt"
)

func (game *Game) CheckCorrectRequest(source string) bool {
	f, t := ParseMessageToMove(source)
	from, to := game.CoordinatesToIndex(f), game.CoordinatesToIndex(t)

	if !game.CheckCellOnBoardByIndex(from) || !game.CheckCellOnBoardByIndex(to) {
		return false
	}

	figure := game.GetFigureByIndex(from)

	possibleMoves := (*figure).GetPossibleMoves(game)
	fmt.Println("Получили массив возможных ходов для запрашиваемой фигуры")

	return CheckMove(possibleMoves, to)
}

func (game *Game) CheckCellOnBoardByIndex(index int) bool {
	coordinates := game.IndexToCoordinates(index)
	if coordinates[0] >= byte('A') && coordinates[0] <= byte('H') {
		if coordinates[1] >= byte('1') && coordinates[1] <= byte('8') {
			return true
		}
	}
	return false
}

func ParseMessageToMove(message string) (string, string) {
	return message[0:2], message[3:]
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
