package main

import (
	"fmt"
	"strconv"
)

type Game struct {
	N             int
	WhiteClientId *int
	BlackClientId *int
	Figures       []*Figure
	History       []*Step
	IsStarted     bool
	IsEnded       bool
	IsCheckWhite  IsCheck
	IsCheckBlack  IsCheck
}

type IsCheck struct {
	IsItCheck  bool
	KingGameID int
}

func CreateGame(N int) Game {
	return Game{
		N:             N,
		WhiteClientId: nil,
		BlackClientId: nil,
		Figures:       CreateDefaultField(),
		History:       []*Step{},
		IsStarted:     false,
		IsEnded:       false,
		IsCheckWhite:  IsCheck{false, 60},
		IsCheckBlack:  IsCheck{false, 4},
	}
}

//

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

func (game *Game) GetFigureByIndex(index int) *Figure {
	return game.Figures[index]
}

func (game *Game) GetFigureByCoordinates(coordinates string) *Figure {
	index := game.CoordinatesToIndex(coordinates)

	return game.Figures[index]
}

func (game *Game) GetBoard() string {
	fieldString := ""

	for index, figure := range game.Figures {
		if (index % game.N) == 0 {
			fieldString += "\n"
		}

		fieldString += FigureToString(figure)
	}

	return fieldString
}

func (game *Game) IsGameMember(clientId int) bool {
	if game.WhiteClientId != nil && clientId == *game.WhiteClientId {
		return true
	}
	if game.BlackClientId != nil && clientId == *game.BlackClientId {
		return true
	}
	return false
}

func (game *Game) GetMoveSide() string {
	if game.IsEnded {
		return ""
	}
	if len(game.History)%2 == 0 {

		return "white"
	}
	return "black"
}

func (game *Game) GetHistoryString() string {
	history := ""

	for i, step := range game.History {
		history += "\n" + "Ход " + strconv.Itoa(i) + ": " + (*(*step).figure).ToString() + " - " + game.IndexToCoordinates((*step).from) + " - " + game.IndexToCoordinates((*step).to)
	}
	return history
}

func (game *Game) DoStep(message string) bool {
	from, to := ParseMessageToMove(message)

	fromIndex := game.CoordinatesToIndex(from)
	toIndex := game.CoordinatesToIndex(to)

	figureFrom := game.Figures[fromIndex]
	figureTo := game.Figures[toIndex]

	game.Figures[toIndex] = game.Figures[fromIndex]
	game.Figures[fromIndex] = nil
	figure := game.GetFigureByIndex(toIndex)
	(*figure).ChangeGameIndex(toIndex)

	game.ChangeKingGameID(figure)
	if game.CheckIsCheck() {
		fmt.Println("Король под шахом")
		game.Figures[toIndex] = figureTo
		game.Figures[fromIndex] = figureFrom
		fmt.Println("поменяли фигуры назад")
		fig := game.GetFigureByIndex(fromIndex)
		(*fig).ChangeGameIndex(fromIndex)
		fmt.Println("поменяли индексы назад")
		game.ChangeKingGameID(figure)
		fmt.Println("поменяли индексы назад")
		return false
	}
	game.AddStepToHistory(from, to)
	return true
}

func (game *Game) AddStepToHistory(from string, to string) {
	figure := game.GetFigureByCoordinates(to)
	//killedFigure := game.GetFigureByCoordinates(to)

	var step = Step{
		figure:       figure,
		from:         game.CoordinatesToIndex(from),
		to:           game.CoordinatesToIndex(to),
		killedFigure: nil,
		newFigure:    nil,
	}

	game.History = append(game.History, &step)
}

