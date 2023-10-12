package main

import (
	"log"

	game "github.com/DimitriSnell/goTest/src/game"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/math/f64"
)

const mapPath = "../../tiles/maps/keepers_forest/keepers_1.tmx"

var screenWidth = 960
var screenHeight = 530

func init() {

	//p := creatEntity("Player", 0, 0, 0)
	//entityList = append(entityList, p)
	game.EntityMap = make(map[int]game.Entity)
}

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Hello, World!")
	c := game.Camera{ViewPort: f64.Vec2{float64(screenWidth), float64(screenHeight)}}
	g := &game.Game{}
	game.CreateEntityLayer(game.NewWall, 0, 20, 20)
	g.Camera = c
	g.SM = game.NewSceneManager(g)
	g.PD = game.NewPlayerData()
	g.PD.SetPlayerClass(game.CLASS(game.MORI))
	g.SM.LoadSceneSpecific(game.KEEPERSFOREST1)
	//g.SM
	//g.Level = game.NewKeepersForest1(g)
	//g.World = ebiten.NewImage(g.Level.GetTileMap().Width*g.Level.GetTileMap().TileWidth, g.Level.GetTileMap().Height*g.Level.GetTileMap().TileHeight)
	//g.M = m
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
