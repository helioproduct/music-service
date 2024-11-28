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

func (s *Song) Validate() error {
	if s == nil {
		return ErrSongIsNil
	}
	if s.Group == nil {
		return ErrGroupIsNil
	}
	if s.Name == "" {
		return ErrNameIsEmpty
	}
	if s.Group.Name == "" {
		return ErrGroupIsEmpty
	}
	return nil
}
