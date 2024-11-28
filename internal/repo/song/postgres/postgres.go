package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"music-service/internal/domain"
	"music-service/internal/repo"
	"strings"
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

	var groupID int
	err = tx.QueryRowContext(ctx, "SELECT id FROM groups WHERE name = $1", song.Group.Name).Scan(&groupID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Group does not exist, insert it
			err = tx.QueryRowContext(ctx, insertGroupQuery, song.Group.Name).Scan(&groupID)
			// fmt.Println()
			if err != nil {
				log.Println()
				return fmt.Errorf("error inserting group: %w", err)
			}
			// update ID in struct
			song.Group.ID = groupID
			log.Println("NEW GROUP ID", groupID)
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
	result, err := s.db.ExecContext(ctx, updateQuery, updatedSong.ReleaseDate, updatedSong.Lyrics, updatedSong.Link, songID)
	affected, _ := result.RowsAffected()
	if affected < 1 {
		return repo.ErrNoSuchSong
	}
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

func (r *PostgresRepo) ListSongs(ctx context.Context, filter *domain.SongFilter) ([]*domain.Song, error) {
	if filter == nil {
		return nil, repo.ErrFilterIsNil
	}
	query := listSongsQuery
	// Conditions and arguments
	conditions := []string{}
	args := []interface{}{}
	argIndex := 1

	if filter.ReleaseDate != nil {
		conditions = append(conditions, fmt.Sprintf("s.release_date = $%d", argIndex))
		args = append(args, *filter.ReleaseDate)
		argIndex++
	}

	if filter.Lyrics != "" {
		conditions = append(conditions, fmt.Sprintf("s.lyrics ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Lyrics+"%")
		argIndex++
	}

	if filter.Link != "" {
		conditions = append(conditions, fmt.Sprintf("s.link ILIKE $%d", argIndex))
		args = append(args, "%"+filter.Link+"%")
		argIndex++
	}

	if filter.GroupName != "" {
		conditions = append(conditions, fmt.Sprintf("g.name ILIKE $%d", argIndex))
		args = append(args, "%"+filter.GroupName+"%")
		argIndex++
	}

	if len(conditions) > 0 {
		query += "\nWHERE " + strings.Join(conditions, " AND ") + " "
	}

	query += fmt.Sprintf("LIMIT $%d OFFSET $%d", argIndex, argIndex+1)
	args = append(args, filter.Limit, filter.Offset)

	// fmt.Println(query)
	// fmt.Println(args...)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %w", err)
	}

	defer rows.Close()

	var songs []*domain.Song
	for rows.Next() {
		song := new(domain.Song)
		group := new(domain.Group)
		song.Group = group

		err := rows.Scan(&song.ID, &song.Name, &song.ReleaseDate, &song.Lyrics, &song.Link,
			&group.ID, &group.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
		songs = append(songs, song)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return songs, nil
}

func (r *PostgresRepo) GetLyrics(ctx context.Context, songID, offset, limit int) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, getLyricsQuery, songID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error fetching lyrics: %w", err)
	}
	defer rows.Close()

	var verses []string
	for rows.Next() {
		var verse string
		if err := rows.Scan(&verse); err != nil {
			return nil, fmt.Errorf("error scanning verse: %w", err)
		}
		verses = append(verses, verse)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}
	return verses, nil
}
