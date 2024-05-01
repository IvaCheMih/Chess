package move_service

import (
	"fmt"
)

type FigurePawn struct {
	BaseFigure
}

type FigureRook struct {
	BaseFigure
	Castling bool
}

type FigureKnight struct {
	BaseFigure
}

type FigureBishop struct {
	BaseFigure
}

type FigureQueen struct {
	BaseFigure
}

type FigureKing struct {
	BaseFigure
	Castling bool
}

type TheoryMoves struct {
	Up       [][]int
	Down     [][]int
	Right    [][]int
	Left     [][]int
	UR       [][]int
	UL       [][]int
	DR       [][]int
	DL       [][]int
	Kn       [][]int
	Castling [][]int
}

func (figure *FigurePawn) GetPossibleMoves(game *Game) *TheoryMoves {
	n := 0

	if figure.IsWhite() {
		n = 1
	} else {
		n = -1
	}

	vert := [][]int{}
	crd := figure.CellCoordinate

	if IsOnRealBoard([]int{crd[0], crd[1] - n}) && game.GetFigureByFieldCoordinates([]int{crd[0], crd[1] - n}) == nil {
		vert = append(vert, []int{crd[0], crd[1] - n})
	}

	if n == 1 && crd[1] == 6 {
		if game.GetFigureByFieldCoordinates([]int{crd[0], crd[1] - n}) == nil && game.GetFigureByFieldCoordinates([]int{crd[0], crd[1] - 2*n}) == nil {
			vert = append(vert, []int{crd[0], crd[1] - 2*n})
		}
	}

	if n == -1 && crd[1] == 1 {
		if game.GetFigureByFieldCoordinates([]int{crd[0], crd[1] - n}) == nil && game.GetFigureByFieldCoordinates([]int{crd[0], crd[1] - 2*n}) == nil {
			vert = append(vert, []int{crd[0], crd[1] - 2*n})
		}
	}

	left := [][]int{}

	if IsOnRealBoard([]int{crd[0] + 1, crd[1] - n}) && game.GetFigureByFieldCoordinates([]int{crd[0] + 1, crd[1] - n}) != nil {
		if figure.IsWhite() != (*game.GetFigureByFieldCoordinates([]int{crd[0] + 1, crd[1] - n})).IsWhite() {
			fmt.Println("пешка пытается кушоц налево")
			left = append(left, []int{crd[0] + 1, crd[1] - n})
		}
	}

	right := [][]int{}

	if IsOnRealBoard([]int{crd[0] - 1, crd[1] - n}) && game.GetFigureByFieldCoordinates([]int{crd[0] - 1, crd[1] - n}) != nil {
		if figure.IsWhite() != (*game.GetFigureByFieldCoordinates([]int{crd[0] - 1, crd[1] - n})).IsWhite() {
			fmt.Println("пешка пытается кушоц направо")
			right = append(right, []int{crd[0] - 1, crd[1] - n})
		}
	}

	var theoryMoves = TheoryMoves{
		Up:    vert,
		Down:  nil,
		Right: nil,
		Left:  nil,
		UR:    right,
		UL:    left,
		DR:    nil,
		DL:    nil,
		Kn:    nil,
	}

	return &theoryMoves
}

func (figure *FigureRook) GetPossibleMoves(game *Game) *TheoryMoves {
	var theoryMoves = TheoryMoves{
		Up:    [][]int{},
		Down:  [][]int{},
		Right: [][]int{},
		Left:  [][]int{},
		UR:    nil,
		UL:    nil,
		DR:    nil,
		DL:    nil,
		Kn:    nil,
	}

	crd := figure.CellCoordinate

	for index := crd[1] + 1; IsOnRealBoard([]int{crd[0], index}); index++ {

		add, _continue := figure.AddMove(game, []int{crd[0], index})

		if add {
			theoryMoves.Up = append(theoryMoves.Up, []int{crd[0], index})
		}

		if !_continue {
			break
		}
	}

	for index := crd[1] - 1; IsOnRealBoard([]int{crd[0], index}); index-- {
		add, _continue := figure.AddMove(game, []int{crd[0], index})

		if add {
			theoryMoves.Down = append(theoryMoves.Down, []int{crd[0], index})
		}

		if !_continue {
			break
		}
	}

	for index := crd[0] + 1; IsOnRealBoard([]int{index, crd[1]}); index++ {
		add, _continue := figure.AddMove(game, []int{index, crd[1]})

		if add {
			theoryMoves.Right = append(theoryMoves.Right, []int{index, crd[1]})
		}

		if !_continue {
			break
		}
	}

	for index := crd[0] - 1; IsOnRealBoard([]int{index, crd[1]}); index-- {
		add, _continue := figure.AddMove(game, []int{index, crd[1]})

		if add {
			theoryMoves.Left = append(theoryMoves.Left, []int{index, crd[1]})
		}

		if !_continue {
			break
		}
	}

	return &theoryMoves
}

