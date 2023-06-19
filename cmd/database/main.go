package main

import (
	"flag"
	"fmt"

	"github.com/bagusyanuar/go-olin-bags/app/config"
	"github.com/bagusyanuar/go-olin-bags/cmd/database/migrations"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}
	db, err := config.NewMySQLConnection(&cfg.MySQL)
	if err != nil {
		panic(err)
	}
	cmd := flag.String("m", "", "unsupport command")
	flag.Parse()
	command := *cmd
	switch command {
	// case "seed":
	// 	migrations.Seed(database)
	// 	return
	case "migrate":
		migrations.Migrate(db)
		return
	default:
		fmt.Println("unknown command")
		return
	}

}
