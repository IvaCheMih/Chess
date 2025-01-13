package move

import (
	"math"
)

type Game struct {
	N                 int
	Figures           map[int]*Figure
	IsCheckWhite      IsCheck
	IsCheckBlack      IsCheck
	WhiteCastling     Castling
	BlackCastling     Castling
	LastPawnMove      *int
	Side              bool
	NewFigureId       int
	newFigures        map[byte]struct{}
	theoryKnightSteps *[]int
	KilledFigure      byte
	LastLoss          int
}

const lastLossLimit = 50

type IsCheck struct {
	IsItCheck  bool
	KingGameID int
}

type Castling struct {
	KingCastling  bool
	RookACastling bool
	RookHCastling bool
}

func (g *Game) GetFigureByIndex(index int) *Figure {
	return g.Figures[index]
}

func (g *Game) MakeKing(index int) Figure {
	var _type byte

	if g.Side {
		_type = 5
	} else {
		_type = 12
	}

	coordinates := IndexToFieldCoordinates(index)

	var bf = BaseFigure{g.Side, _type, coordinates}

	return &FigureKing{bf, false}
}

func (g *Game) GetFigureByFieldCoordinates(crd [2]int) *Figure {
	index := FieldCoordinatesToIndex(crd)
	return g.Figures[index]
}

func (g *Game) GetFigureByCoordinates(coordinates string) *Figure {
	index := g.CoordinatesToIndex(coordinates)

	return g.Figures[index]
}

func (g *Game) CheckCellOnBoardByIndex(index int) bool {
	coordinates := g.IndexToCoordinates(index)
	if coordinates[0] >= byte('A') && coordinates[0] <= byte('H') {
		if coordinates[1] >= byte('1') && coordinates[1] <= byte('8') {
			return true
		}
	}
	return false
}

func (g *Game) ChangeKingGameID(to int) {
	figure := g.GetFigureByIndex(to)

	if (*figure).GetType() != 'K' {
		return
	}
	if (*figure).IsItWhite() {
		g.IsCheckWhite.KingGameID = FieldCoordinatesToIndex((*figure).GetCoordinates())
	} else {
		g.IsCheckBlack.KingGameID = FieldCoordinatesToIndex((*figure).GetCoordinates())
	}
}

// ChangeIsItChecks for both players
func (g *Game) ChangeIsItChecks() {
	if g.IsKingCheck(g.IsCheckWhite.KingGameID) {
		g.IsCheckWhite.IsItCheck = true
	}

	if g.IsKingCheck(g.IsCheckBlack.KingGameID) {
		g.IsCheckBlack.IsItCheck = true
	}
}

func (g *Game) ChangeSide() {
	g.Side = !g.Side
}

// CheckToMovingPlayer to the moving player
func (g *Game) CheckToMovingPlayer() bool {
	if g.Side && g.IsCheckWhite.IsItCheck {
		return true
	}

	if !g.Side && g.IsCheckBlack.IsItCheck {
		return true
	}

	return false
}

func (g *Game) ChangeCastlingFlag(to int) {
	figure := g.GetFigureByIndex(to)

	switch (*figure).GetType() {
	case 'K':
		if (*figure).IsItWhite() {
			g.WhiteCastling.KingCastling = true
		} else {
			g.BlackCastling.KingCastling = true
		}
	case 'a':
		if (*figure).IsItWhite() {
			g.WhiteCastling.RookACastling = true
		} else {
			g.BlackCastling.RookACastling = true
		}
	case 'h':
		if (*figure).IsItWhite() {
			g.WhiteCastling.RookHCastling = true
		} else {
			g.BlackCastling.RookHCastling = true
		}
	}
}

func (g *Game) ChangeLastPawnMove(from int, to int) {
	figure := g.GetFigureByIndex(to)

	if (*figure).GetType() == 'p' && math.Abs(float64(from-to)) == 2*8 {
		g.LastPawnMove = &to
		return
	}

	g.LastPawnMove = nil
}

func (g *Game) IsKingCheck(kingCellIndex int) bool {
	if g.CheckKnightAttack(kingCellIndex) {
		return true
	}
	if g.CheckDiagonalAttack(kingCellIndex) {
		return true
	}

	if g.CheckVertGorAttack(kingCellIndex) {
		return true
	}

	if g.CheckPawnAttack(kingCellIndex) {
		return true
	}

	return false
}

