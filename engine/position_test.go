package engine

import (
	"testing"
)

func TestPosition(t *testing.T) {
	a := 1
	b := 2
	p := newPosition(a, b)

	exp := "1-2"
	if p.String() != exp {
		t.Errorf("unexpected value, exp [%v] got [%v]", exp, p.String())
	}
	if p.X() != a {
		t.Errorf("unexpected value, exp [%v] got [%v]", a, p.X())
	}
	if p.Y() != b {
		t.Errorf("unexpected value, exp [%v] got [%v]", b, p.Y())
	}
}
