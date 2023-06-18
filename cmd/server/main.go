package main

import (
	"fmt"

	"github.com/bagusyanuar/go-olin-bags/app/config"
	"github.com/bagusyanuar/go-olin-bags/app/http/builder"
	"github.com/bagusyanuar/go-olin-bags/common"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	routers := builder.BuildRoute()
	common.BuildRoute(r, routers)
	r.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}
