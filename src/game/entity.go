package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity interface {
	Draw(screen *ebiten.Image)
	Step()
	GetPosition() (x float64, y float64)
	GetCurrentSprite() *Sprite
	getHspeed() float64
	getVspeed() float64
}
type EntityConstructor func(x, y float64) Entity

/*func CreatEntity(entityC EntityConstructor, layer int, x, y float64) {
/*var result entity
switch entityType {
case "Player":
	result = newPlayer(x, y)
	entityList = append(entityList, result)
}*/

//	result := entityC(x, y)

//	EntityList = append(EntityList, result)
//}
