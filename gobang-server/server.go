package server

import (
  "github.com/gin-gonic/gin"
  "github.com/olahol/melody"
  "github.com/Sirupsen/logrus"
)

func NewApp() *App {
  app := &App {
    Melody: melody.New(),
    Logger: logrus.New(),
  }
  app.Start()

  return app
}

type App struct {
  *melody.Melody
  *logrus.Logger
}

func (app *App) Start() {
  app.Melody.HandleConnect(app.OnConnect)
  app.Melody.HandleDisconnect(app.OnDisconnect)
  app.Melody.HandleMessage(app.OnMessage)
}

func (app *App) Handler(c *gin.Context) {
  app.Melody.HandleRequest(c.Writer, c.Request)
}

func (app *App) OnConnect(s *melody.Session) {
  app.Info("connected")
}

func (app *App) OnDisconnect(s *melody.Session) {
  app.Info("disconnected")
}

func (app *App) OnMessage(s *melody.Session, msg []byte) {
  app.Info("recived message")
  app.Melody.Broadcast(msg)
}
