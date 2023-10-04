package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/image/math/f64"
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
	//sprite     *ebiten.Image
	keys          []ebiten.Key
	hspeed        float64
	dir           int
	moveSpeed     float64
	currentSprite *Sprite
	Sprites       map[string]*Sprite
	subPixelX     float64
	subPixelY     float64
}

func newPlayer(x float64, y float64) *Player {
	//img, _, err := ebitenutil.NewImageFromFile("../assets/mori_jump1.png")
	//if err != nil {
	//	log.Fatal(err)
	//}
	keys := []ebiten.Key{}
	var ms float64 = 2
	s := newSprite("../assets/mori_idle_2.png", 5, 8, 64, 64, "mori_idle", f64.Vec2{-31, -23}, f64.Vec2{-40, -48})
	m := make(map[string]*Sprite)
	m[s.name] = s
	p := Player{x, y, keys, 0, 0, ms, s, m, 0, 0}
	return &p
}

func (p *Player) Draw(screen *ebiten.Image) {
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(math.Floor(p.translateX), math.Floor(p.translateY))
	//screen.DrawImage(p.sprite, op)
	p.currentSprite.draw(screen, p.translateX, p.translateY)
}

func (p *Player) Step() {
	p.currentSprite = p.Sprites["mori_idle"]
	p.currentSprite.step()
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
	targetSpeed := float64(p.dir) * p.moveSpeed
	p.hspeed = targetSpeed

	integerPart := int(p.hspeed)
	fractionalPart := p.hspeed - float64(integerPart)
	p.subPixelX += fractionalPart

	p.translateX += p.hspeed

	//fmt.Println(p.translateX)
	p.hspeed = 0
	rightdir = 0
	leftdir = 0
	p.keys = nil
}
