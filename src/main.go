package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/math/f64"
)

// var p entity
var entityList []entity

const mapPath = "../tiles/maps/keepers_forest/keepers_1.tmx"

var globalUIDCounter uint32 = 0
var m []map[uint32]*ebiten.Image
var scene *tiled.Map
var screenWidth = 960
var screenHeight = 530

func init() {

	p := creatEntity("Player", 0, 0, 0)
	entityList = append(entityList, p)

	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Println("ERROR:")
		log.Fatal(err)
	}
	scene = gameMap
	//fmt.Println(gameMap.ImageLayers[0])

	//populate map of images
	layerImageNames := make(map[int]string)
	layerImageNames[0] = "Keeper_forest_4.png"
	layerImageNames[1] = "Keepers_forest_9.png"
	layerImageNames[2] = "Keeper_forst_3.png"
	layerImageNames[3] = "Keepers_forest_8.png"
	layerImageNames[4] = "Keepers_forest_2.png"
	layerImageNames[5] = "Keeper_forest_6.png"
	layerImageNames[6] = "Keepers_forest_1.png"
	layerImageNames[7] = "Keepers_forest_7.png"
	for i, _ := range gameMap.Layers {
		temp := make(map[uint32]*ebiten.Image)
		m = append(m, temp)
		fmt.Println(gameMap.Layers[i].Name)
		holdImg, _, err := ebitenutil.NewImageFromFile("../tiles/tilesets/" + layerImageNames[i])
		if err != nil {
			log.Fatal(err)
		}
		tilesImage := ebiten.NewImageFromImage(holdImg)

		for _, tile := range gameMap.Layers[i].Tiles {
			if tile.Nil == false {
				//xfmt.Println(gameMap.Layers[0].Name)
				spriteRect := tile.Tileset.GetTileRect(tile.ID)
				tileImage := tilesImage.SubImage(spriteRect).(*ebiten.Image)
				m[i][tile.ID] = tileImage
			}
		}
	}

}

type Game struct {
	world  *ebiten.Image
	camera Camera
}

func (g *Game) Update() error {
	for _, entity := range entityList {
		entity.Step()
	}
	g.camera.Target()
	//fmt.Println(ebiten.CurrentTPS())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World")
	g.world.Fill(color.RGBA{149, 209, 197, 0xff})
	for c, layer := range scene.Layers {
		for i, tile := range layer.Tiles {
			if tile.Nil == false {
				//frustum culling
				if float64((i%scene.Width)*scene.TileWidth) > g.camera.Position[0]-32 && float64((i%scene.Width)*scene.TileWidth) < g.camera.Position[0]+g.camera.ViewPort[0] &&
					float64((i/scene.Width)*scene.TileHeight) > g.camera.Position[1]-32 && float64((i/scene.Width)*scene.TileHeight) < g.camera.Position[1]+g.camera.ViewPort[1] {
					op := &ebiten.DrawImageOptions{}
					//scene.TileWidth
					op.GeoM.Translate(float64((i%scene.Width)*scene.TileWidth), float64((i/scene.Width)*scene.TileHeight))
					//fmt.Println(i)
					//screen.DrawImage(m[tile.ID], op)
					g.world.DrawImage(m[c][tile.ID], op)
				}
			}
		}
	}

	//worldX, worldY := g.camera.ScreenToWorld(ebiten.CursorPosition())
	for _, entity := range entityList {
		entity.Draw(g.world)
	}
	g.camera.Render(g.world, screen)
	//ebitenutil.DebugPrint(screen, ebiten.CurrentTPS())
	g.world.Clear()
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Hello, World!")
	//w := ebiten.NewImage(scene.Width, scene.Height)
	c := Camera{ViewPort: f64.Vec2{float64(screenWidth), float64(screenHeight)}}
	fmt.Println(len(entityList))
	for _, e := range entityList {
		//fmt.Println("TESTETSETESTESTE")
		fmt.Printf("Type: %T\n", e)
		if player, ok := e.(*Player); ok {
			c.target = player
		}
	}

	g := &Game{camera: c}
	g.world = ebiten.NewImage(scene.Width*scene.TileWidth, scene.Height*scene.TileHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
