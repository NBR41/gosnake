package sprite

import (
	"image"
	"image/color"
	"math"
)

type triangle struct {
	p image.Point
	h int
}

func (t *triangle) ColorModel() color.Model {
	return color.AlphaModel
}

func (t *triangle) Bounds() image.Rectangle {
	return image.Rect(t.p.X-t.h, t.p.Y, t.p.X+t.h, t.p.Y+t.h)
}

func (t *triangle) At(x, y int) color.Color {
	if math.Abs(float64(t.p.X-x)) <= float64(y-t.p.Y) {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

type circle struct {
	p image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.p.X-c.r, c.p.Y-c.r, c.p.X+c.r, c.p.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.p.X)+0.5, float64(y-c.p.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}
