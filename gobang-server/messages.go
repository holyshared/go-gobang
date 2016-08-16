package server

import (
  "github.com/holyshared/learn-golang/gobang"
  "encoding/json"
)

const (
  GameStart = "start"
)

type Message struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}

func GameStartMessage(gobang *gobang.Gobang) []byte {
  message := &Message {
    Type: GameStart,
    Body: gobang,
  }
  res, _ := json.Marshal(message)
  return res
}
