package song

import (
	"encoding/json"
	"errors"
	"fmt"
	"music-service/internal/domain"
	"music-service/pkg/timex"
	"music-service/pkg/urlcheck"
	"net/http"
	"time"
)

type UpdateSongRequest struct {
	SongID         int    `json:"song_id"`
	NewName        string `json:"new_name"`
	NewReleaseDate string `json:"release_date"`
	NewLyrics      string `json:"lyrics"`
	NewLink        string `json:"link"`
}

func (req *UpdateSongRequest) Validate() error {
	if req.SongID <= 0 {
		return fmt.Errorf("song_id must be positive")
	}
	if req.NewReleaseDate != "" {
		if _, err := timex.ParseDateOnly(req.NewReleaseDate); err != nil {
			return err
		}
	}
	if req.NewLink != "" {
		if err := urlcheck.IsValidURL(req.NewLink); err != nil {
			return err
		}
	}
	return nil
}

func (h *SongHandler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Update song handler hit")

	var req UpdateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("validation error: %v", err), http.StatusBadRequest)
		return
	}

	updatedSong := new(domain.Song)
	if req.NewName != "" {
		updatedSong.Name = req.NewName
	}
	if req.NewReleaseDate != "" {
		newTime, _ := time.Parse(time.DateOnly, req.NewReleaseDate)
		updatedSong.ReleaseDate = newTime
	}
	if req.NewLyrics != "" {
		updatedSong.Lyrics = req.NewLyrics
	}
	if req.NewLink != "" {
		updatedSong.Link = req.NewLink
	}

	err := h.songService.UpdateSong(r.Context(), req.SongID, updatedSong)
	if err != nil {
		h.logger.Error("failed to update song", "error", err)
		if errors.Is(err, domain.ErrNoSuchSong) {
			h.logger.Error("failed to delete song", err)
			if errors.Is(err, domain.ErrNoSuchSong) {
				http.Error(w, fmt.Sprintf("failed to delete song: %v", err), http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNotFound)
		}

		http.Error(w, "failed to update song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Song updated successfully")
}
