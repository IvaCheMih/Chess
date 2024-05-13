package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"math"
)

type Game struct {
	N                     int
	M                     int
	WhiteClientId         *int
	BlackClientId         *int
	Figures               map[int]*Figure
	IsCheckWhite          IsCheck
	IsCheckBlack          IsCheck
	WhiteCastling         WhiteCastling
	BlackCastling         BlackCastling
	LastPawnMove          int
	Side                  int
	RookNewIdIfItCastling int
	RookOldIdIfItCastling int
}

type IsCheck struct {
	IsItCheck  bool
	KingGameID int
}

type WhiteCastling struct {
	WhiteKingCastling  bool
	WhiteRookACastling bool
	WhiteRookHCastling bool
}

type BlackCastling struct {
	BlackKingCastling  bool
	BlackRookACastling bool
	BlackRookHCastling bool
}

var FigureRepo = make(map[int]byte)

func CreateGameStruct(game dto.CreateGameResponse, board models.Board) Game {

	figures, blackKingCell, whiteKingCell := CreateField(board, game)

	return Game{
		N:             8,
		M:             12,
		WhiteClientId: &game.WhiteUserId,
		BlackClientId: &game.BlackUserId,
		Figures:       figures,
		IsCheckWhite:  IsCheck{game.IsCheckWhite, whiteKingCell},
		IsCheckBlack:  IsCheck{game.IsCheckBlack, blackKingCell},
		WhiteCastling: WhiteCastling{game.WhiteKingCastling, game.WhiteRookACastling, game.WhiteRookHCastling},
		BlackCastling: BlackCastling{game.BlackKingCastling, game.BlackRookACastling, game.BlackRookHCastling},
		LastPawnMove:  game.LastPawnMove,
		Side:          game.Side,
	}
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

func (game *Game) GetFigureByIndex(index int) *Figure {
	return game.Figures[index]
}

func (game *Game) GetFigureByFieldCoordinates(crd []int) *Figure {
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

func (game *Game) ChangeKingGameID(figure *Figure) {
	if (*figure).GetType() != 'K' {
		return
	}
	if (*figure).IsWhite() {
		game.IsCheckWhite.KingGameID = FieldCoordinatesToIndex((*figure).GetGameIndex())
	} else {
		game.IsCheckBlack.KingGameID = FieldCoordinatesToIndex((*figure).GetGameIndex())
	}
}

func (game *Game) CheckIsCheck() bool {
	if game.Side == *game.WhiteClientId && game.IsKingCheck(game.IsCheckWhite.KingGameID) {
		return true
	}

	if game.Side == *game.BlackClientId && game.IsKingCheck(game.IsCheckBlack.KingGameID) {
		return true
	}

	return false
}

func (game *Game) ChangeCastlingFlag(figure *Figure) {
	switch (*figure).GetType() {
	case 'K':
		if (*figure).IsWhite() {
			game.WhiteCastling.WhiteKingCastling = true
		} else {
			game.BlackCastling.BlackKingCastling = true
		}
	case 'a':
		if (*figure).IsWhite() {
			game.WhiteCastling.WhiteRookACastling = true
		} else {
			game.BlackCastling.BlackRookACastling = true
		}
	case 'h':
		if (*figure).IsWhite() {
			game.WhiteCastling.WhiteRookHCastling = true
		} else {
			game.BlackCastling.BlackRookHCastling = true
		}
	}
}

func (game *Game) ChangeLastPawnMove(figure *Figure, from int, to int) {
	if (*figure).GetType() == 'p' && math.Abs(float64(from-to)) > 9 {

		game.LastPawnMove = to

		return
	}

	game.LastPawnMove = -1
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
				if (*fig).IsWhite() != (*king).IsWhite() {
					if (*king).IsWhite() {
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

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1] + i}); i++ {
		fmt.Println([]int{crd[0] + i, crd[1] + i})
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] + i, crd[1] + i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1] - i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] + i, crd[1] - i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1] + i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] - i, crd[1] + i}, 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1] - i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] - i, crd[1] - i}, 'b')
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

	for i := 1; IsOnRealBoard([]int{crd[0], crd[1] + i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0], crd[1] + i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0], crd[1] - i}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0], crd[1] - i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1]}); i++ {

		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] + i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1]}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] - i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}
	return false
}

