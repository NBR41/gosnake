package engine

//Segment struct for snake segment
type Segment struct {
	dir        Direction
	start, end *Position
}

func newSegment(dir Direction, start, end *Position) *Segment {
	return &Segment{dir: dir, start: start, end: end}
}

//IsHorizontal return true if the segment is horizontal
func (s Segment) IsHorizontal() bool {
	return s.start.y == s.end.y
}
