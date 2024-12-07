package move

import (
	"fmt"
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
}

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

func (g *Game) GetFigureByFieldCoordinates(crd [2]int) *Figure {
	index := FieldCoordinatesToIndex(crd)
	return g.Figures[index]
}

func (g *Game) GetFigureByCoordinates(coordinates string) *Figure {
	index := g.CoordinatesToIndex(coordinates)

	return g.Figures[index]
}

func (g *Game) IndexToCoordinates(index int) string {
	y := int('8') - (index / g.N)
	x := (index % g.N) + int('A')

	return string(byte(x)) + string(byte(y))
}

func (g *Game) CoordinatesToIndex(coordinates string) int {
	x := int(coordinates[0]) - int('A')
	y := int('8') - int(coordinates[1])

	return (y * g.N) + x
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

func (g *Game) Check() bool {
	if g.Side && g.IsKingCheck(g.IsCheckWhite.KingGameID) {
		return true
	}

	if !g.Side && g.IsKingCheck(g.IsCheckBlack.KingGameID) {
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

	if (*figure).GetType() == 'p' && math.Abs(float64(from-to)) > 9 {

		g.LastPawnMove = &to

		return
	}

	g.LastPawnMove = nil
}

func (g *Game) IsKingCheck(index int) bool {
	if g.CheckKnightAttack(index) {
		return true
	}

	if g.CheckDiagonalAttack(index) {
		return true
	}

	if g.CheckVertGorAttack(index) {
		return true
	}

	if g.CheckPawnAttack(index) {
		return true
	}

	return false
}

func (g *Game) CheckKnightAttack(index int) bool {
	king := g.GetFigureByIndex(index)
	for _, knPosition := range *g.theoryKnightSteps {
		if g.CheckCellOnBoardByIndex(index + knPosition) {
			if fig := g.GetFigureByIndex(index + knPosition); fig != nil && (*fig).GetType() == 'k' {
				if (*fig).IsItWhite() != (*king).IsItWhite() {
					if (*king).IsItWhite() {
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
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1] + i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1] - i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1] - i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1] + i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1] + i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1] - i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1] - i}, 'b')
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
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0], crd[1] + i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0], crd[1] - i}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0], crd[1] - i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1]}); i++ {

		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1]}); i++ {
		isCheck, endFor := g.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}
	return false
}

func (g *Game) CheckAttackCell(kingCoordinate [2]int, cellCoordinate [2]int, triggerFigure byte) (bool, bool) {

	var king *Figure

	if g.Side {
		king = g.GetFigureByIndex(g.IsCheckWhite.KingGameID)
	} else {
		king = g.GetFigureByIndex(g.IsCheckBlack.KingGameID)
	}

	fig := g.GetFigureByFieldCoordinates(cellCoordinate)

	if fig == nil {
		return false, false
	}
	if (*fig).IsItWhite() == (*king).IsItWhite() {
		return false, true
	}
	if (*fig).IsItWhite() != (*king).IsItWhite() {
		if (*fig).GetType() == triggerFigure || (*fig).GetType() == 'q' {
			return true, true
		}
		return false, true
	}
	return false, false
}

