package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

type Figure interface {
	IsWhite() bool
	ToString() string
	GetType() byte
	GetPossibleMoves(*Game) *TheoryMoves
	ChangeGameIndex(int)
	GetGameIndex() int
	Delete()
}

func CreateDefaultField(board models.Board) map[int]*Figure {

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

	field := map[int]*Figure{}

	for i := 0; i < 144; i++ {
		fmt.Println("for: ", i)
		if realIndex, ok := VirtualFieldMap[i]; ok {
			fmt.Println("realInd: ", realIndex)
			cell, oK := board.Cells[realIndex]

			if !oK {
				continue
			}

			isWhite := cell.FigureId <= 6

			field[i] = CreateFigure(FigureRepo[cell.FigureId], isWhite, FromRealToVirtualIndex(cell.IndexCell))
		}
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

var VirtualFieldMap = map[int]int{}

func CreateVirtualFieldMap() {

	for i := 26; i < 34; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}

	for i := 38; i < 46; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}

	for i := 50; i < 58; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}

	for i := 62; i < 70; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}

	for i := 74; i < 82; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}

	for i := 86; i < 94; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}

	for i := 98; i < 106; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}

	for i := 110; i < 118; i++ {
		VirtualFieldMap[i] = FromVirtualToReal(i)
	}
}
