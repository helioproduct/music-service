package postgres

var (
	insertQuery = `
		INSERT INTO songs (release_date, lyrics, link, group_id)
		VALUES ($1, $2, $3, $4) RETURNING id
	`

	getSongQuery = `
		SELECT s.id, s.release_date, s.lyrics, s.link, g.id, g.name
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		WHERE s.id = $1
	`

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
