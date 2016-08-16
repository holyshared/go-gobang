package server

import (
  "github.com/olahol/melody"
  "github.com/holyshared/learn-golang/gobang"
)

func NewGameContainer() *GameContainer {
  return &GameContainer {
    sessions: map[*melody.Session]*gobang.Gobang{},
  }
}

type GameContainer struct {
  sessions map[*melody.Session]*gobang.Gobang
}

func (c *GameContainer) Register(s *melody.Session, game *gobang.Gobang) {
  c.sessions[s] = game
}

func (c *GameContainer) Unregister(s *melody.Session) {
  delete(c.sessions, s)
}
