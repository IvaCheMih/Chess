package move_tests

import (
	"github.com/IvaCheMih/chess/src/domains/game"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckPawnAttacks(t *testing.T) {
	boardRepo := game.CreateBoardCellsRepository(nil)
	gameService := game.CreateGamesService(&boardRepo, nil, nil)
	m := gameService.GetMoveService()

	t.Run("Test black check 1", func(t *testing.T) {

		board := boardRepo.MakeBoardCells(1, makeBlackCheck1())

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

		checkBlack := game.IsKingCheck(game.IsCheckBlack.KingGameID)
		require.True(t, checkBlack)
	})

	t.Run("Test black check 2", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeBlackCheck2())

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

		checkBlack := game.IsKingCheck(game.IsCheckBlack.KingGameID)
		require.True(t, checkBlack)
	})

	t.Run("Test black check 3", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeBlackCheck3())

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

		checkBlack := game.IsKingCheck(game.IsCheckBlack.KingGameID)
		require.True(t, checkBlack)
	})

	t.Run("Test white check 1", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeWhiteCheck1())

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

		check := game.IsKingCheck(game.IsCheckWhite.KingGameID)
		require.True(t, check)
	})

	t.Run("Test white check 2", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeWhiteCheck2())

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

		check := game.IsKingCheck(game.IsCheckWhite.KingGameID)
		require.True(t, check)
	})

	t.Run("Test white not check 1", func(t *testing.T) {
		m := gameService.GetMoveService()

		board := boardRepo.MakeBoardCells(1, makeWhiteNotCheck1())

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

		check := game.IsKingCheck(game.IsCheckWhite.KingGameID)
		require.False(t, check)
	})
}

func makeBlackCheck1() [][]int {
	return [][]int{
		{9, 12},
		{16, 6},
	}
}

func makeBlackCheck2() [][]int {
	return [][]int{
		{9, 12},
		{18, 6},
	}
}

func makeBlackCheck3() [][]int {
	return [][]int{
		{9, 12},
		{16, 6}, {17, 6}, {18, 6},
	}
}

func makeWhiteCheck1() [][]int {
	return [][]int{
		{50, 13},
		{59, 5},
	}
}

func makeWhiteCheck2() [][]int {
	return [][]int{
		{52, 13},
		{59, 5},
	}
}

func makeWhiteNotCheck1() [][]int {
	return [][]int{
		{51, 13},
		{59, 5},
	}
}
