package song

import (
	"music-service/internal/services"
	"music-service/pkg/logger"
)

type SongHanlder struct {
	logger      logger.Logger
	songService services.SongService
}

func NewHandler(svc services.SongService) *SongHanlder {
	return &SongHanlder{
		songService: svc,
	}
}
