package config

import "os"

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
