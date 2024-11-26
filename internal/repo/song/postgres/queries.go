package postgres

var (
	insertSongQuery = `
		INSERT INTO songs (name, release_date, lyrics, link, group_id)
		VALUES ($1, $2, $3, $4, $5)`

	insertGroupQuery = "INSERT INTO groups (name) VALUES ($1) RETURNING id"

	getSongQuery = `
		SELECT s.id, s.release_date, s.lyrics, s.link, g.id, g.name
		FROM songs s
		LEFT JOIN groups g ON s.group_id = g.id
		WHERE s.id = $1`

	updateQuery = `
		UPDATE songs
		SET release_date = $1, lyrics = $2, link = $3
		WHERE id = $4`

	deleteQuery = `DELETE FROM songs WHERE id = $1`

	listSongsQuery = `
		SELECT s.id, s.name, s.release_date, s.lyrics, s.link, g.id, g.name as group_name
		FROM songs s
		JOIN groups g ON s.group_id = g.id`

	getLyricsQuery = `
        SELECT verse
        FROM regexp_split_to_table(
            (SELECT lyrics FROM songs WHERE id = $1),
            '\n\n'
        ) AS verse
        LIMIT $2 OFFSET $3;`
)
