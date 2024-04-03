package move_service

import (
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

type Figure interface {
	IsWhite() bool
	ToString() string
	GetType() byte
	GetPossibleMoves(*Game) *TheoryMoves
	ChangeGameIndex(int)
	GetGameIndex() int
}

func CreateDefaultField(cells []models.BoardCell) []*Figure {
	//startField := []byte{
	//	'r', 'k', 'b', 'q', 'K', 'b', 'k', 'r',
	//	'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p',
	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
	//	' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ',
	//	'p', 'p', 'p', 'p', 'p', 'p', 'p', 'p',
	//	'r', 'k', 'b', 'q', 'K', 'b', 'k', 'r',
	//}
	field := []*Figure{}

	for _, cell := range cells {
		isWhite := false
		if cell.FigureId <= 6 {
			isWhite = true
		}

		field = append(field, CreateFigure(FigureRepo[cell.FigureId], isWhite, cell.IndexCell))
	}

	return field
}

func FigureToString(figure *Figure) string {
	if figure == nil {
		return "0"
	}

	return (*figure).ToString()
}

func CreateFigure(_type byte, isWhite bool, index int) *Figure {
	figure := CreateFigure1(_type, isWhite, index)

	if figure == nil {
		return nil
	}

	return &figure
}

func CreateFigure1(_type byte, isWhite bool, index int) Figure {
	var bf = BaseFigure{isWhite, _type, index}
	//var tm = TheoryMoves{nil, nil, nil, nil, nil, nil, nil, nil, nil}
	switch _type {
	case 'p':
		return &FigurePawn{bf}
	case 'r':
		return &FigureRook{bf}
	case 'k':
		return &FigureKnight{bf}
	case 'b':
		return &FigureBishop{bf}
	case 'q':
		return &FigureQueen{bf}
	case 'K':
		return &FigureKing{bf}
	}
	return nil
}