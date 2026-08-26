package store

import (
	"cellworld/model"
	"encoding/json"
)

func decodeRecord(v []byte) (model.Record, error) {
	var r model.Record
	e := json.Unmarshal(v, &r)
	return r, e
}
func (s *Store) HasRecord(id string) bool { _, e := s.GetRecord(id); return e == nil }
func (s *Store) SaveAll(records []model.Record) error {
	for _, r := range records {
		if e := s.SaveRecord(r); e != nil {
			return e
		}
	}
	return nil
}
