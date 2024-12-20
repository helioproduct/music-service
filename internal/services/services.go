package services

import (
	"context"
	"music-service/internal/domain"
)

type SongService interface {
	AddSong(ctx context.Context, song *domain.Song) (*domain.Song, error)
	UpdateSong(ctx context.Context, songID int, updatedSong *domain.Song) error
	GetSong(ctx context.Context, songID int) (*domain.Song, error)
	DeleteSong(ctx context.Context, songID int) error
	GetSongs(ctx context.Context, filter *domain.SongFilter) ([]*domain.Song, error)
	GetLyrics(ctx context.Context, songID, offset, limit int) ([]string, error)
}
