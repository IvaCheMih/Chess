package test

import (
	gamedto "github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/services"
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

	{From: "B7", To: "C8", NewFigure: 113},
}

var boardFirst = [][]int{
	{0, 8}, {0, 0}, {2, 4}, {0, 0}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
	{8, 13}, {0, 0}, {10, 13}, {11, 13}, {12, 13}, {13, 13}, {14, 13},
	{15, 13}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
	{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
	{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
	{0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0}, {0, 0},
	{0, 0}, {48, 6}, {49, 6}, {0, 0}, {51, 6}, {52, 6}, {53, 6}, {54, 6},
	{55, 6}, {56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
}

func TestGame(t *testing.T) {
	t.Run("test user, session, game", func(t *testing.T) {
		log.Println(1)
		viper.AutomaticEnv()
		viper.SetConfigFile(".env")
		err := viper.ReadInConfig()
		require.NoError(t, err)

		envs := services.NewEnvService()

		expectedFirst := MakeExpected(boardFirst)

		DoTestChessGame(t, movesFirst, expectedFirst, envs.AppURL)
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

func DoTestChessGame(t *testing.T, moves []gamedto.DoMoveBody, expected gamedto.GetBoardResponse, appURL string) {
	var user1password = userdto.CreateUserRequest{Password: "password"}
	var user2password = userdto.CreateUserRequest{Password: "password"}

	err, user1response := CreateUser(user1password, appURL)
	require.NoError(t, err)

	err, user2response := CreateUser(user2password, appURL)
	require.NoError(t, err)

	var session1 = userdto.CreateSessionRequest{ //nolint:gosimple
		Id:       user1response.Id,
		Password: user1response.Password,
	}

	var session2 = userdto.CreateSessionRequest{ //nolint:gosimple
		Id:       user2response.Id,
		Password: user2response.Password,
	}

	err, session1response := CreateSession(session1, appURL)
	require.NoError(t, err)

	err, session2response := CreateSession(session2, appURL)
	require.NoError(t, err)

	var game1user = gamedto.CreateGameBody{IsWhite: true}
	var game2user = gamedto.CreateGameBody{
		IsWhite: false}

	err, game1response := CreateGame(game1user, session1response.Token, appURL)
	require.NoError(t, err)

	err, game2response := CreateGame(game2user, session2response.Token, appURL)
	require.NoError(t, err)

	if game1response.GameId != game2response.GameId {
		log.Println("Game ID is not correct")
		t.Errorf("[!] Game Ids is not equal")
	}

	for i, move := range moves {
		if i%2 == 0 {
			err, _ = CreateMove(move, session1response.Token, game1response.GameId, appURL)
			require.NoError(t, err)
		} else {
			err, _ = CreateMove(move, session2response.Token, game2response.GameId, appURL)
			require.NoError(t, err)
		}
	}

	err, board := GetBoard(session1response.Token, game1response.GameId, appURL)
	require.NoError(t, err)

	if !assert.Equal(t, board, expected) {
		t.Errorf("[!] Boards are not equal")
	}
}
