package archive

import (
	"cellworld/ingest"
	"cellworld/process"
	"cellworld/store"
	"path/filepath"
	"testing"
)

func TestArchive(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	ingest.New(s).Register("r", "a", "b", 1)
	process.New(s).Consume("r")
	if _, e := New(s).Archive("r"); e != nil {
		t.Fatal(e)
	}
}
