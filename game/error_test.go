package game

import (
  "testing"
  "encoding/json"
)

func TestCellNotFoundErrorToJSON(t *testing.T) {
  error := NewCellNotFoundError(NewPoint(1, 1))
  bytes, err := json.Marshal(error)

  if err != nil {
    t.Errorf("error: %v", err)
  }

  actual := string(bytes)
  expected := "{\"message\":\"You have specified a not exist cell\",\"point\":{\"x\":1,\"y\":1}}"

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}

func TestAlreadyPlacedErrorToJSON(t *testing.T) {
  error := NewAlreadyPlacedError(NewPoint(1, 1))
  bytes, err := json.Marshal(error)

  if err != nil {
    t.Errorf("error: %v", err)
  }

  actual := string(bytes)
  expected := "{\"message\":\"Already the stone is placed\",\"point\":{\"x\":1,\"y\":1}}"

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}
