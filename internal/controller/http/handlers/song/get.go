package song

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type GetSongsRequest struct {
	ReleaseDate string `json:"release_date"` // Using string for flexibility in parsing
	Lyrics      string `json:"lyrics"`
	Link        string `json:"link"`
	GroupName   string `json:"group"`
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
	if r.ReleaseDate != "" {
		_, err := time.Parse("2006-01-02", r.ReleaseDate)
		if err != nil {
			return fmt.Errorf("invalid release date format, expected YYYY-MM-DD")
		}
	}
	return nil
}

func (h *SongHanlder) GetSongs(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Get songs handler hit")

	// Decode request
	var req GetSongsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate request
	if err := req.Validate(); err != nil {
		h.logger.Info("handler.GetSongs request validate error", "error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	h.logger.Info("get songs", "request", req)

	// Build filter parameters for the service
	filter := map[string]interface{}{}
	if req.ReleaseDate != "" {
		filter["release_date"] = req.ReleaseDate
	}
	if req.Lyrics != "" {
		filter["lyrics"] = req.Lyrics
	}
	if req.Link != "" {
		filter["link"] = req.Link
	}
	if req.GroupName != "" {
		filter["group_name"] = req.GroupName
	}

	// // Fetch songs from the service
	// songs, err := h.songService.GetSongs(r.Context(), filter)
	// if err != nil {
	// 	h.logger.Error("failed to fetch songs", err)
	// 	http.Error(w, "failed to fetch songs", http.StatusInternalServerError)
	// 	return
	// }

	// Respond with fetched songs
	// w.Header().Set("Content-Type", "application/json")
	// if err := json.NewEncoder(w).Encode(songs); err != nil {
	// 	h.logger.Error("failed to write response", err)
	// 	http.Error(w, "failed to write response", http.StatusInternalServerError)
	// 	return
	// }
}
