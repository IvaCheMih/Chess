package game

import (
	"errors"
	"fmt"
	"github.com/IvaCheMih/chess/server/domains/game/dto"
	"github.com/IvaCheMih/chess/server/domains/game/models"
	"github.com/IvaCheMih/chess/server/domains/game/services/move_service"
	"gorm.io/gorm"
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

func (g *GamesService) CreateGame(userId int, userRequestedColor bool) (dto.CreateGameResponse, error) {
	var createGameResponse dto.CreateGameResponse
	createNewBoard := false

	userColor := "white_user_id"
	gameSide := userId

	if !userRequestedColor {
		userColor = "black_user_id"
	}

	response, err := g.gamesRepo.FindNotStartedGame(userColor)
	if err != nil && err.Error() != "record not found" {
		return dto.CreateGameResponse{}, err
	}

	if userColor == "black_user_id" {
		gameSide = response.WhiteUserId
	}

	tx := g.gamesRepo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return dto.CreateGameResponse{}, err
	}

	if err != nil && err.Error() == "record not found" {
		response, err = g.gamesRepo.CreateGame(userId, userRequestedColor, tx)
		createNewBoard = true
		if err != nil {
			return dto.CreateGameResponse{}, err
		}
	} else {

		response, err = g.gamesRepo.UpdateColorUserIdByColor(response.Id, userColor, gameSide, userId, tx)
	}

	FromModelsToDtoCreateGame(response, &createGameResponse)

	if err != nil {
		return dto.CreateGameResponse{}, err
	}

	if createNewBoard {
		err = g.boardRepo.CreateNewBoardCells(createGameResponse.GameId, tx)
	}
	if err != nil {
		return dto.CreateGameResponse{}, err
	}

	tx.Commit()

	return createGameResponse, err
}

func (g *GamesService) GetBoard(gameId int, userId any) (dto.GetBoardResponse, error) {

	game, err := g.gamesRepo.GetById(gameId)

	if userId != game.WhiteUserId && userId != game.BlackUserId {
		return dto.GetBoardResponse{}, err
	}

	board, err := g.boardRepo.Find(gameId)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

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
	responseGetGame, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return dto.GetHistoryResponse{}, err
	}

	if userId != responseGetGame.WhiteUserId && userId != responseGetGame.BlackUserId {
		return dto.GetHistoryResponse{}, errors.New("This is not your game")
	}

	moves, err := g.movesRepo.Find(gameId)
	if err != nil {
		return dto.GetHistoryResponse{}, err
	}

	var responseGetHistory = dto.GetHistoryResponse{
		Moves: moves,
	}

	return responseGetHistory, err
}

