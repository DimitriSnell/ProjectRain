package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var imgTranslateX float64
var imgTranslateY float64
var p Player

func init() {
	fmt.Println("test")
	imgTranslateX = 0
	imgTranslateY = 0
	p = *newPlayer(0, 0, "./assets/mori_jump1.png")

}

type Game struct {
	p Player
}

func (g *Game) Update() error {
	p.Step()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World")
	p.Draw(screen)
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	g := Game{p}
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