func (g *Game) CheckKnightAttack(index int) bool {
	//king := g.GetFigureByIndex(index)
	king := g.MakeKing(index)
	for _, knPosition := range *g.theoryKnightSteps {
		if g.CheckCellOnBoardByIndex(index + knPosition) {
			if fig := g.GetFigureByIndex(index + knPosition); fig != nil && (*fig).GetType() == 'k' {
				if (*fig).IsItWhite() != king.IsItWhite() {
					if king.IsItWhite() {
						g.IsCheckWhite.IsItCheck = true
						return true
					} else {
						g.IsCheckBlack.IsItCheck = true
						return true
					}
				}
			}
		}
	}
	return false
}

func (g *Game) CheckDiagonalAttack(index int) bool {
	crd := IndexToFieldCoordinates(index)

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1] + i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1] + i}, []byte{'b', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1] - i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1] - i}, []byte{'b', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1] + i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1] + i}, []byte{'b', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1] - i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1] - i}, []byte{'b', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	return false
}

func (g *Game) CheckVertGorAttack(index int) bool {
	crd := IndexToFieldCoordinates(index)

	for i := 1; IsOnRealBoard([2]int{crd[0], crd[1] + i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0], crd[1] + i}, []byte{'a', 'h', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0], crd[1] - i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0], crd[1] - i}, []byte{'a', 'h', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1]}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1]}, []byte{'a', 'h', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1]}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1]}, []byte{'a', 'h', 'q'})
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	return false
}

func (g *Game) CheckAttackCell(kingCoordinate [2]int, cellCoordinate [2]int, triggerFigures []byte) (bool, bool) {
	//king := g.GetFigureByFieldCoordinates(kingCoordinate)
	king := g.MakeKing(FieldCoordinatesToIndex(kingCoordinate))
	fig := g.GetFigureByFieldCoordinates(cellCoordinate)

	if fig == nil {
		return false, false
	}

	if (*fig).IsItWhite() == king.IsItWhite() {
		return false, true
	}
	if (*fig).IsItWhite() != king.IsItWhite() {
		if isTriggerFigure((*fig).GetType(), triggerFigures) {
			return true, true
		}
		return false, true
	}
	return false, false
}

func isTriggerFigure(_type byte, triggerFigures []byte) bool {
	for _, fig := range triggerFigures {
		if _type == fig {
			return true
		}
	}
	return false
}

func (g *Game) CheckPawnAttack(indexKing int) bool {
	//	king := g.GetFigureByIndex(indexKing)
	king := g.MakeKing(indexKing)
	crd := IndexToFieldCoordinates(indexKing)

	if king.IsItWhite() && IsOnRealBoard([2]int{crd[0] + 1, crd[1] - 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] + 1, crd[1] - 1}); fig != nil {
			if (*fig).IsItWhite() != king.IsItWhite() {
				return true
			}
		}
	}

	if king.IsItWhite() && IsOnRealBoard([2]int{crd[0] - 1, crd[1] - 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] - 1, crd[1] - 1}); fig != nil {
			if (*fig).IsItWhite() != king.IsItWhite() {
				return true
			}
		}
	}

	if !king.IsItWhite() && IsOnRealBoard([2]int{crd[0] + 1, crd[1] + 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] + 1, crd[1] + 1}); fig != nil {
			if (*fig).IsItWhite() != king.IsItWhite() {
				return true
			}
		}
	}

	if !king.IsItWhite() && IsOnRealBoard([2]int{crd[0] - 1, crd[1] + 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] - 1, crd[1] + 1}); fig != nil {
			if (*fig).IsItWhite() != king.IsItWhite() {
				return true
			}
		}
	}
	return false
}

func (g *Game) ChangeToAndFrom(to int, from int) {
	coordinateTo := IndexToFieldCoordinates(to)
	coordinateFrom := IndexToFieldCoordinates(from)

	figureTo := g.GetFigureByFieldCoordinates(coordinateTo)
	figureFrom := g.GetFigureByFieldCoordinates(coordinateFrom)

	if figureTo != nil {
		g.KilledFigure = (*figureTo).GetType()
		(*figureTo).Delete()
	}

	(*figureFrom).ChangeCoordinates(coordinateTo)

	g.Figures[to] = g.Figures[from]
	g.Figures[from] = nil

	figureTo = g.GetFigureByIndex(to) //nolint:ineffassign,staticcheck
}

//func (g *Game) ChangeRookIfCastling(to int) {
//	switch to {
//	case 2:
//		g.ChangeToAndFrom(3, 0)
//	case 6:
//		g.ChangeToAndFrom(5, 7)
//	case 57:
//		g.ChangeToAndFrom(59, 56)
//	case 62:
//		g.ChangeToAndFrom(61, 63)
//	}
//}

