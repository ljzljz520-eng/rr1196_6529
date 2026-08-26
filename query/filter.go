package query

import (
	"cellworld/model"
	"strings"
)

func FilterByStatus(records []model.Record, status string) []model.Record {
	out := []model.Record{}
	for _, r := range records {
		if strings.EqualFold(r.Status, status) {
			out = append(out, r)
		}
	}
	return out
}
func TotalScore(profiles []model.Profile) int {
	n := 0
	for _, p := range profiles {
		n += p.Score
	}
	return n
}
func Find(records []model.Record, id string) (model.Record, bool) {
	for _, r := range records {
		if r.ID == id {
			return r, true
		}
	}
	return model.Record{}, false
}
