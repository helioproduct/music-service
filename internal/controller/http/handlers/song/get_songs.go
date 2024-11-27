package song

import (
	"encoding/json"
	"fmt"
	"music-service/internal/domain"
	"music-service/pkg/timex"
	"net/http"
	"net/url"
	"strconv"
)

type GetSongsRequest struct {
	ReleaseDate string `json:"release_date"`
	Lyrics      string `json:"lyrics"`
	Link        string `json:"link"`
	GroupName   string `json:"group"`
	Limit       int
	Offset      int
}

// adapter between layers
func BuildSongFilterFromRequest(req GetSongsRequest) (*domain.SongFilter, error) {
	builder := domain.NewSongFilter()

	if req.ReleaseDate != "" {
		releaseDate, err := timex.ParseDateOnly(req.ReleaseDate)
		if err != nil {
			return nil, err
		}
		builder.SetReleaseDate(releaseDate)
	}

	if req.Lyrics != "" {
		builder.SetLyrics(req.Lyrics)
	}
	if req.Link != "" {
		builder.SetLink(req.Link)
	}
	if req.GroupName != "" {
		builder.SetGroupName(req.GroupName)
	}

	builder.SetLimit(req.Limit)
	builder.SetOffset(req.Offset)

	return builder.Build(), nil
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
		_, err := timex.ParseDateOnly(r.ReleaseDate)
		if err != nil {
			return err
		}
	}
	if r.Limit < 0 {
		return fmt.Errorf("limit must be greater than or equal to 0")
	}
	if r.Offset < 0 {
		return fmt.Errorf("offset must be greater than or equal to 0")
	}
	return nil
}

func (h *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Get songs handler hit")

	query := r.URL.Query()

	req := GetSongsRequest{
		Limit:  domain.DefaulLimit,
		Offset: domain.DefaulOffset,
	}

	if limitStr := query.Get("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 0 {
			http.Error(w, "invalid limit, must be a non-negative integer", http.StatusBadRequest)
			return
		}
		req.Limit = limit
	}

	if offsetStr := query.Get("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil || offset < 0 {
			http.Error(w, "invalid offset, must be a non-negative integer", http.StatusBadRequest)
			return
		}
		req.Offset = offset
	}

	req.ReleaseDate = query.Get("release_date")
	req.Lyrics = query.Get("lyrics")
	req.Link = query.Get("link")
	req.GroupName = query.Get("group")

	if err := req.Validate(); err != nil {
		h.logger.Error("GetSongs", "validate error", err, "req", req)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter, err := BuildSongFilterFromRequest(req)
	if err != nil {
		h.logger.Error("GetSongs", "build filter error", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	songs, err := h.songService.GetSongs(r.Context(), filter)
	if err != nil {
		h.logger.Error("failed to fetch songs", err)
		http.Error(w, "failed to fetch songs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(songs); err != nil {
		h.logger.Error("failed to write response", err)
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
