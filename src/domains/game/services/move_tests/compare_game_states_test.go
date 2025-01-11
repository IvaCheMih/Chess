package move_tests

import (
	"github.com/IvaCheMih/chess/src/domains/game"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompareGames(t *testing.T) {
	boardRepo := game.CreateBoardCellsRepository(nil)
	gameService := game.CreateGamesService(&boardRepo, nil, nil)

	t.Run("Compare start positions", func(t *testing.T) {
		m := gameService.GetMoveService()

		board1 := boardRepo.NewStartBoardCells(1)

		var cells1 = map[int]*models.BoardCell{}

		for i := range board1 {
			cells1[board1[i].IndexCell] = &board1[i]
		}

		game1 := m.CreateGameStruct(models.Game{
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
		}, models.Board{Cells: cells1})

		board2 := boardRepo.NewStartBoardCells(1)

		var cells2 = map[int]*models.BoardCell{}

		for i := range board2 {
			cells2[board2[i].IndexCell] = &board2[i]
		}

		game2 := m.CreateGameStruct(models.Game{
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
		}, models.Board{Cells: cells2})

		isEqual := game1.CompareGamesStates(game2)
		require.True(t, isEqual)
	})

	t.Run("Compare games false 1", func(t *testing.T) {
		m := gameService.GetMoveService()

		board1 := boardRepo.NewStartBoardCells(1)

		var cells1 = map[int]*models.BoardCell{}

		for i := range board1 {
			cells1[board1[i].IndexCell] = &board1[i]
		}

		game1 := m.CreateGameStruct(models.Game{
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
		}, models.Board{Cells: cells1})

		board2 := boardRepo.MakeBoardCells(1, makeBlackPatBoard1())

		var cells2 = map[int]*models.BoardCell{}

		for i := range board2 {
			cells2[board2[i].IndexCell] = &board2[i]
		}

		game2 := m.CreateGameStruct(models.Game{
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
		}, models.Board{Cells: cells2})

		isEqual := game1.CompareGamesStates(game2)
		require.False(t, isEqual)
	})

	t.Run("Compare games true", func(t *testing.T) {
		m := gameService.GetMoveService()

		board1 := boardRepo.MakeBoardCells(1, makeBlackCheckBoard1())

		var cells1 = map[int]*models.BoardCell{}

		for i := range board1 {
			cells1[board1[i].IndexCell] = &board1[i]
		}

		game1 := m.CreateGameStruct(models.Game{
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
		}, models.Board{Cells: cells1})

		board2 := boardRepo.MakeBoardCells(1, makeBlackCheckBoard1())

		var cells2 = map[int]*models.BoardCell{}

		for i := range board2 {
			cells2[board2[i].IndexCell] = &board2[i]
		}

		game2 := m.CreateGameStruct(models.Game{
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
		}, models.Board{Cells: cells2})

		isEqual := game1.CompareGamesStates(game2)
		require.True(t, isEqual)
	})
}
