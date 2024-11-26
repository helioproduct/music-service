package postgres

import (
	"context"
	"database/sql"
	"music-service/internal/domain"
	"music-service/internal/repo"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (s *PostgresRepo) AddSong(ctx context.Context, song *domain.Song) error {
	if song == nil {
		return repo.ErrSongIsNil
	} else if song.Group == nil {
		return repo.ErrGroupIsNil
	}

	err := s.db.QueryRowContext(ctx, insertQuery, song.ReleaseDate, song.Lyrics, song.Link, song.Group.ID).Scan(&song.ID)
	return err
}

func (s *PostgresRepo) UpdateSong(ctx context.Context, songID int, updatedSong *domain.Song) error {
	_, err := s.db.ExecContext(ctx, updateQuery, updatedSong.ReleaseDate, updatedSong.Lyrics, updatedSong.Link, updatedSong.Group.ID, songID)
	return err
}

func (s *PostgresRepo) GetSong(ctx context.Context, songID int) (*domain.Song, error) {
	var song domain.Song
	var group domain.Group
	err := s.db.QueryRowContext(ctx, getSongQuery, songID).Scan(
		&song.ID, &song.ReleaseDate, &song.Lyrics, &song.Link,
		&group.ID, &group.Name,
	)
	if err != nil {
		return nil, err
	}
	song.Group = &group
	return &song, nil
}

func (s *PostgresRepo) DeleteSong(ctx context.Context, songID int) error {
	_, err := s.db.ExecContext(ctx, deleteQuery, songID)
	return err
}

func (s *PostgresRepo) ListSongs(ctx context.Context, filter *repo.SongFilter) ([]*domain.Song, error) {
	rows, err := s.db.QueryContext(ctx, listSongsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []*domain.Song
	for rows.Next() {
		var song domain.Song
		var group domain.Group
		err = rows.Scan(&song.ID, &song.ReleaseDate, &song.Lyrics, &song.Link, &group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		song.Group = &group
		songs = append(songs, &song)
	}

	return songs, rows.Err()
}

func (s *PostgresRepo) GetLyricsByVerse(ctx context.Context, id int, verse int) (string, error) {
	var verseText string
	err := s.db.QueryRowContext(ctx, getLyricsQuery, id, verse).Scan(&verseText)
	return verseText, err
}
