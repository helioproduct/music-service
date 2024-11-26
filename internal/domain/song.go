package domain

import "time"

type Song struct {
	ID          int
	Name        string
	ReleaseDate time.Time
	Lyrics      string
	Link        string
	Group       *Group
}
