package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
)

type Scene struct {
	tileMap   *tiled.Map
	imageMaps []map[uint32]*ebiten.Image
}
