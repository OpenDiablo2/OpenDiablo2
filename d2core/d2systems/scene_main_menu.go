package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	sceneKeyMainMenu = "Main Menu"
)

// NewMainMenuScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewMainMenuScene() *MainMenuScene {
	scene := &MainMenuScene{
		BaseScene: NewBaseScene(sceneKeyMainMenu),
	}

	return scene
}

// static check that MainMenuScene implements the scene interface
var _ d2interface.Scene = &MainMenuScene{}

// MainMenuScene represents the game's main menu, where users can select single or multi player,
// or start the map engine test.
type MainMenuScene struct {
	*BaseScene
	booted bool
}

// Init the main menu scene
func (s *MainMenuScene) Init(_ *akara.World) {
	s.Info("initializing ...")
}

func (s *MainMenuScene) boot() {
	if !s.BaseScene.booted {
		return
	}

	s.createBackground()
	s.createButtons()
	s.createTrademarkScreen()

	s.booted = true
}

func (s *MainMenuScene) createBackground() {
	s.Info("creating background")
	s.Add.Sprite(0, 0, d2resource.GameSelectScreen, d2resource.PaletteSky)
}

func (s *MainMenuScene) createButtons() {
	s.Info("creating buttons")
}

func (s *MainMenuScene) createTrademarkScreen() {
	s.Info("creating trademark screen")
	s.Add.Sprite(0, 0, d2resource.TrademarkScreen, d2resource.PaletteSky)
}

// Update the main menu scene
func (s *MainMenuScene) Update() {
	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	s.BaseScene.Update()
}
