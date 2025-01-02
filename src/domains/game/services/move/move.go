package move

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
)

type MoveService struct {
	figureRepo        map[int]byte
	newFigures        map[byte]struct{}
	theoryKnightSteps *[]int
}

type EndgameReason string

const (
	NotEndgame EndgameReason = ""
	Mate       EndgameReason = "Mate"
	Pat        EndgameReason = "Pat"
	Repetition EndgameReason = "Repetition"
	NoLosses   EndgameReason = "NoLosses"
)

func NewMoveService(figureRepo map[int]byte) *MoveService {
	return &MoveService{
		figureRepo:        figureRepo,
		newFigures:        MakeNewFigures(),
		theoryKnightSteps: MakeTheoryKnightSteps(),
	}
}

func MakeNewFigures() map[byte]struct{} {
	return map[byte]struct{}{
		'k': {},
		'h': {},
		'a': {},
		'q': {},
		'b': {},
	}
}

func MakeTheoryKnightSteps() *[]int {
	return &[]int{
		(2 * 8) + 1,
		(2 * 8) - 1,
		(-1)*(2*8) + 1,
		(-1)*(2*8) - 1,
		8 + 2,
		-8 + 2,
		8 - 2,
		-8 - 2,
	}
}

func (m *MoveService) CreateGameStruct(gameModel models.Game, board models.Board) Game {
	figures, blackKingCell, whiteKingCell := m.createField(board, gameModel)

	side := gameModel.Side

	return Game{
		N: 8,
		//WhiteClientId: &gameModel.WhiteUserId,
		//BlackClientId: &gameModel.BlackUserId,
		Figures:           figures,
		IsCheckWhite:      IsCheck{gameModel.IsCheckWhite, whiteKingCell},
		IsCheckBlack:      IsCheck{gameModel.IsCheckBlack, blackKingCell},
		WhiteCastling:     Castling{gameModel.WhiteKingCastling, gameModel.WhiteRookACastling, gameModel.WhiteRookHCastling},
		BlackCastling:     Castling{gameModel.BlackKingCastling, gameModel.BlackRookACastling, gameModel.BlackRookHCastling},
		LastPawnMove:      gameModel.LastPawnMove,
		Side:              side,
		NewFigureId:       0,
		newFigures:        m.newFigures,
		theoryKnightSteps: m.theoryKnightSteps,
		LastLoss:          gameModel.LastLoss,
	}
}

func (m *MoveService) IsMoveCorrect(gameModel models.Game, board models.Board, from int, to int, newFigure byte) ([]int, Game) {
	game := m.CreateGameStruct(gameModel, board)

	figure := game.GetFigureByIndex(from)

	if !game.IsItYourFigure(figure) {
		return []int{}, Game{}
	}

	// theory moves for this figure ("to" is on board)
	possibleMoves := (*figure).GetPossibleMoves(&game)

	// requested move is possible (is in possibleMoves)
	isCorrect, indexesToChange := checkMove(possibleMoves, []int{from, to})
	if !isCorrect {
		return []int{}, Game{}
	}

	if (*figure).GetType() == 'p' && game.isNewFigureCorrect(newFigure) {
		if !game.NewFigureRequestCorrect(to, (*figure).IsItWhite()) {
			return []int{}, Game{}
		}
	}

	return indexesToChange, game
}

func (m *MoveService) createField(board models.Board, gameModel models.Game) (map[int]*Figure, int, int) {
	blackKingCell, whiteKingCell := 0, 0
	field := map[int]*Figure{}

	for _, cell := range board.Cells {
		if cell.FigureId == 5 {
			whiteKingCell = cell.IndexCell

		}
		if cell.FigureId == 12 {
			blackKingCell = cell.IndexCell

		}

		isWhite := cell.FigureId <= 7

		field[cell.IndexCell] = createFigureI(m.figureRepo[cell.FigureId], isWhite, cell.IndexCell, gameModel)

	}

	return field, blackKingCell, whiteKingCell
}

func (m *MoveService) GetFigureRepo() map[int]byte {
	return m.figureRepo
}

func (m *MoveService) GetFigureID(b byte) int {
	for i := range m.figureRepo {
		if m.figureRepo[i] == b {
			return i
		}
	}

	return 0
}

