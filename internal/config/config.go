package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PostgresConfig struct {
	URI        string `env:"POSTGRES_URI"`
	Migrations string `env:"POSTGRES_MIGRATIONS"` // path to postgres migrations folder
}

type HTTPConfig struct {
	Address string
	Port    string `env:"HTTP_PORT"`
}

type Config struct {
	Env          string `env:"ENV"`
	Postgres     PostgresConfig
	HTTP         HTTPConfig
	SongsInfoURL string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("err loading .env file, %w", err)
	}

	config := &Config{
		Postgres: PostgresConfig{
			URI:        os.Getenv("POSTGRES_URI"),
			Migrations: os.Getenv("POSTGRES_MIGRATIONS"),
		},
		HTTP: HTTPConfig{
			Port:    os.Getenv("HTTP_Port"),
			Address: os.Getenv("HTTP_Address"),
		},
		SongsInfoURL: os.Getenv("SongsInfoURL"),
	}

	if config.Postgres.URI == "" {
		return nil, fmt.Errorf("POSTGRES_URI is not set")
	}
	if config.Postgres.Migrations == "" {
		return nil, fmt.Errorf("POSTGRES_MIGRATIONS_PATH is not set")
	}
	if config.HTTP.Port == "" {
		config.HTTP.Port = "8080" // default value
	}

	if config.Env == "" {
		config.Env = "local"
	}

	return config, nil
}
