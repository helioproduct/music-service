package domain

import "time"

type Song struct {
	ID int
	SongDetail
	Group *Group
}

type SongDetail struct {
	ReleaseDate time.Time
	Lyrics      string
	Link        string
}
