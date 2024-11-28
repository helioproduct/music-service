package services

import "errors"

var (
	ErrParsingDate = errors.New("error parsing date")
	ErrInvalidSong = errors.New("invalid song")
)
