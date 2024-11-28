package service

import (
	"context"
	"music-service/internal/domain"
	"music-service/internal/repo"
	"music-service/internal/services"
	"time"
)

type songService struct {
	repo repo.SongRepo
}

func NewSongService(repo repo.SongRepo) services.SongService {
	return &songService{repo: repo}
}

func (s *songService) AddSong(ctx context.Context, song *domain.Song) error {
	if song == nil {
		return repo.ErrSongIsNil
	}

	var err error
	song.ReleaseDate, err = time.Parse(time.DateOnly, song.ReleaseDate.Format(time.DateOnly))
	if err != nil {
		return services.ErrParsingDate
	}

	return s.repo.AddSong(ctx, song)
}

func (s *songService) UpdateSong(ctx context.Context, songID int, updatedSong *domain.Song) error {
	return s.repo.UpdateSong(ctx, songID, updatedSong)
}

func (s *songService) GetSong(ctx context.Context, songID int) (*domain.Song, error) {
	return s.repo.GetSong(ctx, songID)
}

func (s *songService) DeleteSong(ctx context.Context, songID int) error {
	return s.repo.DeleteSong(ctx, songID)
}

func (s *songService) GetSongs(ctx context.Context, filter *domain.SongFilter) ([]*domain.Song, error) {
	return s.repo.ListSongs(ctx, filter)
}

func (s *songService) GetLyrics(ctx context.Context, songID, offset, limit int) ([]string, error) {
	return s.repo.GetLyrics(ctx, songID, offset, limit)
}
