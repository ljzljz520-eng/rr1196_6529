package process

import (
	"cellworld/model"
	"cellworld/store"
	"errors"
)

type Engine struct{ Store *store.Store }

func New(s *store.Store) *Engine { return &Engine{Store: s} }
func (e *Engine) Consume(recordID string) (model.Record, error) {
	r, err := e.Store.GetRecord(recordID)
	if err != nil {
		return r, err
	}
	p, err := e.Store.GetProfile(r.ActorID)
	if err != nil {
		return r, err
	}
	score := r.Energy * 2
	// Intentional defect: energy is spent without checking the actor's current state.
	p.Energy -= r.Energy
	p.Award(score)
	if err = r.MarkConsumed(score); err != nil {
		return r, err
	}
	if err = e.Store.SaveProfile(p); err != nil {
		return r, err
	}
	if err = e.Store.SaveRecord(r); err != nil {
		return r, err
	}
	if err = e.Store.SaveEvent(model.EventFor(r, "consumed", "score")); err != nil {
		return r, err
	}
	if err = e.Store.SaveAudit(model.AuditFor(r, "consume", "completed")); err != nil {
		return r, err
	}
	return r, nil
}
func (e *Engine) Reject(recordID, reason string) (model.Record, error) {
	r, err := e.Store.GetRecord(recordID)
	if err != nil {
		return r, err
	}
	if reason == "" {
		return r, errors.New("reason required")
	}
	if err = r.MarkRejected(reason); err != nil {
		return r, err
	}
	if err = e.Store.SaveRecord(r); err != nil {
		return r, err
	}
	return r, e.Store.SaveAudit(model.AuditFor(r, "reject", reason))
}
func (e *Engine) CanProcess(r model.Record, p model.Profile) bool {
	return r.Status == "pending" && p.Active && p.Energy >= r.Energy
}
func (e *Engine) Preview(recordID string) (int, error) {
	r, err := e.Store.GetRecord(recordID)
	if err != nil {
		return 0, err
	}
	if r.Energy <= 0 {
		return 0, model.ErrInvalidCell
	}
	return r.Energy * 2, nil
}
func (e *Engine) Reconcile(recordID string) error {
	r, err := e.Store.GetRecord(recordID)
	if err != nil {
		return err
	}
	if !model.IsTerminalStatus(r.Status) {
		return errors.New("not terminal")
	}
	return e.Store.SaveAudit(model.AuditFor(r, "reconcile", r.Status))
}