func (m *MoveService) IsItEndgame(g *Game, history []models.Move, board []models.BoardCell) (bool, EndgameReason) {
	var cells = map[int]*models.BoardCell{}

	for i := range board {
		cells[board[i].IndexCell] = &board[i]
	}

	gameHistory := m.CreateGameStruct(models.Game{
		Id:                 1,
		WhiteUserId:        1,
		BlackUserId:        1,
		IsStarted:          true,
		IsEnded:            false,
		IsCheckWhite:       false,
		WhiteKingCastling:  false,
		WhiteRookACastling: false,
		WhiteRookHCastling: false,
		IsCheckBlack:       false,
		BlackKingCastling:  false,
		BlackRookACastling: false,
		BlackRookHCastling: false,
		LastLoss:           0,
		LastPawnMove:       nil,
		Side:               true,
	}, models.Board{Cells: cells})

	count := 0

	for i := range history {
		gameHistory.DoMoveFromHistory(history[i], m.figureRepo[history[i].NewFigureId])

		if g.CompareGamesStates(gameHistory) {
			count++
			if count == 2 {
				return true, Repetition
			}
		}
	}

	for _, figure := range g.Figures {
		if !g.IsItYourFigure(figure) {
			continue
		}

		fromCrd := (*figure).GetCoordinates()
		theoryMoves := (*figure).GetPossibleMoves(g)

		if g.movesExist(theoryMoves, fromCrd) {
			if g.LastLoss+1 == lastLossLimit {
				return true, NoLosses
			}

			return false, NotEndgame
		}
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
}

func (g *Game) DoMoveFromHistory(move models.Move, newFigure byte) {
	from := move.FromId
	to := move.ToId

	figure := g.GetFigureByIndex(from)

	possibleMoves := (*figure).GetPossibleMoves(g)

	_, indexesToChange := checkMove(possibleMoves, []int{from, to})

	DoMove(indexesToChange, g, newFigure)
}

func DoMove(indexesToChange []int, game *Game, newFigure byte) {
	from := indexesToChange[0]
	to := indexesToChange[1]

	//game := CreateGameStruct(gameModel, board)

	//for i := 0; i < 64; i++ {
	//	if i%8 == 0 {
	//		fmt.Println()
	//	}
	//	if game.Figures[i] != nil {
	//		fmt.Print(string((*game.Figures[i]).GetType()))
	//	} else {
	//		fmt.Printf("0")
	//	}
	//
	//}
	//fmt.Println()

	game.ChangeToAndFrom(to, from)

	if len(indexesToChange) > 2 {
		game.DeletePawn(indexesToChange)
		game.ChangeRookField(indexesToChange)
	}

	game.ChangeKingGameID(to)

	game.ChangePawnToNewFigure(to, newFigure)

	game.ChangeCastlingFlag(to)

	game.ChangeLastPawnMove(from, to)
}

func checkMove(possibleMoves *TheoryMoves, coordinatesToChange []int) (bool, []int) {
	crdFrom := IndexToFieldCoordinates((coordinatesToChange)[0])
	crdTo := IndexToFieldCoordinates((coordinatesToChange)[1])

	if possibleMoves.Up != nil {
		for _, pm := range possibleMoves.Up {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Down != nil {
		for _, pm := range possibleMoves.Down {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Right != nil {
		for _, pm := range possibleMoves.Right {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Left != nil {
		for _, pm := range possibleMoves.Left {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.UR != nil {
		for _, pm := range possibleMoves.UR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.UL != nil {
		for _, pm := range possibleMoves.UL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.DR != nil {
		for _, pm := range possibleMoves.DR {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.DL != nil {
		for _, pm := range possibleMoves.DL {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}
	if possibleMoves.Kn != nil {
		for _, pm := range possibleMoves.Kn {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				return true, coordinatesToChange
			}
		}
	}

	if possibleMoves.Castling != nil {
		for _, pm := range possibleMoves.Castling {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {
				crdRook := getNewRookCoordinatesIfCastling((coordinatesToChange)[1])
				coordinatesToChange = append(coordinatesToChange, crdRook[0])
				coordinatesToChange = append(coordinatesToChange, crdRook[1])
				return true, coordinatesToChange
			}
		}
	}

	if possibleMoves.EnPass != nil {
		for _, pm := range possibleMoves.EnPass {
			if pm[0] == crdTo[0] && pm[1] == crdTo[1] {

				coordinatesToChange = append(coordinatesToChange, -1)
				coordinatesToChange = append(coordinatesToChange, FieldCoordinatesToIndex([2]int{crdTo[0], crdFrom[1]}))

				return true, coordinatesToChange
			}
		}
	}

	return false, []int{}
}

func getNewRookCoordinatesIfCastling(to int) []int {
	crd := []int{}

	switch to {
	case 2:
		crd = append(crd, 0)
		crd = append(crd, 3)
	case 6:
		crd = append(crd, 7)
		crd = append(crd, 5)
	case 57:
		crd = append(crd, 56)
		crd = append(crd, 59)
	case 62:

		crd = append(crd, 63)
		crd = append(crd, 61)
	}

	return crd
}
