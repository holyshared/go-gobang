package server

import (
  "github.com/gin-gonic/gin"
  "github.com/olahol/melody"
  "github.com/Sirupsen/logrus"
  "github.com/holyshared/learn-golang/gobang"
  "encoding/json"
)

func NewApp() *App {
  app := &App {
    Melody: melody.New(),
    Logger: logrus.New(),
    GameContainer: NewGameContainer(),
  }
  app.Start()

  return app
}

type App struct {
  *melody.Melody
  *logrus.Logger
  *GameContainer
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
  var message Message

  app.Info("recived message")

  err := json.Unmarshal(msg, &message)

  if err != nil {
    app.Warnf("recived message error: %v", err)
    return
  }

  if (message.Type == GameStart) {
    game := gobang.NewGobang(
      gobang.DefaultGameRule(),
      gobang.Black,
      gobang.White,
    )
    app.Register(s, game)
    s.Write(GameStartMessage(game))
  }
}
