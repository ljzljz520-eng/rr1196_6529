package ingest

import "cellworld/model"

func NormalizeRequest(actor, target string, energy int) (string, string, int) {
	return model.NormalizeName(actor), model.NormalizeName(target), energy
}
func IsAffordable(p model.Profile, cost int) bool { return p.CanConsume(cost) }
func RegistrationSummary(r model.Record) string {
	if r.Status == "pending" {
		return "awaiting processing"
	}
	return r.Status
}
