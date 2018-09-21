package engine

//Data struct for game data
type Data struct {
	score int
	dir   Direction
	snake *Snake
	grid  *Grid
	fruit *Position
}

func NewData(colnb, rownb int) *Data {
	return &Data{
		score: 0,
		dir:   East,
		snake: newSnake(rownb),
		grid:  newGrid(colnb, rownb),
	}
}

//SetFruit set fruit position on the grid
func (d *Data) SetFruit() error {
	var err error
	d.fruit, err = d.grid.getFreePosition(d.snake.body)
	return err
}

func (d *Data) Move() error {
	return d.move(d.Direction())
}

func (d *Data) MoveNorth() error {
	return d.move(North)
}

func (d *Data) MoveSouth() error {
	return d.move(South)
}

func (d *Data) MoveWest() error {
	return d.move(West)
}

func (d *Data) MoveEast() error {
	return d.move(East)
}

func (d *Data) Direction() Direction {
	return d.dir
}

func (d *Data) Score() int {
	return d.score
}

func (d *Data) move(dir Direction) error {
	head := *d.snake.body[0].start
	next := d.grid.getNextPosition(dir, &head)
	chomp := equalPosition(next, d.fruit)
	segs, err := d.grid.move(dir, d.snake.body, next, chomp)
	if err != nil {
		return err
	}
	if chomp {
		d.score += 10
	}
	d.snake.body = segs
	d.dir = dir
	return d.SetFruit()
}
