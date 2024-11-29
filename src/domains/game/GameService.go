package game

import (
	"errors"
	"fmt"
	"github.com/IvaCheMih/chess/src/domains/game/dto"
	"github.com/IvaCheMih/chess/src/domains/game/models"
	moveservice "github.com/IvaCheMih/chess/src/domains/game/services/move"
	"gorm.io/gorm"
	"log"
)

type GamesService struct {
	boardRepo *BoardCellsRepository
	gamesRepo *GamesRepository
	movesRepo *MovesRepository
}

func CreateGamesService(boardRepo *BoardCellsRepository, gamesRepo *GamesRepository, movesRepo *MovesRepository) GamesService {
	return GamesService{
		boardRepo: boardRepo,
		gamesRepo: gamesRepo,
		movesRepo: movesRepo,
	}
}

func (g *GamesService) CreateGame(userId int, userRequestedColor bool) (dto.CreateGameResponse, error) {
	var createGameResponse dto.CreateGameResponse
	createNewBoard := false

	userColor := "white_user_id"
	gameSide := false

	if !userRequestedColor {
		userColor = "black_user_id"
	}

	notStartedGame, err := g.gamesRepo.FindNotStartedGame(userColor)
	if err != nil && err.Error() != "record not found" {
		return dto.CreateGameResponse{}, err
	}

	if userColor == "black_user_id" {
		gameSide = true
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
		var game = models.Game{}

		if userRequestedColor {
			game.WhiteUserId = userId
		} else {
			game.BlackUserId = userId
		}

		notStartedGame, err = g.gamesRepo.CreateGame(tx, game)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}
		createNewBoard = true
	} else {
		err = g.gamesRepo.UpdateColorUserIdByColor(notStartedGame.Id, userColor, gameSide, userId, tx)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}

		notStartedGame, err = g.gamesRepo.GetById(notStartedGame.Id)
		if err != nil {
			return dto.CreateGameResponse{}, err
		}
	}

	FromModelsToDtoCreateGame(notStartedGame, &createGameResponse)

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
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	if userId != game.WhiteUserId && userId != game.BlackUserId {
		return dto.GetBoardResponse{}, err
	}

	board, err := g.boardRepo.Find(gameId)
	if err != nil {
		return dto.GetBoardResponse{}, err
	}

	responseBoard := make([]dto.BoardCellEntity, 64)

	for index, cell := range board.Cells {
		responseBoard[index] = dto.BoardCellEntity{
			IndexCell: cell.IndexCell,
			FigureId:  cell.FigureId,
		}
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

	if !CheckCorrectRequest(requestFromTo) {
		return models.Move{}, errors.New("Move is not correct")
	}

	gameModel, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Move{}, err
	}

	if err = CheckCorrectRequestSideUser(userId, gameModel); err != nil {
		log.Println(err)
		return models.Move{}, err
	}

	from := CoordinatesToIndex(requestFromTo.From)
	to := CoordinatesToIndex(requestFromTo.To)

	indexesToChange, game := moveservice.IsMoveCorrect(gameModel, board, from, to)

	if len(indexesToChange) == 0 {
		return models.Move{}, errors.New("Move is not possible (IsMoveCorrect)")
	}

	if !moveservice.IsItCheck(indexesToChange, &game, requestFromTo.NewFigure) {
		return models.Move{}, errors.New("Move is not possible (IsItCheck)")
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
		log.Println(err)
		return models.Move{}, err
	}

	killedFigureId := 0
	if board.Cells[to] != nil {
		killedFigureId = board.Cells[to].FigureId
	}

	var move = models.Move{
		GameId:         gameId,
		MoveNumber:     maxNumber,
		From_id:        from,
		To_id:          to,
		FigureId:       board.Cells[from].FigureId,
		KilledFigureId: killedFigureId,
		NewFigureId:    0,
		IsCheckWhite:   game.IsCheckWhite.IsItCheck,
		IsCheckBlack:   game.IsCheckBlack.IsItCheck,
	}

	responseMove, err := g.movesRepo.AddMove(move, tx)
	if err != nil {
		return models.Move{}, err
	}

	game.Side = !game.Side

	err = g.gamesRepo.UpdateGame(gameId, game, tx)
	if err != nil {
		return models.Move{}, err
	}

	err = UpdateBoardAfterMove(g, board, game.NewFigureId, indexesToChange, tx)
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
		responseBoard[index] = dto.BoardCellEntity{
			IndexCell: cell.IndexCell,
			FigureId:  cell.FigureId,
		}

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
			fmt.Print(string(moveservice.FigureRepo[getBoardResponse.BoardCells[i].FigureId]))
		}
	}

	return responseMove, err
}

func (g *GamesService) GiveUp(gameId int) (models.Game, error) {
	err := g.gamesRepo.UpdateIsEnded(gameId)
	if err != nil {
		return models.Game{}, err
	}

	game, err := g.gamesRepo.GetById(gameId)
	if err != nil {
		return models.Game{}, err
	}

	return game, err

}

func CheckCorrectRequestSideUser(userId any, game models.Game) error {
	if userId != game.WhiteUserId && userId != game.BlackUserId {
		return errors.New("This is not your game")
	}

	if !game.IsStarted || game.IsEnded {
		return errors.New("Game is not active")
	}

	if game.Side && userId != game.WhiteUserId {
		return errors.New("Its not your move now")
	}

	if !game.Side && userId != game.BlackUserId {
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

func CheckCorrectRequest(move dto.DoMoveBody) bool {
	from, to := CoordinatesToIndex(move.From), CoordinatesToIndex(move.To)

	if !CheckCellOnBoardByIndex(from) || !CheckCellOnBoardByIndex(to) {
		return false
	}

	return true
}

func UpdateBoardAfterMove(g *GamesService, board models.Board, newFigureId int, indexesToChange []int, tx *gorm.DB) error {
	var err error
	from := indexesToChange[0]
	to := indexesToChange[1]

	if board.Cells[to] != nil {
		err = g.boardRepo.Delete(board.Cells[to].Id, tx)
		if err != nil {
			return err
		}
	}

	if newFigureId != 0 {
		err = g.boardRepo.UpdateNewFigure(board.Cells[from].Id, to, newFigureId, tx)
	} else {
		err = g.boardRepo.Update(board.Cells[from].Id, to, tx)
	}

	if err != nil {
		return err
	}

	if len(indexesToChange) > 2 {
		if indexesToChange[2] == -1 {
			err = g.boardRepo.Delete(board.Cells[indexesToChange[3]].Id, tx)
		} else {
			err = g.boardRepo.Update(board.Cells[indexesToChange[2]].Id, indexesToChange[3], tx)
		}
	}

	return err
}

func FromModelsToDtoCreateGame(response models.Game, createGameResponse *dto.CreateGameResponse) {
	createGameResponse.GameId = response.Id

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
	createGameResponse.Side = response.Side
}

var startField = [][]int{
	{0, 8}, {1, 9}, {2, 10}, {3, 11}, {4, 12}, {5, 10}, {6, 9}, {7, 14},
	{8, 13}, {9, 13}, {10, 13}, {11, 13}, {12, 13}, {13, 13}, {14, 13}, {15, 13},
	{48, 6}, {49, 6}, {50, 6}, {51, 6}, {52, 6}, {53, 6}, {54, 6}, {55, 6},
	{56, 1}, {57, 2}, {58, 3}, {59, 4}, {60, 5}, {61, 3}, {62, 2}, {63, 7},
}
