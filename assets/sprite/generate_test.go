package sprite

import (
	"image"
	"image/png"
	"log"
	"os"
	"testing"
)

func TestWrite(t *testing.T) {
	g := NewGenerator(20, 50, 25)
	writeFile("straight.png", g.GetBodyStraight())
	writeFile("curve.png", g.GetBodyCurve())
	writeFile("tail.png", g.GetBodyTail())
	writeFile("head.png", g.GetBodyHead())
	writeFile("fruit.png", g.GetFruit())
	writeFile("skin.png", g.GetSkin())
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
