package service

import (
	"context"
	"fmt"
	"music-service/internal/domain"
	"music-service/internal/repo"
	"music-service/internal/services"
	"time"
)

type songService struct {
	apiURL string
	repo   repo.SongRepo
}

func NewSongService(apiURL string, repo repo.SongRepo) services.SongService {
	return &songService{repo: repo}
}

func (s *songService) AddSong(ctx context.Context, song *domain.Song) (*domain.Song, error) {
	// Validate the input
	if song == nil {
		return nil, repo.ErrSongIsNil
	}

	if song.Group == nil || song.Group.Name == "" || song.Name == "" {
		return nil, services.ErrInvalidSong
	}

	details, err := FetchSongDetails(s.apiURL, song.Group.Name, song.Name)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch song details from external API: %w", err)
	}

	if details.ReleaseDate != "" {
		song.ReleaseDate, err = time.Parse(time.DateOnly, details.ReleaseDate)
		if err != nil {
			return nil, services.ErrParsingDate
		}
	}
	if details.Text != "" {
		song.Lyrics = details.Text
	}
	if details.Link != "" {
		song.Link = details.Link
	}

	return song, s.repo.AddSong(ctx, song)
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
