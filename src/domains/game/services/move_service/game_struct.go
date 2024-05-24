package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"math"
)

type Game struct {
	N int
	//WhiteClientId *int
	//BlackClientId *int
	Figures       map[int]*Figure
	IsCheckWhite  IsCheck
	IsCheckBlack  IsCheck
	WhiteCastling Castling
	BlackCastling Castling
	LastPawnMove  *int
	Side          bool
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

func (game *Game) ChangeCastlingFlag(figure *Figure) {
	switch (*figure).GetType() {
	case 'K':
		if (*figure).IsWhite() {
			game.WhiteCastling.KingCastling = true
		} else {
			game.BlackCastling.KingCastling = true
		}
	case 'a':
		if (*figure).IsWhite() {
			game.WhiteCastling.RookACastling = true
		} else {
			game.BlackCastling.RookACastling = true
		}
	case 'h':
		if (*figure).IsWhite() {
			game.WhiteCastling.RookHCastling = true
		} else {
			game.BlackCastling.RookHCastling = true
		}
	}
}

func (game *Game) ChangeLastPawnMove(figure *Figure, from int, to int) {
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

	if game.Side {
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

	if game.Side {
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

	if g.Side && !(*figure).IsWhite() {
		return false
	}

	if !g.Side && (*figure).IsWhite() {
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

	g.ChangeToAndFrom(indexesToChange[2], indexesToChange[3])
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
