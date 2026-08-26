package archive

import (
	"cellworld/model"
	"cellworld/store"
	"errors"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Archive(id string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if !r.IsConsumed() {
		return r, errors.New("only consumed records archive")
	}
	r.Status = "archived"
	if e = s.Store.SaveRecord(r); e != nil {
		return r, e
	}
	return r, s.Store.SaveAudit(model.AuditFor(r, "archive", "stored"))
}
func (s *Service) IsArchived(id string) bool {
	r, e := s.Store.GetRecord(id)
	return e == nil && r.Status == "archived"
}
func (s *Service) Restore(id string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.Status != "archived" {
		return r, errors.New("not archived")
	}
	r.Status = "consumed"
	e = s.Store.SaveRecord(r)
	return r, e
}
func (s *Service) ArchiveBatch(ids []string) int {
	n := 0
	for _, id := range ids {
		if _, e := s.Archive(id); e == nil {
			n++
		}
	}
	return n
}
