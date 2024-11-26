package main

import (
	"errors"
	"log"
	"music-service/internal/config"
	"music-service/pkg/logger"
	"music-service/pkg/migrations"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.New(cfg.Env)
	logger.Info("config read")
	logger.Info("config:", *cfg)

	migrator, err := migrations.NewMigrator(cfg.Postgres.Migrations, cfg.Postgres.URI)
	if err != nil {
		logger.Fatal("error initing migrator", err)
	}

	err = migrator.Up()
	if err != nil {
		if errors.Is(err, migrations.ErrNoChange) {
			logger.Info("no migrations to apply")
		} else {
			logger.Fatal("error making migrations", err)
		}
	}

}
