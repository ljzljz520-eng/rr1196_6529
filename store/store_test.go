package store

import (
	"cellworld/model"
	"path/filepath"
	"testing"
)

func TestStoreRoundTrip(t *testing.T) {
	s, e := Open(filepath.Join(t.TempDir(), "db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, _ := model.NewRecord("r", "a", "b", 2)
	if e = s.SaveRecord(r); e != nil {
		t.Fatal(e)
	}
	if _, e = s.GetRecord("r"); e != nil {
		t.Fatal(e)
	}
}
