package d2systems

import (
	"image/color"
	"math"

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

// LoadingScene represents the game's loading screen, where loading progress is displayed
type LoadingScene struct {
	*BaseScene
	loadingSprite akara.EID
	filesToLoad   *akara.Subscription
	d2components.TextureFactory
	booted bool
}

// Init the loading scene
func (s *LoadingScene) Init(world *akara.World) {
	s.World = world

	s.Info("initializing ...")

	s.backgroundColor = color.Black

	s.setupFactories()
	s.setupSubscriptions()
}

func (s *LoadingScene) setupFactories() {
	renderableID := s.RegisterComponent(&d2components.Texture{})
	s.Texture = s.GetComponentFactory(renderableID)
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

func (s *LoadingScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.createLoadingScreen()

	s.booted = true
}

func (s *LoadingScene) createLoadingScreen() {
	s.Info("creating loading screen")
	s.loadingSprite = s.Add.Sprite(0, 0, d2resource.LoadingScreen, d2resource.PaletteLoading)
}

// Update the loading scene
func (s *LoadingScene) Update() {
	for _, id := range s.Viewports {
		s.AddPriority(id).Priority = scenePriorityLoading
	}

	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	s.updateViewportAlpha()
	s.updateLoadingSpritePosition()
	s.updateLoadingSpriteFrame()

	s.BaseScene.Update()
}

func (s *LoadingScene) updateViewportAlpha() {
	if len(s.Viewports) < 1 {
		return
	}

	alpha, found := s.GetAlpha(s.Viewports[0])
	if !found {
		return
	}

	isLoading := len(s.filesToLoad.GetEntities()) > 0

	if isLoading {
		alpha.Alpha = math.Min(alpha.Alpha+0.1, 1)
	} else {
		alpha.Alpha = math.Max(alpha.Alpha-0.1, 0)
	}
}

func (s *LoadingScene) updateLoadingSpritePosition() {
	if len(s.Viewports) < 1 {
		return
	}

	viewport, found := s.GetViewport(s.Viewports[0])
	if !found {
		return
	}

	sprite, found := s.GetSprite(s.loadingSprite)
	if !found {
		return
	}

	position, found := s.GetPosition(s.loadingSprite)
	if !found {
		return
	}

	centerX, centerY := viewport.Width/2, viewport.Height/2
	frameW, frameH := sprite.GetCurrentFrameSize()

	// we add the frameH in the Y because sprites are supposed to be drawn from bottom to top
	position.X, position.Y = float64(centerX-(frameW/2)), float64(centerY+(frameH/2))
}

func (s *LoadingScene) updateLoadingSpriteFrame() {
	//sprite, found := s.GetSprite(s.loadingSprite)
	//if !found {
	//	return
	//}

}
