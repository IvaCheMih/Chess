package move

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"log"
)

type MoveService struct {
	figureRepo map[int]byte
}

func NewMoveService(figureRepo map[int]byte) *MoveService {
	return &MoveService{
		figureRepo: figureRepo,
	}
}

func (m *MoveService) createGameStruct(gameModel models.Game, board models.Board) Game {
	figures, blackKingCell, whiteKingCell := m.CreateField(board, gameModel)

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

func (m *MoveService) IsMoveCorrect(gameModel models.Game, board models.Board, from int, to int) ([]int, Game) {
	game := m.createGameStruct(gameModel, board)

	figure := game.GetFigureByIndex(from)

	if !game.IsItYourFigure(figure) {
		return []int{}, Game{}
	}

	possibleMoves := (*figure).GetPossibleMoves(&game)

	isCorrect, indexesToChange := CheckMove(possibleMoves, []int{from, to})
	if !isCorrect {
		return []int{}, Game{}
	}

	return indexesToChange, game
}

func IsItCheck(indexesToChange []int, game *Game, newFigure byte) bool {
	from := indexesToChange[0]
	to := indexesToChange[1]

	//game := CreateGameStruct(gameModel, board)

	game.ChangeToAndFrom(to, from)

	if len(indexesToChange) > 2 {
		game.DeletePawn(indexesToChange)
		game.ChangeRookField(indexesToChange)
	}

	game.ChangeKingGameID(to)

	if !game.NewFigure(to, newFigure) {
		return false
	}

	if game.Check() {
		return false
	}

	game.ChangeCastlingFlag(to)

	game.ChangeLastPawnMove(from, to)

	return true
}

func CheckMove(possibleMoves *TheoryMoves, coordinatesToChange []int) (bool, []int) {
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
				crdRook := GetNewRookCoordinatesIfCastling((coordinatesToChange)[1])
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

	log.Println("Запрашиваемого хода нет в массиве")
	return false, []int{}
}

func GetNewRookCoordinatesIfCastling(to int) []int {
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
