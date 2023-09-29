package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func contains(s []ebiten.Key, e ebiten.Key) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Player struct {
	translateX float64
	translateY float64
	sprite     *ebiten.Image
	keys       []ebiten.Key
	hspeed     float64
	dir        int
	moveSpeed  float64
}

func newPlayer(x float64, y float64, imgFile string) *Player {
	img, _, err := ebitenutil.NewImageFromFile("./assets/mori_jump1.png")
	if err != nil {
		log.Fatal(err)
	}
	keys := []ebiten.Key{}
	var ms float64 = 1.5
	p := Player{x, y, img, keys, 0, 0, ms}
	return &p
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.translateX, p.translateY)
	screen.DrawImage(p.sprite, op)
}

func (p *Player) Step() {

	p.keys = inpututil.AppendPressedKeys(p.keys)
	rightdir := 0
	leftdir := 0
	if contains(p.keys, ebiten.KeyArrowRight) {
		rightdir = 1
	}
	if contains(p.keys, ebiten.KeyArrowLeft) {
		leftdir = -1
	}
	//left and right movement
	p.dir = rightdir + leftdir
	p.hspeed = float64(p.dir) * p.moveSpeed

	p.translateX += p.hspeed
	p.hspeed = 0
	rightdir = 0
	leftdir = 0
	p.keys = nil
}
