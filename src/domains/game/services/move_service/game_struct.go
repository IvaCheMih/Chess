package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"math"
)

type Game struct {
	N             int
	Figures       map[int]*Figure
	IsCheckWhite  IsCheck
	IsCheckBlack  IsCheck
	WhiteCastling Castling
	BlackCastling Castling
	LastPawnMove  *int
	Side          bool
	NewFigureId   int
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

var FigureRepo = make(map[int]byte)

func CreateGameStruct(gameModel models.Game, board models.Board) Game {

	figures, blackKingCell, whiteKingCell := CreateField(board, gameModel)

	side := gameModel.Side

	return Game{
		N: 8,
		//WhiteClientId: &gameModel.WhiteUserId,
		//BlackClientId: &gameModel.BlackUserId,
		Figures:       figures,
		IsCheckWhite:  IsCheck{gameModel.IsCheckWhite, whiteKingCell},
		IsCheckBlack:  IsCheck{gameModel.IsCheckBlack, blackKingCell},
		WhiteCastling: Castling{gameModel.WhiteKingCastling, gameModel.WhiteRookACastling, gameModel.WhiteRookHCastling},
		BlackCastling: Castling{gameModel.BlackKingCastling, gameModel.BlackRookACastling, gameModel.BlackRookHCastling},
		LastPawnMove:  gameModel.LastPawnMove,
		Side:          side,
		NewFigureId:   0,
	}
}

func (game *Game) GetFigureByIndex(index int) *Figure {
	return game.Figures[index]
}

func (game *Game) GetFigureByFieldCoordinates(crd [2]int) *Figure {
	index := FieldCoordinatesToIndex(crd)
	return game.Figures[index]
}

func (game *Game) GetFigureByCoordinates(coordinates string) *Figure {
	index := game.CoordinatesToIndex(coordinates)

	return game.Figures[index]
}

func (game *Game) IndexToCoordinates(index int) string {
	y := int('8') - (index / game.N)
	x := (index % game.N) + int('A')

	return string(byte(x)) + string(byte(y))
}

func (game *Game) CoordinatesToIndex(coordinates string) int {
	x := int(coordinates[0]) - int('A')
	y := int('8') - int(coordinates[1])

	return (y * game.N) + x
}

func (game *Game) CheckCellOnBoardByIndex(index int) bool {
	coordinates := game.IndexToCoordinates(index)
	if coordinates[0] >= byte('A') && coordinates[0] <= byte('H') {
		if coordinates[1] >= byte('1') && coordinates[1] <= byte('8') {
			return true
		}
	}
	return false
}

func (game *Game) ChangeKingGameID(to int) {
	figure := game.GetFigureByIndex(to)

	if (*figure).GetType() != 'K' {
		return
	}
	if (*figure).IsItWhite() {
		game.IsCheckWhite.KingGameID = FieldCoordinatesToIndex((*figure).GetCoordinates())
	} else {
		game.IsCheckBlack.KingGameID = FieldCoordinatesToIndex((*figure).GetCoordinates())
	}
}

func (game *Game) Check() bool {
	if game.Side && game.IsKingCheck(game.IsCheckWhite.KingGameID) {
		return true
	}

	if !game.Side && game.IsKingCheck(game.IsCheckBlack.KingGameID) {
		return true
	}

	return false
}

func (game *Game) ChangeCastlingFlag(to int) {
	figure := game.GetFigureByIndex(to)

	switch (*figure).GetType() {
	case 'K':
		if (*figure).IsItWhite() {
			game.WhiteCastling.KingCastling = true
		} else {
			game.BlackCastling.KingCastling = true
		}
	case 'a':
		if (*figure).IsItWhite() {
			game.WhiteCastling.RookACastling = true
		} else {
			game.BlackCastling.RookACastling = true
		}
	case 'h':
		if (*figure).IsItWhite() {
			game.WhiteCastling.RookHCastling = true
		} else {
			game.BlackCastling.RookHCastling = true
		}
	}
}

func (game *Game) ChangeLastPawnMove(from int, to int) {
	figure := game.GetFigureByIndex(to)

	if (*figure).GetType() == 'p' && math.Abs(float64(from-to)) > 9 {

		game.LastPawnMove = &to

		return
	}

	game.LastPawnMove = nil
}

func (game *Game) IsKingCheck(index int) bool {
	if game.CheckKnightAttack(index) {
		return true
	}

	if game.CheckDiagonalAttack(index) {
		return true
	}

	if game.CheckVertGorAttack(index) {
		return true
	}

	if game.CheckPawnAttack(index) {
		return true
	}

	return false
}

func (game *Game) CheckKnightAttack(index int) bool {
	king := game.GetFigureByIndex(index)
	for _, knPosition := range TheoryKnightSteps {
		if game.CheckCellOnBoardByIndex(index + knPosition) {
			if fig := game.GetFigureByIndex(index + knPosition); fig != nil && (*fig).GetType() == 'h' {
				if (*fig).IsItWhite() != (*king).IsItWhite() {
					if (*king).IsItWhite() {
						game.IsCheckWhite.IsItCheck = true
						return true
					} else {
						game.IsCheckBlack.IsItCheck = true
						return true
					}
				}
			}

		}
	}
	return false
}

func (game *Game) CheckDiagonalAttack(index int) bool {
	crd := IndexToFieldCoordinates(index)

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1] + i}); i++ {
		fmt.Println([]int{crd[0] + i, crd[1] + i})
		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1] + i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1] - i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1] - i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1] + i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1] + i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1] - i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1] - i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	return false
}

