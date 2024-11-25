package song

import (
	"music-service/internal/services"
)

type SongRoutes struct {
	songService services.SongService
}

func NewSongRoutes(svc services.SongService) *SongRoutes {
	return &SongRoutes{
		songService: svc,
	}
}
