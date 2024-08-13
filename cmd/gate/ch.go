package main

import (
	"cmd/gate/main.go/internal/ym/dal"
	"context"
	"crypto/tls"
	"fmt"
	"gorm.io/gorm"

	ch "github.com/ClickHouse/clickhouse-go/v2"
	"gorm.io/driver/clickhouse"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/sl/dal",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	host := "ch.platina360.ru"
	port := 9443
	username := "zosimenko%40adventum.ru"
	password := "hi%2BRCVE%23LR9qxACE"
	database := "sberbank"

	dsn := fmt.Sprintf(
		"https://%s:%d?username=%s&password=%s&database=%s&secure=true",
		host, port, username, password, database,
	)

	//dsn := "https://zosimenko@adventum.ru:hi+RCVE#LR9qxACE@ch.platina360.ru:9443?database=sberbank&tls=custom"
	//dsn := "https://ch.platina360.ru:9443?database=sberbank&username=zosimenko@adventum.ru&password=hi+RCVE#LR9qxACE&secure=true&tls=false"

	dsn1 := "zosimenko%40adventum.ru:hi%2BRCVE%23LR9qxACE@ch.platina360.ru:9443?sberbank"

	conn, err := ch.Open(&ch.Options{
		TLS:  tlsConfig,
		Addr: []string{dsn1},
	})

	err = conn.Ping(context.Background())
	if err != nil {
		fmt.Println(err)
	}

	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	dal.SetDefault(db)

	g.UseDB(db) // reuse your gorm db
	g.ApplyBasic(g.GenerateModel("smartlink"))
	g.Execute()

}
