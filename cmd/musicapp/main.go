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
	middleware "music-service/internal/controller/http/middleware"
	songsrepo "music-service/internal/repo/song/postgres"
	songservice "music-service/internal/services/song"

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
		logger.Fatal("error initing migrator", "error", err)
	}

	err = migrator.Up()
	if err != nil {
		if errors.Is(err, migrations.ErrNoChange) {
			logger.Info("no migrations to apply")
		} else {
			logger.Fatal("error making migrations", "error", err)
		}
	}

	db, err := sql.Open("postgres", cfg.Postgres.URI)
	if err != nil {
		logger.Fatal("error connecting to db", "error", err)
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		logger.Fatal("error db ping", "error", err)
	} else {
		logger.Info("connected to db")
	}

	songsRepo := songsrepo.NewPostgres(db)
	songService := songservice.NewSongService(cfg.SongsInfoURL, songsRepo)
	songsHandler := songshandler.NewHandler(songService, logger)

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/songs", songsHandler.GetSongs).Methods("GET")
	api.HandleFunc("/songs", songsHandler.AddSong).Methods("PUT")
	api.HandleFunc("/lyrics", songsHandler.GetLyrics).Methods("GET")
	api.HandleFunc("/songs", songsHandler.DeleteSong).Methods("DELETE")
	api.HandleFunc("/songs", songsHandler.UpdateSong).Methods("PATCH")

	r.Use(middleware.Logging(logger))
	r.Use(middleware.PanicRecoverer(logger))

	done := make(chan bool)

	go func() {
		err := http.ListenAndServe(cfg.HTTP.Address+":"+cfg.HTTP.Port, r)
		if err != nil {
			logger.Error("server", "error", err)
		}
		done <- true
	}()

	logger.Info("started server at", "PORT", cfg.HTTP.Port)
	<-done
}
