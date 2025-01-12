package test

import (
	"bytes"
	"encoding/json"
	gameDto "github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	userDto "github.com/IvaCheMih/chess/src/domains/user/dto"
	"io"
	"net/http"
	"strconv"
)

func CreateUser(user1password userDto.CreateUserRequest, appURL string) (userDto.CreateUserResponse, error) {
	body, err := json.Marshal(user1password)
	if err != nil {
		return userDto.CreateUserResponse{}, err
	}

	url := appURL + "user/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return userDto.CreateUserResponse{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return userDto.CreateUserResponse{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return userDto.CreateUserResponse{}, err
	}

	var user1response = userDto.CreateUserResponse{}

	err = json.Unmarshal(resBody, &user1response)

	return user1response, err
}

func TelegramAuth(req userDto.TelegramSignInRequest, appURL string) (userDto.TelegramSignInResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return userDto.TelegramSignInResponse{}, err
	}

	url := appURL + "user/sign-in/telegram/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return userDto.TelegramSignInResponse{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return userDto.TelegramSignInResponse{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return userDto.TelegramSignInResponse{}, err
	}

	var user1response = userDto.TelegramSignInResponse{}

	err = json.Unmarshal(resBody, &user1response)
	if err != nil {
		return userDto.TelegramSignInResponse{}, err
	}

	return user1response, nil
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

func CreateGame(game gameDto.CreateGameBody, token string, appURL string) (gameDto.CreateGameResponse, error) {
	body, err := json.Marshal(game)
	if err != nil {
		return gameDto.CreateGameResponse{}, err
	}

	url := appURL + "game/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return gameDto.CreateGameResponse{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	request.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return gameDto.CreateGameResponse{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return gameDto.CreateGameResponse{}, err
	}

	var gameResponse = gameDto.CreateGameResponse{}

	err = json.Unmarshal(resBody, &gameResponse)
	if err != nil {
		return gameDto.CreateGameResponse{}, err
	}

	return gameResponse, nil
}

func GetGame(gameId int, token string, appURL string) (gameDto.GetGameResponse, error) {
	url := appURL + "game/" + strconv.Itoa(gameId)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return gameDto.GetGameResponse{}, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	request.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return gameDto.GetGameResponse{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return gameDto.GetGameResponse{}, err
	}

	var gameResponse = gameDto.GetGameResponse{}

	err = json.Unmarshal(resBody, &gameResponse)
	if err != nil {
		return gameDto.GetGameResponse{}, err
	}

	return gameResponse, nil
}

func CreateMove(move gameDto.DoMoveBody, token string, gameId int, appURL string) (models.Move, error) {
	body, err := json.Marshal(move)
	if err != nil {
		return models.Move{}, err
	}

	url := appURL + "game/" + strconv.Itoa(gameId) + "/move/"

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))

	if err != nil {
		return models.Move{}, err
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return models.Move{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Move{}, err
	}

	var doMove = models.Move{}

	err = json.Unmarshal(resBody, &doMove)
	if err != nil {
		return models.Move{}, err
	}

	err = json.Unmarshal(resBody, &doMove)
	if err != nil {
		return models.Move{}, err
	}

	return doMove, nil
}

func GetBoard(gameId int, token string, appURL string) (gameDto.GetBoardResponse, error) {
	url := appURL + "game/" + strconv.Itoa(gameId) + "/board/"

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return gameDto.GetBoardResponse{}, err
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return gameDto.GetBoardResponse{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return gameDto.GetBoardResponse{}, err
	}

	var board = gameDto.GetBoardResponse{}

	err = json.Unmarshal(resBody, &board)
	if err != nil {
		return gameDto.GetBoardResponse{}, err
	}

	return board, nil
}

func EndGame(endgame gameDto.EndGameRequest, token string, appURL string) (models.Game, error) {
	url := appURL + "game/endgame/"

	body, err := json.Marshal(endgame)
	if err != nil {
		return models.Game{}, err
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return models.Game{}, err
	}

	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("accept", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return models.Game{}, err
	}

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return models.Game{}, err
	}

	var board = models.Game{}

	err = json.Unmarshal(resBody, &board)
	if err != nil {
		return models.Game{}, err
	}

	return board, nil
}
