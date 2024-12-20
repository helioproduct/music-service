package song

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *SongHandler) GetLyrics(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Get lyrics handler hit")

	songIDStr := r.URL.Query().Get("song_id")
	offsetStr := r.URL.Query().Get("offset")
	limitStr := r.URL.Query().Get("limit")

	songID, err := strconv.Atoi(songIDStr)
	if err != nil || songID <= 0 {
		http.Error(w, "invalid song_id parameter", http.StatusBadRequest)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		h.logger.Error("GetLyrics", "error", err)
		http.Error(w, "invalid offset parameter", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		h.logger.Error("DeleteSongs", "error", err)
		http.Error(w, "invalid limit parameter", http.StatusBadRequest)
		return
	}

	verses, err := h.songService.GetLyrics(r.Context(), songID, offset, limit)
	if err != nil {
		h.logger.Error("failed to get lyrics", err)
		http.Error(w, "failed to get lyrics", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(verses); err != nil {
		h.logger.Error("failed to write response", err)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
	}
}