func (g *GamesService) Move(gameId int, userId any, requestFromTo dto.DoMoveBody) (models.Move, error) {
	board, err := g.boardRepo.Find(gameId)
	if err != nil {
		return models.Move{}, err
	}

	if !CheckCorrectRequest(requestFromTo.From, requestFromTo.To) {
		return models.Move{}, errors.New("Move is not correct")
	}

	response, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Move{}, err
	}

	var responseGetGame dto.CreateGameResponse

	FromModelsToDtoCreateGame(response, &responseGetGame)

	if err = CheckCorrectRequestSideUser(userId, responseGetGame); err != nil {
		fmt.Println(err)
		return models.Move{}, err
	}

	from := CoordinatesToIndex(requestFromTo.From)
	to := CoordinatesToIndex(requestFromTo.To)

	isCorrect, indexesToChange := move_service.CheckCorrectMove(responseGetGame, board, from, to)

	if !isCorrect {
		return models.Move{}, errors.New("Move is not possible (CheckCorrectMove)")
	}

	game, check := move_service.CheckIsItCheck(responseGetGame, board, indexesToChange)

	if !check {
		return models.Move{}, errors.New("Move is not possible (CheckIsItCheck)")
	}

	tx := g.gamesRepo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if tx.Error != nil {
		return models.Move{}, err
	}

	maxNumber, err := g.movesRepo.FindMaxMoveNumber(gameId)

	if err != nil {
		fmt.Println(err)
		return models.Move{}, err
	}

	responseMove, err := g.movesRepo.AddMove(gameId, from, to, board, game.IsCheckWhite, game.IsCheckBlack, maxNumber+1, tx)
	if err != nil {
		return models.Move{}, err
	}

	if game.Side == *game.WhiteClientId {
		game.Side = *game.BlackClientId
	} else {
		game.Side = *game.WhiteClientId
	}

	err = g.gamesRepo.UpdateGame(gameId, game, tx)
	if err != nil {
		return models.Move{}, err
	}

	err = UpdateBoardAfterMove(g, board, from, to, game, isCastling, tx)
	if err != nil {
		return models.Move{}, err
	}

	tx.Commit()

	responseBoard := make([]dto.BoardCellEntity, 64)

	cells, err := g.boardRepo.Find(gameId)
	if err != nil {
		return models.Move{}, err
	}

	for index, cell := range cells.Cells {
		responseBoard[index] = dto.BoardCellEntity{cell.IndexCell, cell.FigureId}

	}

	getBoardResponse := dto.GetBoardResponse{
		BoardCells: responseBoard,
	}

	for i := 0; i < 64; i++ {
		if i%8 == 0 {
			fmt.Print("\n")
		}

		if getBoardResponse.BoardCells[i].FigureId == 0 {
			fmt.Print(0)
		} else {
			fmt.Print(string(move_service.FigureRepo[getBoardResponse.BoardCells[i].FigureId]))
		}
	}

	return responseMove, err
}

func (g *GamesService) GiveUp(gameId int, userId any) (models.Game, error) {

	game, err := g.gamesRepo.UpdateIsEnded(gameId)
	if err != nil {
		return models.Game{}, err
	}

	return game, err

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

func CheckCorrectNewFigure(figureId int) bool {
	figureType, ok := move_service.FigureRepo[figureId]
	if !ok || figureType == 'p' || figureType == 'K' {
		return false
	}
	return true
}

func UpdateBoardAfterMove(g *GamesService, board models.Board, from int, to int, game move_service.Game, isCastling bool, tx *gorm.DB) error {
	if board.Cells[to] != nil {
		err := g.boardRepo.Delete(board.Cells[to].Id, tx)

		if err != nil {
			return err
		}
	}

	err := g.boardRepo.Update(board.Cells[from].Id, to, tx)
	if err != nil {
		return err
	}

	if isCastling {
		err = g.boardRepo.Update(board.Cells[game.RookOldIdIfItCastling].Id, game.RookNewIdIfItCastling, tx)
	}

	return err
}

func FromModelsToDtoCreateGame(response models.Game, createGameResponse *dto.CreateGameResponse) {

	createGameResponse.GameId = response.Id
	createGameResponse.Side = response.Side

	createGameResponse.IsCheckWhite = response.IsCheckWhite
	createGameResponse.IsCheckBlack = response.IsCheckBlack

	createGameResponse.IsStarted = response.IsStarted
	createGameResponse.IsEnded = response.IsEnded

	createGameResponse.WhiteKingCastling = response.WhiteKingCastling
	createGameResponse.BlackKingCastling = response.BlackKingCastling

	createGameResponse.WhiteRookACastling = response.WhiteRookACastling
	createGameResponse.WhiteRookHCastling = response.WhiteRookHCastling

	createGameResponse.BlackRookACastling = response.BlackRookACastling
	createGameResponse.BlackRookHCastling = response.BlackRookHCastling

	createGameResponse.BlackUserId = response.BlackUserId
	createGameResponse.WhiteUserId = response.WhiteUserId

	createGameResponse.LastPawnMove = response.LastPawnMove

}

var startField = [][]int{
	{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
	{8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13}, {13, 13}, {14, 13}, {15, 13},
	{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
}
