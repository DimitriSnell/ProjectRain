package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var screenWidth = 960
var screenHeight = 530
var EntityList []Entity
var globalUIDCounter = 0

type Game struct {
	World  *ebiten.Image
	Camera Camera
	Level  *Scene
	M      []map[uint32]*ebiten.Image
}

func CreateEntityLayer(entityC EntityConstructor, layer int, x, y float64) {

	result := entityC(x, y, globalUIDCounter)
	globalUIDCounter++
	EntityList = append(EntityList, result)
}

func (g *Game) Update() error {
	for _, entity := range EntityList {
		entity.Step()
	}
	g.Camera.Target()
	//fmt.Println(ebiten.CurrentTPS())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World")
	//g.World.Fill(color.RGBA{149, 209, 197, 0xff})

	g.Level.DrawTiles(g)

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
