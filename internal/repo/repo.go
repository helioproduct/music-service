package repo

import (
	"context"
	"music-service/internal/domain"
)

type SongRepo interface {
	AddSong(ctx context.Context, song *domain.Song) error
	UpdateSong(ctx context.Context, songID int, updatedSong *domain.Song) error
	GetSong(ctx context.Context, songID int) (*domain.Song, error)
	DeleteSong(ctx context.Context, songID int) error
	ListSongs(ctx context.Context, filter *SongFilter) ([]*domain.Song, error)
	GetLyrics(ctx context.Context, offset, limit int) ([]string, error)
}
