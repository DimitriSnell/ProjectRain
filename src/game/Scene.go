package game

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

type Scene struct {
	tileMap   *tiled.Map
	imageMaps []map[uint32]*ebiten.Image
}

func (s *Scene) DrawTiles(g *Game) {
	g.World.Fill(color.RGBA{149, 209, 197, 0xff})
	fmt.Println("LAYER LENGTH!!!!!!!!!!!!")
	fmt.Println(len(s.tileMap.Layers))
	fmt.Println("MAP LENGTH!!!!!!!!!!!!")
	fmt.Println(len(s.imageMaps))
	for c, layer := range s.tileMap.Layers {
		for i, tile := range layer.Tiles {
			if tile.Nil == false {
				//frustum culling
				if float64((i%s.tileMap.Width)*s.tileMap.TileWidth) > g.Camera.Position[0]-32 && float64((i%s.tileMap.Width)*s.tileMap.TileWidth) < g.Camera.Position[0]+g.Camera.ViewPort[0] &&
					float64((i/s.tileMap.Width)*s.tileMap.TileHeight) > g.Camera.Position[1]-32 && float64((i/s.tileMap.Width)*s.tileMap.TileHeight) < g.Camera.Position[1]+g.Camera.ViewPort[1] {
					op := &ebiten.DrawImageOptions{}
					//scene.TileWidth
					op.GeoM.Translate(float64((i%s.tileMap.Width)*s.tileMap.TileWidth), float64((i/s.tileMap.Width)*s.tileMap.TileHeight))
					//fmt.Println(i)
					//screen.DrawImage(m[tile.ID], op)
					g.World.DrawImage(s.imageMaps[c][tile.ID], op)
				}
			}
		}
	}
}

func NewKeepersForest1() *Scene {
	result := Scene{}
	result.imageMaps = make([]map[uint32]*ebiten.Image, 0)
	gameMap, err := tiled.LoadFile("../../tiles/maps/keepers_forest/keepers_1.tmx")
	if err != nil {
		fmt.Println("ERROR:")
		log.Fatal(err)
	}
	result.tileMap = gameMap
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
		if gameMap.Layers[i].Name == "オブジェクトレイヤー1" {
			for _, object := range gameMap.Layers[i].Tiles {
				fmt.Println(gameMap.Layers[i].GetTilePosition(int(object.ID)))
			}
			break
		}
		temp := make(map[uint32]*ebiten.Image)
		result.imageMaps = append(result.imageMaps, temp)
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
				result.imageMaps[i][tile.ID] = tileImage
			}
		}
	}
	return &result
}
