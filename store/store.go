package store

import (
	"cellworld/model"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
	"time"
)

var ErrNotFound = errors.New("not found")
var buckets = [][]byte{[]byte("records"), []byte("profiles"), []byte("events"), []byte("audits")}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	if path == "" {
		return nil, errors.New("path required")
	}
	db, e := bbolt.Open(filepath.Clean(path), 0600, &bbolt.Options{Timeout: time.Second})
	if e != nil {
		return nil, e
	}
	s := &Store{db: db, path: path}
	e = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, x := tx.CreateBucketIfNotExists(b); x != nil {
				return x
			}
		}
		return nil
	})
	if e != nil {
		db.Close()
		return nil, e
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	e := s.db.Close()
	s.db = nil
	return e
}
func (s *Store) Path() string { return s.path }
func put[T any](s *Store, b []byte, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(key), data) })
}
func get[T any](s *Store, b []byte, key string, out *T) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return errors.New("closed")
	}
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(b).Get([]byte(key))
		if v == nil {
			return ErrNotFound
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) SaveRecord(v model.Record) error {
	if !v.Valid() {
		return model.ErrInvalidCell
	}
	return put(s, buckets[0], v.ID, v)
}
func (s *Store) GetRecord(id string) (model.Record, error) {
	var v model.Record
	e := get(s, buckets[0], id, &v)
	return v, e
}
func (s *Store) SaveProfile(v model.Profile) error {
	if v.ID == "" {
		return model.ErrInvalidCell
	}
	return put(s, buckets[1], v.ID, v)
}
func (s *Store) GetProfile(id string) (model.Profile, error) {
	var v model.Profile
	e := get(s, buckets[1], id, &v)
	return v, e
}
func (s *Store) SaveEvent(v model.Event) error {
	if !v.Valid() {
		return model.ErrInvalidCell
	}
	return put(s, buckets[2], v.ID, v)
}
func (s *Store) SaveAudit(v model.Audit) error {
	if !v.Valid() {
		return model.ErrInvalidCell
	}
	return put(s, buckets[3], v.ID, v)
}
func (s *Store) ListRecords() ([]model.Record, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return nil, errors.New("closed")
	}
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[0]).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) Count(bucket string) (int, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if s.db == nil {
		return 0, errors.New("closed")
	}
	var n int
	e := s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf("bucket %s", bucket)
		}
		n = b.Stats().KeyN
		return nil
	})
	return n, e
}
