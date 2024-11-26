package song

import (
	"fmt"
	"net/url"
	"time"
)

type AddSongRequest struct {
	Song  string `json:"song"`
	Group string `json:"group"`
}

func (r *AddSongRequest) Validate() error {
	if r.Song == "" {
		return fmt.Errorf("song cannot be empty")
	}
	if r.Group == "" {
		return fmt.Errorf("group cannot be empty")
	}
	return nil
}

type DeleteSongRequest struct {
	SongID int `json:"songID"`
}

type GetSongsRequest struct {
	ReleaseDate time.Time `json:"release-date"`
	Lyrics      string    `json:"lyrics"`
	Link        string    `json:"link"`
	GroupName   string    `json:"group"`
}

func (r *GetSongsRequest) Validate() error {
	if r.Link != "" {
		u, err := url.Parse(r.Link)
		if err != nil {
			return fmt.Errorf("not a valid URL")
		} else if u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("not a valid URL")
		}
	}
	return nil
}

type UpdateSongRequest struct {
}
