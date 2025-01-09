package game

import (
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEndgame(t *testing.T) {
	boardRepo := CreateBoardCellsRepository(nil)
	gameService := CreateGamesService(&boardRepo, nil, nil)

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
			LastLoss:           0,
			LastPawnMove:       nil,
			Side:               true,
		}, models.Board{Cells: cells})

		isEnd, endgameReason := gameService.moveService.IsItEndgame(&game, nil, boardRepo.NewStartBoardCells(1))
		require.False(t, isEnd)
		require.Equal(t, endgameReason, models.NotEndgame)
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

		isEnd, endgameReason := gameService.moveService.IsItEndgame(&game, nil, board)
		require.True(t, isEnd)
		require.Equal(t, endgameReason, models.Mate)
	})

	t.Run("Test black mate 1", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeBlackCheckMatBoard1())

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
			IsCheckBlack:       true,
			BlackKingCastling:  false,
			BlackRookACastling: false,
			BlackRookHCastling: false,
			LastPawnMove:       nil,
			Side:               false,
		}, models.Board{Cells: cells})

		isEnd, endgameReason := gameService.moveService.IsItEndgame(&game, nil, board)
		require.True(t, isEnd)
		require.Equal(t, endgameReason, models.Mate)
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

		isEnd, endgameReason := gameService.moveService.IsItEndgame(&game, nil, board)
		require.True(t, isEnd)
		require.Equal(t, endgameReason, models.Pat)
	})

	t.Run("Test no losses", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeNoMateNoPat())

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
			IsCheckBlack:       true,
			BlackKingCastling:  false,
			BlackRookACastling: false,
			BlackRookHCastling: false,
			LastPawnMove:       nil,
			LastLoss:           49,
			Side:               false,
		}, models.Board{Cells: cells})

		isEnd, endgameReason := gameService.moveService.IsItEndgame(&game, nil, boardRepo.NewStartBoardCells(1))
		require.True(t, isEnd)
		require.Equal(t, endgameReason, models.NoLosses)
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
func makeBlackCheckMatBoard1() [][]int {
	return [][]int{
		{0, 0}, {1, 4}, {2, 0}, {3, 0}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
		{8, 13}, {9, 0}, {10, 13}, {11, 13}, {12, 13}, {13, 13}, {14, 13}, {15, 13},
		{16, 0}, {17, 0}, {18, 0}, {19, 0}, {20, 0}, {21, 0}, {22, 0}, {23, 0},
		{24, 0}, {25, 0}, {26, 0}, {27, 0}, {28, 0}, {29, 0}, {30, 0}, {31, 0},
		{32, 0}, {33, 0}, {34, 0}, {35, 0}, {36, 0}, {37, 0}, {38, 0}, {39, 0},
		{40, 0}, {41, 0}, {42, 0}, {43, 0}, {44, 0}, {45, 0}, {46, 0}, {47, 0},
		{48, 6}, {49, 6}, {50, 0}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
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

func makeNoMateNoPat() [][]int {
	return [][]int{
		{0, 12},
		{13, 7},
		{58, 1}, {60, 5},
	}
}
