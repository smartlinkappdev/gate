package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env     string
	Port    string
	Storage Storage
	Auth    Auth
}

type Storage struct {
	DSN string
}

type Auth struct {
	DSN string
}

func New() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	return &Config{
		Env:  os.Getenv("ENV"),
		Port: os.Getenv("PORT"),
		Storage: Storage{
			DSN: os.Getenv("DATABASE_DSN"),
		},
		Auth: Auth{
			DSN: os.Getenv("AUTH_DSN"),
		},
	}
}
