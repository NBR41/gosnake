package assets

import (
	"github.com/NBR41/gosnake/assets/fonts"
	"github.com/NBR41/gosnake/assets/sounds"
	"github.com/NBR41/gosnake/assets/sprite"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
)

//Assets struct for assets list
type Assets struct {
	ArcadeFont *truetype.Font
	Skin       *ebiten.Image
	Body       *Body
	Fruit      *ebiten.Image
}

//Body struct for snake's body
type Body struct {
	Head     *ebiten.Image
	Straight *ebiten.Image
	Curve    *ebiten.Image
	Tail     *ebiten.Image
}

// LoadAssets converts the character images(png, jpg, ...) to
// ebiten image format and loads fonts.
func LoadAssets(size, colnb, rownb int) (*Assets, error) {
	var (
		a   = &Assets{}
		err error
	)
	g := sprite.NewGenerator(size, colnb, rownb)

	if a.Skin, err = ebiten.NewImageFromImage(g.GetSkin(), ebiten.FilterDefault); err != nil {
		return nil, err
	}
	if a.ArcadeFont, err = truetype.Parse(fonts.ArcadeTTF); err != nil {
		return nil, err
	}
	if a.Body, err = loadBody(g); err != nil {
		return nil, err
	}
	if a.Fruit, err = ebiten.NewImageFromImage(g.GetFruit(), ebiten.FilterDefault); err != nil {
		return nil, err
	}
	return a, nil
}

func loadBody(g *sprite.Generator) (*Body, error) {
	var (
		b   = &Body{}
		err error
	)
	if b.Head, err = ebiten.NewImageFromImage(g.GetBodyHead(), ebiten.FilterDefault); err != nil {
		return nil, err
	}
	if b.Straight, err = ebiten.NewImageFromImage(g.GetBodyStraight(), ebiten.FilterDefault); err != nil {
		return nil, err
	}
	if b.Curve, err = ebiten.NewImageFromImage(g.GetBodyCurve(), ebiten.FilterDefault); err != nil {
		return nil, err
	}
	if b.Tail, err = ebiten.NewImageFromImage(g.GetBodyTail(), ebiten.FilterDefault); err != nil {
		return nil, err
	}
	return b, nil
}

//Sounds struct for all sounds
type Sounds struct {
	Beginning *wav.Stream
	Chomp     *wav.Stream
	Death     *wav.Stream
}

// LoadSounds returns a struct with wav files decoded
// for the provided audio context.
func LoadSounds(ctx *audio.Context) (*Sounds, error) {
	var (
		s   = &Sounds{}
		err error
	)
	if s.Beginning, err = wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.BeginningWav)); err != nil {
		return nil, err
	}
	if s.Chomp, err = wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.ChompWav)); err != nil {
		return nil, err
	}
	if s.Death, err = wav.Decode(ctx, audio.BytesReadSeekCloser(sounds.DeathWav)); err != nil {
		return nil, err
	}
	return s, nil
}
