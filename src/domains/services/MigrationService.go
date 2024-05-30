package services

import (
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "time/tzdata"
)

type MigrationService struct {
}

func CreateMigrationService() MigrationService {
	return MigrationService{}
}

func (s *MigrationService) RunUp(postgresqlUrl string, pathToMigrations string) error {
	migration, err := migrate.New(pathToMigrations, postgresqlUrl)
	if err != nil {
		return err
	}

	defer migration.Close()

	err = migration.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}
