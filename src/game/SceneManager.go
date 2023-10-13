package game

import "github.com/hajimehoshi/ebiten/v2"

type SCENENAME int

const (
	MAINMENU1 SCENENAME = iota
	KEEPERSFOREST1
)

type SceneManager struct {
	currentSceneIndex SCENENAME
	currentScene      *Scene
	g                 *Game
}

func NewSceneManager(g *Game) *SceneManager {
	result := &SceneManager{}
	result.g = g

	g.EM.Subscribe("SceneTransition", result.LoadSceneSpecific)
	result.LoadSceneSpecific(MAINMENU1)
	return result

}

func (sm *SceneManager) LoadSceneSpecific(index SCENENAME) {
	clearScene()
	switch index {
	case MAINMENU1:
		sm.currentScene = NewMainMenu1(sm.g)
		sm.currentSceneIndex = MAINMENU1
	case KEEPERSFOREST1:
		AssetsLoaded = false
		go func() {
			sm.currentScene = NewKeepersForest1(sm.g)
			AssetsLoaded = true
		}()
		sm.currentSceneIndex = KEEPERSFOREST1
	}
	sm.g.World = ebiten.NewImage(sm.currentScene.tileMap.Width*sm.currentScene.tileMap.TileWidth, sm.currentScene.tileMap.Height*sm.currentScene.tileMap.TileHeight)
}

func clearScene() {
	for _, e := range EntityList {
		DestroyEntity(e.getUID())
	}
}
