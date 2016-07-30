package gobang

import (
  "testing"
  "encoding/json"
)

func TestPointToJSON(t *testing.T) {
  point := Point { X: 1, Y: 1 }
  bytes, err := json.Marshal(point)

  if err != nil {
    t.Errorf("error: %v", err)
  }

  actual := string(bytes)
  expected := "{\"x\":1,\"y\":1}"

  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}
