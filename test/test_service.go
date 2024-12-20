package test

import (
	"bytes"
	"encoding/json"
	gameDto "github.com/IvaCheMih/chess/src/domains/game/dto"
	userDto "github.com/IvaCheMih/chess/src/domains/user/dto"
	"io"
	"net/http"
	"strconv"
)

func CreateUser(user1password userDto.CreateUserRequest, appURL string) (error, userDto.CreateUserResponse) {
	body, err := json.Marshal(user1password)
	if err != nil {
		return err, userDto.CreateUserResponse{}
	}

	url := appURL + "user/"

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

func CreateSession(session userDto.CreateSessionRequest, appURL string) (error, userDto.CreateSessionResponse) {
	body, err := json.Marshal(session)
	if err != nil {
		return err, userDto.CreateSessionResponse{}
	}

	url := appURL + "session/"

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

func CreateGame(game gameDto.CreateGameBody, token string, appURL string) (error, gameDto.CreateGameResponse) {
	body, err := json.Marshal(game)
	if err != nil {
		return err, gameDto.CreateGameResponse{}
	}

	url := appURL + "game/"

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

func CreateMove(move gameDto.DoMoveBody, token string, gameId int, appURL string) (error, gameDto.DoMoveResponse) {
	body, err := json.Marshal(move)
	if err != nil {
		return err, gameDto.DoMoveResponse{}
	}

	url := appURL + "game/" + strconv.Itoa(gameId) + "/move/"

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

func GetBoard(token string, gameId int, appURL string) (error, gameDto.GetBoardResponse) {

	url := appURL + "game/" + strconv.Itoa(gameId) + "/board/"

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
