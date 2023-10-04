package main

import "github.com/hajimehoshi/ebiten/v2"

type entity interface {
	Draw(screen *ebiten.Image)
	Step()
}

func creatEntity(entityType string, layer int, x, y float64) entity {
	var result entity
	switch entityType {
	case "Player":
		result = newPlayer(x, y)
	}
	return result
}
