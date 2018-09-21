package engine

import (
	"testing"
)

func TestPosition(t *testing.T) {
	p := newPosition(1, 2)

	exp := "1-2"
	if p.String() != exp {
		t.Errorf("unexpected value, exp [%v] got [%v]", exp, p.String())
	}
}
