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

	Conn2DSN string

	CHHost     string
	CHPort     string
	CHUsername string
	CHPassword string
	CHDatabase string
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

		Conn2DSN: os.Getenv("CONN2_DSN"),

		CHHost:     os.Getenv("CH_HOST"),
		CHPort:     os.Getenv("CH_PORT"),
		CHUsername: os.Getenv("CH_USERNAME"),
		CHPassword: os.Getenv("CH_PASSWORD"),
		CHDatabase: os.Getenv("CH_DATABASE"),
	}
}
