package d2systems

import (
	"image/color"
	"math/rand"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2components"
)

const (
	sceneKeyLabelTest = "Label Test Scene"
)

// NewLabelTestScene creates a new main menu scene. This is the first screen that the user
// will see when launching the game.
func NewLabelTestScene() *LabelTestScene {
	scene := &LabelTestScene{
		BaseScene: NewBaseScene(sceneKeyLabelTest),
	}

	return scene
}

// static check that LabelTestScene implements the scene interface
var _ d2interface.Scene = &LabelTestScene{}

// LabelTestScene represents the game's main menu, where users can select single or multi player,
// or start the map engine test.
type LabelTestScene struct {
	*BaseScene
	booted   bool
	labels *akara.Subscription
}

// Init the main menu scene
func (s *LabelTestScene) Init(world *akara.World) {
	s.World = world

	labels := s.World.NewComponentFilter().Require(&d2components.Label{}).Build()
	s.labels = s.World.AddSubscription(labels)

	s.Debug("initializing ...")
}

func (s *LabelTestScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.createLabels()

	s.booted = true
}

func (s *LabelTestScene) createLabels() {
	for idx := 0; idx < 1000; idx++ {
		l := s.Add.Label("LOLWUT", d2resource.Font24, d2resource.PaletteStatic)
		trs := s.AddTransform(l)
		trs.Translation.Set(rand.Float64()*800, rand.Float64()*600, 1)
	}
}

// Update the main menu scene
func (s *LabelTestScene) Update() {
	if s.Paused() {
		return
	}

	for _, eid := range s.labels.GetEntities() {
		//s.setLabelBackground(eid)
		s.updatePosition(eid)
	}

	if !s.booted {
		s.boot()
	}

	s.BaseScene.Update()
}

func (s *LabelTestScene) setLabelBackground(eid akara.EID) {
	label, found := s.GetLabel(eid)
	if !found {
		return
	}

	label.SetBackgroundColor(color.Black)
}

func (s *LabelTestScene) updatePosition(eid akara.EID) {
	trs, found := s.GetTransform(eid)
	if !found {
		return
	}

	x, y, z := trs.Translation.AddScalar(1).XYZ()

	if x > 800 {
		x -= 800
	}

	if y > 600 {
		y -= 600
	}

	trs.Translation.Set(x, y, z)
}