func (game *Game) CheckIsCheck() bool {
	color := game.GetMoveSide()
	if color == "white" && game.IsKingCheck(game.IsCheckWhite.KingGameID) {
		return true
	}
	if color == "black" && game.IsKingCheck(game.IsCheckBlack.KingGameID) {
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
	for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N+1)); i++ {
		isCheck, endFor := game.CheckAttackCell(index, index+i*(game.N+1), 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; game.CheckCellOnBoardByIndex(index + i*(game.N-1)); i++ {
		isCheck, endFor := game.CheckAttackCell(index, index+i*(game.N-1), 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N-1)); i++ {
		isCheck, endFor := game.CheckAttackCell(index, index-i*(game.N-1), 'b')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	for i := 1; game.CheckCellOnBoardByIndex(index - i*(game.N+1)); i++ {
		isCheck, endFor := game.CheckAttackCell(index, index-i*(game.N+1), 'b')
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
	fmt.Println("Начало верт проверки")
	for i := 1; game.CheckCellOnBoardByIndex(index + i); i++ {
		fmt.Println(index + i)
		isCheck, endFor := game.CheckAttackCell(index, index+i, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	fmt.Println("направо нет шаха")
	for i := 1; game.CheckCellOnBoardByIndex(index - i); i++ {
		isCheck, endFor := game.CheckAttackCell(index, index-i, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}

	fmt.Println("налево нет шаха")
	for i := 1; game.CheckCellOnBoardByIndex(index + i*game.N); i++ {
		isCheck, endFor := game.CheckAttackCell(index, index+i*game.N, 'r')
		if isCheck {
			return true
		}
		if endFor {
			break
		}
	}
	fmt.Println("вверх нет шаха")
	for i := 1; game.CheckCellOnBoardByIndex(index - i*game.N); i++ {
		isCheck, endFor := game.CheckAttackCell(index, index-i*game.N, 'r')
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

func (game *Game) CheckAttackCell(indexKing int, indexCell int, triggerFigure byte) (bool, bool) {

	king := game.GetFigureByIndex(indexKing)

	if game.GetFigureByIndex(indexCell) == nil {
		return false, false
	}
	if fig := game.GetFigureByIndex(indexCell); (*fig).IsWhite() == (*king).IsWhite() {
		return false, true
	}
	if fig := game.GetFigureByIndex(indexCell); (*fig).IsWhite() != (*king).IsWhite() {
		if (*fig).GetType() == triggerFigure || (*fig).GetType() == 'q' {
			return true, true
		}
		return false, true
	}
	return false, false
}

func (game *Game) ChangeKingGameID(figure *Figure) {
	if (*figure).GetType() != 'K' {
		return
	}
	if (*figure).IsWhite() {
		game.IsCheckWhite.KingGameID = (*figure).GetGameIndex()
	} else {
		game.IsCheckBlack.KingGameID = (*figure).GetGameIndex()
	}
}

func (game *Game) CheckPawnAttack(indexKing int) bool {
	fmt.Println("Начало пешечной проверки")
	king := game.GetFigureByIndex(indexKing)

	if (*king).IsWhite() && game.CheckCellOnBoardByIndex(indexKing+game.N+1) {
		if fig := game.GetFigureByIndex(indexKing + game.N + 1); fig != nil {
			if (*fig).IsWhite() != (*king).IsWhite() {
				return true
			}
		}
	}

	if (*king).IsWhite() && game.CheckCellOnBoardByIndex(indexKing+game.N-1) {
		if fig := game.GetFigureByIndex(indexKing + game.N - 1); fig != nil {
			if (*fig).IsWhite() != (*king).IsWhite() {
				return true
			}
		}
	}

	if !(*king).IsWhite() && game.CheckCellOnBoardByIndex(indexKing-game.N-1) {
		if fig := game.GetFigureByIndex(indexKing - game.N - 1); fig != nil {
			if (*fig).IsWhite() != (*king).IsWhite() {
				return true
			}
		}
	}

	if !(*king).IsWhite() && game.CheckCellOnBoardByIndex(indexKing-game.N+1) {
		if fig := game.GetFigureByIndex(indexKing - game.N + 1); fig != nil {
			if (*fig).IsWhite() != (*king).IsWhite() {
				return true
			}
		}
	}
	fmt.Println("Начало пешечной проверки")
	return false
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
