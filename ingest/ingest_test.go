package ingest

import (
	"cellworld/store"
	"path/filepath"
	"testing"
)

func TestRegister(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r, e := New(s).RegisterValidated("r", "a", "b", 4)
	if e != nil || r.Status != "pending" {
		t.Fatalf("%v", e)
	}
}