func (figure *FigureKnight) GetPossibleMoves(game *Game) *TheoryMoves {
	crd := figure.CellCoordinate

	theorySteps := [][]int{
		{crd[0] + 2, crd[1] + 1},
		{crd[0] + 2, crd[1] - 1},
		{crd[0] - 2, crd[1] + 1},
		{crd[0] - 2, crd[1] - 1},
		{crd[0] - 1, crd[1] + 2},
		{crd[0] - 1, crd[1] - 2},
		{crd[0] + 1, crd[1] + 2},
		{crd[0] + 1, crd[1] - 2},
	}
	kn := [][]int{}

	for _, coordinates := range theorySteps {
		if !IsOnRealBoard(coordinates) {
			continue
		} else {
			fig := game.GetFigureByFieldCoordinates(coordinates)
			if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
				continue
			}
			kn = append(kn, coordinates)
		}
	}

	var theoryMoves = TheoryMoves{
		Up:    nil,
		Down:  nil,
		Right: nil,
		Left:  nil,
		UR:    nil,
		UL:    nil,
		DR:    nil,
		DL:    nil,
		Kn:    kn,
	}

	return &theoryMoves
}

func (figure *FigureBishop) GetPossibleMoves(game *Game) *TheoryMoves {
	crd := figure.CellCoordinate

	var theoryMoves = TheoryMoves{
		Up:    nil,
		Down:  nil,
		Right: nil,
		Left:  nil,
		UR:    [][]int{},
		UL:    [][]int{},
		DR:    [][]int{},
		DL:    [][]int{},
		Kn:    nil,
	}

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1] + i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] + i, crd[1] + i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] + i, crd[1] + i})
		}

		if !_continue {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1] + i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] - i, crd[1] + i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] - i, crd[1] + i})
		}

		if !_continue {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1] - i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] + i, crd[1] - i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] + i, crd[1] - i})
		}

		if !_continue {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1] - i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] - i, crd[1] - i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] - i, crd[1] - i})
		}

		if !_continue {
			break
		}
	}

	return &theoryMoves
}

func (figure *FigureQueen) GetPossibleMoves(game *Game) *TheoryMoves {
	var theoryMoves = TheoryMoves{
		Up:    [][]int{},
		Down:  [][]int{},
		Right: [][]int{},
		Left:  [][]int{},
		UR:    [][]int{},
		UL:    [][]int{},
		DR:    [][]int{},
		DL:    [][]int{},
		Kn:    nil,
	}

	crd := figure.CellCoordinate

	for index := crd[1] + 1; IsOnRealBoard([]int{crd[0], index}); index++ {

		add, _continue := figure.AddMove(game, []int{crd[0], index})

		if add {
			theoryMoves.Up = append(theoryMoves.Up, []int{crd[0], index})
		}

		if !_continue {
			break
		}
	}

	for index := crd[1] - 1; IsOnRealBoard([]int{crd[0], index}); index-- {
		add, _continue := figure.AddMove(game, []int{crd[0], index})

		if add {
			theoryMoves.Down = append(theoryMoves.Down, []int{crd[0], index})
		}

		if !_continue {
			break
		}
	}

	for index := crd[0] + 1; IsOnRealBoard([]int{index, crd[1]}); index++ {
		add, _continue := figure.AddMove(game, []int{index, crd[1]})

		if add {
			theoryMoves.Right = append(theoryMoves.Right, []int{index, crd[1]})
		}

		if !_continue {
			break
		}
	}

	for index := crd[0] - 1; IsOnRealBoard([]int{index, crd[1]}); index-- {
		add, _continue := figure.AddMove(game, []int{index, crd[1]})

		if add {
			theoryMoves.Left = append(theoryMoves.Left, []int{index, crd[1]})
		}

		if !_continue {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1] + i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] + i, crd[1] + i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] + i, crd[1] + i})
		}

		if !_continue {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1] + i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] - i, crd[1] + i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] - i, crd[1] + i})
		}

		if !_continue {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1] - i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] + i, crd[1] - i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] + i, crd[1] - i})
		}

		if !_continue {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1] - i}); i++ {
		add, _continue := figure.AddMove(game, []int{crd[0] - i, crd[1] - i})

		if add {
			theoryMoves.UR = append(theoryMoves.UR, []int{crd[0] - i, crd[1] - i})
		}

		if !_continue {
			break
		}
	}

	return &theoryMoves
}

