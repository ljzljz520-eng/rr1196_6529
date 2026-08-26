package query

import (
	"cellworld/model"
	"cellworld/store"
	"path/filepath"
	"testing"
)

func TestQueries(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	r, _ := model.NewRecord("r", "a", "b", 1)
	s.SaveRecord(r)
	if _, e := New(s).Record("r"); e != nil {
		t.Fatal(e)
	}
}
