package gosnake

import (
	"image/color"
	"strconv"

	"github.com/NBR41/gosnake/engine"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"

	"golang.org/x/image/font"
)

//MaxScoreView maximum score
const MaxScoreView = 999999999

func skinView(skin *ebiten.Image, arcadeFont *truetype.Font) (fView, error) {
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

	text.Draw(skin, "GoSnake", fontface, 10, 34, color.White)

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
				score := data.Score()
				if score > MaxScoreView {
					score = MaxScoreView
				}
				numstr := strconv.Itoa(score)
				text.Draw(view, numstr, fontface, width-((len(numstr)*27)+6), 34, color.White)
			}
		}
		return view, nil
	}, nil
}
