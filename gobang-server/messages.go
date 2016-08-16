package server

const (
  GameStart = "start"
)

type Message struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}
