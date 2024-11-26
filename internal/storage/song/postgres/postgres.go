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
	query := `
		INSERT INTO songs (release_date, lyrics, link, group_id)
		VALUES ($1, $2, $3, $4) RETURNING id
	`
	err := s.db.QueryRowContext(ctx, query, song.ReleaseDate, song.Lyrics, song.Link, song.Group.ID).Scan(&song.ID)
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
	query := `
		SELECT s.id, s.release_date, s.lyrics, s.link, g.id, g.name
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		WHERE s.id = $1
	`
	var song domain.Song
	var group domain.Group
	err := s.db.QueryRowContext(ctx, query, songID).Scan(
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
	query := `
		DELETE FROM songs WHERE id = $1
	`
	_, err := s.db.ExecContext(ctx, query, songID)
	return err
}

func (s *PostgresSongStorage) ListSongs(ctx context.Context, filter *storage.SongFilter) ([]*domain.Song, error) {
	query := `
		SELECT s.id, s.release_date, s.lyrics, s.link, g.id, g.name
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		ORDER BY s.release_date DESC
	`
	rows, err := s.db.QueryContext(ctx, query)
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
	query := `
		SELECT COALESCE(SPLIT_PART(lyrics, '\n', $2), '')
		FROM songs
		WHERE id = $1
	`
	var verseText string
	err := s.db.QueryRowContext(ctx, query, id, verse).Scan(&verseText)
	return verseText, err
}