func (figure *FigureKing) GetPossibleMoves(game *Game) *TheoryMoves {
	crd := figure.CellCoordinate

	theorySteps := GetTheorySteps(crd)

	k := [][]int{}

	for _, move := range theorySteps {
		if IsOnRealBoard(move) && figure.AddMove(game, move) {
			for _, move1 := range GetTheorySteps(move) {
				if move1[0] != move[0] && move1[1] != move[1] &&
					(*game.GetFigureByFieldCoordinates(move1)).GetType() != 'K' &&
					(*game.GetFigureByFieldCoordinates(move1)).GetType() != 'k' {
					k = append(k, move)
				}
			}
		}
	}

	castling := [][]int{}

	if !figure.Castling {
		if figure.IsWhite() && crd[0] == 4 && crd[1] == 7 {
			rookA := game.GetFigureByFieldCoordinates([]int{0, 7})
			if rookA != nil && (*rookA).IsWhite() && (*rookA).GetType() == 'a' && !game.WhiteCastling.WhiteRookACastling {
				if !game.IsKingCheck(60) &&
					!game.IsKingCheck(59) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(59)) == nil &&
					!game.IsKingCheck(58) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(58)) == nil &&
					!game.IsKingCheck(57) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(57)) == nil {
					castling = append(castling, []int{2, 7})
				}

			}

			rookH := game.GetFigureByFieldCoordinates([]int{7, 7})
			if rookH != nil && (*rookH).IsWhite() && (*rookH).GetType() == 'h' && !game.WhiteCastling.WhiteRookHCastling {
				if !game.IsKingCheck(60) &&
					!game.IsKingCheck(61) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(61)) == nil &&
					!game.IsKingCheck(62) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(62)) == nil {
					castling = append(castling, []int{6, 7})
				}
			}
		}

		if !figure.IsWhite() && crd[0] == 4 && crd[1] == 0 {
			rookA := game.GetFigureByFieldCoordinates([]int{0, 0})
			if rookA != nil && !(*rookA).IsWhite() && (*rookA).GetType() == 'a' && !game.BlackCastling.BlackRookACastling {
				if !game.IsKingCheck(4) &&
					!game.IsKingCheck(3) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(3)) == nil &&
					!game.IsKingCheck(2) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(2)) == nil &&
					!game.IsKingCheck(1) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(1)) == nil {
					castling = append(castling, []int{2, 0})
				}
			}

			rookH := game.GetFigureByFieldCoordinates([]int{7, 0})
			if rookH != nil && !(*rookH).IsWhite() && (*rookH).GetType() == 'h' && !game.BlackCastling.BlackRookHCastling {
				if !game.IsKingCheck(4) &&
					!game.IsKingCheck(5) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(5)) == nil &&
					!game.IsKingCheck(6) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(6)) == nil {
					castling = append(castling, []int{6, 0})
				}
			}
		}
	}

	var theoryMoves = TheoryMoves{
		Up:       nil,
		Down:     nil,
		Right:    nil,
		Left:     nil,
		UR:       nil,
		UL:       nil,
		DR:       nil,
		DL:       nil,
		Kn:       k,
		Castling: castling,
	}

	return &theoryMoves
}

func (figure *FigureRook) ToString() string {
	if figure.IsWhite() {
		return "R"
	}
	return "r"
}

func (figure *FigureKnight) ToString() string {
	if figure.IsWhite() {
		return "H"
	}
	return "h"
}

func (figure *FigureBishop) ToString() string {
	if figure.IsWhite() {
		return "B"
	}
	return "b"
}

func (figure *FigureQueen) ToString() string {
	if figure.IsWhite() {
		return "Q"
	}
	return "q"
}

func (figure *FigureKing) ToString() string {
	if figure.IsWhite() {
		return "K"
	}
	return "k"
}

func (figure *FigurePawn) ToString() string {
	if figure.IsWhite() {
		return "P"
	}
	return "p"
}

func (figure *FigureRook) AddMove(game *Game, crd []int) (bool, bool) {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false, false
	}

	if fig != nil && (*fig).IsWhite() != (*figure).IsWhite() {
		return true, false
	}

	return true, true
}

func (figure *FigureBishop) AddMove(game *Game, crd []int) (bool, bool) {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false, false
	}

	if fig != nil && (*fig).IsWhite() != (*figure).IsWhite() {
		return true, false
	}

	return true, true
}

func (figure *FigureQueen) AddMove(game *Game, crd []int) (bool, bool) {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false, false
	}

	if fig != nil && (*fig).IsWhite() != (*figure).IsWhite() {
		return true, false
	}

	return true, true
}

func (figure *FigureKing) AddMove(game *Game, crd []int) bool {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false
	}
	return true
}

func IsOnRealBoard(coordinates []int) bool {
	if coordinates[0] < 0 || 7 < coordinates[0] {
		return false
	}

	if coordinates[1] < 0 || 7 < coordinates[1] {
		return false
	}

	return true
}

func GetTheorySteps(crd []int) [][]int {
	return [][]int{
		{crd[0] + 1, crd[1] + 1},
		{crd[0], crd[1] + 1},
		{crd[0] - 1, crd[1] + 1},

		{crd[0] + 1, crd[1]},
		{crd[0], crd[1]},
		{crd[0] - 1, crd[1]},

		{crd[0] + 1, crd[1] - 1},
		{crd[0], crd[1] - 1},
		{crd[0] - 1, crd[1] - 1},
	}
}

func (figure *FigureKing) GetCastling() bool {
	return figure.Castling
}

func (figure *FigureRook) GetCastling() bool {
	return figure.Castling
}

func (figure *FigureQueen) GetCastling() bool {
	return true
}

func (figure *FigureKnight) GetCastling() bool {
	return true
}

func (figure *FigurePawn) GetCastling() bool {
	return true
}

func (figure *FigureBishop) GetCastling() bool {
	return true
}