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
	Port string `env:"HTTP_PORT"`
}

type Config struct {
	Postgres PostgresConfig
	HTTP     HTTPConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("err loading .env file, %w", err)
	}

	config := &Config{
		Postgres: PostgresConfig{
			URI:        os.Getenv("POSTGRES_PATH"),
			Migrations: os.Getenv("POSTGRES_MIGRATIONS_PATH"),
		},
		HTTP: HTTPConfig{
			Port: os.Getenv("HTTP_PORT"),
		},
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

	return config, nil
}
