package main

import (
	"fmt"
	"log"

	game "github.com/DimitriSnell/goTest/src/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
	"golang.org/x/image/math/f64"
)

// var p entity

const mapPath = "../../tiles/maps/keepers_forest/keepers_1.tmx"

var globalUIDCounter uint32 = 0
var m []map[uint32]*ebiten.Image
var scene *tiled.Map
var screenWidth = 960
var screenHeight = 530

func init() {

	//p := creatEntity("Player", 0, 0, 0)
	//entityList = append(entityList, p)
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
		holdImg, _, err := ebitenutil.NewImageFromFile("../../tiles/tilesets/" + layerImageNames[i])
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

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Hello, World!")
	//w := ebiten.NewImage(scene.Width, scene.Height)
	c := game.Camera{ViewPort: f64.Vec2{float64(screenWidth), float64(screenHeight)}}
	g := &game.Game{}
	g.CreateEntityLayer(game.NewPlayer, 0, 0, 0)
	g.CreateEntityLayer(game.NewWall, 0, 20, 20)
	fmt.Println(len(game.EntityList))
	for _, e := range game.EntityList {
		//fmt.Println("TESTETSETESTESTE")
		fmt.Printf("Type: %T\n", e)
		if player, ok := e.(*game.Player); ok {
			c.SetTarget(player)
		}
	}
	g.Camera = c
	g.World = ebiten.NewImage(scene.Width*scene.TileWidth, scene.Height*scene.TileHeight)
	g.Scene = scene
	g.M = m
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
