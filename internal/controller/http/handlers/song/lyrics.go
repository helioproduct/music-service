package song

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type GetLyricsRequest struct {
}

type GetLyricsResponse struct {
	Verses []string `json:"verses"`
}

func (h *SongHandler) GetLyrics(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Get lyrics handler hit")

	// Parse query parameters
	songIDStr := r.URL.Query().Get("song_id")
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	// Validate and convert query parameters
	songID, err := strconv.Atoi(songIDStr)
	if err != nil || songID <= 0 {
		http.Error(w, "invalid song_id parameter", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		http.Error(w, "invalid offset parameter", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		http.Error(w, "invalid limit parameter", http.StatusBadRequest)
		return
	}

	lyricsResponse := new(GetLyricsResponse)

	// Call the service method to get lyrics
	lyricsResponse.Verses, err = h.songService.GetLyrics(r.Context(), songID, offset, limit)
	if err != nil {
		h.logger.Error("failed to get lyrics", err)
		http.Error(w, "failed to get lyrics", http.StatusInternalServerError)
		return
	}

	// Respond with the lyrics
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(lyricsResponse); err != nil {
		h.logger.Error("failed to write response", err)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
