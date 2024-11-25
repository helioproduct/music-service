package postgres

import (
	"context"
	"database/sql"
	"music-service/internal/domain"
)

type SongPostgres struct {
	db *sql.DB
}

func AddSong(ctx context.Context, song *domain.Song) error {
	return nil
}
