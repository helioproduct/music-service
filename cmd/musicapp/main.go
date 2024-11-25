package main

import (
	"fmt"
	"log"
	"music-service/internal/config"
)

func main() {

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
