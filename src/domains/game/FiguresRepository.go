package game

import (
	"gorm.io/gorm"
)

type FiguresRepository struct {
	db *gorm.DB
}

func CreateFiguresRepository(db *gorm.DB) FiguresRepository {
	return FiguresRepository{
		db: db,
	}
}
