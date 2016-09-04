package gobang

import (
	"encoding/json"
	"testing"
)

func TestPointToJSON(t *testing.T) {
	point := NewPoint(1, 1)
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

func TestJSONToPoint(t *testing.T) {
	point := DefaultPoint()
	err := json.Unmarshal([]byte("{\"x\":1,\"y\":1}"), &point)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if point.X() != 1 {
		t.Errorf("got %v\nwant %v", point.X(), 1)
	}

	if point.Y() != 1 {
		t.Errorf("got %v\nwant %v", point.Y(), 1)
	}
}
