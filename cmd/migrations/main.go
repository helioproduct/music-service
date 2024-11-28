package main

import (
	"errors"
	"log"
	"music-service/internal/config"
	"music-service/pkg/logger"
	"music-service/pkg/migrations"
	"os"
)

// CLI tool for applying migrations on debug
func main() {
	if len(os.Args) < 2 {
		log.Println("Usage: main <command> (up | down)")
		return
	}

	command := os.Args[1]

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.New(cfg.Env)

	migrator, err := migrations.NewMigrator(cfg.Postgres.Migrations, cfg.Postgres.URI)
	if err != nil {
		logger.Fatal("error initializing migrator", err)
	}

	switch command {
	case "up":
		err = migrator.Up()
		if err != nil {
			if errors.Is(err, migrations.ErrNoChange) {
				logger.Info(err.Error())
				return
			} else {
				logger.Fatal("error making migrations", err)
			}
		}
		logger.Info("Up migrations applied successfully")

	case "down":
		err = migrator.Down()
		if err != nil {
			if errors.Is(err, migrations.ErrNoChange) {
				logger.Info(err.Error())
				return
			} else {
				logger.Fatal("error making migrations", err)
			}
		}
		logger.Info("Down migrations applied successfully")

	default:
		log.Println("Unknown command. Usage: main <command> (up | down)")
		return
	}
}
