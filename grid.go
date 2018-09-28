package gosnake

import (
	"image/color"
	"math"

	"github.com/NBR41/gosnake/assets"
	"github.com/NBR41/gosnake/assets/sprite"
	"github.com/NBR41/gosnake/engine"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"

	"golang.org/x/image/font"
)

var (
	uTurnRadian     = math.Pi
	rightTurnRadian = math.Pi / 2
	leftTurnRadian  = -1 * math.Pi / 2
)

func gridView(
	size, colnb, rownb int, body *assets.Body, imgfruit *ebiten.Image, arcadeFont *truetype.Font,
) (fView, error) {
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
		if err := view.Fill(sprite.BackgroundColor); err != nil {
			return nil, err
		}

		switch state {
		case GameLoading:
			text.Draw(view, "PRESS SPACE", fontface, (gridWith/2)-176, (gridHeight/2)-(10+32), color.White)
			text.Draw(view, "TO BEGIN", fontface, (gridWith/2)-128, (gridHeight/2)+(10), color.White)
		case GameStart, GamePause, GameOver:
			ops := &ebiten.DrawImageOptions{}
			// fruit
			if err := drawSprite(view, imgfruit, ops, data.GetFruit(), size, nil); err != nil {
				return nil, err
			}

			// body
			parts := data.GetBodyParts()
			var turn *float64
			var img *ebiten.Image
			for i := range parts {
				switch parts[i].GetImage() {

				case engine.HeadNorth:
					img = body.Head
					turn = nil
				case engine.HeadEast:
					img = body.Head
					turn = &rightTurnRadian
				case engine.HeadSouth:
					img = body.Head
					turn = &uTurnRadian
				case engine.HeadWest:
					img = body.Head
					turn = &leftTurnRadian

				case engine.TailNorth:
					img = body.Tail
					turn = &uTurnRadian
				case engine.TailEast:
					img = body.Tail
					turn = &leftTurnRadian
				case engine.TailSouth:
					img = body.Tail
					turn = nil
				case engine.TailWest:
					img = body.Tail
					turn = &rightTurnRadian

				case engine.BodyHorizontal:
					img = body.Straight
					turn = nil
				case engine.BodyVertical:
					img = body.Straight
					turn = &rightTurnRadian

				case engine.BodyNorthWest:
					img = body.Curve
					turn = nil
				case engine.BodyNorthEast:
					img = body.Curve
					turn = &rightTurnRadian
				case engine.BodySouthEast:
					img = body.Curve
					turn = &uTurnRadian
				case engine.BodySouthWest:
					img = body.Curve
					turn = &leftTurnRadian
				}
				if err := drawSprite(view, img, ops, parts[i].GetPosition(), size, turn); err != nil {
					return nil, err
				}
			}

			if state == GamePause {
				img, err := getMessageImage()
				if err != nil {
					return nil, err
				}
				text.Draw(img, "GAME PAUSED", fontface, 24, 65-(10), color.White)
				text.Draw(img, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)
				if err = pastMessageImage(view, img, gridWith, gridHeight); err != nil {
					return nil, err
				}
			} else if state == GameOver {
				img, err := getMessageImage()
				if err != nil {
					return nil, err
				}
				text.Draw(img, "GAME OVER", fontface, 56, 65-(10), color.White)
				text.Draw(img, "PRESS SPACE", fontface, 24, 65+(10+31), color.White)
				if err = pastMessageImage(view, img, gridWith, gridHeight); err != nil {
					return nil, err
				}
			}
		}

		return view, nil
	}, nil
}

func drawSprite(
	view, img *ebiten.Image, ops *ebiten.DrawImageOptions, pos *engine.Position,
	size int, turn *float64,
) error {
	ops.GeoM.Reset()
	if turn != nil {
		ops.GeoM.Rotate(*turn)
		switch *turn {
		case uTurnRadian:
			ops.GeoM.Translate(float64((pos.X()+1)*size), float64((pos.Y()+1)*size))
		case rightTurnRadian:
			ops.GeoM.Translate(float64((pos.X()+1)*size), float64(pos.Y()*size))
		case leftTurnRadian:
			ops.GeoM.Translate(float64(pos.X()*size), float64((pos.Y()+1)*size))
		}
	} else {
		ops.GeoM.Translate(float64(pos.X()*size), float64(pos.Y()*size))
	}
	return view.DrawImage(img, ops)
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

func pastMessageImage(view *ebiten.Image, img *ebiten.Image, gridWidth, gridHeight int) error {
	ops := &ebiten.DrawImageOptions{}
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64((gridWidth/2)-(389/2)), float64((gridHeight/2)-(130/2)))
	return view.DrawImage(img, ops)
}