func (g *Game) IsItYourFigure(figure *Figure) bool {
	if figure == nil {
		return false
	}

	if g.Side && !(*figure).IsItWhite() {
		return false
	}

	if !g.Side && (*figure).IsItWhite() {
		return false
	}

	return true
}

func (g *Game) DeletePawn(indexesToChange []int) {
	if indexesToChange[1] != -1 {
		return
	}

	figure := g.GetFigureByIndex(indexesToChange[3])

	g.KilledFigure = (*figure).GetType()

	(*figure).Delete()
}

func (g *Game) ChangeRookField(indexesToChange []int) {
	if indexesToChange[1] == -1 {
		return
	}

	g.ChangeToAndFrom(indexesToChange[3], indexesToChange[2])
}

func (g *Game) NewFigureRequestCorrect(to int, pawnColor bool) bool {
	if pawnColor && to < 8 {
		return true
	}

	if !pawnColor && to > 55 {
		return true
	}

	return false
}

func (g *Game) ChangePawnToNewFigure(to int, newFigure byte) {
	figure := g.GetFigureByIndex(to)

	if (*figure).GetType() == 'p' {
		if (*figure).IsItWhite() {
			if to < 8 && g.isNewFigureCorrect(newFigure) {
				(*figure).ChangeType(newFigure)

				g.NewFigureId = mutateNewFigureId(newFigure, (*figure).IsItWhite())
			}
		} else {
			if to > 55 && g.isNewFigureCorrect(newFigure) {
				(*figure).ChangeType(newFigure)

				g.NewFigureId = mutateNewFigureId(newFigure, (*figure).IsItWhite())
			}
		}
	}
}

func (g *Game) isNewFigureCorrect(newFigure byte) bool {
	_, ok := g.newFigures[newFigure]
	return ok
}

func mutateNewFigureId(newFigure byte, color bool) int {
	if color {
		switch rune(newFigure) {
		case 'a':
			return 1
		case 'h':
			return 7
		case 'k':
			return 2
		case 'b':
			return 3
		case 'q':
			return 4
		}
	}

	switch rune(newFigure) {
	case 'a':
		return 8
	case 'h':
		return 14
	case 'k':
		return 9
	case 'b':
		return 10
	case 'q':
		return 11
	}
	return 0
}

func CreateFigureRepo() map[int]byte {
	var figureRepo = make(map[int]byte)

	figureRepo[1] = 'a'
	figureRepo[2] = 'k'
	figureRepo[3] = 'b'
	figureRepo[4] = 'q'
	figureRepo[5] = 'K'
	figureRepo[6] = 'p'
	figureRepo[7] = 'h'

	figureRepo[8] = 'a'
	figureRepo[9] = 'k'
	figureRepo[10] = 'b'
	figureRepo[11] = 'q'
	figureRepo[12] = 'K'
	figureRepo[13] = 'p'
	figureRepo[14] = 'h'

	return figureRepo
}

func (g *Game) copyGame() *Game {
	var lastPawnMove int
	if g.LastPawnMove != nil {
		lastPawnMove = *g.LastPawnMove
	}

	var figures = make(map[int]*Figure, len(g.Figures))

	for i, figure := range g.Figures {
		var f Figure
		if figure != nil {
			f = *figure

			figures[i] = &f
		}
	}

	var newGame = Game{
		N: g.N,
		IsCheckWhite: IsCheck{
			IsItCheck:  g.IsCheckWhite.IsItCheck,
			KingGameID: g.IsCheckWhite.KingGameID,
		},
		IsCheckBlack: IsCheck{
			IsItCheck:  g.IsCheckBlack.IsItCheck,
			KingGameID: g.IsCheckBlack.KingGameID,
		},
		WhiteCastling: Castling{
			KingCastling:  g.WhiteCastling.KingCastling,
			RookACastling: g.WhiteCastling.RookACastling,
			RookHCastling: g.WhiteCastling.RookHCastling,
		},
		BlackCastling: Castling{
			KingCastling:  g.BlackCastling.KingCastling,
			RookACastling: g.BlackCastling.RookACastling,
			RookHCastling: g.BlackCastling.RookHCastling,
		},
		LastPawnMove:      &lastPawnMove,
		Figures:           figures,
		Side:              g.Side,
		NewFigureId:       g.NewFigureId,
		newFigures:        g.newFigures,
		theoryKnightSteps: g.theoryKnightSteps,
		KilledFigure:      g.KilledFigure,
		LastLoss:          g.LastLoss,
	}

	return &newGame
}

func (g *Game) CompareGamesStates(gameState Game) bool {
	return checkCompares(*g, gameState, compares)
}

