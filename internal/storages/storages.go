package storages

import (
	"context"
	"music-service/internal/domain"
)

type SongStorage interface {
	AddSong(ctx context.Context, song *domain.Song) error
	UpdateSong(ctx context.Context, songID int, updatedSong *domain.Song) error
	GetSong(ctx context.Context, songID int) (*domain.Song, error)
	DeleteSong(ctx context.Context, songID int) error
	ListSongs(ctx context.Context, filter *SongFilter) ([]*domain.Song, error)
	GetLyricsByVerse(ctx context.Context, id int, verse int) (string, error)
}

type GroupStorage interface {
	AddGroup(ctx context.Context, group *domain.Group) (int, error)
	GetGroupByNanme(ctx context.Context, groupName string) (*domain.Group, error)
	DeleteGroup(ctx context.Context, groupName string) error
}