func (game *Game) CheckVertGorAttack(index int) bool {
	crd := IndexToFieldCoordinates(index)

	for i := 1; IsOnRealBoard([2]int{crd[0], crd[1] + i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0], crd[1] + i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0], crd[1] - i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0], crd[1] - i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] + i, crd[1]}); i++ {

		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0] + i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([2]int{crd[0] - i, crd[1]}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, [2]int{crd[0] - i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}
	return false
}

func (game *Game) CheckAttackCell(kingCoordinate [2]int, cellCoordinate [2]int, triggerFigure byte) (bool, bool) {

	var king *Figure

	if game.Side {
		king = game.GetFigureByIndex(game.IsCheckWhite.KingGameID)
	} else {
		king = game.GetFigureByIndex(game.IsCheckBlack.KingGameID)
	}

	fig := game.GetFigureByFieldCoordinates(cellCoordinate)

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

func (game *Game) CheckPawnAttack(indexKing int) bool {

	var king *Figure

	if game.Side {
		king = game.GetFigureByIndex(game.IsCheckWhite.KingGameID)
	} else {
		king = game.GetFigureByIndex(game.IsCheckBlack.KingGameID)
	}

	crd := IndexToFieldCoordinates(indexKing)

	if (*king).IsItWhite() && IsOnRealBoard([2]int{crd[0], crd[1] + 1}) {

		if fig := game.GetFigureByFieldCoordinates([2]int{crd[0], crd[1] + 1}); fig != nil {

			if (*fig).IsItWhite() != (*king).IsItWhite() {
				return true
			}
		}
	}

	if (*king).IsItWhite() && IsOnRealBoard([2]int{crd[0], crd[1] - 1}) {
		if fig := game.GetFigureByFieldCoordinates([2]int{crd[0], crd[1] - 1}); fig != nil {
			if (*fig).IsItWhite() != (*king).IsItWhite() {
				return true
			}
		}
	}

	if !(*king).IsItWhite() && IsOnRealBoard([2]int{crd[0] + 1, crd[1]}) {
		if fig := game.GetFigureByFieldCoordinates([2]int{crd[0] + 1, crd[1]}); fig != nil {
			if (*fig).IsItWhite() != (*king).IsItWhite() {
				return true
			}
		}
	}

	if !(*king).IsItWhite() && IsOnRealBoard([2]int{crd[0] - 1, crd[1]}) {
		if fig := game.GetFigureByFieldCoordinates([2]int{crd[0] - 1, crd[1]}); fig != nil {
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

	figureTo = g.GetFigureByIndex(to)
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

func (g *Game) NewFigure(to int, newFigure byte) bool {
	figure := g.GetFigureByIndex(to)

	if (*figure).GetType() == 'p' {
		if (*figure).IsItWhite() {
			if to < 8 {
				if !isNewFigureCorrect(newFigure) {
					return false
				}

				(*g.Figures[to]).ChangeType(newFigure)

				(*figure).ChangeType(newFigure)

				g.NewFigureId = mutateNewFigureId(newFigure, (*figure).IsItWhite())
			}
		} else {
			if to > 55 {
				if !isNewFigureCorrect(newFigure) {
					return false
				}
				(*g.Figures[to]).ChangeType(newFigure)

				(*figure).ChangeType(newFigure)

				g.NewFigureId = mutateNewFigureId(newFigure, (*figure).IsItWhite())
			}
		}
	}

	return true
}

func isNewFigureCorrect(newFigure byte) bool {
	switch newFigure {
	case 'k':
		return true
	case 'h':
		return true
	case 'a':
		return true
	case 'q':
		return true
	case 'b':
		return true
	default:
		return false
	}
}

var TheoryKnightSteps = []int{
	(2 * 8) + 1,
	(2 * 8) - 1,
	(-1)*(2*8) + 1,
	(-1)*(2*8) - 1,
	8 + 2,
	-8 + 2,
	8 - 2,
	-8 - 2,
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
