package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	// Mocked JSON response for development or testing
	mockedResponse = `{
		"releaseDate": "2000-08-08",
		"text": "Breathe, breathe in the air\nDon't be afraid to care",
		"link": "https://example.com/breathe"
	}`
)

type SongDetail struct {
	ReleaseDate string `json:"releaseDate"`
	Text        string `json:"text"`
	Link        string `json:"link"`
}

func FetchSongDetails(url, group, song string) (*SongDetail, error) {
	// reqURL := fmt.Sprintf("%s/info?group=%s&song=%s", songsApiURL, group, song)
	// resp, err := http.Get(reqURL)

	// Mocking an HTTP response
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(mockedResponse)),
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch song details: status %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var songDetail SongDetail
	if err := json.NewDecoder(resp.Body).Decode(&songDetail); err != nil {
		return nil, fmt.Errorf("failed to decode song details: %w", err)
	}

	return &songDetail, nil
}
