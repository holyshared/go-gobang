package server

import (
  "fmt"
  "encoding/json"
  "github.com/holyshared/go-gobang/gobang"
)

const (
  GameStart = "start"
  SelectCell = "selectCell"
  NexTurn = "nextTurn"
  GameEnd = "finish"
  PutFailed = "putFailed"
)

type OutgoingMessage struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}

type ReceiveMessage struct {
  Type string `json:"type"`
  Body json.RawMessage `json:"body"`
}

type UnkownMessageError struct {
  *ReceiveMessage
}

func (msg *UnkownMessageError) Error() string {
  return fmt.Sprintf("'%s' is not a valid message type", msg.Type)
}

type GameStartMessage struct {
  Body interface{} `json:"body"`
}

type SelectCellMessage struct {
  gobang.Point2D
}

type CurrentGame struct {
  Game *gobang.Gobang `json:"game"`
}

type GameResult struct {
  Game *gobang.Gobang `json:"game"`
  Result gobang.GameProgressResult `json:"result"`
}

func DecodeMessage(data []byte) (interface{}, error) {
  message := &ReceiveMessage{}
  err := json.Unmarshal(data, message)

  if err != nil {
    return nil, err
  }

  return message.decodeBody()
}

func (m *ReceiveMessage) decodeBody() (interface{}, error) {
  switch m.Type {
  default:
    return nil, &UnkownMessageError { ReceiveMessage: m }
  case GameStart:
    body := &GameStartMessage{}
    err := json.Unmarshal(m.Body, body)
    return body, err
  case SelectCell:
    body := &SelectCellMessage{ Point2D: gobang.DefaultPoint() }
    err := json.Unmarshal(m.Body, body.Point2D)
    return body, err
  }
}

func SendGameStartMessage(gobang *gobang.Gobang) []byte {
  message := &OutgoingMessage {
    Type: GameStart,
    Body: &CurrentGame {
      Game: gobang,
    },
  }
  res, _ := json.Marshal(message)
  return res
}

func SendNextTurnMessage(gobang *gobang.Gobang) []byte {
  message := &OutgoingMessage {
    Type: NexTurn,
    Body: &CurrentGame {
      Game: gobang,
    },
  }
  res, _ := json.Marshal(message)
  return res
}

func SendGameEndMessage(result gobang.GameProgressResult, gobang *gobang.Gobang) []byte {
  message := &OutgoingMessage {
    Type: GameEnd,
    Body: &GameResult {
      Game: gobang,
      Result: result,
    },
  }
  res, _ := json.Marshal(message)
  return res
}

func SendPutFailedMessage(reason error) []byte {
  message := &OutgoingMessage {
    Type: PutFailed,
    Body: reason,
  }
  res, _ := json.Marshal(message)
  return res
}
