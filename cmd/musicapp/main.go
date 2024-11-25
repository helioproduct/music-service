package main

import (
	"errors"
	"log"
	"music-service/internal/config"
	"music-service/pkg/logger"

	"github.com/golang-migrate/migrate"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.New(cfg.Env)
	logger.Info("config read")

	m, err := migrate.New("file://"+cfg.Postgres.Migrations, cfg.Postgres.URI)
	if err != nil {
		logger.Fatal(err.Error())
	}

	if err := m.Up(); err != nil {
		// already up
		if errors.Is(err, migrate.ErrNoChange) {
			logger.Info("no migrations to apply")
		} else {
			logger.Fatal(err.Error())
		}
	}

}
