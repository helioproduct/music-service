package main

import (
	"database/sql"
	"errors"
	"log"
	"music-service/internal/config"
	"music-service/pkg/logger"
	"music-service/pkg/migrations"
	"net/http"

	songshandler "music-service/internal/controller/http/handlers/song"
	songsrepo "music-service/internal/repo/song/postgres"
	songservice "music-service/internal/services/song"

	// _ "github.com/golang-migrate/migrate/v4/database/postgres"
	// _ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gorilla/mux"
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

	songsRepo := songsrepo.NewPostgres(db)
	songService := songservice.NewSongService(cfg.SongsInfoURL, songsRepo)
	songsHandler := songshandler.NewHandler(songService, logger)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/songs", songsHandler.GetSongs).Methods("GET")
	api.HandleFunc("/songs", songsHandler.AddSong).Methods("PUT")
	api.HandleFunc("/lyrics", songsHandler.GetLyrics).Methods("GET")

	done := make(chan bool)

	go func() {
		err := http.ListenAndServe("localhost:"+cfg.HTTP.Port, r)
		if err != nil {
			logger.Error("server", "error", err)
		}
		done <- true
	}()

	logger.Info("started server at", cfg.HTTP.Port)
	<-done
}
