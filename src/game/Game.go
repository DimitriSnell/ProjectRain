package game

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

var screenWidth = 960
var screenHeight = 530
var EntityList []Entity

type Game struct {
	World  *ebiten.Image
	Camera Camera
	Scene  *tiled.Map
	M      []map[uint32]*ebiten.Image
}

func (g *Game) CreateEntityLayer(entityC EntityConstructor, layer int, x, y float64) {

	result := entityC(x, y)

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
	g.World.Fill(color.RGBA{149, 209, 197, 0xff})
	fmt.Println(len(g.Scene.Layers))
	for c, layer := range g.Scene.Layers {
		for i, tile := range layer.Tiles {
			if tile.Nil == false {
				//frustum culling
				if float64((i%g.Scene.Width)*g.Scene.TileWidth) > g.Camera.Position[0]-32 && float64((i%g.Scene.Width)*g.Scene.TileWidth) < g.Camera.Position[0]+g.Camera.ViewPort[0] &&
					float64((i/g.Scene.Width)*g.Scene.TileHeight) > g.Camera.Position[1]-32 && float64((i/g.Scene.Width)*g.Scene.TileHeight) < g.Camera.Position[1]+g.Camera.ViewPort[1] {
					op := &ebiten.DrawImageOptions{}
					//scene.TileWidth
					op.GeoM.Translate(float64((i%g.Scene.Width)*g.Scene.TileWidth), float64((i/g.Scene.Width)*g.Scene.TileHeight))
					//fmt.Println(i)
					//screen.DrawImage(m[tile.ID], op)
					g.World.DrawImage(g.M[c][tile.ID], op)
				}
			}
		}
	}

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
