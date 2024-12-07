package services

import (
	gameservice "github.com/IvaCheMih/chess/src/domains/game"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/IvaCheMih/chess/src/domains/game/services/move"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEndgame(t *testing.T) {
	boardRepo := gameservice.CreateBoardCellsRepository(nil)
	gameService := gameservice.CreateGamesService(&boardRepo, nil, nil)

	t.Run("Test start field", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.NewStartBoardCells(1)

		var cells = map[int]*models.BoardCell{}

		for i := range board {
			cells[board[i].IndexCell] = &board[i]
		}

		game := m.CreateGameStruct(models.Game{
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
			LastPawnMove:       nil,
			Side:               true,
		}, models.Board{Cells: cells})

		isEnd, endgameReason := game.IsItEndgame()
		require.False(t, isEnd)
		require.Equal(t, endgameReason, move.NotEndgame)
	})

	t.Run("Test white mate 1", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeWhiteCheckMatBoard1())

		var cells = map[int]*models.BoardCell{}

		for i := range board {
			cells[board[i].IndexCell] = &board[i]
		}

		game := m.CreateGameStruct(models.Game{
			Id:                 1,
			WhiteUserId:        1,
			BlackUserId:        1,
			IsStarted:          true,
			IsEnded:            false,
			IsCheckWhite:       true,
			WhiteKingCastling:  false,
			WhiteRookACastling: false,
			WhiteRookHCastling: false,
			IsCheckBlack:       false,
			BlackKingCastling:  false,
			BlackRookACastling: false,
			BlackRookHCastling: false,
			LastPawnMove:       nil,
			Side:               true,
		}, models.Board{Cells: cells})

		isEnd, endgameReason := game.IsItEndgame()
		require.True(t, isEnd)
		require.Equal(t, endgameReason, move.Mate)
	})

	t.Run("Test black pat 1", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeBlackPatBoard1())

		var cells = map[int]*models.BoardCell{}

		for i := range board {
			cells[board[i].IndexCell] = &board[i]
		}

		game := m.CreateGameStruct(models.Game{
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
			LastPawnMove:       nil,
			Side:               false,
		}, models.Board{Cells: cells})

		isEnd, endgameReason := game.IsItEndgame()
		require.True(t, isEnd)
		require.Equal(t, endgameReason, move.Pat)
	})
}

func makeWhiteCheckMatBoard1() [][]int {
	return [][]int{
		{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
		{8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13}, {13, 11}, {14, 13}, {15, 13},
		{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 13}, {54, 6}, {55, 6},
		{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
	}
}

func makeBlackPatBoard1() [][]int {
	return [][]int{
		{0, 12},
		{13, 7},
		{57, 1}, {60, 5},
	}
}
