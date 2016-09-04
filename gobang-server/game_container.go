package server

import (
	"github.com/holyshared/go-gobang/gobang"
	"github.com/olahol/melody"
)

func NewGameContainer() *GameContainer {
	return &GameContainer{
		sessions: map[*melody.Session]*gobang.Gobang{},
	}
}

type GameContainer struct {
	sessions map[*melody.Session]*gobang.Gobang
}

func (c *GameContainer) Lookup(s *melody.Session) *gobang.Gobang {
	return c.sessions[s]
}

func (c *GameContainer) Register(s *melody.Session, game *gobang.Gobang) {
	c.sessions[s] = game
}

func (c *GameContainer) Unregister(s *melody.Session) {
	delete(c.sessions, s)
}
