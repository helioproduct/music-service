package song

import (
	"errors"
	"fmt"
	"music-service/internal/domain"
	"net/http"
	"strconv"
)

func (h *SongHandler) DeleteSong(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Delete song handler hit")

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

	err = h.songService.DeleteSong(r.Context(), songID)
	if err != nil {
		if errors.Is(err, domain.ErrNoSuchSong) {
			http.Error(w, fmt.Sprintf("failed to delete song: %v", err), http.StatusNotFound)
			return
		}

		h.logger.Error("failed to delete song", err)
		http.Error(w, fmt.Sprintf("failed to delete song: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Song deleted successfully")
}
