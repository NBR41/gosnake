package engine

import (
	"fmt"
)

//Position struct for point position
type Position struct {
	x, y int
}

//NewPosition returns new instance of position
func NewPosition(x, y int) *Position {
	return &Position{x: x, y: y}
}

//String returns position string representation
func (p Position) String() string {
	return fmt.Sprintf("%d-%d", p.x, p.y)
}

//X returns X axis position
func (p Position) X() int {
	return p.x
}

//Y returns Y axis position
func (p Position) Y() int {
	return p.y
}
