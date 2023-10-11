package game

import (
	"fmt"
	"reflect"

	util "github.com/DimitriSnell/goTest/src/utils"
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
	vspeed        float64
	dir           int
	moveSpeed     float64
	currentSprite *Sprite
	Sprites       map[string]*Sprite
	subPixelX     float64
	subPixelY     float64
	gravity       float32
	UID           int
}

func NewPlayer(x float64, y float64, UID int) Entity {
	//img, _, err := ebitenutil.NewImageFromFile("../assets/mori_jump1.png")
	//if err != nil {
	//	log.Fatal(err)
	//}
	keys := []ebiten.Key{}
	var ms float64 = 2
	s := NewSprite("../../assets/mori_idle_2.png", 5, 8, 64, 64, "mori_idle", f64.Vec2{31, 23}, f64.Vec2{40, 48})
	m := make(map[string]*Sprite)
	m[s.Name] = s
	//m[] = s
	gv := .2
	p := Player{x, y, keys, 0, 0, 0, ms, s, m, 0, 0, float32(gv), UID}
	return &p
}

func (p *Player) Draw(screen *ebiten.Image) {
	//op := &ebiten.DrawImageOptions{}
	//op.GeoM.Translate(math.Floor(p.translateX), math.Floor(p.translateY))
	//screen.DrawImage(p.sprite, op)
	p.currentSprite.Draw(screen, p.translateX, p.translateY)
}

func (p *Player) Step() {
	p.currentSprite = p.Sprites["mori_idle"]
	p.currentSprite.Step()
	p.keys = inpututil.AppendPressedKeys(p.keys)

	rightdir := 0
	leftdir := 0
	if contains(p.keys, ebiten.KeyArrowRight) {
		rightdir = 1
	}
	if contains(p.keys, ebiten.KeyArrowLeft) {
		leftdir = -1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		p.vspeed = -6
	}
	//placeMeeting(p, int(p.translateX), int(p.translateY), reflect.TypeOf(&Wall{}))
	//left and right movement
	p.dir = rightdir + leftdir
	targetSpeed := float64(p.dir) * p.moveSpeed
	p.hspeed = targetSpeed

	p.vspeed += float64(p.gravity)
	//collision and vertical/horizontal movement
	p.playerWallCollision()

	p.translateX += p.hspeed

	p.translateY += p.vspeed
	p.hspeed = 0
	rightdir = 0
	leftdir = 0
	p.keys = nil
}

func (p *Player) GetPosition() (x float64, y float64) {
	return p.translateX, p.translateY
}

func (p *Player) GetCurrentSprite() *Sprite {
	return p.currentSprite
}

func (p *Player) getHspeed() float64 {
	return p.hspeed
}

func (p *Player) getVspeed() float64 {
	return p.vspeed
}

func (p *Player) playerWallCollision() {
	if (placeMeeting(p, p.translateX+p.hspeed, p.translateY, reflect.TypeOf(&Wall{}))) {
		fmt.Println("TEST COLLISION")
		for !placeMeeting(p, p.translateX+float64(util.Sign(p.vspeed)), p.translateY, reflect.TypeOf(&Wall{})) {
			p.translateX += float64(util.Sign(p.hspeed))
		}
		p.hspeed = 0
	}

	if (placeMeeting(p, p.translateX, p.translateY+p.vspeed, reflect.TypeOf(&Wall{}))) {
		//fmt.Println("TEST COLLISION")
		for !placeMeeting(p, p.translateX, p.translateY+float64(util.Sign(p.vspeed)), reflect.TypeOf(&Wall{})) {
			fmt.Println(p.translateY)
			p.translateY += float64(util.Sign(p.vspeed))
		}
		p.vspeed = 0
	}
}

func (p *Player) getUID() int {

	return p.UID
}
