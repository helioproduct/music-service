package main

import (
	"log"
	"music-service/internal/config"
	"music-service/pkg/logger"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := logger.New(cfg.Env)
	logger.Info("config read")

}
