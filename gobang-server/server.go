package server

import (
  "github.com/gin-gonic/gin"
  "github.com/olahol/melody"
  "github.com/Sirupsen/logrus"
  "github.com/holyshared/go-gobang/gobang"
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
  app.Unregister(s)
}

func (app *App) OnMessage(s *melody.Session, msg []byte) {
  app.Info("recived message")
  message, err := DecodeMessage(msg)

  if err != nil {
    app.Warnf("recived message error: %v", err)
    return
  }

  switch recivedMessage := message.(type) {
  default:
    app.Warnf("recived message is not support: %v", recivedMessage)
  case *GameStartMessage:
    app.startGame(s, recivedMessage)
  case *SelectCellMessage:
    app.selectCell(s, recivedMessage)
  }
}

func (app *App) startGame(s *melody.Session, message *GameStartMessage) {
  app.Infof("game started: ", message)
  game := gobang.NewGobang(
    gobang.DefaultGameRule(),
    gobang.Black,
    gobang.White,
  )
  app.Register(s, game)
  s.Write(SendGameStartMessage(game))
}

func (app *App) selectCell(s *melody.Session, message *SelectCellMessage) {
  var err error
  var result gobang.GameProgressResult

  app.Infof("player cell selected: ", message)

  game := app.Lookup(s)
  result, err = game.PlayerPutStoneTo(message)

  if err != nil {
    app.Warnf("put stone faild: ", err)
    s.Write(SendPutFailedMessage(err))
    return
  }

  app.Infof("player selected done")

  if result == gobang.Win || result == gobang.Draw {
    app.Unregister(s)
    s.Write(SendGameEndMessage(result, game))
    return
  }

  result, err = game.NpcPlayerPutStone()

  if err != nil {
    app.Fatalf("put stone faild: ", err)
    return
  }

  if result == gobang.Lose || result == gobang.Draw {
    app.Unregister(s)
    s.Write(SendGameEndMessage(result, game))
    return
  }
  app.Infof("player turn")

  s.Write(SendNextTurnMessage(game))
}
