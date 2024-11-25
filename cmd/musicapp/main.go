package main

import (
	"fmt"
	"log"
	"music-service/internal/config"
	"music-service/pkg/logger"
)

func main() {

	logger := logger.New("local")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	logger.Info("hello")

}
