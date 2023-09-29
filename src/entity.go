package main

import "github.com/hajimehoshi/ebiten/v2"

type entity interface {
	Draw(screen *ebiten.Image)
	Step()
}
