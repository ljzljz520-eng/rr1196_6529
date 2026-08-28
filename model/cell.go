package model

import (
	"errors"
	"time"
)

var ErrInvalidCell = errors.New("invalid cell")
var ErrInsufficientEnergy = errors.New("insufficient energy")
var ErrAlreadyConsumed = errors.New("cell already consumed")

type Record struct {
	ID        string    `json:"id"`
	TargetID  string    `json:"target_id"`
	ActorID   string    `json:"actor_id"`
	Energy    int       `json:"energy"`
	Score     int       `json:"score"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Profile struct {
	ID     string
	Name   string
	Energy int
	Score  int
	Rank   int
	Active bool
}
type Event struct {
	ID       string
	RecordID string
	Kind     string
	Payload  string
	At       time.Time
}
type Audit struct {
	ID       string
	RecordID string
	Action   string
	Detail   string
	At       time.Time
}

func NewRecord(id, actor, target string, energy int) (Record, error) {
	if id == "" || actor == "" || target == "" || energy < 0 {
		return Record{}, ErrInvalidCell
	}
	now := time.Now().UTC()
	return Record{ID: id, ActorID: actor, TargetID: target, Energy: energy, Status: "pending", CreatedAt: now, UpdatedAt: now}, nil
}
func (r Record) Valid() bool {
	return r.ID != "" && r.ActorID != "" && r.TargetID != "" && r.Energy >= 0
}
func (r Record) IsConsumed() bool { return r.Status == "consumed" }
func (r *Record) MarkConsumed(score int) error {
	if !r.Valid() {
		return ErrInvalidCell
	}
	if r.IsConsumed() {
		return ErrAlreadyConsumed
	}
	r.Status = "consumed"
	r.Score = score
	r.UpdatedAt = time.Now().UTC()
	return nil
}
func (r *Record) MarkRejected(reason string) error {
	if reason == "" || !r.Valid() {
		return ErrInvalidCell
	}
	r.Status = "rejected"
	r.UpdatedAt = time.Now().UTC()
	return nil
}
func (p Profile) CanConsume(cost int) bool { return p.Active && cost > 0 && p.Energy >= cost }
func (p *Profile) Spend(cost int) error {
	if !p.CanConsume(cost) {
		return ErrInsufficientEnergy
	}
	p.Energy -= cost
	return nil
}
func (p *Profile) Award(points int) {
	if points > 0 {
		p.Score += points
	}
}
func (p *Profile) RecomputeRank(total int) {
	if total < 0 {
		total = 0
	}
	p.Rank = 1 + total/100
}
func (e Event) Valid() bool { return e.ID != "" && e.RecordID != "" && e.Kind != "" }
func (a Audit) Valid() bool { return a.ID != "" && a.RecordID != "" && a.Action != "" }
