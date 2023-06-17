package main

import (
	"fmt"

	"github.com/bagusyanuar/go-olin-bags/app/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.NewConfig(".env")
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})
	r.Run(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
}
