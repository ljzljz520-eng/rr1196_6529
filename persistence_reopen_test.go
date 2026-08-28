package main

import (
	"cellworld/model"
	"cellworld/store"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "db")
	s, e := store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	r, _ := model.NewRecord("r", "a", "b", 1)
	s.SaveRecord(r)
	s.Close()
	s, e = store.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetRecord("r"); e != nil {
		t.Fatal(e)
	}
}
