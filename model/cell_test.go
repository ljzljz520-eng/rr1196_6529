package model

import "testing"

func TestRecordLifecycle(t *testing.T) {
	r, e := NewRecord("r", "a", "b", 3)
	if e != nil {
		t.Fatal(e)
	}
	if e = r.MarkConsumed(6); e != nil {
		t.Fatal(e)
	}
	if !r.IsConsumed() {
		t.Fatal("status")
	}
}
