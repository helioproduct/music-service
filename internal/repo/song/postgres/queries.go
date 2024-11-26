package postgres

// PostgreSQL format
var (
	insertSongQuery = `
		INSERT INTO songs (name, release_date, lyrics, link, group_id)
		VALUES ($1, $2, $3, $4, $5)
	`

	insertGroupQuery = "INSERT INTO groups (name) VALUES ($1) RETURNING id"

	getSongQuery = `
		SELECT s.id, s.release_date, s.lyrics, s.link, g.id, g.name
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		WHERE s.id = $1`

	updateQuery = `
		UPDATE songs
		SET release_date = $1, lyrics = $2, link = $3, group_id = $4
		WHERE id = $5`

	deleteQuery = `DELETE FROM songs WHERE id = $1`

	listSongsQuery = `
		SELECT s.id, s.release_date, s.lyrics, s.link, g.id, g.name
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		ORDER BY s.release_date DESC
	`

	getLyricsQuery = `
		SELECT COALESCE(SPLIT_PART(lyrics, '\n', $2), '')
		FROM songs
		WHERE id = $1
	`
)
