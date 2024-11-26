package postgres

import (
	"context"
	"database/sql"
	"music-service/internal/domain"
	"music-service/internal/storage"
)

type PostgresSongStorage struct {
	db *sql.DB
}

func NewPostgresSongStorage(db *sql.DB) *PostgresSongStorage {
	return &PostgresSongStorage{db: db}
}

func (s *PostgresSongStorage) AddSong(ctx context.Context, song *domain.Song) error {
	err := s.db.QueryRowContext(ctx, insertQuery, song.ReleaseDate, song.Lyrics, song.Link, song.Group.ID).Scan(&song.ID)
	return err
}

func (s *PostgresSongStorage) UpdateSong(ctx context.Context, songID int, updatedSong *domain.Song) error {
	query := `
		UPDATE songs
		SET release_date = $1, lyrics = $2, link = $3, group_id = $4
		WHERE id = $5
	`
	_, err := s.db.ExecContext(ctx, query, updatedSong.ReleaseDate, updatedSong.Lyrics, updatedSong.Link, updatedSong.Group.ID, songID)
	return err
}

func (s *PostgresSongStorage) GetSong(ctx context.Context, songID int) (*domain.Song, error) {
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

func (s *PostgresSongStorage) DeleteSong(ctx context.Context, songID int) error {
	_, err := s.db.ExecContext(ctx, deleteQuery, songID)
	return err
}

func (s *PostgresSongStorage) ListSongs(ctx context.Context, filter *storage.SongFilter) ([]*domain.Song, error) {
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

func (s *PostgresSongStorage) GetLyricsByVerse(ctx context.Context, id int, verse int) (string, error) {
	var verseText string
	err := s.db.QueryRowContext(ctx, getLyricsQuery, id, verse).Scan(&verseText)
	return verseText, err
}
