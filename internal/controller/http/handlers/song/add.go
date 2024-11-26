package song

import "fmt"

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
