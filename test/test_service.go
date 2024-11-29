package test

import (
	"bytes"
	"encoding/json"
	gameDto "github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/services"
	userDto "github.com/IvaCheMih/chess/src/domains/user/dto"
	"io"
	"net/http"
	"strconv"
)

func CreateUser(user1password userDto.CreateUserRequest) (error, userDto.CreateUserResponse) {
	body, err := json.Marshal(user1password)
	if err != nil {
		return err, userDto.CreateUserResponse{}
	}

	url := services.APP_URL + "user/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err, userDto.CreateUserResponse{}
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err, userDto.CreateUserResponse{}
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err, userDto.CreateUserResponse{}
	}

	var user1response = userDto.CreateUserResponse{}

	err = json.Unmarshal(resBody, &user1response)

	return err, user1response
}

func CreateSession(session userDto.CreateSessionRequest) (error, userDto.CreateSessionResponse) {
	body, err := json.Marshal(session)
	if err != nil {
		return err, userDto.CreateSessionResponse{}
	}

	url := services.APP_URL + "session/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err, userDto.CreateSessionResponse{}
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err, userDto.CreateSessionResponse{}
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err, userDto.CreateSessionResponse{}
	}

	var session1response = userDto.CreateSessionResponse{}

	err = json.Unmarshal(resBody, &session1response)

	return err, session1response
}

func CreateGame(game gameDto.CreateGameBody, token string) (error, gameDto.CreateGameResponse) {
	body, err := json.Marshal(game)
	if err != nil {
		return err, gameDto.CreateGameResponse{}
	}

	url := services.APP_URL + "game/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return err, gameDto.CreateGameResponse{}
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	request.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err, gameDto.CreateGameResponse{}
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err, gameDto.CreateGameResponse{}
	}

	var gameResponse = gameDto.CreateGameResponse{}

	err = json.Unmarshal(resBody, &gameResponse)
	if err != nil {
		return err, gameDto.CreateGameResponse{}
	}

	return err, gameResponse
}

func CreateMove(move gameDto.DoMoveBody, token string, gameId int) (error, gameDto.DoMoveResponse) {
	body, err := json.Marshal(move)
	if err != nil {
		return err, gameDto.DoMoveResponse{}
	}

	url := services.APP_URL + "game/" + strconv.Itoa(gameId) + "/move/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))

	if err != nil {
		return err, gameDto.DoMoveResponse{}
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err, gameDto.DoMoveResponse{}
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err, gameDto.DoMoveResponse{}
	}

	var doMove = gameDto.DoMoveResponse{}

	err = json.Unmarshal(resBody, &doMove)
	if err != nil {
		return err, gameDto.DoMoveResponse{}
	}

	err = json.Unmarshal(resBody, &doMove)
	if err != nil {
		return err, gameDto.DoMoveResponse{}
	}

	return err, doMove
}

func GetBoard(token string, gameId int) (error, gameDto.GetBoardResponse) {

	url := services.APP_URL + "game/" + strconv.Itoa(gameId) + "/board/"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return err, gameDto.GetBoardResponse{}
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return err, gameDto.GetBoardResponse{}
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err, gameDto.GetBoardResponse{}
	}

	var board = gameDto.GetBoardResponse{}

	err = json.Unmarshal(resBody, &board)

	return err, board
}

//func compareExpectedAndActual(board gameDto.GetBoardResponse, expectedFile string) bool {
//	file, err := os.Open(expectedFile)
//	if err != nil {
//		fmt.Println(err)
//		os.Exit(1)
//	}
//	defer file.Close()
//
//	expected := make([]byte, 2048)
//
//	for {
//		_, er := file.Read(expected)
//		if er == io.EOF {
//			break
//		}
//	}
//
//	actual, _ := json.Marshal(board)
//
//	for i, b := range actual {
//		if b != expected[i] {
//			return false
//		}
//	}
//
//	return true
//}
