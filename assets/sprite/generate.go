package sprite

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
)

const (
	Bound = 4
)

var (
	BackgroundColor = color.RGBA{60, 20, 60, 255}
	BoundColor      = color.RGBA{60, 100, 20, 255}
)

var (
	black = color.RGBA{0, 0, 0, 255}
	red   = color.RGBA{255, 0, 0, 255}
	fruit = color.RGBA{100, 200, 0, 255}

	radius = 4
	tailH  = 9
	headH  = 5
)

//Generator struct to generate images
type Generator struct {
	Size  int
	ColNb int
	RowNb int
}

//NewGenerator returns new instance of generator
func NewGenerator(size, colnb, rownb int) *Generator {
	return &Generator{Size: size, ColNb: colnb, RowNb: rownb}
}

//GetBodyStraight returns body straight image
func (g *Generator) GetBodyStraight() *image.RGBA {
	dst := getBase(g.Size, g.Size)
	r := dst.Bounds()
	for i, c := range []color.RGBA{black, red} {
		r.Min.Y++
		r.Max.Y--
		img := getColorizedImage(g.Size, g.Size-(i+1), c)
		draw.Draw(dst, r, img, img.Bounds().Min, draw.Src)
	}
	return dst
}

//GetBodyCurve returns body curve image
func (g *Generator) GetBodyCurve() *image.RGBA {
	dst := getBase(g.Size, g.Size)

	// bound
	src := getColorizedImage(g.Size, g.Size, black)
	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{image.Point{radius + 1, radius + 1}, radius}, image.ZP, draw.Over)
	draw.Draw(dst, image.Rect(radius+1, 1, dst.Rect.Max.X, dst.Rect.Max.Y-1), src, src.Bounds().Min, draw.Src)
	draw.Draw(dst, image.Rect(1, radius+1, dst.Rect.Max.X-1, dst.Rect.Max.Y), src, src.Bounds().Min, draw.Src)

	// fill
	src = getColorizedImage(g.Size, g.Size, red)
	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{image.Point{radius + 2, radius + 2}, radius - 1}, image.ZP, draw.Over)
	draw.Draw(dst, image.Rect(radius+2, 2, dst.Rect.Max.X, dst.Rect.Max.Y-2), src, src.Bounds().Min, draw.Src)
	draw.Draw(dst, image.Rect(2, radius+2, dst.Rect.Max.X-2, dst.Rect.Max.Y), src, src.Bounds().Min, draw.Src)
	return dst
}

//GetBodyTail returns tail image
func (g *Generator) GetBodyTail() *image.RGBA {
	dst := getBase(g.Size, g.Size)

	// bound
	img := getColorizedImage(g.Size, g.Size, black)
	draw.DrawMask(dst, dst.Bounds(), img, image.ZP, &triangle{image.Point{int((float64(g.Size) / 2) - 0.5), 1}, tailH}, image.ZP, draw.Over)
	draw.Draw(dst, image.Rect(1, tailH, dst.Rect.Max.X-1, dst.Rect.Max.Y), img, img.Bounds().Min, draw.Src)

	//fill
	img = getColorizedImage(g.Size, g.Size, red)
	draw.DrawMask(dst, dst.Bounds(), img, image.ZP, &triangle{image.Point{int((float64(g.Size) / 2) - 0.5), 2}, tailH - 1}, image.ZP, draw.Over)
	draw.Draw(dst, image.Rect(2, tailH+1, dst.Rect.Max.X-2, dst.Rect.Max.Y), img, img.Bounds().Min, draw.Src)
	return dst
}

//GetBodyHead returns head image
func (g *Generator) GetBodyHead() *image.RGBA {
	dst := getBase(g.Size, g.Size)
	firstX := int((float64(g.Size) / 4) - 0.5)

	//bound
	img := getColorizedImage(g.Size, g.Size, black)
	draw.DrawMask(dst, dst.Bounds(), img, image.ZP, &triangle{image.Point{firstX + 1, 1}, headH}, image.ZP, draw.Over)
	draw.DrawMask(dst, dst.Bounds(), img, image.ZP, &triangle{image.Point{(firstX * 3) + 1, 1}, headH}, image.ZP, draw.Over)
	draw.Draw(dst, image.Rect(1, headH+1, dst.Rect.Max.X-1, dst.Rect.Max.Y), img, img.Bounds().Min, draw.Src)

	//fill
	img = getColorizedImage(g.Size, g.Size, red)
	draw.DrawMask(dst, dst.Bounds(), img, image.ZP, &triangle{image.Point{firstX + 1, 2}, headH - 1}, image.ZP, draw.Over)
	draw.DrawMask(dst, dst.Bounds(), img, image.ZP, &triangle{image.Point{(firstX * 3) + 1, 2}, headH - 1}, image.ZP, draw.Over)
	draw.Draw(dst, image.Rect(2, headH+1, dst.Rect.Max.X-2, dst.Rect.Max.Y), img, img.Bounds().Min, draw.Src)
	return dst
}

//GetFruit returns fruit image
func (g *Generator) GetFruit() *image.RGBA {
	dst := getBase(g.Size, g.Size)
	rad := int((float64(g.Size-2) / 2))
	src := getColorizedImage(g.Size, g.Size, black)
	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{image.Point{rad + 1, rad + 1}, rad - 2}, image.ZP, draw.Over)
	src = getColorizedImage(g.Size, g.Size, fruit)
	draw.DrawMask(dst, dst.Bounds(), src, image.ZP, &circle{image.Point{rad + 1, rad + 1}, rad - 3}, image.ZP, draw.Over)
	return dst
}

//GetSkin returns game skin image
func (g *Generator) GetSkin() *image.RGBA {
	gridWidth := g.Size * g.ColNb
	gridHeight := g.Size * g.RowNb
	dst := getColorizedImage(gridWidth+(Bound*2), gridHeight+34+(Bound*3), BoundColor)
	img := getColorizedImage(dst.Rect.Dx(), dst.Rect.Dy(), BackgroundColor)
	draw.Draw(dst, image.Rect(Bound, Bound, dst.Rect.Max.X-Bound, dst.Rect.Max.Y-Bound), img, img.Bounds().Min, draw.Src)
	draw.Draw(dst, image.Rect(0, Bound+34, dst.Rect.Max.X, 34+Bound*2), &image.Uniform{BoundColor}, image.ZP, draw.Src)
	return dst
}

func WriteFiles() {
	g := NewGenerator(20, 25, 50)
	writeFile("straight.png", g.GetBodyStraight())
	writeFile("curve.png", g.GetBodyCurve())
	writeFile("tail.png", g.GetBodyTail())
	writeFile("head.png", g.GetBodyHead())
	writeFile("fruit.png", g.GetFruit())
	writeFile("skin.png", g.GetSkin())
}

func getBase(width, heigth int) *image.RGBA {
	dst := image.NewRGBA(image.Rect(0, 0, width, heigth))
	draw.Draw(dst, dst.Bounds(), image.Transparent, image.ZP, draw.Src)
	return dst
}

func getColorizedImage(width, height int, color color.RGBA) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(img, img.Bounds(), &image.Uniform{color}, image.ZP, draw.Src)
	return img
}

func writeFile(name string, dst *image.RGBA) {
	f, err := os.Create(name)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, dst); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
