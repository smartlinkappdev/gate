package main

import (
	"cmd/gate/main.go/internal/auth"
	"cmd/gate/main.go/internal/entity"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	conn, err := gorm.Open(postgres.Open("host=62.109.5.203 user=akaletr password=adelaida2011 dbname=gate port=5432"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn.Migrator().DropTable(entity.User{}, entity.Group{}, entity.Feedback{}, entity.Link{})
	err = conn.AutoMigrate(entity.User{}, entity.Group{}, entity.Feedback{}, entity.Link{})
	if err != nil {
		fmt.Println("============", err)
	}

	conn2, err := gorm.Open(postgres.Open("host=62.109.5.203 user=akaletr password=adelaida2011 dbname=auth port=5432"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	conn2.Migrator().DropTable(auth.UserAuth{})
	conn2.AutoMigrate(auth.UserAuth{})

}
