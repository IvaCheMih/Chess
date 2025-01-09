package test

import (
	gamedto "github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/services/env"
	"github.com/IvaCheMih/chess/src/domains/services/test"
	userdto "github.com/IvaCheMih/chess/src/domains/user/dto"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

var movesFirst = []gamedto.DoMoveBody{
	{From: "C2", To: "C4", NewFigure: 0},
	{From: "B7", To: "B5", NewFigure: 0},

	{From: "C4", To: "B5", NewFigure: 0},
	{From: "B8", To: "A6", NewFigure: 0},

	{From: "B5", To: "A6", NewFigure: 0},
	{From: "C8", To: "B7", NewFigure: 0},

	{From: "A6", To: "B7", NewFigure: 0},
	{From: "D8", To: "C8", NewFigure: 0},

	{From: "B7", To: "A8", NewFigure: 113},
	{From: "C8", To: "B8", NewFigure: 0},

	{From: "A8", To: "B8", NewFigure: 0},
}

var board1 = [][]int{
	{0, 0}, {1, 4}, {2, 0}, {3, 0}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
	{8, 13}, {9, 0}, {10, 13}, {11, 13}, {12, 13}, {13, 13}, {14, 13}, {15, 13},
	{16, 0}, {17, 0}, {18, 0}, {19, 0}, {20, 0}, {21, 0}, {22, 0}, {23, 0},
	{24, 0}, {25, 0}, {26, 0}, {27, 0}, {28, 0}, {29, 0}, {30, 0}, {31, 0},
	{32, 0}, {33, 0}, {34, 0}, {35, 0}, {36, 0}, {37, 0}, {38, 0}, {39, 0},
	{40, 0}, {41, 0}, {42, 0}, {43, 0}, {44, 0}, {45, 0}, {46, 0}, {47, 0},
	{48, 6}, {49, 6}, {50, 0}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
}

var game1 = gamedto.GetGameResponse{
	IsEnded:   false,
	EndReason: "Mat",
}

func TestGame(t *testing.T) {
	t.Run("test user, session, game", func(t *testing.T) {
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		require.NoError(t, err)

		envs := env.NewEnvService()

		expected1 := MakeExpected(board1)

		DoTestChessGame(t, movesFirst, expected1, game1, envs.AppURL)
	})
}

func MakeExpected(boardMas [][]int) gamedto.GetBoardResponse {
	var board = gamedto.GetBoardResponse{
		BoardCells: []gamedto.BoardCellEntity{},
	}

	for _, b := range boardMas {
		board.BoardCells = append(board.BoardCells, gamedto.BoardCellEntity{
			IndexCell: b[0],
			FigureId:  b[1],
		})
	}

	return board
}

func DoTestChessGame(t *testing.T, moves []gamedto.DoMoveBody, expected gamedto.GetBoardResponse, expectedGame gamedto.GetGameResponse, appURL string) {
	var user1password = userdto.CreateUserRequest{TelegramId: 1, Password: "password"}
	var user2password = userdto.CreateUserRequest{TelegramId: 2, Password: "password"}

	user1response, err := test.CreateUser(user1password, appURL)
	require.NoError(t, err)

	user2response, err := test.CreateUser(user2password, appURL)
	require.NoError(t, err)

	var session1 = userdto.CreateSessionRequest{ //nolint:gosimple
		Id:       user1response.Id,
		Password: user1response.Password,
	}

	var session2 = userdto.CreateSessionRequest{ //nolint:gosimple
		Id:       user2response.Id,
		Password: user2response.Password,
	}

	err, session1response := test.CreateSession(session1, appURL)
	require.NoError(t, err)

	err, session2response := test.CreateSession(session2, appURL)
	require.NoError(t, err)

	var game1user = gamedto.CreateGameBody{IsWhite: true}
	var game2user = gamedto.CreateGameBody{
		IsWhite: false}

	game1response, err := test.CreateGame(game1user, session1response.Token, appURL)
	require.NoError(t, err)

	game2response, err := test.CreateGame(game2user, session2response.Token, appURL)
	require.NoError(t, err)

	if game1response.GameId != game2response.GameId {
		log.Println("Game ID is not correct")
		t.Errorf("[!] Game Ids is not equal")
	}

	for i, move := range moves {
		if i%2 == 0 {
			_, err = test.CreateMove(move, session1response.Token, game1response.GameId, appURL)
			require.NoError(t, err)
		} else {
			_, err = test.CreateMove(move, session2response.Token, game2response.GameId, appURL)
			require.NoError(t, err)
		}
	}

	board, err := test.GetBoard(session1response.Token, game1response.GameId, appURL)
	require.NoError(t, err)

	if !assert.Equal(t, expected, board) {
		t.Errorf("[!] Boards are not equal")
	}

	game, err := test.GetGame(
		game1response.GameId,
		session1response.Token,
		appURL,
	)

	if !assert.Equal(t, expectedGame.IsEnded, game.IsEnded) {
		t.Errorf("[!] IsEnded are not equal")
	}

	if !assert.Equal(t, expectedGame.EndReason, game.EndReason) {
		t.Errorf("[!] IsEnded are not equal")
	}
}
