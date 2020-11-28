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

// static check that LoadingScene implements the scene interface
var _ d2interface.Scene = &LoadingScene{}

// NewLoadingScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewLoadingScene() *LoadingScene {
	scene := &LoadingScene{
		BaseScene: NewBaseScene(sceneKeyLoading),
	}

	return scene
}

// LoadingScene represents the game's main menu, where users can select single or multi player,
// or start the map engine test.
type LoadingScene struct {
	*BaseScene
	loadingSprite akara.EID
	filesToLoad   *akara.Subscription
	booted        bool
}

func (s *LoadingScene) setupSubscriptions() {
	s.Info("setting up component subscriptions")

	filesToLoad := s.NewComponentFilter().
		Require(
			&d2components.FileHandle{},
		).
		Forbid( // but we forbid files that are already loaded
			&d2components.FileSource{},
			&d2components.GameConfig{},
			&d2components.StringTable{},
			&d2components.DataDictionary{},
			&d2components.Palette{},
			&d2components.PaletteTransform{},
			&d2components.Cof{},
			&d2components.Dc6{},
			&d2components.Dcc{},
			&d2components.Ds1{},
			&d2components.Dt1{},
			&d2components.Wav{},
			&d2components.AnimationData{},
		).
		Build()

	s.filesToLoad = s.World.AddSubscription(filesToLoad)
}

// Init the main menu scene
func (s *LoadingScene) Init(world *akara.World) {
	s.World = world

	s.Info("initializing ...")

	s.setupSubscriptions()
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
