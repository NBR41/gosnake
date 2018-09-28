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

// List of game state
const (
	GameLoading gameState = iota
	GameStart
	GamePause
	GameOver
)

//Game whole game struct
type Game struct {
	state      gameState
	rand       *rand.Rand
	data       *engine.Data
	dir        *engine.Direction
	size       int
	colnb      int
	rownb      int
	fps        int
	frameCount int
	speed      int
	skinView   fView
	gridView   fView
	audio      *audioReg
}

//NewGame returns new instance of game
func NewGame(size, colnb, rownb int) (*Game, error) {
	assets, err := assets.LoadAssets(size, colnb, rownb)
	if err != nil {
		return nil, err
	}

	gridView, err := gridView(size, colnb, rownb, assets.Body, assets.Fruit, assets.ArcadeFont)
	if err != nil {
		return nil, err
	}

	skinView, err := skinView(assets.Skin, assets.ArcadeFont)
	if err != nil {
		return nil, err
	}

	audio, err := newAudio()
	if err != nil {
		return nil, err
	}

	return &Game{
		rand:       rand.New(rand.NewSource(time.Now().UnixNano())),
		state:      GameLoading,
		size:       size,
		colnb:      colnb,
		rownb:      rownb,
		fps:        30,
		frameCount: 0,
		speed:      0,
		skinView:   skinView,
		gridView:   gridView,
		audio:      audio,
	}, nil
}

//Run start the game
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
			g.incFrame()
			var err error
			var chomp bool
			switch {
			case upKeyPressed():
				chomp, err = g.data.MoveNorth()
			case downKeyPressed():
				chomp, err = g.data.MoveSouth()
			case leftKeyPressed():
				chomp, err = g.data.MoveWest()
			case rightKeyPressed():
				chomp, err = g.data.MoveEast()
			default:
				if g.canMove() {
					chomp, err = g.data.Move()
				}
			}
			if err != nil {
				if err == engine.ErrColision {
					g.state = GameOver
				} else {
					return err
				}
			} else {
				if chomp {
					g.audio.players.Chomp.Play()
					g.audio.players.Chomp.Rewind()
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

func (g *Game) incFrame() {
	if g.frameCount >= g.fps {
		g.frameCount = 0
	} else {
		g.frameCount++
	}
}

func (g *Game) canMove() bool {
	return g.frameCount%g.fps == g.speed
}
