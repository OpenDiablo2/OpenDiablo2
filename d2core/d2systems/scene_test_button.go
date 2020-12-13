package d2systems

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2button"
	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	sceneKeyButtonTest = "Button Test Scene"
)

// NewButtonTestScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewButtonTestScene() *ButtonTestScene {
	scene := &ButtonTestScene{
		BaseScene: NewBaseScene(sceneKeyButtonTest),
	}

	return scene
}

// static check that ButtonTestScene implements the scene interface
var _ d2interface.Scene = &ButtonTestScene{}

// ButtonTestScene represents the game's main menu, where users can select single or multi player,
// or start the map engine test.
type ButtonTestScene struct {
	*BaseScene
	booted   bool
	buttons *akara.Subscription
}

// Init the main menu scene
func (s *ButtonTestScene) Init(world *akara.World) {
	s.World = world

	buttons := s.World.NewComponentFilter().
		Require(&d2components.Button{}).
		Require(&d2components.Ready{}).
		Build()

	s.buttons = s.World.AddSubscription(buttons)

	s.Debug("initializing ...")
}

func (s *ButtonTestScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.AddSystem(NewMouseCursorScene())

	s.createButtons()

	s.booted = true
}

func (s *ButtonTestScene) createButtons() {
	s.Add.Button(100, 100, d2button.ButtonTypeBuy, "Test")
}

// Update the main menu scene
func (s *ButtonTestScene) Update() {
	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	for _,  eid := range s.buttons.GetEntities() {
		s.updateButtonPosition(eid)
	}

	s.BaseScene.Update()
}

func (s *ButtonTestScene) updateButtonPosition(eid akara.EID) {
	trs, found := s.Components.Transform.Get(eid)
	if !found {
		return
	}

	trs.Translation.AddScalar(1)
}
