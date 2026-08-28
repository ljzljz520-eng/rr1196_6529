package main

import (
	"cellworld/archive"
	"cellworld/ingest"
	"cellworld/process"
	"cellworld/query"
	"cellworld/store"
	"path/filepath"
	"testing"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	if _, e := ingest.New(s).RegisterValidated("r", "a", "b", 2); e != nil {
		t.Fatal(e)
	}
	if _, e := query.New(s).Record("r"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowTwo(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	ingest.New(s).Register("r", "a", "b", 2)
	process.New(s).Consume("r")
	if _, e := archive.New(s).Archive("r"); e != nil {
		t.Fatal(e)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	if _, e := ingest.New(s).Register("r", "a", "b", 2); e != nil {
		t.Fatal(e)
	}
}
func TestBusinessChain10(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	ingest.New(s).Register("r", "a", "b", 100)
	p, _ := s.GetProfile("a")
	p.Energy = 1
	s.SaveProfile(p)
	r, e := process.New(s).Consume("r")
	if e == nil && r.Status == "consumed" {
		t.Fatal("insufficient energy must preserve target and score")
	}
}
