package migrations

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migrator struct {
	PostgresURI string
	Source      string
	migrate     *migrate.Migrate
}

func (m *Migrator) Up() error {
	err := m.migrate.Up()
	if errors.Is(err, migrate.ErrNoChange) {
		return ErrNoChange
	}
	return err
}

func (m *Migrator) Down() error {
	err := m.migrate.Down()
	if errors.Is(err, migrate.ErrNoChange) {
		return ErrNoChange
	}
	return err
}

func NewMigrator(source, postgresURI string) (*Migrator, error) {
	m, err := migrate.New(source, postgresURI)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Migrator{
		Source:      source,
		PostgresURI: postgresURI,
		migrate:     m,
	}, nil
}
