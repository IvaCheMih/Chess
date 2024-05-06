package move_service

import (
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

type Figure interface {
	IsWhite() bool
	ToString() string
	GetType() byte
	GetPossibleMoves(*Game) *TheoryMoves
	ChangeGameIndex([]int)
	GetGameIndex() []int
	Delete()
	GetCastling() bool
}

func CreateField(board models.Board, game dto.CreateGameResponse) (map[int]*Figure, int, int) {
	blackKingCell, whiteKingCell := 0, 0
	field := map[int]*Figure{}

	for _, cell := range board.Cells {
		if cell.FigureId == 5 {
			whiteKingCell = cell.IndexCell

		}
		if cell.FigureId == 12 {
			blackKingCell = cell.IndexCell

		}

		isWhite := cell.FigureId <= 7

		field[cell.IndexCell] = CreateFigure(FigureRepo[cell.FigureId], isWhite, cell.IndexCell, game)

	}

	return field, blackKingCell, whiteKingCell
}

func FigureToString(figure *Figure) string {
	if figure == nil {
		return "0"
	}

	return (*figure).ToString()
}

func CreateFigure(_type byte, isWhite bool, index int, game dto.CreateGameResponse) *Figure {

	coordinates := IndexToFieldCoordinates(index)

	figure := CreateFigure1(_type, isWhite, coordinates, game)

	if figure == nil {
		return nil
	}

	return &figure
}

func CreateFigure1(_type byte, isWhite bool, coordinates []int, game dto.CreateGameResponse) Figure {
	var bf = BaseFigure{isWhite, _type, coordinates}
	//var tm = TheoryMoves{nil, nil, nil, nil, nil, nil, nil, nil, nil}
	switch _type {
	case 'p':
		return &FigurePawn{bf}
	case 'a':

		castling := false
		if isWhite {
			if game.WhiteRookACastling {
				castling = true
			}
		} else {
			if game.BlackRookACastling {
				castling = true
			}
		}

		return &FigureRook{bf, castling}
	case 'h':
		castling := false
		if isWhite {
			if game.WhiteRookHCastling {
				castling = true
			}
		} else {
			if game.BlackRookHCastling {
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
			if game.WhiteKingCastling {
				castling = true
			}
		} else {
			if game.BlackKingCastling {
				castling = true
			}
		}
		return &FigureKing{bf, castling}
	}
	return nil
}

func IndexToFieldCoordinates(ind int) []int {
	x := ind % 8
	y := ind / 8

	return []int{x, y}
}

func FieldCoordinatesToIndex(coordinates []int) int {
	return coordinates[1]*8 + coordinates[0]
}
