package main

import (
	"cmd/gate/main.go/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	conn, err := gorm.Open(postgres.Open("host=62.109.5.203 user=akaletr password=adelaida2011 dbname=gate port=5432"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.AutoMigrate(entity.User{})
}
