package main

import (
	"errors"
	"log"
	"music-service/internal/config"
	"music-service/pkg/logger"

	"github.com/golang-migrate/migrate/v4"
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

	m, err := migrate.New("file://"+cfg.Postgres.Migrations, cfg.Postgres.URI+"?sslmode=disable")
	if err != nil {
		logger.Fatal("error creating migrations", err.Error())
	}

	if err := m.Up(); err != nil {
		// already up
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("no migrations to apply")
		} else {
			logger.Fatal(err.Error())
		}
	}

	// if err := m.Down(); err != nil {
	// 	logger.Error("error migrations down", err)
	// }
}
