package song

import (
	"encoding/json"
	"fmt"
	"music-service/internal/domain"
	"net/http"
	"strconv"
)

type UpdateSongRequest struct {
	Name        string `json:"name"`
	ReleaseDate string `json:"release_date"`
	Lyrics      string `json:"lyrics"`
	Link        string `json:"link"`
	GroupID     int    `json:"group_id"`
}

func (req *UpdateSongRequest) Validate() error {
	if req.Name == "" {
		return fmt.Errorf("name cannot be empty")
	}
	if req.ReleaseDate == "" {
		return fmt.Errorf("release date cannot be empty")
	}
	if req.GroupID <= 0 {
		return fmt.Errorf("group_id must be a positive integer")
	}
	return nil
}

func (h *SongHandler) UpdateSong(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Update song handler hit")

	songIDStr := r.URL.Query().Get("songID")
	if songIDStr == "" {
		http.Error(w, "missing songID query parameter", http.StatusBadRequest)
		return
	}

	songID, err := strconv.Atoi(songIDStr)
	if err != nil || songID <= 0 {
		http.Error(w, "invalid songID query parameter", http.StatusBadRequest)
		return
	}

	var req UpdateSongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, fmt.Sprintf("validation error: %v", err), http.StatusBadRequest)
		return
	}

	// Prepare the updated song object
	updatedSong := &domain.Song{
		Name:        req.Name,
		ReleaseDate: req.ReleaseDate,
		Lyrics:      req.Lyrics,
		Link:        req.Link,
		Group: &domain.Group{
			ID: req.GroupID,
		},
	}

	err = h.songService.UpdateSong(r.Context(), songID, updatedSong)
	if err != nil {
		h.logger.Error("failed to update song", err)
		http.Error(w, "failed to update song", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Song updated successfully")
}
