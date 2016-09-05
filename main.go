package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	gobang "github.com/holyshared/go-gobang/gobang-server"
)

func main() {
	server := gin.Default()
	gobang := gobang.NewApp()

	server.Use(static.Serve("/", static.LocalFile("./public", true)))
	server.GET("/ws", gobang.Handler)
	server.Run(":3000")
}
