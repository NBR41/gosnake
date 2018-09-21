package engine

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
