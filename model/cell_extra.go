package model

import "strings"

func NormalizeName(name string) string { return strings.TrimSpace(strings.ToLower(name)) }
func IsTerminalStatus(status string) bool {
	switch status {
	case "consumed", "rejected", "archived":
		return true
	default:
		return false
	}
}
func AllowedTransition(from, to string) bool {
	if from == "pending" && (to == "consumed" || to == "rejected") {
		return true
	}
	if from == "consumed" && to == "archived" {
		return true
	}
	return false
}
func DefaultProfile(id, name string) Profile {
	return Profile{ID: id, Name: NormalizeName(name), Energy: 100, Active: true, Rank: 1}
}
func EventFor(r Record, kind, payload string) Event {
	return Event{ID: r.ID + ":" + kind, RecordID: r.ID, Kind: kind, Payload: payload, At: r.UpdatedAt}
}
func AuditFor(r Record, action, detail string) Audit {
	return Audit{ID: r.ID + ":" + action, RecordID: r.ID, Action: action, Detail: detail, At: r.UpdatedAt}
}
