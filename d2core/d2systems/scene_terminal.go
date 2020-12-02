package d2systems

import (
	"image/color"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	sceneKeyTerminal = "Terminal"
)

// static check that TerminalScene implements the scene interface
var _ d2interface.Scene = &TerminalScene{}

// NewTerminalScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewTerminalScene() *TerminalScene {
	scene := &TerminalScene{
		BaseScene: NewBaseScene(sceneKeyTerminal),
	}

	return scene
}

// TerminalScene represents the game's loading screen, where loading progress is displayed
type TerminalScene struct {
	*BaseScene
	booted bool
}

// Init the loading scene
func (s *TerminalScene) Init(world *akara.World) {
	s.World = world

	s.Info("initializing ...")

	s.backgroundColor = color.Black

	s.setupFactories()
	s.setupSubscriptions()
}

func (s *TerminalScene) setupFactories() {
	texture := s.RegisterComponent(&d2components.Texture{})
	s.Texture = s.GetComponentFactory(texture)
}

func (s *TerminalScene) setupSubscriptions() {
	s.Info("setting up component subscriptions")

	//stage1 := s.NewComponentFilter().
	//	Require(
	//		&d2components.FilePath{},
	//	).
	//	Forbid( // but we forbid files that are already loaded
	//		&d2components.FileType{},
	//		&d2components.FileHandle{},
	//		&d2components.FileSource{},
	//		&d2components.GameConfig{},
	//		&d2components.StringTable{},
	//		&d2components.DataDictionary{},
	//		&d2components.Palette{},
	//		&d2components.PaletteTransform{},
	//		&d2components.Cof{},
	//		&d2components.Dc6{},
	//		&d2components.Dcc{},
	//		&d2components.Ds1{},
	//		&d2components.Dt1{},
	//		&d2components.Wav{},
	//		&d2components.AnimationData{},
	//	).
	//	Build()

	//s.loadStages.stage1 = s.World.AddSubscription(stage1)
}

func (s *TerminalScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	//s.createTerminalScene()

	s.booted = true
}

// Update the loading scene
func (s *TerminalScene) Update() {
	for _, id := range s.Viewports {
		s.AddPriority(id).Priority = scenePriorityLoading
	}

	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	s.BaseScene.Update()
}
