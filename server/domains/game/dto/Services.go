package dto

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetClientId(c *fiber.Ctx) (int, error) {
	headers := c.GetReqHeaders()

	clientIds := headers["X-Client-Id"]
	clientId, err := strconv.ParseInt(clientIds[0], 10, 0)

	return int(clientId), err
}

func GetRequestNewGame(c *fiber.Ctx) (RequestedCreateGame, error) {
	body := c.Body()

	var request RequestedCreateGame

	err := json.Unmarshal(body, &request)

	return request, err
}

func GetGameId(c *fiber.Ctx) (RequestGetBoard, error) {
	gameId, err := c.ParamsInt("gameId")
	if err != nil {
		return RequestGetBoard{}, err
	}

	var request = RequestGetBoard{
		GameId: gameId,
	}

	return request, nil
}

func GetRequestDoMoveFromBody(c *fiber.Ctx) (RequestDoMove, error) {
	body := c.Body()

	var request RequestDoMove

	err := json.Unmarshal(body, &request)

	return request, err
}

//func GetRequestGetBoard(c *fiber.Ctx) (RequestGetBoard, error) {
//	body := c.Body()
//
//	var request RequestGetBoard
//
//	err := json.Unmarshal(body, &request)
//
//	return request, err
//}

func IndexToCoordinates(index int) string {
	y := int('8') - (index / 8)
	x := (index % 8) + int('A')

	return string(byte(x)) + string(byte(y))
}

func CoordinatesToIndex(coordinates string) int {
	x := int(coordinates[0]) - int('A')
	y := int('8') - int(coordinates[1])

	return (y * 8) + x
}

func ParseMessageToMove(message string) (string, string) {
	return message[0:2], message[3:]
}

func CheckCellOnBoardByIndex(index int) bool {
	coordinates := IndexToCoordinates(index)
	if coordinates[0] >= byte('A') && coordinates[0] <= byte('H') {
		if coordinates[1] >= byte('1') && coordinates[1] <= byte('8') {
			return true
		}
	}
	return false
}

func CheckCorrectRequest(f, t string) bool {
	from, to := CoordinatesToIndex(f), CoordinatesToIndex(t)

	if !CheckCellOnBoardByIndex(from) || !CheckCellOnBoardByIndex(to) {
		return false
	}
	return true

}
