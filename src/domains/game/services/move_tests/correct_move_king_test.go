package move_tests

import (
	"github.com/IvaCheMih/chess/src/domains/game"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCorrectKingMove(t *testing.T) {
	boardRepo := game.CreateBoardCellsRepository(nil)
	gameService := game.CreateGamesService(&boardRepo, nil, nil)

	t.Run("Test king move 1", func(t *testing.T) {
		board := boardRepo.MakeBoardCells(1, makeKingMove1())

		var cells = map[int]*models.BoardCell{}

		for i := range board {
			cells[board[i].IndexCell] = &board[i]
		}

		gameModel := models.Game{
			Id:                 1,
			WhiteUserId:        1,
			BlackUserId:        2,
			Status:             models.Active,
			IsCheckWhite:       false,
			WhiteKingCastling:  false,
			WhiteRookACastling: false,
			WhiteRookHCastling: false,
			IsCheckBlack:       false,
			BlackKingCastling:  false,
			BlackRookACastling: false,
			BlackRookHCastling: false,
			LastPawnMove:       nil,
			Side:               false,
		}

		indexes, _ := gameService.MoveService.IsMoveCorrect(
			gameModel,
			models.Board{Cells: cells},
			4,
			6,
			0,
		)
		require.Equal(t, expectedKingMove1, indexes)
	})

	t.Run("Test king move 2", func(t *testing.T) {
		board := boardRepo.MakeBoardCells(1, makeKingMove2())

		var cells = map[int]*models.BoardCell{}

		for i := range board {
			cells[board[i].IndexCell] = &board[i]
		}

		gameModel := models.Game{
			Id:                 1,
			WhiteUserId:        1,
			BlackUserId:        2,
			Status:             models.Active,
			IsCheckWhite:       false,
			WhiteKingCastling:  false,
			WhiteRookACastling: false,
			WhiteRookHCastling: false,
			IsCheckBlack:       false,
			BlackKingCastling:  false,
			BlackRookACastling: false,
			BlackRookHCastling: false,
			LastPawnMove:       nil,
			Side:               false,
		}

		indexes, _ := gameService.MoveService.IsMoveCorrect(
			gameModel,
			models.Board{Cells: cells},
			4,
			2,
			0,
		)
		require.Equal(t, expectedKingMove2, indexes)
	})
}

func makeKingMove1() [][]int {
	return [][]int{
		{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 0}, {6, 0}, {7, 14},
		{8, 13}, {9, 13}, {10, 13}, {0, 13}, {12, 13}, {13, 11}, {14, 13}, {15, 13},
		{48, 6}, {49, 6}, {50, 6}, {51, 6}, {36, 6}, {53, 13}, {54, 6}, {55, 6},
		{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
	}
}

var expectedKingMove1 = []int{4, 6, 7, 5}

func makeKingMove2() [][]int {
	return [][]int{
		{0, 8}, {1, 0}, {2, 0}, {3, 0}, {4, 12}, {5, 0}, {6, 0}, {7, 14},
		{8, 13}, {9, 13}, {10, 13}, {11, 0}, {12, 13}, {13, 11}, {14, 13}, {15, 13},
		{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 13}, {54, 6}, {55, 6},
		{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
	}
}

var expectedKingMove2 = []int{4, 2, 0, 3}
