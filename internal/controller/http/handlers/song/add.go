package song

import (
	"encoding/json"
	"fmt"
	"music-service/internal/domain"
	"net/http"
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

func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Add song handler hit")

	var req AddSongRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepare the domain Song object
	song := &domain.Song{
		Name:  req.Song,
		Group: &domain.Group{Name: req.Group},
	}

	// Call the service layer to add the song
	var err error
	if song, err = h.songService.AddSong(r.Context(), song); err != nil {
		h.logger.Error("failed to add song", err)
		http.Error(w, "failed to add song", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(song); err != nil {
		h.logger.Error("failed to write response", err)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}

	// Respond with success
	// w.WriteHeader(http.StatusCreated)
	// fmt.Fprintln(w, "Song added successfully")
}
