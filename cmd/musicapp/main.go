package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"music-service/internal/config"
	"music-service/internal/domain"
	"music-service/pkg/logger"
	"music-service/pkg/migrations"
	"time"

	songrepo "music-service/internal/repo/song/postgres"

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

	db, err := sql.Open("postgres", cfg.Postgres.URI)
	if err != nil {
		logger.Fatal("error connecting to db", err)
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logger.Fatal("error db ping", err)
	}
	fmt.Println("Connected!")

	song := &domain.Song{
		ReleaseDate: time.Now(),
		Lyrics:      "fuck women",
		Link:        "google.com",
	}

	songStorage := songrepo.NewPostgres(db)
	err = songStorage.AddSong(context.Background(), song)
	if err != nil {
		logger.Info("error adding song", err)
	}

}
