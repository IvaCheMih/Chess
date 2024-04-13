package game

import (
	"errors"
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"github.com/IvaCheMih/chess/server/domains/game/move_service"
)

type GamesService struct {
	boardRepo  *BoardCellsRepository
	figureRepo *FiguresRepository
	gamesRepo  *GamesRepository
	movesRepo  *MovesRepository
}

func CreateGamesService(boardRepo *BoardCellsRepository, figureRepo *FiguresRepository, gamesRepo *GamesRepository, movesRepo *MovesRepository) GamesService {
	return GamesService{
		boardRepo:  boardRepo,
		figureRepo: figureRepo,
		gamesRepo:  gamesRepo,
		movesRepo:  movesRepo,
	}
}

func (g *GamesService) CreateGame(userId any, userRequestedColor bool) (dto.CreateGameResponse, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.CreateGameResponse{}, err
	}

	defer tx.Rollback()

	var createGameResponse dto.CreateGameResponse

	if userRequestedColor {
		response, err := g.gamesRepo.CreateGame(userId, tx)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}

		FromModelsToDtoCreateGame(response, &createGameResponse)
	} else {
		response, err := g.gamesRepo.FindNotStartedGame(tx)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}

		FromModelsToDtoCreateGame(response, &createGameResponse)
		err = g.gamesRepo.JoinBlackToGame(createGameResponse.GameId, userId, tx)
	}

	if err != nil {
		return dto.CreateGameResponse{}, err
	}

	if userRequestedColor {
		err = g.boardRepo.CreateNewBoardCells(createGameResponse.GameId, tx)
	}
	if err != nil {
		return dto.CreateGameResponse{}, err
	}

	err = tx.Commit()

	return createGameResponse, err
}

func (g *GamesService) GetBoard(gameId int, userId any) (dto.GetBoardResponse, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	defer tx.Rollback()

	game, err := g.gamesRepo.GetById(gameId, tx)

	if userId != game.WhiteUserId && userId != game.BlackUserId {
		return dto.GetBoardResponse{}, err
	}

	board, err := g.boardRepo.Find(gameId, tx)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	err = tx.Commit()

	responseBoard := make([]dto.BoardCellEntity, 64)

	for index, cell := range board.Cells {
		responseBoard[index] = dto.BoardCellEntity{cell.IndexCell, cell.FigureId}

	}

	getBoardResponse := dto.GetBoardResponse{
		BoardCells: responseBoard,
	}

	return getBoardResponse, err
}

func (g *GamesService) GetHistory(gameId int, userId any) (dto.GetHistoryResponse, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.GetHistoryResponse{}, err
	}

	defer tx.Rollback()

	responseGetGame, err := g.gamesRepo.GetById(gameId, tx)
	if err != nil {
		return dto.GetHistoryResponse{}, err
	}

	if userId != responseGetGame.WhiteUserId && userId != responseGetGame.BlackUserId {
		return dto.GetHistoryResponse{}, errors.New("This is not your game")
	}

	moves, err := g.movesRepo.Find(gameId, tx)
	if err != nil {
		return dto.GetHistoryResponse{}, err
	}

	err = tx.Commit()

	var responseGetHistory = dto.GetHistoryResponse{
		Moves: moves,
	}

	return responseGetHistory, err
}

func (g *GamesService) Move(gameId int, userId any, requestFromTo dto.DoMoveBody) (models.Move, error) {
	if !CheckCorrectRequest(requestFromTo.From, requestFromTo.To) {
		return models.Move{}, errors.New("Move is not correct")
	}

	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return models.Move{}, err
	}

	defer tx.Rollback()

	fmt.Println(200)

	response, err := g.gamesRepo.GetById(gameId, tx)
	if err != nil {
		return models.Move{}, err
	}

	var responseGetGame dto.CreateGameResponse

	fmt.Println(201)

	FromModelsToDtoCreateGame(response, &responseGetGame)

	if err = CheckCorrectRequestSideUser(userId, responseGetGame); err != nil {
		fmt.Println(err)
		return models.Move{}, err
	}

	fmt.Println(202)

	board, err := g.boardRepo.Find(gameId, tx)
	if err != nil {
		return models.Move{}, err
	}

	if !move_service.CheckCorrectMove(responseGetGame, board, requestFromTo) {
		return models.Move{}, errors.New("Move is not possible (CheckCorrectMove)")
	}

	from := CoordinatesToIndex(requestFromTo.From)
	to := CoordinatesToIndex(requestFromTo.To)

	fmt.Println(207)

	game, check := move_service.CheckIsItCheck(responseGetGame, board, from, to)

	fmt.Println(208)

	if !check {
		return models.Move{}, errors.New("Move is not possible (CheckIsItCheck)")
	}

	responseMove, err := g.movesRepo.AddMove(gameId, from, to, board, game.IsCheckWhite, game.IsCheckBlack, tx)
	if err != nil {
		fmt.Println(err)
		return models.Move{}, err
	}

	fmt.Println(209)

	if game.Side == *game.WhiteClientId {
		game.Side = *game.BlackClientId
	} else {
		game.Side = *game.WhiteClientId
	}

	err = g.gamesRepo.UpdateGame(gameId, game.IsCheckWhite, game.IsCheckBlack, game.Side, tx)
	if err != nil {
		fmt.Print(err)
		return models.Move{}, err
	}

	fmt.Println(210)

	if board.Cells[to] != nil {
		err = g.boardRepo.Delete(board.Cells[to].Id, tx)

		if err != nil {
			return models.Move{}, err
		}
	}

	fmt.Println(211)

	err = g.boardRepo.Update(board.Cells[from].Id, to, tx)

	if err != nil {
		return models.Move{}, err
	}

	return responseMove, err
}

func CheckCorrectRequestSideUser(userId any, responseGetGame dto.CreateGameResponse) error {
	if userId != responseGetGame.WhiteUserId && userId != responseGetGame.BlackUserId {
		return errors.New("This is not your game")
	}

	if !responseGetGame.IsStarted || responseGetGame.IsEnded {
		return errors.New("Game is not active")
	}

	if responseGetGame.WhiteUserId == responseGetGame.Side && userId != responseGetGame.WhiteUserId {
		return errors.New("Its not your move now")
	}

	if responseGetGame.BlackUserId == responseGetGame.Side && userId != responseGetGame.BlackUserId {
		return errors.New("Its not your move now")
	}
	return nil
}

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

func FromModelsToDtoCreateGame(response models.Game, createGameResponse *dto.CreateGameResponse) {
	createGameResponse.GameId = response.GameId
	createGameResponse.Side = response.Side

	createGameResponse.IsCheckWhite = response.IsCheckWhite
	createGameResponse.IsCheckBlack = response.IsCheckBlack

	createGameResponse.IsStarted = response.IsStarted
	createGameResponse.IsEnded = response.IsEnded

	createGameResponse.WhiteKingCell = response.WhiteKingCell
	createGameResponse.BlackKingCell = response.BlackKingCell

	createGameResponse.BlackUserId = response.BlackUserId
	createGameResponse.WhiteUserId = response.WhiteUserId

}

var startField = [][]int{
	{0, 7}, {1, 8}, {2, 9}, {3, 10}, {4, 11}, {5, 9}, {6, 8}, {7, 7},
	{8, 12}, {9, 12}, {10, 12}, {11, 12}, {12, 12}, {13, 12}, {14, 12}, {15, 12},
	{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 1},
}
