package game

import (
	"fmt"
	"image/color"
	"log"
	"math"

	util "github.com/DimitriSnell/goTest/src/utils"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/lafriks/go-tiled"
)

type Scene struct {
	tileMap   *tiled.Map
	imageMaps []map[uint32]*ebiten.Image
	ui        *ebitenui.UI
	PlayerRef *Player
	Parallax  []float64
}

func (s *Scene) GetTileMap() *tiled.Map {
	return s.tileMap
}

func (s *Scene) DrawTiles(g *Game) {
	g.World.Fill(color.RGBA{149, 209, 197, 0xff})
	//fmt.Println("LAYER LENGTH!!!!!!!!!!!!")
	//fmt.Println(len(s.tileMap.Layers))
	//fmt.Println("MAP LENGTH!!!!!!!!!!!!")
	//fmt.Println(len(s.imageMaps))

	for c, layer := range s.tileMap.Layers {
		parallaxFactor := 1.0
		if s.PlayerRef != nil && InstanceExists(s.PlayerRef.getUID()) {
			parallaxFactor = (math.Floor(s.PlayerRef.translateX) / s.Parallax[c])
		}
		for i, tile := range layer.Tiles {
			if tile.Nil == false {
				//frustum culling
				if float64((i%s.tileMap.Width)*s.tileMap.TileWidth)+parallaxFactor > g.Camera.Position[0]-32 && float64((i%s.tileMap.Width)*s.tileMap.TileWidth)+parallaxFactor < g.Camera.Position[0]+g.Camera.ViewPort[0] &&
					float64((i/s.tileMap.Width)*s.tileMap.TileHeight) > g.Camera.Position[1]-32 && float64((i/s.tileMap.Width)*s.tileMap.TileHeight) < g.Camera.Position[1]+g.Camera.ViewPort[1] {
					op := &ebiten.DrawImageOptions{}
					op.Filter = ebiten.FilterNearest
					if s.PlayerRef != nil && InstanceExists(s.PlayerRef.getUID()) && math.Abs(s.PlayerRef.hspeed) > 0 {
						op.Filter = ebiten.FilterLinear
					}
					op.GeoM.Translate((float64((i%s.tileMap.Width)*s.tileMap.TileWidth) + parallaxFactor), float64((i/s.tileMap.Width)*s.tileMap.TileHeight))
					g.World.DrawImage(s.imageMaps[c][tile.ID], op)
				}
			}
		}
	}

	if s.ui != nil {
		s.ui.Draw(g.World)
	}
}

func (s *Scene) step() {
	if s.ui != nil {
		s.ui.Update()
	}

}

func NewMainMenu1(g *Game) *Scene {
	result := Scene{}
	result.imageMaps = make([]map[uint32]*ebiten.Image, 0)
	gameMap, err := tiled.LoadFile("../../tiles/maps/main_menu_1.tmx")
	if err != nil {
		fmt.Println("ERROR:")
		log.Fatal(err)
	}
	result.tileMap = gameMap
	layerImageNames := make(map[int]string)
	layerImageNames[0] = "Keeper_forest_4.png"
	for i, _ := range gameMap.Layers {
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

	//create UI
	buttonImage, _ := util.LoadButtonImage()
	rootContainer := widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image.NewNineSliceColor(color.NRGBA{0x13, 0x1a, 0x22, 0xff})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(2),
			widget.GridLayoutOpts.Padding(widget.NewInsetsSimple(30)),
			widget.GridLayoutOpts.Spacing(20, 20),
		)),
	)
	//mori class select
	button := widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),

		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("button clicked")
			fmt.Println("CREATE MORI")
			g.PD.class = CLASS(MORI)
			g.EM.Publish("SceneTransition", KEEPERSFOREST1)
		}),
	)
	rootContainer.AddChild(button)
	//lui class select
	button2 := widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),

		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
			println("button clicked")
			fmt.Println("CREATE LUI")
			g.PD.class = CLASS(LUI)
			g.EM.Publish("SceneTransition", KEEPERSFOREST1)
		}),
	)
	rootContainer.AddChild(button2)
	result.ui = &ebitenui.UI{Container: rootContainer}
	//face, _ := loadFont(20)
	AssetsLoaded = true
	return &result
}

func NewKeepersForest1(g *Game) *Scene {
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
	pxList := []float64{1, 1, 1.2, 1.2, 1.8, 1.8, 2.5, 2.5}
	result.Parallax = append(result.Parallax, pxList...)
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
	result.loadObjects()

	//create player object and set camera target
	result.SpawnPlayer(g.PD)
	for _, e := range EntityList {
		fmt.Printf("Type: %T\n", e)
		if player, ok := e.(*Player); ok {
			result.PlayerRef = player
			g.Camera.SetTarget(player)
		}
	}
	return &result
}

func (s *Scene) loadObjects() {

	//fmt.Println(gameMap.ImageLayers[0])

	for _, groups := range s.tileMap.ObjectGroups {
		for _, object := range groups.Objects {
			CreateEntityLayer(NewWall, 0, object.X, object.Y)
		}
	}

}

func (s *Scene) SpawnPlayer(pd *PlayerData) {
	switch pd.class {
	case MORI:
		CreateEntityLayer(NewMori, 0, 100, 0)
	case LUI:
		CreateEntityLayer(NewLui, 0, 100, 0)
	}
}
