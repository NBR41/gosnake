package gosnake

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Direction int

//List of directions
const (
	North Direction = iota
	East
	South
	West
)

var (
	errNoPosition = errors.New("no position")
	errInvalidKey = errors.New("invalid key")
)

//Data struct for game data
type Data struct {
	score int
	snake *Snake
	grid  *Grid
	fruit *Position
}

func newData(colnb, rownb int) *Data {
	return &Data{
		score: 0,
		snake: newSnake(rownb),
		grid:  newGrid(colnb, rownb),
	}
}

//SetFruit set fruit position on the grid
func (d *Data) SetFruit() error {
	var err error
	d.fruit, err = d.grid.GetFreePosition(d.snake.body)
	return err
}

//Position struct for point position
type Position struct {
	x, y int
}

func newPosition(x, y int) *Position {
	return &Position{x: x, y: y}
}

func (p Position) getKey() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}

//Segment struct for snake segment
type Segment struct {
	dir        Direction
	start, end *Position
}

func newSegment(dir Direction, start, end *Position) *Segment {

	return &Segment{dir: dir, start: start, end: end}
}

//Snake struct for snake data
type Snake struct {
	dir  Direction
	body []*Segment
	size int
}

func newSnake(rownb int) *Snake {
	return &Snake{
		dir:  East,
		size: 3,
		body: []*Segment{newSegment(East, newPosition(2, rownb/2), newPosition(0, rownb/2))},
	}
}

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
func (g *Grid) GetFreePosition(segs []*Segment) (*Position, error) {
	m := g.getGridMap()
	g.filterGridMap(m, segs)
	if len(m) > 0 {
		for i := range m {
			return g.getPosition(i)
		}
	}
	return nil, errNoPosition
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
		switch {

		case segs[i].start.y == segs[i].end.y: //horizontal
			cur := segs[i].start
		Loop:
			for {

				delete(m, cur.getKey())

			}

		default: //vertical
		}
	}
}

func (g *Grid) getPosition(k string) (*Position, error) {
	parts := strings.Split(k, "-")
	if len(parts) != 2 {
		return nil, errInvalidKey
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
