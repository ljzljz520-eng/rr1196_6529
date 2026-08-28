package main

import (
	"cellworld/api"
	"cellworld/archive"
	"cellworld/ingest"
	"cellworld/process"
	"cellworld/query"
	"cellworld/store"
	"log"
	"net/http"
)

func main() {
	s, e := store.Open("cellworld.db")
	if e != nil {
		log.Fatal(e)
	}
	defer s.Close()
	i := ingest.New(s)
	p := process.New(s)
	q := query.New(s)
	a := archive.New(s)
	log.Println(http.ListenAndServe(":8080", api.New(i, p, q, a).Handler()))
}
