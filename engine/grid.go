package engine

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//List of errors
var (
	ErrNoPosition = errors.New("no position")
	ErrInvalidKey = errors.New("invalid key")
	ErrColision   = errors.New("colision")
)

//Grid struct for grid
type Grid struct {
	colNb int
	rowNb int
	rand  *rand.Rand
}

func newGrid(colnb, rownb int) *Grid {
	return &Grid{
		colNb: colnb,
		rowNb: rownb,
		rand:  rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

//GetFreePosition returns a free position on the grid
func (g *Grid) getFreePosition(segs []*Segment) (*Position, error) {
	m := g.getFreePositions(segs)
	if len(m) > 0 {
		for i := range m {
			return getPosition(i)
		}
	}
	return nil, ErrNoPosition
}

func (g *Grid) getFreePositions(segs []*Segment) map[string]interface{} {
	m := g.getGridMap()
	g.filterGridMap(m, segs)
	return m
}

func (g *Grid) getGridMap() map[string]interface{} {
	ret := make(map[string]interface{})
	for i := 0; i < g.colNb; i++ {
		for j := 0; j < g.rowNb; j++ {
			ret[fmt.Sprintf("%d-%d", i, j)] = nil
		}
	}
	return ret
}

func (g *Grid) filterGridMap(m map[string]interface{}, segs []*Segment) {
	for i := range segs {
		g.filterSegment(m, segs[i])
	}
}

func (g *Grid) filterSegment(m map[string]interface{}, seg *Segment) {
	cur := *seg.end
	pos := &cur
	for {
		delete(m, cur.String())
		if equalPosition(pos, seg.start) {
			return
		}

		pos = g.getNextPosition(seg.dir, pos)
	}
}

//IsFreePosition returns true if the position is free in the grid
func (g *Grid) isFreePosition(p *Position, segs []*Segment) bool {
	if *p == *segs[len(segs)-1].end {
		return true
	}
	m := g.getFreePositions(segs)
	_, ok := m[p.String()]
	return ok
}

//Move returns new list of segments with a move of one
func (g *Grid) move(dir Direction, segs []*Segment, next *Position, movetail bool) ([]*Segment, error) {
	if !g.isFreePosition(next, segs) {
		return nil, ErrColision
	}
	return g.translate(dir, segs, next, movetail), nil
}

//GetNextPosition get next position according to the direction
func (g *Grid) getNextPosition(dir Direction, pos *Position) *Position {
	switch dir {
	case North:
		if pos.y == 0 {
			pos.y = (g.rowNb - 1)
		} else {
			pos.y--
		}
	case South:
		if pos.y == (g.rowNb - 1) {
			pos.y = 0
		} else {
			pos.y++
		}
	case West:
		if pos.x == 0 {
			pos.x = (g.colNb - 1)
		} else {
			pos.x--
		}
	case East:
		if pos.x == (g.colNb - 1) {
			pos.x = 0
		} else {
			pos.x++
		}
	}
	return pos
}

func (g *Grid) translate(dir Direction, segs []*Segment, next *Position, chomp bool) []*Segment {
	// Move Head
	if dir == segs[0].dir {
		segs[0].start = next
	} else {
		segs = append([]*Segment{newSegment(dir, next, next)}, segs...)
	}

	// Move Tail
	if !chomp {
		lastIndex := len(segs) - 1
		if equalPosition(segs[lastIndex].start, segs[lastIndex].end) {
			return segs[:lastIndex]
		}
		segs[lastIndex].end = g.getNextPosition(segs[lastIndex].dir, segs[lastIndex].end)
	}
	return segs
}

func (g *Grid) getBodyParts(segs []*Segment) []*BodyPart {
	last := len(segs) - 1
	ret := []*BodyPart{}

	for i := range segs {
		pos := *segs[i].end
		cur := &pos
	LoopSegment:
		for {
			switch {
			case equalPosition(segs[i].start, cur):
				if i == 0 { // Head
					ret = append(ret, newBodyPart(getHeadImageType(segs[i].dir), *cur))
				} else { // Body Straigth
					ret = append(ret, newBodyPart(getBodyImageType(segs[i].dir), *cur))
				}

			case equalPosition(segs[i].end, cur):
				if i == last { // Tail
					ret = append(ret, newBodyPart(getTailImageType(segs[i].dir), *cur))
				} else { // Body Curve
					ret = append(ret, newBodyPart(getCurveBodyImageType(segs[i].dir, segs[i+1].dir), *cur))
				}

			default:
				ret = append(ret, newBodyPart(getBodyImageType(segs[i].dir), *cur))
			}

			if equalPosition(cur, segs[i].start) {
				break LoopSegment
			}

			cur = g.getNextPosition(segs[i].dir, cur)
		}
	}
	return ret
}

func equalPosition(a, b *Position) bool {
	return a.x == b.x && a.y == b.y
}

func getPosition(k string) (*Position, error) {
	parts := strings.Split(k, "-")
	if len(parts) != 2 {
		return nil, ErrInvalidKey
	}
	p := &Position{}
	var err error
	if p.x, err = strconv.Atoi(parts[0]); err != nil {
		return nil, err
	}
	if p.y, err = strconv.Atoi(parts[1]); err != nil {
		return nil, err
	}
	return p, nil
}
