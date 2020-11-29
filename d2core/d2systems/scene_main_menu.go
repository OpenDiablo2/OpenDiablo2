package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2input"
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
	booted  bool
	sprites struct {
		trademark      akara.EID
		mainBackground akara.EID
	}
}

// Init the main menu scene
func (s *MainMenuScene) Init(world *akara.World) {
	s.World = world

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

	imgPath := d2resource.GameSelectScreen
	palPath := d2resource.PaletteSky

	s.sprites.mainBackground = s.Add.SegmentedSprite(0, 0, imgPath, palPath, 4, 3, 0)
}

func (s *MainMenuScene) createButtons() {
	s.Info("creating buttons")
}

func (s *MainMenuScene) createTrademarkScreen() {
	s.Info("creating trademark screen")

	imgPath := d2resource.TrademarkScreen
	palPath := d2resource.PaletteSky

	s.sprites.trademark = s.Add.SegmentedSprite(0, 0, imgPath, palPath, 4, 3, 0)

	interactive := s.AddInteractive(s.sprites.trademark)

	interactive.InputVector.SetMouseButton(d2input.MouseButtonLeft)

	interactive.Callback = func() bool {
		s.Info("hiding trademark sprite")

		alpha, _ := s.GetAlpha(s.sprites.trademark)
		alpha.Alpha = 0

		interactive.Enabled = false

		return true // prevent propagation
	}
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
