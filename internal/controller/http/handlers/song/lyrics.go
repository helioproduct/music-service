package song

type GetLyricsRequest struct {
}

type Lyrics struct {
	Text string `json:"lyrics"`
}
