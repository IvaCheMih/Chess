package move

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
)

type Figure interface {
	IsItWhite() bool
	GetType() byte
	GetPossibleMoves(*Game) *TheoryMoves
	ChangeCoordinates([2]int)
	GetCoordinates() [2]int
	ChangeType(byte)
	Delete()
}

func createFigureI(_type byte, isWhite bool, index int, gameModel models.Game) Figure {
	coordinates := IndexToFieldCoordinates(index)

	figure := CreateFigure(_type, isWhite, coordinates, gameModel)

	if figure == nil {
		return nil
	}

	return figure
}

func CreateFigure(_type byte, isWhite bool, coordinates [2]int, gameModel models.Game) Figure {
	var bf = BaseFigure{isWhite, _type, coordinates}

	switch _type {
	case 'p':
		return &FigurePawn{bf}
	case 'a':

		castling := false
		if isWhite {
			if gameModel.WhiteRookACastling {
				castling = true
			}
		} else {
			if gameModel.BlackRookACastling {
				castling = true
			}
		}

		return &FigureRook{bf, castling}
	case 'h':
		castling := false
		if isWhite {
			if gameModel.WhiteRookHCastling {
				castling = true
			}
		} else {
			if gameModel.BlackRookHCastling {
				castling = true
			}
		}
		return &FigureRook{bf, castling}
	case 'k':
		return &FigureKnight{bf}
	case 'b':
		return &FigureBishop{bf}
	case 'q':
		return &FigureQueen{bf}
	case 'K':
		castling := false
		if isWhite {
			if gameModel.WhiteKingCastling {
				castling = true
			}
		} else {
			if gameModel.BlackKingCastling {
				castling = true
			}
		}
		return &FigureKing{bf, castling}
	}
	return nil
}

func IndexToFieldCoordinates(ind int) [2]int {
	x := ind % 8
	y := ind / 8

	return [2]int{x, y}
}

func FieldCoordinatesToIndex(coordinates [2]int) int {
	return coordinates[1]*8 + coordinates[0]
}
