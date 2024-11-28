package song

import (
	"music-service/internal/services"
	"music-service/pkg/logger"
)

type SongHandler struct {
	logger      logger.Logger
	songService services.SongService
}

func NewHandler(svc services.SongService, logger logger.Logger) *SongHandler {
	return &SongHandler{
		songService: svc,
		logger:      logger,
	}
}
