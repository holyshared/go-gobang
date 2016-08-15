package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/contrib/static"
  gobang "github.com/holyshared/learn-golang/gobang-server"
)

func main() {
  server := gin.Default()
  gobang := gobang.NewApp()

  server.Use(static.Serve("/", static.LocalFile("./public", true)))
  server.GET("/ws", gobang.Handler)
  server.Run(":3000")
}
