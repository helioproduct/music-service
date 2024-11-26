package repo

import "errors"

var (
	ErrSongIsNil  = errors.New("song is nil")
	ErrGroupIsNil = errors.New("group is nil")
	ErrNoSuchSong = errors.New("song does not exist")
)
