package engine

import (
	"fmt"
)

//Position struct for point position
type Position struct {
	x, y int
}

func newPosition(x, y int) *Position {
	return &Position{x: x, y: y}
}

func (p Position) String() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}
