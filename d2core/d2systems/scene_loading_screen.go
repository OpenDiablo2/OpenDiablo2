package d2systems

import (
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	sceneKeyLoading = "Loading"
)

// NewLoadingScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewLoadingScene() *LoadingScene {
	scene := &LoadingScene{
		BaseScene: NewBaseScene(sceneKeyLoading),
	}

	return scene
}

// static check that LoadingScene implements the scene interface
var _ d2interface.Scene = &LoadingScene{}

// LoadingScene represents the game's main menu, where users can select single or multi player,
// or start the map engine test.
type LoadingScene struct {
	*BaseScene
	loadingSprite akara.EID
	filesToLoad   *akara.Subscription
	booted        bool
}

// Init the main menu scene
func (s *LoadingScene) Init(_ *akara.World) {
	s.Info("initializing ...")

	s.setupSubscription()
}

func (s *LoadingScene) setupSubscription() {
	filesToLoad := akara.NewFilter().
		Require(d2components.FileHandle).
		Forbid(d2components.FileSource). // but we forbid files that are already loaded
		Forbid(d2components.GameConfig).
		Forbid(d2components.StringTable).
		Forbid(d2components.DataDictionary).
		Forbid(d2components.Palette).
		Forbid(d2components.PaletteTransform).
		Forbid(d2components.Cof).
		Forbid(d2components.Dc6).
		Forbid(d2components.Dcc).
		Forbid(d2components.Ds1).
		Forbid(d2components.Dt1).
		Forbid(d2components.Wav).
		Forbid(d2components.AnimData).
		Build()

	s.filesToLoad = s.World.AddSubscription(filesToLoad)
}

func (s *LoadingScene) boot() {
	if !s.BaseScene.booted {
		return
	}

	s.createLoadingScreen()

	s.booted = true
}

func (s *LoadingScene) createLoadingScreen() {
	s.Info("creating loading screen")
	s.loadingSprite = s.Add.Sprite(0, 0, d2resource.LoadingScreen, d2resource.PaletteLoading)
}

// Update the main menu scene
func (s *LoadingScene) Update() {
	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	s.BaseScene.Update()
}
