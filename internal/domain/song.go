package domain

import "time"

type Song struct {
	ID          int
	ReleaseDate time.Time
	Lyrics      string
	Link        string
	Group       *Group
}
