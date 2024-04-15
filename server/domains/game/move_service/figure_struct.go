package move_service

import (
	"fmt"
	"sync"
)

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
	Mu    sync.Mutex
}

func (figure *FigurePawn) GetPossibleMoves(game *Game) *TheoryMoves {
	n := 0
	if figure.IsWhite() {
		n = 1
	} else {
		n = -1
	}
	vert := []int{}
	index := figure.GameIndex
	if game.CheckCellOnBoardByIndex(index-n*game.N) && game.GetFigureByIndex(index-n*game.N) == nil {
		vert = append(vert, index-n*game.N)
	}
	if n == 1 && game.IndexToCoordinates(index)[1] == '2' {
		if game.GetFigureByIndex(index-n*game.N) == nil && game.GetFigureByIndex(index-n*2*game.N) == nil {
			vert = append(vert, index-n*2*game.N)
		}
	}
	if n == -1 && game.IndexToCoordinates(index)[1] == '7' {
		if game.GetFigureByIndex(index-n*game.N) == nil && game.GetFigureByIndex(index-n*2*game.N) == nil {
			vert = append(vert, index-n*2*game.N)
		}
	}
	left := []int{}
	if game.CheckCellOnBoardByIndex(index-n*(game.N+1)) && game.GetFigureByIndex(index-n*(game.N+1)) != nil {
		if figure.IsWhite() != (*game.GetFigureByIndex(index - n*(game.N+1))).IsWhite() {
			fmt.Println("пешка пытается кушоц налево")
			left = append(left, index-n*(game.N+1))
		}
	}
	right := []int{}
	if game.CheckCellOnBoardByIndex(index-n*(game.N-1)) && game.GetFigureByIndex(index-n*(game.N-1)) != nil {
		if figure.IsWhite() != (*game.GetFigureByIndex(index - n*(game.N-1))).IsWhite() {
			fmt.Println("пешка пытается кушоц направо")
			right = append(right, index-n*(game.N-1))
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
		Up:    []int{},
		Down:  []int{},
		Right: []int{},
		Left:  []int{},
		UR:    nil,
		UL:    nil,
		DR:    nil,
		DL:    nil,
		Kn:    nil,
		Mu:    sync.Mutex{},
	}

	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		for index := figure.GameIndex + game.N; game.CheckCellOnBoardByIndex(index); index += game.N {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Up = append(theoryMoves.Up, index)
			theoryMoves.Mu.Unlock()
		}
		wg.Done()

	}()

	go func() {
		wg.Add(1)

		for index := figure.GameIndex - game.N; game.CheckCellOnBoardByIndex(index); index -= game.N {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Down = append(theoryMoves.Down, index)
			theoryMoves.Mu.Unlock()
		}
		wg.Done()

	}()

	go func() {
		wg.Add(1)

		for index := figure.GameIndex + 1; game.CheckCellOnBoardByIndex(index); index++ {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Right = append(theoryMoves.Right, index)
			theoryMoves.Mu.Unlock()
		}
		wg.Done()

	}()

	go func() {
		wg.Add(1)

		for index := figure.GameIndex - 1; game.CheckCellOnBoardByIndex(index); index-- {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Left = append(theoryMoves.Left, index)
			theoryMoves.Mu.Unlock()
		}
		wg.Done()

	}()

	wg.Wait()

	return &theoryMoves
}

func (figure *FigureKnight) GetPossibleMoves(game *Game) *TheoryMoves {
	index := figure.GameIndex

	theorySteps := []int{
		(2 * game.N) + 1,
		(2 * game.N) - 1,
		(-1)*(2*game.N) + 1,
		(-1)*(2*game.N) - 1,
		game.N + 2,
		-game.N + 2,
		game.N - 2,
		-game.N - 2,
	}
	kn := []int{}

	for _, step := range theorySteps {
		if game.CheckCellOnBoardByIndex(index + step) {
			if game.GetFigureByIndex(index+step) != nil && (*game.GetFigureByIndex(index + step)).IsWhite() == (*figure).IsWhite() {
				continue
			}
			kn = append(kn, index+step)
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
	index := figure.GameIndex

	fmt.Println(index)

	var theoryMoves = TheoryMoves{
		Up:    nil,
		Down:  nil,
		Right: nil,
		Left:  nil,
		UR:    []int{},
		UL:    []int{},
		DR:    []int{},
		DL:    []int{},
		Kn:    nil,
		Mu:    sync.Mutex{},
	}

	wg := sync.WaitGroup{}

	wg.Add(1)

	go func() {
		for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N+1)); i++ {
			add, _continue := figure.AddMove(game, index+i*(game.N+1))

			if add {
				theoryMoves.Mu.Lock()
				theoryMoves.UR = append(theoryMoves.UR, index+i*(game.N+1))
				theoryMoves.Mu.Unlock()
			}

			if !_continue {
				break
			}
		}
		wg.Done()

	}()

	wg.Add(1)

	go func() {
		for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N-1)); i++ {
			add, _continue := figure.AddMove(game, index+i*(game.N-1))

			if add {
				theoryMoves.Mu.Lock()
				theoryMoves.UR = append(theoryMoves.UR, index+i*(game.N+1))
				theoryMoves.Mu.Unlock()
			}

			if !_continue {
				break
			}
		}
		wg.Done()
	}()

	wg.Add(1)

	go func() {
		for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N-1)); i++ {
			add, _continue := figure.AddMove(game, index-i*(game.N-1))

			if add {
				theoryMoves.Mu.Lock()
				theoryMoves.UR = append(theoryMoves.UR, index+i*(game.N+1))
				theoryMoves.Mu.Unlock()
			}

			if !_continue {
				break
			}
		}
		wg.Done()

	}()

	wg.Add(1)

	go func() {
		for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N+1)); i++ {
			add, _continue := figure.AddMove(game, index-i*(game.N+1))

			if add {
				theoryMoves.Mu.Lock()
				theoryMoves.UR = append(theoryMoves.UR, index+i*(game.N+1))
				theoryMoves.Mu.Unlock()
			}

			if !_continue {
				break
			}
		}
		wg.Done()

	}()

	wg.Wait()

	fmt.Println(theoryMoves)

	return &theoryMoves
}

