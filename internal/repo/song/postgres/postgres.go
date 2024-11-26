package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Ensure the group exists or create it
	var groupID int
	err = tx.QueryRowContext(ctx, "SELECT id FROM groups WHERE name = $1", song.Group.Name).Scan(&groupID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Group does not exist, insert it
			err = tx.QueryRowContext(ctx, insertGroupQuery, song.Group.Name).Scan(&groupID)
			if err != nil {
				return fmt.Errorf("error inserting group: %w", err)
			}
			// update ID in struct
			song.Group.ID = groupID
		} else {
			return fmt.Errorf("error checking group existence: %w", err)
		}
	}

	_, err = tx.ExecContext(ctx, insertSongQuery,
		song.Name,
		song.ReleaseDate,
		song.Lyrics,
		song.Link,
		groupID)

	if err != nil {
		return fmt.Errorf("error inserting song: %w", err)
	}

	return nil
}

func (s *PostgresRepo) UpdateSong(ctx context.Context, songID int, updatedSong *domain.Song) error {
	if updatedSong == nil {
		return repo.ErrSongIsNil
	}
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrNoSuchSong
		}
		return nil, err
	}
	song.Group = &group
	return &song, nil
}

func (s *PostgresRepo) DeleteSong(ctx context.Context, songID int) error {
	result, err := s.db.ExecContext(ctx, deleteQuery, songID)
	affected, _ := result.RowsAffected()
	if affected < 1 {
		return repo.ErrNoSuchSong
	}
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
