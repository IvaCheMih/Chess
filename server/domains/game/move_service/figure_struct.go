package move_service

type FigurePawn struct {
	BaseFigure
}

type FigureRook struct {
	BaseFigure
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
}

type TheoryMoves struct {
	Up    []int
	Down  []int
	Right []int
	Left  []int
	UR    []int
	UL    []int
	DR    []int
	DL    []int
	Kn    []int
}

//func (figure *FigurePawn) GetPossibleMoves(game *Game) *TheoryMoves {
//	n := 0
//	if figure.IsWhite() {
//		n = 1
//	} else {
//		n = -1
//	}
//	vert := []int{}
//	index := figure.GameIndex
//	if game.CheckCellOnBoardByIndex(index-n*game.N) && game.GetFigureByIndex(index-n*game.N) == nil {
//		vert = append(vert, index-n*game.N)
//	}
//	if n == 1 && game.IndexToCoordinates(index)[1] == '2' {
//		if game.GetFigureByIndex(index-n*game.N) == nil && game.GetFigureByIndex(index-n*2*game.N) == nil {
//			vert = append(vert, index-n*2*game.N)
//		}
//	}
//	if n == -1 && game.IndexToCoordinates(index)[1] == '7' {
//		if game.GetFigureByIndex(index-n*game.N) == nil && game.GetFigureByIndex(index-n*2*game.N) == nil {
//			vert = append(vert, index-n*2*game.N)
//		}
//	}
//	left := []int{}
//	if game.CheckCellOnBoardByIndex(index-n*(game.N+1)) && game.GetFigureByIndex(index-n*(game.N+1)) != nil {
//		if figure.IsWhite() != (*game.GetFigureByIndex(index - n*(game.N+1))).IsWhite() {
//			fmt.Println("пешка пытается кушоц налево")
//			left = append(left, index-n*(game.N+1))
//		}
//	}
//	right := []int{}
//	if game.CheckCellOnBoardByIndex(index-n*(game.N-1)) && game.GetFigureByIndex(index-n*(game.N-1)) != nil {
//		if figure.IsWhite() != (*game.GetFigureByIndex(index - n*(game.N-1))).IsWhite() {
//			fmt.Println("пешка пытается кушоц направо")
//			right = append(right, index-n*(game.N-1))
//		}
//	}
//
//	var theoryMoves = TheoryMoves{
//		Up:    vert,
//		Down:  nil,
//		Right: nil,
//		Left:  nil,
//		UR:    right,
//		UL:    left,
//		DR:    nil,
//		DL:    nil,
//		Kn:    nil,
//	}
//
//	return &theoryMoves
//}
//
//func (figure *FigureRook) GetPossibleMoves(game *Game) *TheoryMoves {
//	up := []int{}
//
//	for index := figure.GameIndex + game.N; game.CheckCellOnBoardByIndex(index); index += game.N {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		up = append(up, index)
//	}
//
//	down := []int{}
//
//	for index := figure.GameIndex - game.N; game.CheckCellOnBoardByIndex(index); index -= game.N {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		down = append(up, index)
//	}
//
//	right := []int{}
//
//	for index := figure.GameIndex + 1; game.CheckCellOnBoardByIndex(index); index++ {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		right = append(up, index)
//	}
//
//	left := []int{}
//
//	for index := figure.GameIndex - 1; game.CheckCellOnBoardByIndex(index); index-- {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		left = append(up, index)
//	}
//
//	var theoryMoves = TheoryMoves{
//		Up:    up,
//		Down:  down,
//		Right: right,
//		Left:  left,
//		UR:    nil,
//		UL:    nil,
//		DR:    nil,
//		DL:    nil,
//		Kn:    nil,
//	}
//
//	return &theoryMoves
//}
//
//func (figure *FigureKnight) GetPossibleMoves(game *Game) *TheoryMoves {
//	theorySteps := []int{
//		(2 * game.N) + 1,
//		(2 * game.N) - 1,
//		(-1)*(2*game.N) + 1,
//		(-1)*(2*game.N) - 1,
//		game.N + 2,
//		-game.N + 2,
//		game.N - 2,
//		-game.N - 2,
//	}
//	kn := []int{}
//
//	for _, step := range theorySteps {
//
//		if game.CheckCellOnBoardByIndex(step) {
//			if (*game.GetFigureByIndex(step)).IsWhite() == (*figure).IsWhite() {
//				continue
//			}
//			kn = append(kn, step)
//		}
//	}
//
//	var theoryMoves = TheoryMoves{
//		Up:    nil,
//		Down:  nil,
//		Right: nil,
//		Left:  nil,
//		UR:    nil,
//		UL:    nil,
//		DR:    nil,
//		DL:    nil,
//		Kn:    kn,
//	}
//
//	return &theoryMoves
//}
//
//func (figure *FigureBishop) GetPossibleMoves(game *Game) *TheoryMoves {
//	index := figure.GameIndex
//
//	upRight := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N+1)); i++ {
//		if figure.AddMove(game, index+i*(game.N+1)) {
//			upRight = append(upRight, index+i*(game.N+1))
//		}
//	}
//
//	upLeft := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N-1)); i++ {
//		if figure.AddMove(game, index+i*(game.N-1)) {
//			upLeft = append(upLeft, index+i*(game.N-1))
//		}
//	}
//
//	downLeft := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N-1)); i++ {
//		if figure.AddMove(game, index-i*(game.N-1)) {
//			downLeft = append(downLeft, index-i*(game.N-1))
//		}
//	}
//
//	downRight := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N+1)); i++ {
//		if figure.AddMove(game, index-i*(game.N+1)) {
//			downRight = append(downRight, index-i*(game.N+1))
//		}
//	}
//
//	var theoryMoves = TheoryMoves{
//		Up:    nil,
//		Down:  nil,
//		Right: nil,
//		Left:  nil,
//		UR:    upRight,
//		UL:    upLeft,
//		DR:    downRight,
//		DL:    downLeft,
//		Kn:    nil,
//	}
//
//	return &theoryMoves
//}
//
//func (figure *FigureQueen) GetPossibleMoves(game *Game) *TheoryMoves {
//	up := []int{}
//
//	for index := figure.GameIndex + game.N; game.CheckCellOnBoardByIndex(index); index += game.N {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		up = append(up, index)
//	}
//
//	down := []int{}
//
//	for index := figure.GameIndex - game.N; game.CheckCellOnBoardByIndex(index); index -= game.N {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		down = append(up, index)
//	}
//
//	right := []int{}
//
//	for index := figure.GameIndex + 1; game.CheckCellOnBoardByIndex(index); index++ {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		right = append(up, index)
//	}
//
//	left := []int{}
//
//	for index := figure.GameIndex - 1; game.CheckCellOnBoardByIndex(index); index-- {
//		if !figure.AddMove(game, index) {
//			break
//		}
//		left = append(up, index)
//	}
//
//	index := figure.GameIndex
//
//	upRight := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N+1)); i++ {
//		if figure.AddMove(game, index+i*(game.N+1)) {
//			upRight = append(upRight, index+i*(game.N+1))
//		}
//	}
//
//	upLeft := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N-1)); i++ {
//		if figure.AddMove(game, index+i*(game.N-1)) {
//			upLeft = append(upLeft, index+i*(game.N-1))
//		}
//	}
//
//	downLeft := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N-1)); i++ {
//		if figure.AddMove(game, index-i*(game.N-1)) {
//			downLeft = append(downLeft, index-i*(game.N-1))
//		}
//	}
//
//	downRight := []int{}
//
//	for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N+1)); i++ {
//		if figure.AddMove(game, index-i*(game.N+1)) {
//			downRight = append(downRight, index-i*(game.N+1))
//		}
//	}
//
//	var theoryMoves = TheoryMoves{
//		Up:    up,
//		Down:  down,
//		Right: right,
//		Left:  left,
//		UR:    upRight,
//		UL:    upLeft,
//		DR:    downRight,
//		DL:    downLeft,
//		Kn:    nil,
//	}
//
//	return &theoryMoves
//}
//
//func (figure *FigureKing) GetPossibleMoves(game *Game) *TheoryMoves {
//	theorySteps := []int{
//		game.N,
//		game.N + 1,
//		game.N - 1,
//		-1,
//		1,
//		-game.N,
//		-game.N - 1,
//		-game.N + 1,
//	}
//
//	index := figure.GameIndex
//	k := []int{}
//
//	for _, move := range theorySteps {
//		if game.CheckCellOnBoardByIndex(index+move) && figure.AddMove(game, index+move) {
//			for _, move1 := range theorySteps {
//				if move1+move != 0 &&
//					(*game.GetFigureByIndex(index + move + move1)).GetType() != 'K' &&
//					(*game.GetFigureByIndex(index + move + move1)).GetType() != 'k' {
//					k = append(k, index+move)
//				}
//			}
//		}
//	}
//
//	var theoryMoves = TheoryMoves{
//		Up:    nil,
//		Down:  nil,
//		Right: nil,
//		Left:  nil,
//		UR:    nil,
//		UL:    nil,
//		DR:    nil,
//		DL:    nil,
//		Kn:    k,
//	}
//
//	return &theoryMoves
//}
//
//func (figure *FigureRook) ToString() string {
//	if figure.IsWhite() {
//		return "R"
//	}
//	return "r"
//}
//
//func (figure *FigureKnight) ToString() string {
//	if figure.IsWhite() {
//		return "H"
//	}
//	return "h"
//}
//
//func (figure *FigureBishop) ToString() string {
//	if figure.IsWhite() {
//		return "B"
//	}
//	return "b"
//}
//
//func (figure *FigureQueen) ToString() string {
//	if figure.IsWhite() {
//		return "Q"
//	}
//	return "q"
//}
//
//func (figure *FigureKing) ToString() string {
//	if figure.IsWhite() {
//		return "K"
//	}
//	return "k"
//}
//
//func (figure *FigurePawn) ToString() string {
//	if figure.IsWhite() {
//		return "P"
//	}
//	return "p"
//}
//
//func (figure *FigureRook) AddMove(game *Game, index int) bool {
//	fig := game.GetFigureByIndex(index)
//	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
//		return false
//	}
//	return true
//}
//
//func (figure *FigureBishop) AddMove(game *Game, index int) bool {
//	fig := game.GetFigureByIndex(index)
//	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
//		return false
//	}
//	return true
//}
//
//func (figure *FigureQueen) AddMove(game *Game, index int) bool {
//	fig := game.GetFigureByIndex(index)
//	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
//		return false
//	}
//	return true
//}
//
//func (figure *FigureKing) AddMove(game *Game, index int) bool {
//	fig := game.GetFigureByIndex(index)
//	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
//		return false
//	}
//	return true
//}
