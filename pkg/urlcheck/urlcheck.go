package urlcheck

import (
	"errors"
	"net/url"
)

var (
	ErrNotValidURL = errors.New("not a valid URL")
)

func IsValidURL(link string) error {
	if link != "" {
		u, err := url.Parse(link)
		if err != nil {
			return ErrNotValidURL
		} else if u.Scheme == "" || u.Host == "" {
			return ErrNotValidURL
		}
	}
	return nil
}
