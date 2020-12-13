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
	velocity d2components.VelocityFactory
}

// Init the main menu scene
func (s *LabelTestScene) Init(world *akara.World) {
	s.World = world

	labels := s.World.NewComponentFilter().Require(&d2components.Label{}).Build()
	s.labels = s.World.AddSubscription(labels)

	s.InjectComponent(&d2components.Velocity{}, &s.velocity.ComponentFactory)

	s.Debug("initializing ...")
}

func (s *LabelTestScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.AddSystem(&MovementSystem{})

	s.createLabels()

	s.booted = true
}

func (s *LabelTestScene) createLabels() {
	fonts := []string{
		d2resource.Font6,
		d2resource.Font8,
		d2resource.Font16,
		d2resource.Font24,
		d2resource.Font30,
		d2resource.Font42,
		d2resource.FontFormal12,
		d2resource.FontFormal11,
		d2resource.FontFormal10,
		d2resource.FontExocet10,
		d2resource.FontExocet8,
		d2resource.FontSucker,
		d2resource.FontRediculous,
	}

	for idx := 0; idx < 1000; idx++ {
		fontIdx := rand.Intn(len(fonts))

		labelEID := s.Add.Label("Open Diablo II", fonts[fontIdx], d2resource.PaletteStatic)

		c := s.Components.Color.Add(labelEID)

		r, g, b, a := uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255))
		c.Color = color.RGBA{r, g, b, a}

		trs := s.Components.Transform.Add(labelEID)
		trs.Translation.Set(rand.Float64()*800, rand.Float64()*600, 1)

		v := s.velocity.Add(labelEID)

		rx, ry := (rand.Float64()-0.5)*2, (rand.Float64()-0.5)*2

		v.Set(rx*20, ry*20, v.Z)
	}
}

// Update the main menu scene
func (s *LabelTestScene) Update() {
	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	s.BaseScene.Update()
}
