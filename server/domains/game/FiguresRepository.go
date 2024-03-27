package game

import "database/sql"

type FiguresRepository struct {
	db *sql.DB
}

func CreateFiguresRepository(db *sql.DB) FiguresRepository {
	return FiguresRepository{
		db: db,
	}
}
