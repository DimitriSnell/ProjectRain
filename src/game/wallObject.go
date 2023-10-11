package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

type Wall struct {
	x      float64
	y      float64
	sprite *Sprite
	UID    int
}

func NewWall(x, y float64, UID int) Entity {

	s := NewSprite("", 1, 1, 0, 0, "wall", f64.Vec2{0, 0}, f64.Vec2{32, 32})
	w := Wall{x, y, s, UID}
	return &w
}

func (w *Wall) Draw(screen *ebiten.Image) {
	w.sprite.Draw(screen, w.x, w.y)
}

func (w *Wall) Step() {
	w.sprite.Step()
}

func (w *Wall) GetPosition() (x float64, y float64) {
	return w.x, w.y
}

func (w *Wall) GetCurrentSprite() *Sprite {
	return w.sprite
}

func (w *Wall) getHspeed() float64 {
	return 0
}

func (w *Wall) getVspeed() float64 {
	return 0
}

func (w *Wall) getUID() int {
	return w.UID
}
