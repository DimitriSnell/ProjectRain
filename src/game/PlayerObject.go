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

type PLAYERSTATE int

const (
	FREE PLAYERSTATE = iota
	ABILITY1
)

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
	STATE         PLAYERSTATE
}

func NewPlayer(x float64, y float64, UID int) Entity {
	keys := []ebiten.Key{}
	var ms float64 = 2
	s := NewSprite("../../assets/mori_idle_2.png", 5, 8, 64, 64, "mori_idle", f64.Vec2{31, 23}, f64.Vec2{40, 48}, 0, 0)
	m := make(map[string]*Sprite)
	m[s.Name] = s
	//m[] = s
	gv := .2
	p := Player{x, y, keys, 0, 0, 0, ms, s, m, 0, 0, float32(gv), UID, FREE}
	return &p
}

func NewMori(x float64, y float64, UID int) Entity {
	keys := []ebiten.Key{}
	var ms float64 = 2
	s := NewSprite("../../assets/mori_idle_2.png", 5, 8, 64, 64, "mori_idle", f64.Vec2{31, 23}, f64.Vec2{40, 48}, 0, 0)
	m := make(map[string]*Sprite)
	m[s.Name] = s
	s2 := NewSprite("../../assets/mori_swing_1_strip10.png", 10, 4, 128, 128, "mori_ability_1", f64.Vec2{63, 55}, f64.Vec2{72, 80}, 32, 32)
	m[s2.Name] = s2
	s2 = NewSprite("../../assets/mori_run_sheet.png", 6, 8, 64, 64, "mori_run", f64.Vec2{31, 23}, f64.Vec2{40, 48}, 0, 0)
	m[s2.Name] = s2
	//m[] = s
	gv := .2
	p := Player{x, y, keys, 0, 0, 0, ms, s, m, 0, 0, float32(gv), UID, FREE}
	return &p
}

func (p *Player) Draw(screen *ebiten.Image) {
	p.currentSprite.Draw(screen, p.translateX, p.translateY)
}

func (p *Player) Step() {
	//p.currentSprite = p.Sprites["mori_idle"]
	p.currentSprite.Step()
	//p.keys = inpututil.AppendPressedKeys(p.keys)
	if p.hspeed > 0 {
		p.dir = 1
	} else if p.hspeed < 0 {
		p.dir = -1
	}
	rightdir := 0
	leftdir := 0
	switch p.STATE {
	case FREE:
		p.currentSprite = p.Sprites["mori_idle"]
		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			rightdir = 1
		}
		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			leftdir = -1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			p.vspeed = -6
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			p.STATE = ABILITY1
		}
		targetSpeed := float64(rightdir+leftdir) * p.moveSpeed
		p.hspeed = targetSpeed
		if p.hspeed > 0 || p.hspeed < 0 {
			p.currentSprite = p.Sprites["mori_run"]
		}
	case ABILITY1:
		p.hspeed = 0
		p.currentSprite = p.Sprites["mori_ability_1"]
		if p.currentSprite.spriteIndex >= p.currentSprite.numOfFrames-1 {
			p.currentSprite.count = 0
			p.STATE = FREE
		}
	}

	p.vspeed += float64(p.gravity)
	//fmt.Println(p.dir)
	//collision and vertical/horizontal movement
	p.playerWallCollision()
	if p.dir > 0 {
		p.currentSprite.imageXScale = 1
	} else if p.dir < 0 {

		p.currentSprite.imageXScale = -1
	}
	p.translateX += p.hspeed
	p.translateY += p.vspeed
	//p.hspeed = 0
	//rightdir = 0
	//leftdir = 0
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
		for !placeMeeting(p, p.translateX+float64(util.Sign(p.hspeed)), p.translateY, reflect.TypeOf(&Wall{})) {
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
