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

func TestGame(t *testing.T) {
	t.Run("test user, session, game", func(t *testing.T) {
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		require.NoError(t, err)

		envs := env.NewEnvService()

		expected1 := MakeExpected(board1)
		//expected2 := MakeExpected(board2)

		DoTestChessGame(t, moves1, expected1, game1, envs.AppURL)
		//DoTestChessGame(t, moves2, expected2, game2, envs.AppURL)
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

	board, err := test.GetBoard(game1response.GameId, session1response.Token, appURL)
	require.NoError(t, err)

	if !assert.Equal(t, expected, board) {
		t.Errorf("[!] Boards are not equal")
	}

	game, err := test.GetGame(
		game1response.GameId,
		session1response.Token,
		appURL,
	)
	require.NoError(t, err)

	if !assert.Equal(t, expectedGame.IsEnded, game.IsEnded) {
		t.Errorf("[!] IsEnded are not equal")
	}

	if !assert.Equal(t, expectedGame.EndReason, game.EndReason) {
		t.Errorf("[!] IsEnded are not equal")
	}
}
