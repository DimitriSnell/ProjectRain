package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/math/f64"
)

var imgTranslateX float64
var imgTranslateY float64
var p Player
var entityList []entity

const mapPath = "../tiles/maps/tiles1..tmx"

var m map[uint32]*ebiten.Image
var scene *tiled.Map
var screenWidth = 960
var screenHeight = 530

func init() {
	fmt.Println("test")
	imgTranslateX = 0
	imgTranslateY = 0
	p = *newPlayer(0, 0, "./assets/mori_jump1.png")
	entityList = append(entityList, &p)

	gameMap, err := tiled.LoadFile(mapPath)
	if err != nil {
		fmt.Println("ERROR:")
		log.Fatal(err)
	}
	scene = gameMap
	fmt.Println(gameMap.ImageLayers[0])

	holdImg, _, err := ebitenutil.NewImageFromFile("../tiles/tilesets/keepers_forest_background_5.png")
	if err != nil {
		log.Fatal(err)
	}
	tilesImage := ebiten.NewImageFromImage(holdImg)
	fmt.Println(tilesImage)
	m = make(map[uint32]*ebiten.Image)

	//populate map of images

	for _, tile := range gameMap.Layers[0].Tiles {
		if tile.Nil == false {
			//xfmt.Println(gameMap.Layers[0].Name)
			spriteRect := tile.Tileset.GetTileRect(tile.ID)
			tileImage := tilesImage.SubImage(spriteRect).(*ebiten.Image)
			m[tile.ID] = tileImage
		}
	}
	p.translateX = 400
	p.translateY = 400
}

type Game struct {
	p      Player
	world  *ebiten.Image
	camera Camera
}

func (g *Game) Update() error {
	g.p.Step()

	g.camera.Target(g.p.translateX+float64(g.p.sprite.imgWidth/2), g.p.translateY+float64(g.p.sprite.imgHeight)/2)
	fmt.Println(g.camera.String())
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World")

	for i, tile := range scene.Layers[0].Tiles {
		if tile.Nil == false {
			op := &ebiten.DrawImageOptions{}
			//scene.TileWidth
			op.GeoM.Translate(float64((i%scene.Width)*scene.TileWidth), float64((i/scene.Width)*scene.TileHeight))
			//fmt.Println(i)
			//screen.DrawImage(m[tile.ID], op)
			g.world.DrawImage(m[tile.ID], op)
		}
	}
	//worldX, worldY := g.camera.ScreenToWorld(ebiten.CursorPosition())
	g.p.Draw(g.world)
	g.camera.Render(g.world, screen)
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
	g := &Game{camera: c}
	g.p = p
	g.world = ebiten.NewImage(scene.Width*scene.TileWidth, scene.Height*scene.TileHeight)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
