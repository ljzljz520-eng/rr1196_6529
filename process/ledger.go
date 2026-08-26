package process

import (
	"cellworld/model"
	"fmt"
)

func ScoreForEnergy(energy int) int {
	if energy < 1 {
		return 0
	}
	return energy * 2
}
func TransitionLabel(r model.Record) string {
	if r.Status == "consumed" {
		return fmt.Sprintf("%s consumed", r.ID)
	}
	return fmt.Sprintf("%s %s", r.ID, r.Status)
}
func VerifyTransition(from, to string) error {
	if !model.AllowedTransition(from, to) {
		return fmt.Errorf("transition %s to %s denied", from, to)
	}
	return nil
}
