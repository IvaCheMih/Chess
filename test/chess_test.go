package test

import (
	"fmt"
	gameDto "github.com/IvaCheMih/chess/src/domains/game/dto"
	userDto "github.com/IvaCheMih/chess/src/domains/user/dto"
	"github.com/stretchr/testify/assert"
	"testing"
)

var movesFirst = []gameDto.DoMoveBody{
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

	expectedFirst := MakeExpected(boardFirst)

	DoTestChessGame(t, movesFirst, expectedFirst)
}

func MakeExpected(boardMas [][]int) gameDto.GetBoardResponse {
	var board = gameDto.GetBoardResponse{
		BoardCells: []gameDto.BoardCellEntity{},
	}

	for _, b := range boardMas {
		board.BoardCells = append(board.BoardCells, gameDto.BoardCellEntity{b[0], b[1]})
	}

	return board
}

func DoTestChessGame(t *testing.T, moves []gameDto.DoMoveBody, expected gameDto.GetBoardResponse) {
	var user1password = userDto.CreateUserRequest{Password: "password"}
	var user2password = userDto.CreateUserRequest{Password: "password"}

	err, user1response := CreateUser(user1password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("[!] Error creating user: %s", user1password)
	}

	err, user2response := CreateUser(user2password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("[!] Error creating user: %s", user1password)
	}

	var session1 = userDto.CreateSessionRequest{
		Id:       user1response.Id,
		Password: user1response.Password,
	}

	var session2 = userDto.CreateSessionRequest{
		Id:       user2response.Id,
		Password: user2response.Password,
	}

	err, session1response := CreateSession(session1)
	if err != nil {
		fmt.Println(err)
		t.Errorf("[!] Error creating session: %+v", session1response)
	}

	err, session2response := CreateSession(session2)
	if err != nil {
		fmt.Println(err)
		t.Errorf("[!] Error creating session: %+v", session2response)
	}

	var game1user = gameDto.CreateGameBody{true}
	var game2user = gameDto.CreateGameBody{false}

	err, game1response := CreateGame(game1user, session1response.Token)
	if err != nil {
		fmt.Println(err)
		t.Errorf("[!] Error creating game: %+v", game1response)
	}

	err, game2response := CreateGame(game2user, session2response.Token)
	if err != nil {
		fmt.Println(err)
		t.Errorf("[!] Error creating game: %+v", game2response)
	}

	if game1response.GameId != game2response.GameId {
		fmt.Println("Game ID is not correct")
		t.Errorf("[!] Game Ids is not equal")
	}

	for i, move := range moves {
		if i%2 == 0 {
			err, _ = CreateMove(move, session1response.Token, game1response.GameId)
			if err != nil {
				fmt.Println(err)
				t.Errorf("[!] Error creating move: %+v", move)
			}
		} else {
			err, _ = CreateMove(move, session2response.Token, game2response.GameId)
			if err != nil {
				fmt.Println(err)
				t.Errorf("[!] Error creating move: %+v", move)
			}
		}
	}

	err, board := GetBoard(session1response.Token, game1response.GameId)
	if err != nil {
		fmt.Println(err)
		t.Errorf("[!] Error getting board: %+v", board)
	}

	if !assert.Equal(t, board, expected) {
		t.Errorf("[!] Boards are not equal")
	}
}
