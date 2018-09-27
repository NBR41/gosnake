package engine

import (
	"testing"
)

func TestSegment(t *testing.T) {
	s := newSegment(East, NewPosition(2, 0), NewPosition(0, 0))
	exp := true
	if exp != s.IsHorizontal() {
		t.Errorf("unexpected value, exp [%v] got [%v]", exp, s.IsHorizontal())
	}

	s = newSegment(South, NewPosition(0, 0), NewPosition(0, 2))
	exp = false
	if exp != s.IsHorizontal() {
		t.Errorf("unexpected value, exp [%v] got [%v]", exp, s.IsHorizontal())
	}
}
