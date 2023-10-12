package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var screenWidth = 960
var screenHeight = 530
var EntityList []Entity
var EntityMap map[int]Entity
var globalUIDCounter = 0

type Game struct {
	World  *ebiten.Image
	Camera Camera
	SM     *SceneManager
	M      []map[uint32]*ebiten.Image
	PD     *PlayerData
}

func CreateEntityLayer(entityC EntityConstructor, layer int, x, y float64) {
	result := entityC(x, y, globalUIDCounter)
	EntityList = append(EntityList, result)
	EntityMap[result.getUID()] = result
	globalUIDCounter++
}

func InstanceExists(UID int) bool {
	_, ok := EntityMap[UID]
	return ok
}

// searches for entity and deletes returns true if successful
func DestroyEntity(UID int) bool {

	result := false

	if !InstanceExists(UID) {
		return result
	}
	result = true
	fmt.Println("TEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEST")
	fmt.Println(len(EntityMap), len(EntityList))
	delete(EntityMap, UID)
	for i, entity := range EntityList {
		if entity.getUID() == UID && i != len(EntityList)-1 {
			EntityList = append(EntityList[:i], EntityList[i+1])
		} else if entity.getUID() == UID {
			EntityList = EntityList[:len(EntityList)-1]
		}

	}
	fmt.Println("TEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEST")
	fmt.Println(len(EntityMap), len(EntityList))
	return result
}

func (g *Game) Update() error {
	for _, entity := range EntityList {
		entity.Step()
	}
	g.Camera.Target()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World")
	g.SM.currentScene.DrawTiles(g)

	//worldX, worldY := g.camera.ScreenToWorld(ebiten.CursorPosition())
	for _, entity := range EntityList {
		entity.Draw(g.World)
	}
	g.Camera.Render(g.World, screen)
	//ebitenutil.DebugPrint(screen, ebiten.CurrentTPS())
	g.World.Clear()
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
