package game

import "database/sql"

type MovesRepository struct {
	db *sql.DB
}

func CreateMovesRepository(db *sql.DB) MovesRepository {
	return MovesRepository{
		db: db,
	}
}
