package move_service

import (
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
)

type Game struct {
	N             int
	M             int
	WhiteClientId *int
	BlackClientId *int
	Figures       map[int]*Figure
	IsCheckWhite  IsCheck
	IsCheckBlack  IsCheck
	Side          int
}

type IsCheck struct {
	IsItCheck  bool
	KingGameID int
}

var FigureRepo = make(map[int]byte)

func CreateGameStruct(game dto.CreateGameResponse, board models.Board) Game {

	return Game{
		N:             8,
		M:             12,
		WhiteClientId: &game.WhiteUserId,
		BlackClientId: &game.BlackUserId,
		Figures:       CreateField(board),
		IsCheckWhite:  IsCheck{game.IsCheckWhite, game.WhiteKingCell},
		IsCheckBlack:  IsCheck{game.IsCheckBlack, game.BlackKingCell},
		Side:          game.Side,
	}
}

func CreateFigureRepo() map[int]byte {
	var figureRepo = make(map[int]byte)

	figureRepo[1] = 'r'
	figureRepo[2] = 'k'
	figureRepo[3] = 'b'
	figureRepo[4] = 'q'
	figureRepo[5] = 'K'
	figureRepo[6] = 'p'

	figureRepo[7] = 'r'
	figureRepo[8] = 'k'
	figureRepo[9] = 'b'
	figureRepo[10] = 'q'
	figureRepo[11] = 'K'
	figureRepo[12] = 'p'

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

func (game *Game) IsKingCheck(index int) bool {
	fmt.Println("начало проверки аттак на короля")
	if game.CheckKnightAttack(index) {
		return true
	}
	fmt.Println("Король не находится под атакой коня")

	if game.CheckDiagonalAttack(index) {
		fmt.Println("проверка диаг атаки")
		return true
	}

	if game.CheckVertGorAttack(index) {
		fmt.Println("проверка верт атаки")
		return true
	}

	if game.CheckPawnAttack(index) {
		fmt.Println("проверка пешечной атаки")
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
	fmt.Println("Начало диаг проверки")

	crd := IndexToFieldCoordinates(index)

	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1] + i}); i++ {
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
	fmt.Println("Конец диаг проверки")
	return false
}

func (game *Game) CheckVertGorAttack(index int) bool {
	crd := IndexToFieldCoordinates(index)

	fmt.Println("Начало верт проверки")
	for i := 1; IsOnRealBoard([]int{crd[0], crd[1] + i}); i++ {
		fmt.Println(index + i)
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0], crd[1] + i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	fmt.Println("направо нет шаха")
	for i := 1; IsOnRealBoard([]int{crd[0], crd[1] - i}); i++ {
		fmt.Println(500)
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0], crd[1] - i}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	fmt.Println("налево нет шаха")
	for i := 1; IsOnRealBoard([]int{crd[0] + i, crd[1]}); i++ {

		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] + i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}
	fmt.Println("вверх нет шаха")
	for i := 1; IsOnRealBoard([]int{crd[0] - i, crd[1]}); i++ {
		isCheck, endFor := game.CheckAttackCell(crd, []int{crd[0] - i, crd[1]}, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}
	fmt.Println("вниз нет шаха")
	fmt.Println("конец верт проверки")
	return false
}

func (game *Game) CheckAttackCell(kingCoordinate []int, cellCoordinate []int, triggerFigure byte) (bool, bool) {

	king := game.GetFigureByFieldCoordinates(kingCoordinate)
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
	fmt.Println("Начало пешечной проверки")
	crd := IndexToFieldCoordinates(indexKing)
	king := game.GetFigureByFieldCoordinates(crd)

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
	fmt.Println("Начало пешечной проверки")
	return false
}

func (g *Game) ChangeToAndFrom(to int, from int) {

	coordinateTo := IndexToFieldCoordinates(to)
	coordinateFrom := IndexToFieldCoordinates(from)

	figureTo := g.GetFigureByFieldCoordinates(coordinateTo)

	if figureTo != nil {
		(*figureTo).Delete()
	}

	figureFrom := g.GetFigureByFieldCoordinates(coordinateFrom)

	(*figureFrom).ChangeGameIndex(coordinateTo)

	g.Figures[to] = g.Figures[from]
	g.Figures[from] = nil

	figureTo = g.GetFigureByIndex(to)
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
