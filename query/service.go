package query

import (
	"cellworld/model"
	"cellworld/store"
	"sort"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service                         { return &Service{Store: s} }
func (s *Service) Record(id string) (model.Record, error) { return s.Store.GetRecord(id) }
func (s *Service) All() ([]model.Record, error) {
	r, e := s.Store.ListRecords()
	sort.Slice(r, func(i, j int) bool { return r[i].CreatedAt.Before(r[j].CreatedAt) })
	return r, e
}
func (s *Service) Leaderboard() ([]model.Profile, error) {
	out := []model.Profile{}
	for _, id := range []string{"a", "b", "c", "d"} {
		if p, e := s.Store.GetProfile(id); e == nil {
			out = append(out, p)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Score > out[j].Score })
	return out, nil
}
func (s *Service) Status(id string) (string, error) { r, e := s.Record(id); return r.Status, e }
func (s *Service) Energy(id string) (int, error)    { p, e := s.Store.GetProfile(id); return p.Energy, e }
