package store

import (
	"cellworld/model"
	"errors"
)

type Transaction struct {
	Store   *Store
	Record  model.Record
	Profile model.Profile
}

func (s *Store) Begin(id string) (Transaction, error) {
	r, e := s.GetRecord(id)
	if e != nil {
		return Transaction{}, e
	}
	p, e := s.GetProfile(r.ActorID)
	if e != nil {
		return Transaction{}, e
	}
	return Transaction{Store: s, Record: r, Profile: p}, nil
}
func (t *Transaction) Commit() error {
	if t.Store == nil {
		return errors.New("no store")
	}
	if e := t.Store.SaveProfile(t.Profile); e != nil {
		return e
	}
	return t.Store.SaveRecord(t.Record)
}
func (t *Transaction) Abort()                     { t.Record = model.Record{}; t.Profile = model.Profile{} }
func (t *Transaction) Valid() bool                { return t.Store != nil && t.Record.Valid() && t.Profile.ID != "" }
func (t *Transaction) ApplyScore(points int)      { t.Profile.Award(points); t.Record.Score = points }
func (t *Transaction) SpendEnergy(cost int) error { return t.Profile.Spend(cost) }
func (t *Transaction) MarkConsumed() error        { return t.Record.MarkConsumed(t.Record.Score) }
