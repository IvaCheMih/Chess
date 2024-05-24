package move_service

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
	EnPass   [][]int
	Castling [][]int
}

func (figure *FigurePawn) GetPossibleMoves(game *Game) *TheoryMoves {
	n := 0

	if figure.IsItWhite() {
		n = 1
	} else {
		n = -1
	}

	vert := [][]int{}
	EnPass := [][]int{}
	crdLastPawnMove := []int{}

	crd := figure.CellCoordinates

	if game.LastPawnMove != nil {

		crdLastPawnMove = IndexToFieldCoordinates(*game.LastPawnMove)

		if crdLastPawnMove[0] == crd[0]+1 || crdLastPawnMove[0] == crd[0]-1 {
			if crd[1] == crdLastPawnMove[1] {
				EnPass = append(EnPass, []int{crdLastPawnMove[0], crdLastPawnMove[1] - n})
			}
		}
	}

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
		if figure.IsItWhite() != (*game.GetFigureByFieldCoordinates([]int{crd[0] + 1, crd[1] - n})).IsItWhite() {
			left = append(left, []int{crd[0] + 1, crd[1] - n})
		}
	}

	right := [][]int{}

	if IsOnRealBoard([]int{crd[0] - 1, crd[1] - n}) && game.GetFigureByFieldCoordinates([]int{crd[0] - 1, crd[1] - n}) != nil {
		if figure.IsItWhite() != (*game.GetFigureByFieldCoordinates([]int{crd[0] - 1, crd[1] - n})).IsItWhite() {
			right = append(right, []int{crd[0] - 1, crd[1] - n})
		}
	}

	var theoryMoves = TheoryMoves{
		Up:     vert,
		Down:   nil,
		Right:  nil,
		Left:   nil,
		UR:     right,
		UL:     left,
		DR:     nil,
		DL:     nil,
		Kn:     nil,
		EnPass: EnPass,
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

	crd := figure.CellCoordinates

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
	crd := figure.CellCoordinates

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
			if fig != nil && (*fig).IsItWhite() == (*figure).IsItWhite() {
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
	crd := figure.CellCoordinates

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

	crd := figure.CellCoordinates

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
	crd := figure.CellCoordinates

	theorySteps := GetTheorySteps(crd)

	var k = [][]int{}

	for _, move := range theorySteps {
		if IsOnRealBoard(move) && figure.AddMove(game, move) {

			canMove := true

			for _, move1 := range GetTheorySteps(move) {
				if move1[0] == crd[0] && move1[1] == crd[1] {
					continue
				}

				fig := (*game).GetFigureByFieldCoordinates(move1)

				if fig != nil &&
					(*fig).GetType() == 'K' {
					canMove = false
					continue
				}

				if fig != nil &&
					(*fig).GetType() == 'k' {
					canMove = false
					continue
				}
			}
			if canMove {
				k = append(k, move)
			}

		}
	}

	castling := [][]int{}

	if !figure.Castling {
		if figure.IsItWhite() && crd[0] == 4 && crd[1] == 7 {
			rookA := game.GetFigureByFieldCoordinates([]int{0, 7})
			if rookA != nil && (*rookA).IsItWhite() && (*rookA).GetType() == 'a' && !game.WhiteCastling.RookACastling {
				if !game.IsKingCheck(60) &&
					!game.IsKingCheck(59) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(59)) == nil &&
					!game.IsKingCheck(58) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(58)) == nil &&
					!game.IsKingCheck(57) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(57)) == nil {
					castling = append(castling, []int{2, 7})
				}
			}

			rookH := game.GetFigureByFieldCoordinates([]int{7, 7})
			if rookH != nil && (*rookH).IsItWhite() && (*rookH).GetType() == 'h' && !game.WhiteCastling.RookHCastling {
				if !game.IsKingCheck(60) &&
					!game.IsKingCheck(61) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(61)) == nil &&
					!game.IsKingCheck(62) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(62)) == nil {
					castling = append(castling, []int{6, 7})
				}
			}
		}

		if !figure.IsItWhite() && crd[0] == 4 && crd[1] == 0 {
			rookA := game.GetFigureByFieldCoordinates([]int{0, 0})
			if rookA != nil && !(*rookA).IsItWhite() && (*rookA).GetType() == 'a' && !game.BlackCastling.RookACastling {
				if !game.IsKingCheck(4) &&
					!game.IsKingCheck(3) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(3)) == nil &&
					!game.IsKingCheck(2) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(2)) == nil &&
					!game.IsKingCheck(1) && game.GetFigureByFieldCoordinates(IndexToFieldCoordinates(1)) == nil {
					castling = append(castling, []int{2, 0})
				}
			}

			rookH := game.GetFigureByFieldCoordinates([]int{7, 0})
			if rookH != nil && !(*rookH).IsItWhite() && (*rookH).GetType() == 'h' && !game.BlackCastling.RookHCastling {
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

func (figure *FigureRook) AddMove(game *Game, crd []int) (bool, bool) {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsItWhite() == (*figure).IsItWhite() {
		return false, false
	}

	if fig != nil && (*fig).IsItWhite() != (*figure).IsItWhite() {
		return true, false
	}

	return true, true
}

func (figure *FigureBishop) AddMove(game *Game, crd []int) (bool, bool) {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsItWhite() == (*figure).IsItWhite() {
		return false, false
	}

	if fig != nil && (*fig).IsItWhite() != (*figure).IsItWhite() {
		return true, false
	}

	return true, true
}

func (figure *FigureQueen) AddMove(game *Game, crd []int) (bool, bool) {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsItWhite() == (*figure).IsItWhite() {
		return false, false
	}

	if fig != nil && (*fig).IsItWhite() != (*figure).IsItWhite() {
		return true, false
	}

	return true, true
}

func (figure *FigureKing) AddMove(game *Game, crd []int) bool {
	fig := game.GetFigureByFieldCoordinates(crd)
	if fig != nil && (*fig).IsItWhite() == (*figure).IsItWhite() {
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
