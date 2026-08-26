package api

import (
	"cellworld/archive"
	"cellworld/ingest"
	"cellworld/process"
	"cellworld/query"
	"encoding/json"
	"net/http"
)

type Server struct {
	Ingest  *ingest.Service
	Process *process.Engine
	Query   *query.Service
	Archive *archive.Service
}

func New(i *ingest.Service, p *process.Engine, q *query.Service, a *archive.Service) *Server {
	return &Server{Ingest: i, Process: p, Query: q, Archive: a}
}
func (s *Server) Handler() http.Handler {
	m := http.NewServeMux()
	m.HandleFunc("/health", s.health)
	m.HandleFunc("/records", s.records)
	m.HandleFunc("/consume", s.consume)
	m.HandleFunc("/archive", s.archive)
	return m
}
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func (s *Server) records(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var x struct {
			ID, Actor, Target string
			Energy            int
		}
		if json.NewDecoder(r.Body).Decode(&x) != nil {
			http.Error(w, "bad", 400)
			return
		}
		v, e := s.Ingest.RegisterValidated(x.ID, x.Actor, x.Target, x.Energy)
		if e != nil {
			http.Error(w, e.Error(), 400)
			return
		}
		json.NewEncoder(w).Encode(v)
		return
	}
	id := r.URL.Query().Get("id")
	v, e := s.Query.Record(id)
	if e != nil {
		http.Error(w, e.Error(), 404)
		return
	}
	json.NewEncoder(w).Encode(v)
}
func (s *Server) consume(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	v, e := s.Process.Consume(id)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(v)
}
func (s *Server) archive(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	v, e := s.Archive.Archive(id)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(v)
}