func (figure *FigureQueen) GetPossibleMoves(game *Game) *TheoryMoves {
	var theoryMoves = TheoryMoves{
		Up:    []int{},
		Down:  []int{},
		Right: []int{},
		Left:  []int{},
		UR:    []int{},
		UL:    []int{},
		DR:    []int{},
		DL:    []int{},
		Kn:    nil,
		Mu:    sync.Mutex{},
	}

	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		for index := figure.GameIndex + game.N; game.CheckCellOnBoardByIndex(index); index += game.N {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Up = append(theoryMoves.Up, index)
			theoryMoves.Mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for index := figure.GameIndex - game.N; game.CheckCellOnBoardByIndex(index); index -= game.N {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Down = append(theoryMoves.Down, index)
			theoryMoves.Mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for index := figure.GameIndex + 1; game.CheckCellOnBoardByIndex(index); index++ {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Right = append(theoryMoves.Right, index)
			theoryMoves.Mu.Unlock()

		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for index := figure.GameIndex - 1; game.CheckCellOnBoardByIndex(index); index-- {
			if !figure.AddMove(game, index) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.Left = append(theoryMoves.Left, index)
			theoryMoves.Mu.Unlock()
		}
		wg.Done()
	}()

	index := figure.GameIndex

	go func() {
		wg.Add(1)
		for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N+1)); i++ {
			if !figure.AddMove(game, index+i*(game.N+1)) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.UR = append(theoryMoves.UR, index+i*(game.N+1))
			theoryMoves.Mu.Unlock()
		}
		wg.Done()

	}()

	go func() {
		wg.Add(1)
		for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N-1)); i++ {
			if !figure.AddMove(game, index+i*(game.N-1)) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.UL = append(theoryMoves.UL, index+i*(game.N-1))
			theoryMoves.Mu.Unlock()
		}
		wg.Done()

	}()

	go func() {
		wg.Add(1)
		for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N-1)); i++ {
			if !figure.AddMove(game, index-i*(game.N-1)) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.DL = append(theoryMoves.DL, index-i*(game.N-1))
			theoryMoves.Mu.Unlock()
		}
		wg.Done()
	}()

	go func() {
		wg.Add(1)
		for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N+1)); i++ {
			if !figure.AddMove(game, index-i*(game.N+1)) {
				break
			}
			theoryMoves.Mu.Lock()
			theoryMoves.DR = append(theoryMoves.DR, index-i*(game.N+1))
			theoryMoves.Mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()

	return &theoryMoves
}

func (figure *FigureKing) GetPossibleMoves(game *Game) *TheoryMoves {
	theorySteps := []int{
		game.N,
		game.N + 1,
		game.N - 1,
		-1,
		1,
		-game.N,
		-game.N - 1,
		-game.N + 1,
	}

	index := figure.GameIndex
	k := []int{}

	for _, move := range theorySteps {
		if game.CheckCellOnBoardByIndex(index+move) && figure.AddMove(game, index+move) {
			for _, move1 := range theorySteps {
				if move1+move != 0 &&
					(*game.GetFigureByIndex(index + move + move1)).GetType() != 'K' &&
					(*game.GetFigureByIndex(index + move + move1)).GetType() != 'k' {
					k = append(k, index+move)
				}
			}
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
		Kn:    k,
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

func (figure *FigureRook) AddMove(game *Game, index int) bool {
	fig := game.GetFigureByIndex(index)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false
	}
	return true
}

func (figure *FigureBishop) AddMove(game *Game, index int) (bool, bool) {
	fig := game.GetFigureByIndex(index)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false, false
	}

	if fig != nil && (*fig).IsWhite() != (*figure).IsWhite() {
		return true, false
	}

	return true, true
}

func (figure *FigureQueen) AddMove(game *Game, index int) bool {
	fig := game.GetFigureByIndex(index)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false
	}
	return true
}

func (figure *FigureKing) AddMove(game *Game, index int) bool {
	fig := game.GetFigureByIndex(index)
	if fig != nil && (*fig).IsWhite() == (*figure).IsWhite() {
		return false
	}
	return true
}
