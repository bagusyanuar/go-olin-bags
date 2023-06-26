package main

import (
	"flag"
	"fmt"

	"github.com/bagusyanuar/go-olin-bags/cmd/database/migrations"
	"github.com/bagusyanuar/go-olin-bags/cmd/database/seeder"
	"github.com/bagusyanuar/go-olin-bags/config"
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
		fmt.Println("successfully migrating database")
		return
	case "fresh":
		migrations.Drop(db)
		migrations.Migrate(db)
		fmt.Println("successfully fresh database")
	case "seed":
		seeder.Seed(db)
		fmt.Println("successfully seed database")
	default:
		fmt.Println("unknown command")
		return
	}

}
