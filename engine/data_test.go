package engine

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
)

func TestData(t *testing.T) {
	d := NewData(3, 3)

	exp := &Data{
		score: 0,
		dir:   East,
		body:  []*Segment{newSegment(East, NewPosition(2, 1), NewPosition(0, 1))},
		grid: &Grid{
			colNb: 3,
			rowNb: 3,
			rand:  d.grid.rand,
		},
	}

	if diff := pretty.Compare(exp, d); diff != "" {
		t.Errorf("unexpected value\n%s", diff)
	}

	if d.Direction() != East {
		t.Errorf("unexpected value, exp [%d] got [%d]", East, d.Direction())
	}
	if d.Score() != 0 {
		t.Errorf("unexpected value, exp [%d] got [%d]", 0, d.Score())
	}
}

func TestSetFruit1(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{
			colNb: 2,
			rowNb: 2,
		},
	}

	err := d.SetFruit()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(NewPosition(0, 1), d.fruit); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestSetFruit2(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
			newSegment(East, NewPosition(0, 1), NewPosition(0, 1)),
		},
		grid: &Grid{
			colNb: 2,
			rowNb: 2,
		},
	}

	err := d.SetFruit()
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != ErrNoPosition {
			t.Errorf("unexpected error, exp [%v] got [%v]", ErrNoPosition, err)
		}
	}
}

func TestDataMove1(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{
			colNb: 2,
			rowNb: 2,
		},
	}
	err := d.Move()
	if err == nil {
		t.Error("expecting error")
	} else {
		if err != ErrColision {
			t.Errorf("unexpected error, exp [%v] got [%v]", ErrColision, err)
		}
	}
}

func TestDataMove2(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   South,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		fruit: NewPosition(0, 1),
		grid: &Grid{
			colNb: 2,
			rowNb: 2,
		},
	}

	exp := &Data{
		score: 10,
		dir:   South,
		body: []*Segment{
			newSegment(South, NewPosition(0, 1), NewPosition(0, 0)),
			newSegment(West, NewPosition(1, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		fruit: nil,
		grid: &Grid{
			colNb: 2,
			rowNb: 2,
		},
	}

	err := d.Move()
	if err == nil {
		t.Error("expecting error")
	} else {
		if diff := pretty.Compare(exp, *d); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
		if err != ErrNoPosition {
			t.Errorf("unexpected error, exp [%v] got [%v]", ErrNoPosition, err)
		}
	}
}

func TestMoveNorth1(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   North,
		body: []*Segment{
			newSegment(North, NewPosition(0, 1), NewPosition(0, 0)),
			newSegment(West, NewPosition(1, 0), NewPosition(1, 0)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}

	err := d.MoveNorth()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMoveNorth2(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   South,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   South,
		body: []*Segment{
			newSegment(South, NewPosition(0, 1), NewPosition(0, 0)),
			newSegment(West, NewPosition(1, 0), NewPosition(1, 0)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}

	err := d.MoveNorth()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMoveWest1(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   North,
		body: []*Segment{
			newSegment(North, NewPosition(1, 0), NewPosition(1, 1)),
			newSegment(East, NewPosition(0, 1), NewPosition(0, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}

	err := d.MoveWest()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMoveWest2(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   East,
		body: []*Segment{
			newSegment(East, NewPosition(1, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
			newSegment(East, NewPosition(0, 1), NewPosition(0, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   East,
		body: []*Segment{
			newSegment(East, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}

	err := d.MoveWest()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMoveEast1(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   North,
		body: []*Segment{
			newSegment(North, NewPosition(0, 0), NewPosition(0, 1)),
			newSegment(West, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   East,
		body: []*Segment{
			newSegment(East, NewPosition(1, 0), NewPosition(0, 0)),
			newSegment(North, NewPosition(0, 1), NewPosition(0, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	err := d.MoveEast()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMoveEast2(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(1, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
			newSegment(East, NewPosition(0, 1), NewPosition(0, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	err := d.MoveEast()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMoveSouth1(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   North,
		body: []*Segment{
			newSegment(North, NewPosition(1, 0), NewPosition(1, 0)),
			newSegment(East, NewPosition(0, 0), NewPosition(0, 0)),
			newSegment(North, NewPosition(0, 1), NewPosition(0, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   North,
		body: []*Segment{
			newSegment(North, NewPosition(1, 1), NewPosition(1, 0)),
			newSegment(East, NewPosition(0, 0), NewPosition(0, 0)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}

	err := d.MoveSouth()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestMoveSouth2(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}
	exp := &Data{
		score: 0,
		dir:   South,
		body: []*Segment{
			newSegment(South, NewPosition(0, 1), NewPosition(0, 0)),
			newSegment(West, NewPosition(1, 0), NewPosition(1, 0)),
		},
		grid: &Grid{colNb: 2, rowNb: 2},
	}

	err := d.MoveSouth()
	if err != nil {
		t.Errorf("unexpected error, %v", err)
	} else {
		if diff := pretty.Compare(d, exp); diff != "" {
			t.Errorf("unexpected value\n%s", diff)
		}
	}
}

func TestGetBodyParts(t *testing.T) {
	d := &Data{
		score: 0,
		dir:   West,
		body: []*Segment{
			newSegment(West, NewPosition(0, 0), NewPosition(1, 0)),
			newSegment(North, NewPosition(1, 1), NewPosition(1, 1)),
		},
		grid: &Grid{
			colNb: 2,
			rowNb: 2,
		},
	}

	p := d.GetBodyParts()
	exp := []*BodyPart{
		&BodyPart{img: BodyNorthEast, pos: &Position{x: 1, y: 0}},
		&BodyPart{img: HeadWest, pos: &Position{x: 0, y: 0}},
		&BodyPart{img: TailNorth, pos: &Position{x: 1, y: 1}},
	}
	if diff := pretty.Compare(exp, p); diff != "" {
		t.Errorf("unexpected value\n%s", diff)
	}
}
