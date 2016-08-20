package server

import (
  "github.com/holyshared/learn-golang/gobang"
  "encoding/json"
)

const (
  GameStart = "start"
  SelectCell = "selectCell"
  NexTurn = "nexTurn"
  GameEnd = "finish"
)

type Message struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}

type UnkownMessage struct {
  *Message
}

func (msg *UnkownMessage) Error() string {
  return msg.Type
}

type GameStartMessage struct {
  Type string `json:"type"`
  Body interface{} `json:"body"`
}

type SelectCellMessage struct {
  Type string `json:"type"`
  Body *gobang.Point `json:"body"`
}

type GameEndMessage struct {
  Type string `json:"type"`
  Body *GameResult `json:"body"`
}

type CurrentGame struct {
  Game *gobang.Gobang `json:"game"`
}

type GameResult struct {
  Game *gobang.Gobang `json:"game"`
  Result gobang.GameProgressResult `json:"result"`
}

func UnmarshalMessage(msg []byte) (interface{}, error) {
  msgType, err := UnmarshalMessageType(msg)

  if err != nil {
    return nil, err
  }

  if msgType.Type == GameStart {
    s := new(GameStartMessage)
    err := json.Unmarshal(msg, s)
    return s, err
  } else if msgType.Type == SelectCell {
    s := new(SelectCellMessage)
    err := json.Unmarshal(msg, s)
    return s, err
  }

  return nil, &UnkownMessage {
    Message: msgType,
  }
}

func UnmarshalMessageType(msg []byte) (*Message, error) {
  message := new(Message)
  err := json.Unmarshal(msg, message)

  if err != nil {
    return nil, err
  }

  return message, nil
}

func SendGameStartMessage(gobang *gobang.Gobang) []byte {
  message := &Message {
    Type: GameStart,
    Body: &CurrentGame {
      Game: gobang,
    },
  }
  res, _ := json.Marshal(message)
  return res
}

func SendNextTurnMessage(gobang *gobang.Gobang) []byte {
  message := &Message {
    Type: NexTurn,
    Body: &CurrentGame {
      Game: gobang,
    },
  }
  res, _ := json.Marshal(message)
  return res
}

func SendGameEndMessage(result gobang.GameProgressResult, gobang *gobang.Gobang) []byte {
  message := &GameEndMessage {
    Type: GameEnd,
    Body: &GameResult {
      Game: gobang,
      Result: result,
    },
  }
  res, _ := json.Marshal(message)
  return res
}
