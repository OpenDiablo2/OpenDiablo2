package d2systems

import (
	"image/color"
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	sceneKeyCheckboxTest = "Checkbox Test Scene"
)

// NewCheckboxTestScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewCheckboxTestScene() *CheckboxTestScene {
	scene := &CheckboxTestScene{
		BaseScene: NewBaseScene(sceneKeyCheckboxTest),
	}

	return scene
}

// static check that CheckboxTestScene implements the scene interface
var _ d2interface.Scene = &CheckboxTestScene{}

// CheckboxTestScene represents the game's main menu, where users can select single or multi player,
// or start the map engine test.
type CheckboxTestScene struct {
	*BaseScene
	state      d2enum.SceneState
	checkboxes *akara.Subscription
}

// Init the main menu scene
func (s *CheckboxTestScene) Init(world *akara.World) {
	s.World = world

	checkboxes := s.World.NewComponentFilter().
		Require(&d2components.Checkbox{}).
		Require(&d2components.Ready{}).
		Build()

	s.checkboxes = s.World.AddSubscription(checkboxes)

	s.Debug("initializing ...")
}

func (s *CheckboxTestScene) boot() {
	if !s.BaseScene.Booted() {
		s.BaseScene.boot()
		return
	}

	viewport, found := s.Components.Viewport.Get(s.Viewports[0])
	if !found {
		return
	}

	s.AddSystem(NewMouseCursorScene())

	s.Add.Rectangle(0, 0, viewport.Width, viewport.Height, color.White)

	s.createCheckboxes()

	s.state = d2enum.SceneStateBooted
}

//nolint:gomnd // arbitrary example numbers for test
func (s *CheckboxTestScene) createCheckboxes() {
	s.Add.Checkbox(100, 100, true, true, "Expansion character", checkboxClickCallback)
	s.Add.Checkbox(100, 120, false, true, "Hardcore", checkboxClickCallback)
	s.Add.Checkbox(100, 140, true, false, "disabled checked test", checkboxClickCallback)
	s.Add.Checkbox(100, 160, false, false, "disabled unchecked test",
		checkboxClickCallback)
}

// Update the main menu scene
func (s *CheckboxTestScene) Update() {
	if s.Paused() {
		return
	}

	if s.state == d2enum.SceneStateUninitialized {
		s.boot()
	}

	if s.state != d2enum.SceneStateBooted {
		return
	}

	s.BaseScene.Update()
}

func checkboxClickCallback(thisComponent akara.Component) bool {
	this := thisComponent.(*d2components.Checkbox)
	if this.Checkbox.GetEnabled() {
		text := this.Checkbox.Label.GetText()

		if this.Checkbox.GetPressed() {
			log.Printf("%s enabled", text)
		} else {
			log.Printf("%s disabled", text)
		}
	}

	return false
}
