package engine

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

func (d *Data) move(dir Direction) error {
	segs, err := d.grid.Move(dir, d.snake.body)
	if err != nil {
		return err
	}
	d.snake.body = segs
	return nil
}
