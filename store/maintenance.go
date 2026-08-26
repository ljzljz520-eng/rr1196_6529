package store

import (
	"cellworld/model"
	"errors"
	"go.etcd.io/bbolt"
)

func (s *Store) DeleteRecord(id string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(buckets[0]).Delete([]byte(id)) })
}
func (s *Store) ReplaceRecord(r model.Record) error {
	if _, e := s.GetRecord(r.ID); e != nil {
		return e
	}
	return s.SaveRecord(r)
}
func (s *Store) UpdateStatus(id, status string) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	if !model.AllowedTransition(r.Status, status) {
		return errors.New("invalid transition")
	}
	r.Status = status
	return s.SaveRecord(r)
}
