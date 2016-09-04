package gobang

import (
	"encoding/json"
	"testing"
)

func TestCellToJSON(t *testing.T) {
	cell := NewCell(1, 1, Black)
	bytes, err := json.Marshal(cell)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	actual := string(bytes)
	expected := "{\"point\":{\"x\":1,\"y\":1},\"stone\":1}"

	if actual != expected {
		t.Errorf("got %v\nwant %v", actual, expected)
	}
}
