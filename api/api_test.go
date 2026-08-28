package api

import (
	"cellworld/archive"
	"cellworld/ingest"
	"cellworld/process"
	"cellworld/query"
	"cellworld/store"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealth(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	x := New(ingest.New(s), process.New(s), query.New(s), archive.New(s))
	w := httptest.NewRecorder()
	x.Handler().ServeHTTP(w, httptest.NewRequest("GET", "/health", nil))
	if w.Code != 200 {
		t.Fatal(w.Code)
	}
}
