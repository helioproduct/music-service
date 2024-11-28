package domain

import "errors"

var (
	// group
	ErrSongIsNil  = errors.New("song is nil")
	ErrGroupIsNil = errors.New("group is nil")
	// song
	ErrNoSuchSong = errors.New("song does not exist")
	// filter
	ErrFilterIsNil = errors.New("filter is nil")
)
