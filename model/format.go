package model

import "fmt"

func Describe(r Record) string       { return fmt.Sprintf("%s:%s:%d", r.ID, r.Status, r.Score) }
func CloneProfile(p Profile) Profile { return p }
func SameRecord(a, b Record) bool    { return a.ID == b.ID && a.Status == b.Status && a.Score == b.Score }
