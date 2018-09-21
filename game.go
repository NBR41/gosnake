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
	state    gameState
	rand     *rand.Rand
	data     *engine.Data
	skinView fView
	gridView fView
	audio    *Audio
}

func NewGame(size, colnb, rownb int) (*Game, error) {
	lAssets, err := assets.LoadAssets(size, colnb, rownb)
	if err != nil {
		return nil, err
	}

	gridView, err := GridView(size, colnb, rownb, lAssets.ArcadeFont)
	if err != nil {
		return nil, err
	}

	skinView, err := SkinView(lAssets.Skin, lAssets.ArcadeFont)
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
		1008,
		546,
		1,
		"GoSnake",
	) // scale is kept to 0.5, for good rendering in retina.
}

func (g *Game) update(screen *ebiten.Image) error {
	switch g.state {
	case GameLoading:
		if spaceReleased() {

			g.audio.players.Beginning.Pause()
			g.audio.players.Beginning.Rewind()
			g.state = GameStart
			break
		}
		g.data = nil
		g.audio.players.Beginning.Play()
	case GameStart:
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
