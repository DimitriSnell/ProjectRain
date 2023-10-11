package main

import (
	"fmt"
	"log"

	game "github.com/DimitriSnell/goTest/src/game"
	"github.com/hajimehoshi/ebiten/v2"
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

	for _, groups := range gameMap.ObjectGroups {
		for _, object := range groups.Objects {
			game.CreateEntityLayer(game.NewWall, 0, object.X, object.Y)
		}
	}
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Hello, World!")
	//w := ebiten.NewImage(scene.Width, scene.Height)
	c := game.Camera{ViewPort: f64.Vec2{float64(screenWidth), float64(screenHeight)}}
	g := &game.Game{}
	game.CreateEntityLayer(game.NewPlayer, 0, 100, 20)
	game.CreateEntityLayer(game.NewWall, 0, 20, 20)
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
	g.Level = game.NewKeepersForest1()
	//g.M = m
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
