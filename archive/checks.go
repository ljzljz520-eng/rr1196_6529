package archive

import "cellworld/model"

func Archivable(r model.Record) bool { return r.Status == "consumed" }
func ArchiveLabel(r model.Record) string {
	if r.Status == "archived" {
		return "stored"
	}
	return "ready"
}
func ValidateBatch(records []model.Record) int {
	n := 0
	for _, r := range records {
		if Archivable(r) {
			n++
		}
	}
	return n
}
