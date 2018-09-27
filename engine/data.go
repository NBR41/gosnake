package engine

//Data struct for game data
type Data struct {
	score int
	dir   Direction
	body  []*Segment
	grid  *Grid
	fruit *Position
}

//NewData returns new instance of data
func NewData(colnb, rownb int) *Data {
	return &Data{
		score: 0,
		dir:   East,
		body:  newBody(rownb),
		grid:  newGrid(colnb, rownb),
	}
}

func newBody(rownb int) []*Segment {
	return []*Segment{newSegment(East, NewPosition(2, rownb/2), NewPosition(0, rownb/2))}
}

//SetFruit set fruit position on the grid
func (d *Data) SetFruit() error {
	var err error
	d.fruit, err = d.grid.getFreePosition(d.body)
	return err
}

func (d *Data) GetFruit() *Position {
	return d.fruit
}

//Move move snake in the current direction
func (d *Data) Move() error {
	return d.move(d.Direction())
}

//MoveNorth move snake to the North if the current dir is not South
func (d *Data) MoveNorth() error {
	if d.dir == South {
		return d.Move()
	}
	return d.move(North)
}

//MoveSouth move snake to the South if the current dir is not North
func (d *Data) MoveSouth() error {
	if d.dir == North {
		return d.Move()
	}
	return d.move(South)
}

//MoveWest move snake to the West if the current dir is not East
func (d *Data) MoveWest() error {
	if d.dir == East {
		return d.Move()
	}
	return d.move(West)
}

//MoveEast move snake to the East if the current dir is not West
func (d *Data) MoveEast() error {
	if d.dir == West {
		return d.Move()
	}
	return d.move(East)
}

//Direction returns current Direction
func (d *Data) Direction() Direction {
	return d.dir
}

//Score returns current score
func (d *Data) Score() int {
	return d.score
}

//GetBodyParts returns snke body parts
func (d *Data) GetBodyParts() []*BodyPart {
	return d.grid.getBodyParts(d.body)
}

func (d *Data) move(dir Direction) error {
	head := *d.body[0].start
	next := d.grid.getNextPosition(dir, &head)

	var chomp bool
	if d.fruit != nil {
		chomp = equalPosition(next, d.fruit)
	}

	segs, err := d.grid.move(dir, d.body, next, chomp)
	if err != nil {
		return err
	}
	if chomp {
		d.score += 10
	}
	d.body = segs
	d.dir = dir
	if chomp {
		err = d.SetFruit()
	}
	return err
}
