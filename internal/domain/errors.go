package domain

import "errors"

var (
	ErrNameIsEmpty  = errors.New("name is empty")
	ErrGroupIsEmpty = errors.New("group is empty")
	// group
	ErrSongIsNil  = errors.New("song is nil")
	ErrGroupIsNil = errors.New("group is nil")
	// song
	ErrNoSuchSong = errors.New("song does not exist")
	// filter
	ErrFilterIsNil = errors.New("filter is nil")
)
