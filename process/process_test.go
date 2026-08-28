package process

import (
	"cellworld/ingest"
	"cellworld/store"
	"path/filepath"
	"testing"
)

func TestConsume(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	ingest.New(s).Register("r", "a", "b", 4)
	r, e := New(s).Consume("r")
	if e != nil || r.Status != "consumed" {
		t.Fatal(e)
	}
}