func (game *Game) CheckAttackCell(kingCoordinate []int, cellCoordinate []int, triggerFigure byte) (bool, bool) {

	var king *Figure

	if game.Side == *game.WhiteClientId {
		king = game.GetFigureByIndex(game.IsCheckWhite.KingGameID)
	} else {
		king = game.GetFigureByIndex(game.IsCheckBlack.KingGameID)
	}

	fig := game.GetFigureByFieldCoordinates(cellCoordinate)

	if fig == nil {
		return false, false
	}
	if (*fig).IsWhite() == (*king).IsWhite() {
		return false, true
	}
	if (*fig).IsWhite() != (*king).IsWhite() {
		if (*fig).GetType() == triggerFigure || (*fig).GetType() == 'q' {
			return true, true
		}
		return false, true
	}
	return false, false
}

func (game *Game) CheckPawnAttack(indexKing int) bool {

	var king *Figure

	if game.Side == *game.WhiteClientId {
		king = game.GetFigureByIndex(game.IsCheckWhite.KingGameID)
	} else {
		king = game.GetFigureByIndex(game.IsCheckBlack.KingGameID)
	}

	crd := IndexToFieldCoordinates(indexKing)

	if (*king).IsWhite() && IsOnRealBoard([]int{crd[0], crd[1] + 1}) {

		if fig := game.GetFigureByFieldCoordinates([]int{crd[0], crd[1] + 1}); fig != nil {

			if (*fig).IsWhite() != (*king).IsWhite() {
				return true
			}
		}
	}

	if (*king).IsWhite() && IsOnRealBoard([]int{crd[0], crd[1] - 1}) {
		if fig := game.GetFigureByFieldCoordinates([]int{crd[0], crd[1] - 1}); fig != nil {
			if (*fig).IsWhite() != (*king).IsWhite() {
				return true
			}
		}
	}

	if !(*king).IsWhite() && IsOnRealBoard([]int{crd[0] + 1, crd[1]}) {
		if fig := game.GetFigureByFieldCoordinates([]int{crd[0] + 1, crd[1]}); fig != nil {
			if (*fig).IsWhite() != (*king).IsWhite() {
				return true
			}
		}
	}

	if !(*king).IsWhite() && IsOnRealBoard([]int{crd[0] - 1, crd[1]}) {
		if fig := game.GetFigureByFieldCoordinates([]int{crd[0] - 1, crd[1]}); fig != nil {
			if (*fig).IsWhite() != (*king).IsWhite() {
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

	crdLPM := []int{}

	if g.LastPawnMove != -1 {
		crdLPM = IndexToFieldCoordinates(g.LastPawnMove)
	}

	if figureTo != nil {
		(*figureTo).Delete()
	} else {
		if (*figureFrom).GetType() == 'p' && g.LastPawnMove != -1 && coordinateFrom[0] == crdLPM[0] && coordinateTo[1] == crdLPM[1] {
			pawn := g.GetFigureByFieldCoordinates(crdLPM)
			(*pawn).Delete()
		}
	}

	(*figureFrom).ChangeGameIndex(coordinateTo)

	g.Figures[to] = g.Figures[from]
	g.Figures[from] = nil

	figureTo = g.GetFigureByIndex(to)
}

func (g *Game) ChangeRookIfCastling(to int) {
	switch to {
	case 2:
		g.ChangeToAndFrom(3, 0)
		g.RookNewIdIfItCastling = 3
		g.RookOldIdIfItCastling = 0
	case 6:
		g.ChangeToAndFrom(5, 7)
		g.RookNewIdIfItCastling = 5
		g.RookOldIdIfItCastling = 7
	case 57:
		g.ChangeToAndFrom(59, 56)
		g.RookNewIdIfItCastling = 59
		g.RookOldIdIfItCastling = 56
	case 62:
		g.ChangeToAndFrom(61, 63)
		g.RookNewIdIfItCastling = 61
		g.RookOldIdIfItCastling = 63
	}
}

func (g *Game) IsItYourFigure(figure *Figure) bool {
	if figure == nil {
		return false
	}

	if *g.WhiteClientId == g.Side && !(*figure).IsWhite() {
		return false
	}

	if *g.BlackClientId == g.Side && (*figure).IsWhite() {
		return false
	}

	return true
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
