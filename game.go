package gosnake

import (
	"math/rand"
	"time"

	"github.com/NBR41/gosnake/assets"
	"github.com/NBR41/gosnake/engine"
	"github.com/hajimehoshi/ebiten"
)

type gameState int

type fView func(gameState, *engine.Data) (*ebiten.Image, error)

const (
	GameLoading gameState = iota
	GameStart
	GamePause
	GameOver
)

type Game struct {
	state gameState
	rand  *rand.Rand
	data  *engine.Data
	dir   *engine.Direction
	size  int
	colnb int
	rownb int

	skinView fView
	gridView fView
	audio    *Audio
}

func NewGame(size, colnb, rownb int) (*Game, error) {
	assets, err := assets.LoadAssets(size, colnb, rownb)
	if err != nil {
		return nil, err
	}

	gridView, err := GridView(size, colnb, rownb, assets.Body, assets.Fruit, assets.ArcadeFont)
	if err != nil {
		return nil, err
	}

	skinView, err := SkinView(assets.Skin, assets.ArcadeFont)
	if err != nil {
		return nil, err
	}

	audio, err := NewAudio()
	if err != nil {
		return nil, err
	}

	return &Game{
		rand:     rand.New(rand.NewSource(time.Now().UnixNano())),
		state:    GameLoading,
		size:     size,
		colnb:    colnb,
		rownb:    rownb,
		skinView: skinView,
		gridView: gridView,
		audio:    audio,
	}, nil
}

func (g *Game) Run() error {
	return ebiten.Run(
		func(screen *ebiten.Image) error {
			return g.update(screen)
		},
		(g.colnb*g.size)+8,
		(g.rownb*g.size)+46,
		1,
		"GoSnake",
	) // scale is kept to 0.5, for good rendering in retina.
}

func (g *Game) update(screen *ebiten.Image) error {
	switch g.state {
	case GameLoading:
		if spaceReleased() {
			g.data = engine.NewData(g.colnb, g.rownb)
			if err := g.data.SetFruit(); err != nil {
				return err
			}
			g.audio.players.Beginning.Pause()
			g.audio.players.Beginning.Rewind()
			g.state = GameStart
			break
		} else {
			g.data = nil
			g.audio.players.Beginning.Play()
		}
	case GameStart:
		if spaceReleased() {
			g.state = GamePause
		} else {
			var err error
			switch {
			case upKeyPressed():
				err = g.data.MoveNorth()
			case downKeyPressed():
				err = g.data.MoveSouth()
			case leftKeyPressed():
				err = g.data.MoveWest()
			case rightKeyPressed():
				err = g.data.MoveEast()
			default:
				//err = g.data.Move()
			}
			if err != nil {
				if err == engine.ErrColision {
					g.state = GameOver
				} else {
					return err
				}
			}
		}
	case GamePause:
		if spaceReleased() {
			g.state = GameStart
		}
	case GameOver:
		if spaceReleased() {
			g.state = GameLoading
			g.audio.players.Death.Pause()
			g.audio.players.Death.Rewind()
		} else {
			g.audio.players.Death.Play()
		}
	default:
		// reset state to GameLoading
		// dont return error
		g.state = GameLoading
	}

	if ebiten.IsRunningSlowly() {
		return nil
	}

	sview, err := g.skinView(g.state, g.data)
	if err != nil {
		return err
	}

	gview, err := g.gridView(g.state, g.data)
	if err != nil {
		return err
	}

	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	if err := screen.DrawImage(sview, ops); err != nil {
		return err
	}

	ops.GeoM.Reset()
	ops.GeoM.Translate(4, 42)
	if err := screen.DrawImage(gview, ops); err != nil {
		return err
	}

	return nil
}
