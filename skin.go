package gosnake

import (
	"image/color"
	"strconv"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"

	//"github.com/NBR41/gosnake/assets"
	"github.com/NBR41/gosnake/engine"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	//"github.com/skatiyar/pacman/spritetools"
)

const MaxScoreView = 999999999

func SkinView(skin *ebiten.Image, arcadeFont *truetype.Font) (fView, error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    30,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	width, height := skin.Size()
	view, err := ebiten.NewImage(width, height, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	return func(state gameState, data *engine.Data) (*ebiten.Image, error) {
		if err := view.Clear(); err != nil {
			return nil, err
		}
		if err := view.DrawImage(skin, &ebiten.DrawImageOptions{}); err != nil {
			return nil, err
		}

		switch state {
		case GameStart:
			fallthrough
		case GamePause:
			fallthrough
		case GameOver:
			if data != nil {
				score := data.Score
				if score > MaxScoreView {
					score = MaxScoreView
				}
				numstr := strconv.Itoa(score)
				text.Draw(view, numstr, fontface, 682-(len(numstr)*27), 64, color.White)
			}
		}
		return view, nil
	}, nil
}