func (g *Game) CheckPawnAttack(indexKing int) bool {
	var king *Figure

	if g.Side {
		king = g.GetFigureByIndex(indexKing)
	} else {
		king = g.GetFigureByIndex(indexKing)
	}

	crd := IndexToFieldCoordinates(indexKing)

	if (*king).IsItWhite() && IsOnRealBoard([2]int{crd[0] + 1, crd[1] - 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] + 1, crd[1] - 1}); fig != nil {

			if (*fig).IsItWhite() != (*king).IsItWhite() {
				return true
			}
		}
	}

	if (*king).IsItWhite() && IsOnRealBoard([2]int{crd[0] - 1, crd[1] - 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] - 1, crd[1] - 1}); fig != nil {
			if (*fig).IsItWhite() != (*king).IsItWhite() {
				return true
			}
		}
	}

	if !(*king).IsItWhite() && IsOnRealBoard([2]int{crd[0] + 1, crd[1] + 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] + 1, crd[1] + 1}); fig != nil {
			if (*fig).IsItWhite() != (*king).IsItWhite() {
				return true
			}
		}
	}

	if !(*king).IsItWhite() && IsOnRealBoard([2]int{crd[0] - 1, crd[1] + 1}) {
		if fig := g.GetFigureByFieldCoordinates([2]int{crd[0] - 1, crd[1] + 1}); fig != nil {
			if (*fig).IsItWhite() != (*king).IsItWhite() {
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

	(*figure).Delete()
}

func (g *Game) ChangeRookField(indexesToChange []int) {
	if indexesToChange[1] == -1 {
		return
	}

	g.ChangeToAndFrom(indexesToChange[3], indexesToChange[2])
}

func (g *Game) NewFigureRequestCorrect(to int, newFigure byte) bool {
	figure := g.GetFigureByIndex(to)

	if (*figure).GetType() == 'p' {
		if (*figure).IsItWhite() {
			if to < 8 {
				return g.isNewFigureCorrect(newFigure)
			}
		} else {
			if to > 55 {
				return g.isNewFigureCorrect(newFigure)
			}
		}
	}

	return false
}

func (g *Game) ChangePawnToNewFigure(to int, newFigure byte) {
	figure := g.GetFigureByIndex(to)

	(*figure).ChangeType(newFigure)

	g.NewFigureId = mutateNewFigureId(newFigure, (*figure).IsItWhite())
}

func (g *Game) isNewFigureCorrect(newFigure byte) bool {
	_, ok := g.newFigures[newFigure]
	return ok
}

func IndexToCoordinates(index int) string {
	y := int('8') - (index / 8)
	x := (index % 8) + int('A')

	return string(byte(x)) + string(byte(y))
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
	}

	return &newGame
}

func (g *Game) IsItEndgame() (bool, EndgameReason) {
	for _, figure := range g.Figures {
		if !g.IsItYourFigure(figure) {
			continue
		}

		fmt.Println(string((*figure).GetType()))

		fromCrd := (*figure).GetCoordinates()
		theoryMoves := (*figure).GetPossibleMoves(g)

		if g.movesExist(theoryMoves, fromCrd) {
			return false, ""
		}
		fmt.Println()
	}

	if g.Side {
		if g.IsCheckWhite.IsItCheck {
			return true, Mate
		}
		return true, Pat
	} else {
		if g.IsCheckBlack.IsItCheck {
			return true, Mate
		}
		return true, Pat
	}

	return false, ""
}

func (g *Game) movesExist(theoryMoves *TheoryMoves, fromCrd [2]int) bool {
	if theoryMoves.Up != nil {
		for _, move := range theoryMoves.Up {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))
			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Down != nil {
		for _, move := range theoryMoves.Down {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Right != nil {
		for _, move := range theoryMoves.Right {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Left != nil {
		for _, move := range theoryMoves.Left {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.UR != nil {
		for _, move := range theoryMoves.UR {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.UL != nil {
		for _, move := range theoryMoves.UL {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.DR != nil {
		for _, move := range theoryMoves.DR {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.DL != nil {
		for _, move := range theoryMoves.DL {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Kn != nil {
		for _, move := range theoryMoves.Kn {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.EnPass != nil {
		for _, move := range theoryMoves.EnPass {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

			if g.moveExists(theoryMoves, move, fromCrd) {
				return true
			}
		}
	}
	if theoryMoves.Castling != nil {
		for _, move := range theoryMoves.Castling {
			fmt.Println(IndexToCoordinates(FieldCoordinatesToIndex(fromCrd)), IndexToCoordinates(FieldCoordinatesToIndex(move)))

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
			if !IsItCheck(indexesToChange, newGame, figureByte) {
				fmt.Println("no check on board")
				return true
			}
		}
	} else {
		newGame := g.copyGame()
		if !IsItCheck(indexesToChange, newGame, '0') {
			return true
		}
	}

	return false
}
