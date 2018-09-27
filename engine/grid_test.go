package engine

import (
	"errors"
	"reflect"
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestGetFreePosition1(t *testing.T) {
	g := newGrid(3, 3)
	p, err := g.getFreePosition(
		[]*Segment{
			newSegment(East, NewPosition(0, 1), NewPosition(2, 1)),
			newSegment(South, NewPosition(2, 0), NewPosition(2, 0)),
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(South, NewPosition(1, 2), NewPosition(1, 2)),
			newSegment(East, NewPosition(0, 2), NewPosition(2, 2)),
		},
	)
	if err != nil {
		t.Errorf("unexpected error, [%v]", err)
	} else {
		exp := NewPosition(1, 1)
		if !reflect.DeepEqual(exp, p) {
			t.Errorf("unexpected value, exp [%+v] got [%+v]", *exp, *p)
		}
	}
}

func TestGetFreePosition2(t *testing.T) {
	g := newGrid(3, 3)
	p, err := g.getFreePosition(
		[]*Segment{
			newSegment(South, NewPosition(1, 0), NewPosition(1, 2)),
			newSegment(West, NewPosition(2, 2), NewPosition(2, 2)),
			newSegment(North, NewPosition(2, 0), NewPosition(2, 1)),
			newSegment(West, NewPosition(0, 1), NewPosition(0, 1)),
			newSegment(South, NewPosition(0, 0), NewPosition(0, 2)),
		},
	)
	if err != nil {
		t.Errorf("unexpected error, [%v]", err)
	} else {
		exp := NewPosition(1, 1)
		if !reflect.DeepEqual(exp, p) {
			t.Errorf("unexpected value, exp [%+v] got [%+v]", *exp, *p)
		}
	}
}

func TestGetFreePosition3(t *testing.T) {
	g := newGrid(2, 2)
	_, err := g.getFreePosition(
		[]*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(East, NewPosition(1, 1), NewPosition(0, 1)),
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
	v := g.isFreePosition(
		NewPosition(0, 0),
		[]*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(East, NewPosition(1, 1), NewPosition(0, 1)),
		},
	)
	if v {
		t.Errorf("unexpected value")
	}
	v = g.isFreePosition(
		NewPosition(0, 0),
		[]*Segment{
			newSegment(West, NewPosition(1, 0), NewPosition(1, 0)),
			newSegment(East, NewPosition(1, 1), NewPosition(0, 1)),
		},
	)
	if !v {
		t.Errorf("unexpected value")
	}
	v = g.isFreePosition(
		NewPosition(0, 1),
		[]*Segment{
			newSegment(West, NewPosition(1, 0), NewPosition(1, 0)),
			newSegment(East, NewPosition(1, 1), NewPosition(0, 1)),
		},
	)
	if !v {
		t.Errorf("unexpected value")
	}
}

func TestMove1(t *testing.T) {
	g := newGrid(3, 3)
	p, err := g.move(
		East,
		[]*Segment{
			newSegment(East, NewPosition(0, 1), NewPosition(2, 1)),
			newSegment(South, NewPosition(2, 0), NewPosition(2, 0)),
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(South, NewPosition(1, 2), NewPosition(1, 2)),
			newSegment(East, NewPosition(0, 2), NewPosition(2, 2)),
		},
		NewPosition(1, 1),
		false,
	)
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		exp := []*Segment{
			newSegment(East, NewPosition(1, 1), NewPosition(2, 1)),
			newSegment(South, NewPosition(2, 0), NewPosition(2, 0)),
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(South, NewPosition(1, 2), NewPosition(1, 2)),
			newSegment(East, NewPosition(0, 2), NewPosition(0, 2)),
		}
		if diff := pretty.Compare(exp, p); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMove2(t *testing.T) {
	g := newGrid(3, 3)
	_, err := g.move(
		North,
		[]*Segment{
			newSegment(East, NewPosition(0, 1), NewPosition(2, 1)),
			newSegment(West, NewPosition(2, 0), NewPosition(1, 0)),
			newSegment(East, NewPosition(1, 2), NewPosition(2, 2)),
		},
		NewPosition(0, 0),
		false,
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
	p, err := g.move(
		South,
		[]*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		NewPosition(0, 1),
		false,
	)
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		exp := []*Segment{
			newSegment(South, NewPosition(0, 1), NewPosition(0, 1)),
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
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
		{"0-0", NewPosition(0, 0), nil},
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

func TestGetBodyParts1(t *testing.T) {
	g := newGrid(3, 3)
	v := g.getBodyParts(
		[]*Segment{
			newSegment(East, NewPosition(0, 1), NewPosition(2, 1)),
			newSegment(South, NewPosition(2, 0), NewPosition(2, 0)),
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(South, NewPosition(1, 2), NewPosition(1, 2)),
			newSegment(East, NewPosition(0, 2), NewPosition(2, 2)),
		},
	)
	exp := []*BodyPart{
		&BodyPart{BodySouthWest, NewPosition(2, 1)},
		&BodyPart{HeadEast, NewPosition(0, 1)},
		&BodyPart{BodyNorthWest, NewPosition(2, 0)},
		&BodyPart{BodySouthEast, NewPosition(1, 0)},
		&BodyPart{BodyHorizontal, NewPosition(0, 0)},
		&BodyPart{BodyNorthEast, NewPosition(1, 2)},
		&BodyPart{TailEast, NewPosition(2, 2)},
		&BodyPart{BodyHorizontal, NewPosition(0, 2)},
	}

	if diff := pretty.Compare(exp, v); diff != "" {
		t.Errorf("unexpected value\n%s", diff)
	}
}

func TestGetBodyParts2(t *testing.T) {
	g := newGrid(3, 3)
	v := g.getBodyParts(
		[]*Segment{
			newSegment(South, NewPosition(1, 0), NewPosition(1, 2)),
			newSegment(West, NewPosition(2, 2), NewPosition(2, 2)),
			newSegment(North, NewPosition(2, 0), NewPosition(2, 1)),
			newSegment(West, NewPosition(0, 1), NewPosition(0, 1)),
			newSegment(South, NewPosition(0, 0), NewPosition(0, 2)),
		},
	)
	exp := []*BodyPart{
		&BodyPart{BodyNorthWest, NewPosition(1, 2)},
		&BodyPart{HeadSouth, NewPosition(1, 0)},
		&BodyPart{BodyNorthEast, NewPosition(2, 2)},
		&BodyPart{BodySouthWest, NewPosition(2, 1)},
		&BodyPart{BodyVertical, NewPosition(2, 0)},
		&BodyPart{BodySouthEast, NewPosition(0, 1)},
		&BodyPart{TailSouth, NewPosition(0, 2)},
		&BodyPart{BodyVertical, NewPosition(0, 0)},
	}

	if diff := pretty.Compare(exp, v); diff != "" {
		t.Errorf("unexpected value\n%s", diff)
	}
}
