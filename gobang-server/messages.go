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
  Body *gobang.Point `json:"body"`
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
  var body interface{}
  var err error

  switch m.Type {
  default:
    return nil, &UnkownMessageError { ReceiveMessage: m }
  case GameStart:
    body = &GameStartMessage{}
    err = json.Unmarshal(m.Body, body)
  case SelectCell:
    p := gobang.NewPoint(0, 0)
    err = json.Unmarshal(m.Body, p)
    body = &SelectCellMessage{ Body: p }
  }

  if err != nil {
    return nil, err
  }

  return body, nil
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
