package ingest

import (
	"cellworld/model"
	"cellworld/store"
	"errors"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Register(id, actor, target string, energy int) (model.Record, error) {
	r, e := model.NewRecord(id, actor, target, energy)
	if e != nil {
		return r, e
	}
	if _, e = s.Store.GetProfile(actor); e != nil {
		if !errors.Is(e, store.ErrNotFound) {
			return r, e
		}
		p := model.DefaultProfile(actor, actor)
		if e = s.Store.SaveProfile(p); e != nil {
			return r, e
		}
	}
	if e = s.Store.SaveRecord(r); e != nil {
		return r, e
	}
	e = s.Store.SaveEvent(model.EventFor(r, "registered", "pending"))
	return r, e
}
func (s *Service) Validate(r model.Record) error {
	if !r.Valid() {
		return model.ErrInvalidCell
	}
	if r.ActorID == r.TargetID {
		return errors.New("self target")
	}
	return nil
}
func (s *Service) RegisterValidated(id, actor, target string, energy int) (model.Record, error) {
	r, e := s.Register(id, actor, target, energy)
	if e != nil {
		return r, e
	}
	if e = s.Validate(r); e != nil {
		_ = r.MarkRejected(e.Error())
		_ = s.Store.SaveRecord(r)
		return r, e
	}
	return r, nil
}
