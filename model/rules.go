package model

import "time"

type Snapshot struct {
	Record  Record
	Profile Profile
	Event   Event
	Audit   Audit
}

func (s Snapshot) Ready() bool                     { return s.Record.Valid() && s.Profile.ID != "" }
func (s Snapshot) Age(now time.Time) time.Duration { return now.Sub(s.Record.CreatedAt) }
func (s Snapshot) ScoreDelta() int                 { return s.Record.Score }
func (s Snapshot) HasEvent(kind string) bool       { return s.Event.Kind == kind }
func (s Snapshot) HasAudit(action string) bool     { return s.Audit.Action == action }
func (s Snapshot) Terminal() bool                  { return IsTerminalStatus(s.Record.Status) }
func (s Snapshot) ActorMatches() bool              { return s.Record.ActorID == s.Profile.ID }
func (s Snapshot) TargetDistinct() bool            { return s.Record.ActorID != s.Record.TargetID }
func (s Snapshot) Affordable() bool                { return s.Profile.Energy >= s.Record.Energy }
func (s Snapshot) ValidLifecycle() bool {
	return AllowedTransition("pending", s.Record.Status) || s.Record.Status == "pending"
}
func (s Snapshot) Summary() string { return Describe(s.Record) }
