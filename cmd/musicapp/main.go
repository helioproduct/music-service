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
	songservice "music-service/internal/services/song"

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

	// updateSong := &domain.Song{
	// 	Name:        "NEW NAME OLD TAPES",
	// 	Lyrics:      song.Lyrics,
	// 	ReleaseDate: time.Now().Add(time.Duration(time.Now().Year())),
	// 	Link:        "helio.com",
	// 	Group:       song.Group,
	// }

	// err = songStorage.AddSong(context.Background(), song)
	// if err != nil {
	// 	logger.Info("error adding song", "error", err)
	// }

	// song, err := songStorage.GetSong(context.Background(), 1)
	// if err != nil {
	// 	logger.Error("error getting song by id", "error", err)
	// 	return
	// }

	// fmt.Println(song)
	// fmt.Println(song.Group)

	// err = songStorage.DeleteSong(context.Background(), 2)
	// if err != nil {
	// 	logger.Error("error deleting song", "error", err)
	// 	return
	// }

	// err = songStorage.UpdateSong(context.Background(), 3, updateSong)
	// if err != nil {
	// 	logger.Error("error updateing song", "error", err)
	// }

	// filter := &repo.SongFilter{
	// 	Lyrics:    "fuck",
	// 	GroupName: "popov",
	// 	Limit:     5,
	// }

	// songs, err := songStorage.ListSongs(context.Background(), filter)
	// if err != nil {
	// 	logger.Error("error filtering", "error", err)
	// 	return
	// }

	// for _, song := range songs {
	// 	fmt.Println(song)
	// 	fmt.Println(song.Group)
	// 	fmt.Println()
	// }

	// verses, err := songsRepo.GetLyrics(context.Background(), 3, 3, 1)
	// if err != nil {
	// 	log.Fatalf("error retrieving lyrics: %v", err)
	// }

	// for i, verse := range verses {
	// 	fmt.Printf("Verse %d:\n%s\n\n", i+1, verse)
	// }

	songsRepo := songrepo.NewPostgres(db)
	songSvc := songservice.NewSongService(songsRepo)

	group := &domain.Group{
		Name: "helioproduct",
	}

	song := &domain.Song{
		Name:        "helio2",
		Lyrics:      "fuck this wo,an",
		Group:       group,
		ReleaseDate: time.Now(),
		Link:        "ya.ru",
	}

	err = songSvc.AddSong(context.Background(), song)
	if err != nil {
		logger.Error("service error adding song", "error", err)
	}

}
