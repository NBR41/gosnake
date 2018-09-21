package gosnake

import (
	"image/color"

	"github.com/NBR41/gosnake/engine"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"

	"golang.org/x/image/font"
)

var (
	GrayColor = color.RGBA{236, 240, 241, 255.0}
)

func GridView(size, colnb, rownb int, arcadeFont *truetype.Font) (fView, error) {
	fontface := truetype.NewFace(arcadeFont, &truetype.Options{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})

	gridWith := size * colnb
	gridHeight := size * rownb

	view, err := ebiten.NewImage(gridWith, gridHeight, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	return func(state gameState, data *engine.Data) (*ebiten.Image, error) {
		if err := view.Clear(); err != nil {
			return nil, err
		}
		if err := view.Fill(color.Black); err != nil {
			return nil, err
		}

		switch state {
		case GameLoading:
			text.Draw(view, "PRESS SPACE", fontface, (gridWith/2)-176, (gridHeight/2)-(10+32), color.White)
			text.Draw(view, "TO BEGIN", fontface, (gridWith/2)-128, (gridHeight/2)+(10), color.White)
		case GameStart, GamePause, GameOver:
			if state == GamePause {
				img, err := getMessageImage()
				if err != nil {
					return nil, err
				}
				text.Draw(img, "GAME PAUSED", fontface, 24, 65-(10), color.White)
				text.Draw(img, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)
				if err = pastMessageImage(view, img, gridHeight); err != nil {
					return nil, err
				}
			} else if state == GameOver {
				img, err := getMessageImage()
				if err != nil {
					return nil, err
				}
				text.Draw(img, "GAME OVER", fontface, 56, 65-(10), color.White)
				text.Draw(img, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)
				if err = pastMessageImage(view, img, gridHeight); err != nil {
					return nil, err
				}
			}
		}

		return view, nil
	}, nil
}

func getMessageImage() (*ebiten.Image, error) {
	img, err := ebiten.NewImage(389, 130, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	if err := img.Fill(color.Black); err != nil {
		return nil, err
	}
	return img, nil
}

func pastMessageImage(view *ebiten.Image, img *ebiten.Image, gridHeight int) error {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64((gridHeight/2)-(389/2)), float64((gridHeight/2)-(130/2)))
	return view.DrawImage(img, ops)
}
