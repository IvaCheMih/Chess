package game

import "github.com/IvaCheMih/chess/server/domains/game/dto"

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

func (g *GamesService) CreateGame(userId int, userRequestedColor dto.RequestedColor) (dto.ResponseGetGame, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	defer tx.Rollback()

	var requestCreateGame dto.ResponseGetGame

	if userRequestedColor.IsWhite {
		requestCreateGame, err = g.gamesRepo.CreateGame(userId, tx)
	} else {
		requestCreateGame, err = g.gamesRepo.FindNotStartedGame(tx)
		if err != nil {
			return dto.ResponseGetGame{}, err
		}

		err = g.gamesRepo.JoinBlackToGame(requestCreateGame.GameId, userId, tx)
	}

	if err != nil {
		return dto.ResponseGetGame{}, err
	}

	if !g.boardRepo.CreateNewBoardCells(requestCreateGame.GameId, tx) {
		return dto.ResponseGetGame{}, err
	}

	err = tx.Commit()

	return requestCreateGame, err
}

func (g *GamesService) GetBoard(gameId int, userId int) (dto.ResponseGetBoard, error) {
	tx, err := g.gamesRepo.db.Begin()
	if err != nil {
		return dto.ResponseGetBoard{}, err
	}

	defer tx.Rollback()

	responseGetGame, err := g.gamesRepo.GetGame(gameId, tx)

	if userId != responseGetGame.WhiteUserId || userId != responseGetGame.BlackUserId {
		return dto.ResponseGetBoard{}, err
	}

	boardCells, err := g.boardRepo.GetBoardCells(gameId, tx)
	if err != nil {
		return dto.ResponseGetBoard{}, err
	}

	boardMap := map[int]int{}

	for _, cell := range boardCells {
		boardMap[cell.IndexCell] = cell.FigureId
	}

	var responseBoard []dto.BoardCell

	for i := 0; i < 64; i++ {
		if boardMap[i] > 0 {
			var cell = dto.BoardCell{i, boardMap[i]}
			responseBoard = append(responseBoard, cell)
		} else {
			var cell = dto.BoardCell{i, 0}
			responseBoard = append(responseBoard, cell)
		}
	}

	var responseGetBoard = dto.ResponseGetBoard{
		BoardCells: responseBoard,
	}

	err = tx.Commit()

	return responseGetBoard, err
}
