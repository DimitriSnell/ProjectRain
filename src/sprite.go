package main

import (
	"image"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Sprite struct {
	img             *ebiten.Image
	numOfFrames     int
	framesPerSecond int
	imgWidth        int
	imgHeight       int
	name            string
	count           int
}

func newSprite(filename string, numOfFrames int, framesPerSecond int, imgWidth int, imgHeigt int, name string) *Sprite {
	img, _, err := ebitenutil.NewImageFromFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	s := Sprite{img, numOfFrames, framesPerSecond, imgWidth, imgHeigt, name, 0}
	return &s
}

func (s *Sprite) step() {

	s.count++
}

func (s *Sprite) draw(img *ebiten.Image, x, y float64) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(math.Floor(x), math.Floor(y))
	i := (s.count / int(s.framesPerSecond)) % int(s.numOfFrames)
	//fmt.Println(i)
	sx, sy := i*int(s.imgWidth), 0
	img.DrawImage(s.img.SubImage(image.Rect(sx, sy, sx+int(s.imgWidth), sy+int(s.imgHeight))).(*ebiten.Image), op)
	//img.DrawImage(s.img.SubImage(image.Rect(sx, sy, sx+64, sy+64)).(*ebiten.Image), op)

}
