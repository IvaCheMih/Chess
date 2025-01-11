package move_tests

import (
	"github.com/IvaCheMih/chess/src/domains/game"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCorrectEnPass(t *testing.T) {
	boardRepo := game.CreateBoardCellsRepository(nil)
	gameService := game.CreateGamesService(&boardRepo, nil, nil)

	t.Run("Test en pass 1", func(t *testing.T) {
		board := boardRepo.MakeBoardCells(1, makeEnPass1())

		var cells = map[int]*models.BoardCell{}

		for i := range board {
			cells[board[i].IndexCell] = &board[i]
		}

		gameModel := models.Game{
			Id:                 1,
			WhiteUserId:        1,
			BlackUserId:        2,
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
			LastPawnMove:       &lastPawnMove1,
			Side:               false,
		}

		indexes, _ := gameService.MoveService.IsMoveCorrect(
			gameModel,
			models.Board{Cells: cells},
			35,
			44,
			0,
		)
		require.Equal(t, expected1, indexes)
	})

	t.Run("Test en pass 2", func(t *testing.T) {
		board := boardRepo.MakeBoardCells(1, makeEnPass2())

		var cells = map[int]*models.BoardCell{}

		for i := range board {
			cells[board[i].IndexCell] = &board[i]
		}

		gameModel := models.Game{
			Id:                 1,
			WhiteUserId:        1,
			BlackUserId:        2,
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
			LastPawnMove:       nil,
			Side:               false,
		}

		indexes, _ := gameService.MoveService.IsMoveCorrect(gameModel, models.Board{Cells: cells}, 35, 44, 0)
		require.Equal(t, expected2, indexes)
	})
}

var lastPawnMove1 = 36

func makeEnPass1() [][]int {
	return [][]int{
		{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
		{8, 13}, {9, 13}, {10, 13}, {35, 13}, {12, 13}, {13, 11}, {14, 13}, {15, 13},
		{48, 6}, {49, 6}, {50, 6}, {51, 6}, {36, 6}, {53, 13}, {54, 6}, {55, 6},
		{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
	}
}

var expected1 = []int{35, 44, -1, 36}

//var lastPawnMove2 = 0

func makeEnPass2() [][]int {
	return [][]int{
		{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
		{8, 13}, {9, 13}, {10, 13}, {35, 13}, {12, 13}, {13, 11}, {14, 13}, {15, 13},
		{48, 6}, {49, 6}, {50, 6}, {51, 6}, {36, 6}, {53, 13}, {54, 6}, {55, 6},
		{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
	}
}

var expected2 = []int{}
