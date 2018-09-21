package engine

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestGetFreePosition1(t *testing.T) {
	g := newGrid(3, 3)
	p, err := g.GetFreePosition(
		[]*Segment{
			newSegment(East, newPosition(0, 1), newPosition(2, 1)),
			newSegment(West, newPosition(2, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 2), newPosition(2, 2)),
		},
	)
	if err != nil {
		t.Errorf("unexpected error, [%v]", err)
	} else {
		exp := newPosition(1, 1)
		if !reflect.DeepEqual(exp, p) {
			t.Errorf("unexpected value, exp [%+v] got [%+v]", *exp, *p)
		}
	}
}

func TestGetFreePosition2(t *testing.T) {
	g := newGrid(3, 3)
	p, err := g.GetFreePosition(
		[]*Segment{
			newSegment(South, newPosition(1, 0), newPosition(1, 2)),
			newSegment(North, newPosition(2, 2), newPosition(2, 1)),
			newSegment(South, newPosition(0, 1), newPosition(0, 2)),
		},
	)
	if err != nil {
		t.Errorf("unexpected error, [%v]", err)
	} else {
		exp := newPosition(1, 1)
		if !reflect.DeepEqual(exp, p) {
			t.Errorf("unexpected value, exp [%+v] got [%+v]", *exp, *p)
		}
	}
}

func TestGetFreePosition3(t *testing.T) {
	g := newGrid(2, 2)
	_, err := g.GetFreePosition(
		[]*Segment{
			newSegment(West, newPosition(0, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 1), newPosition(0, 1)),
		},
	)
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != ErrNoPosition {
			t.Errorf("unexpected error, exp [%v] got [%v]", ErrNoPosition, err)
		}
	}
}

func TestIsFreePosition(t *testing.T) {
	g := newGrid(2, 2)
	v := g.IsFreePosition(
		newPosition(0, 0),
		[]*Segment{
			newSegment(West, newPosition(0, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 1), newPosition(0, 1)),
		},
	)
	if v {
		t.Errorf("unexpected value")
	}
	v = g.IsFreePosition(
		newPosition(0, 0),
		[]*Segment{
			newSegment(West, newPosition(1, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 1), newPosition(0, 1)),
		},
	)
	if !v {
		t.Errorf("unexpected value")
	}
	v = g.IsFreePosition(
		newPosition(0, 1),
		[]*Segment{
			newSegment(West, newPosition(1, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 1), newPosition(0, 1)),
		},
	)
	if !v {
		t.Errorf("unexpected value")
	}
}

func TestMove1(t *testing.T) {
	g := newGrid(3, 3)
	p, err := g.Move(
		East,
		[]*Segment{
			newSegment(East, newPosition(0, 1), newPosition(2, 1)),
			newSegment(West, newPosition(2, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 2), newPosition(2, 2)),
		},
	)
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		exp := []*Segment{
			newSegment(East, newPosition(1, 1), newPosition(2, 1)),
			newSegment(West, newPosition(2, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 2), newPosition(0, 2)),
		}
		if diff := pretty.Compare(exp, p); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMove2(t *testing.T) {
	g := newGrid(3, 3)
	_, err := g.Move(
		North,
		[]*Segment{
			newSegment(East, newPosition(0, 1), newPosition(2, 1)),
			newSegment(West, newPosition(2, 0), newPosition(1, 0)),
			newSegment(East, newPosition(1, 2), newPosition(2, 2)),
		},
	)
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != ErrColision {
			t.Errorf("unexpected error, exp [%v] got [%v]", ErrColision, err)
		}
	}
}

func TestMove3(t *testing.T) {
	g := newGrid(2, 2)
	p, err := g.Move(
		South,
		[]*Segment{
			newSegment(West, newPosition(0, 0), newPosition(1, 0)),
			newSegment(North, newPosition(1, 1), newPosition(1, 1)),
		},
	)
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		exp := []*Segment{
			newSegment(South, newPosition(0, 1), newPosition(0, 1)),
			newSegment(West, newPosition(0, 0), newPosition(1, 0)),
		}
		if diff := pretty.Compare(exp, p); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestGetPosition(t *testing.T) {
	tests := []struct {
		p   string
		exp *Position
		err error
	}{
		{"A", nil, ErrInvalidKey},
		{"A-A", nil, errors.New(`strconv.Atoi: parsing "A": invalid syntax`)},
		{"0-A", nil, errors.New(`strconv.Atoi: parsing "A": invalid syntax`)},
		{"0-0", newPosition(0, 0), nil},
	}

	for i := range tests {
		v, err := getPosition(tests[i].p)
		if err != nil {
			if tests[i].err == nil {
				t.Errorf("unexpected error at %d, [%v]", i, err)
			} else {
				if err.Error() != tests[i].err.Error() {
					t.Errorf("unexpected error at %d, exp [%v] got [%v]", i, tests[i].err, err)
				}
			}
		} else {
			if tests[i].err != nil {
				t.Errorf("expecting error at %d", i)
			}
		}

		if !reflect.DeepEqual(v, tests[i].exp) {
			t.Errorf("unexpected value at %d, exp [%v] got [%v]", i, tests[i].exp, v)
		}
	}
}
