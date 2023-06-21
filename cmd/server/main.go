package main

import (
	"github.com/bagusyanuar/go-olin-bags/app/server"
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
	server.Serve(cfg, db)
}
