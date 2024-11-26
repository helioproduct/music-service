package song

import (
	"fmt"
	"net/url"
	"time"
)

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