func (g *Game) movesExist(theoryMoves *TheoryMoves, fromCrd [2]int) bool {
	if theoryMoves.Up != nil {
		for _, move := range theoryMoves.Up {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Down != nil {
		for _, move := range theoryMoves.Down {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Right != nil {
		for _, move := range theoryMoves.Right {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Left != nil {
		for _, move := range theoryMoves.Left {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.UR != nil {
		for _, move := range theoryMoves.UR {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.UL != nil {
		for _, move := range theoryMoves.UL {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.DR != nil {
		for _, move := range theoryMoves.DR {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.DL != nil {
		for _, move := range theoryMoves.DL {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Kn != nil {
		for _, move := range theoryMoves.Kn {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.EnPass != nil {
		for _, move := range theoryMoves.EnPass {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Castling != nil {
		for _, move := range theoryMoves.Castling {
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}

	return false
}

func (g *Game) moveExists(theoryMoves *TheoryMoves, toCrd [2]int, fromCrd [2]int) bool {
	moveInIndexes := []int{FieldCoordinatesToIndex(fromCrd), FieldCoordinatesToIndex(toCrd)}

	isCorrect, indexesToChange := checkMove(theoryMoves, moveInIndexes)
	if !isCorrect {
		return false
	}

	if (*g.GetFigureByFieldCoordinates(fromCrd)).GetType() == 'p' {
		for figureByte := range g.newFigures {
			newGame := g.copyGame()

			newGame.DoMove(indexesToChange, figureByte)

			if !newGame.CheckToMovingPlayer() {
				return true
			}
		}
	} else {
		newGame := g.copyGame()

		newGame.DoMove(indexesToChange, '0')

		if !newGame.CheckToMovingPlayer() {
			return true
		}
	}

	return false
}

type compareFunctions[T any] []compareFunction[T]

type compareFunction[T any] func(a1, a2 T) bool

func checkCompares(g1 Game, g2 Game, compares compareFunctions[Game]) bool {
	for _, compare := range compares {
		if !compare(g1, g2) {
			return false
		}
	}

	return true
}

var compares = compareFunctions[Game]{
	func(g1, g2 Game) bool {
		for i := range g1.N*g1.N - 1 {
			if g1.Figures[i] == nil && g2.Figures[i] == nil {
				continue
			}

			if g1.Figures[i] != nil && g2.Figures[i] != nil &&
				(*g1.Figures[i]).IsItWhite() == (*g2.Figures[i]).IsItWhite() &&
				(*g1.Figures[i]).GetType() == (*g2.Figures[i]).GetType() {
				continue
			}

			return false
		}
		return true
	},

	func(g1, g2 Game) bool {
		return g1.WhiteCastling.RookHCastling == g2.WhiteCastling.RookHCastling
	},
	func(g1, g2 Game) bool {
		return g1.WhiteCastling.RookACastling == g2.WhiteCastling.RookACastling
	},
	func(g1, g2 Game) bool {
		return g1.WhiteCastling.KingCastling == g2.WhiteCastling.KingCastling
	},

	func(g1, g2 Game) bool {
		return g1.BlackCastling.RookHCastling == g2.BlackCastling.RookHCastling
	},
	func(g1, g2 Game) bool {
		return g1.BlackCastling.RookACastling == g2.BlackCastling.RookACastling
	},
	func(g1, g2 Game) bool {
		return g1.BlackCastling.KingCastling == g2.BlackCastling.KingCastling
	},

	func(g1, g2 Game) bool {
		return g1.N == g2.N
	},

	func(g1, g2 Game) bool {
		return g1.IsCheckBlack.KingGameID == g2.IsCheckBlack.KingGameID
	},
	func(g1, g2 Game) bool {
		return g1.IsCheckBlack.IsItCheck == g2.IsCheckBlack.IsItCheck
	},
	func(g1, g2 Game) bool {
		return g1.IsCheckWhite.KingGameID == g2.IsCheckWhite.KingGameID
	},
	func(g1, g2 Game) bool {
		return g1.IsCheckWhite.IsItCheck == g2.IsCheckWhite.IsItCheck
	},

	func(g1, g2 Game) bool {
		return g1.LastLoss == g2.LastLoss
	},
	func(g1, g2 Game) bool {
		return g1.KilledFigure == g2.KilledFigure
	},
	func(g1, g2 Game) bool {
		return g1.NewFigureId == g2.NewFigureId
	},
	func(g1, g2 Game) bool {
		return g1.Side == g2.Side
	},
	func(g1, g2 Game) bool {
		return g1.LastPawnMove == g2.LastPawnMove
	},
	func(g1, g2 Game) bool {
		return g1.NewFigureId == g2.NewFigureId
	},
}
