package d2systems

import (
	"time"

	"github.com/gravestench/akara"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
)

const (
	sceneKeyMainMenu = "Main Menu"
)

const (
	viewportMainBackground = iota + 1
	viewportTrademark
	viewport
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
	booted   bool
	logoInit bool
	sprites  struct {
		trademark      akara.EID
		logoFireLeft   akara.EID
		logoBlackLeft  akara.EID
		logoFireRight  akara.EID
		logoBlackRight akara.EID
		mainBackground akara.EID
	}
}

// Init the main menu scene
func (s *MainMenuScene) Init(world *akara.World) {
	s.World = world

	s.Debug("initializing ...")
}

func (s *MainMenuScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.setupViewports()
	s.createBackground()
	s.createButtons()
	s.createTrademarkScreen()
	s.createLogo()

	s.booted = true
}

func (s *MainMenuScene) setupViewports() {
	s.Debug("setting up viewports")

	imgPath := d2resource.GameSelectScreen
	palPath := d2resource.PaletteSky

	s.sprites.mainBackground = s.Add.SegmentedSprite(0, 0, imgPath, palPath, 4, 3, 0)
}

func (s *MainMenuScene) createBackground() {
	s.Debug("creating background")

	imgPath := d2resource.GameSelectScreen
	palPath := d2resource.PaletteSky

	s.sprites.mainBackground = s.Add.SegmentedSprite(0, 0, imgPath, palPath, 4, 3, 0)
}

func (s *MainMenuScene) createLogo() {
	s.Debug("creating logo")

	const (
		logoX, logoY = 400, 120
	)

	s.sprites.logoBlackLeft = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoBlackLeft, d2resource.PaletteUnits)
	s.sprites.logoBlackRight = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoBlackRight, d2resource.PaletteUnits)

	s.sprites.logoFireLeft = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoFireLeft, d2resource.PaletteUnits)
	s.sprites.logoFireRight = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoFireRight, d2resource.PaletteUnits)

	s.Components.DrawEffect.Add(s.sprites.logoFireLeft).DrawEffect = d2enum.DrawEffectModulate
	s.Components.DrawEffect.Add(s.sprites.logoFireRight).DrawEffect = d2enum.DrawEffectModulate
}

func (s *MainMenuScene) createButtons() {
	s.Debug("creating buttons")
}

func (s *MainMenuScene) createTrademarkScreen() {
	s.Debug("creating trademark screen")

	imgPath := d2resource.TrademarkScreen
	palPath := d2resource.PaletteSky

	s.sprites.trademark = s.Add.SegmentedSprite(0, 0, imgPath, palPath, 4, 3, 0)

	interactive := s.Components.Interactive.Add(s.sprites.trademark)

	interactive.InputVector.SetMouseButton(d2input.MouseButtonLeft)

	interactive.Callback = func() bool {
		if !s.Active() {
			interactive.Enabled = false
			return false
		}

		s.Debug("hiding trademark sprite")

		alpha := s.Components.Alpha.Add(s.sprites.trademark)

		go func() {
			alpha.Alpha = 1.0

			for alpha.Alpha > 0 {
				alpha.Alpha *= 0.725

				if alpha.Alpha <= 1e-3 {
					alpha.Alpha = 0
					return
				}

				time.Sleep(time.Second / 25)
			}

			alpha.Alpha = 0
		}()

		interactive.Enabled = false

		return true // prevent propagation
	}
}

// Update the main menu scene
func (s *MainMenuScene) Update() {
	for _, id := range s.Viewports {
		s.Components.Priority.Add(id).Priority = scenePriorityMainMenu
	}

	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	if !s.logoInit {
		s.Debug("attempting logo sprite init")
		s.initLogoSprites()
	}

	s.BaseScene.Update()
}

func (s *MainMenuScene) initLogoSprites() {
	logoSprites := []akara.EID{
		s.sprites.logoBlackLeft,
		s.sprites.logoBlackRight,
		s.sprites.logoFireLeft,
		s.sprites.logoFireRight,
	}

	for _, id := range logoSprites {
		sprite, found := s.Components.Sprite.Get(id)
		if !found {
			return
		}

		if sprite.Sprite == nil {
			return
		}

		texture, found := s.Components.Texture.Get(id)
		if !found {
			return
		}

		if texture.Texture == nil {
			return
		}
	}

	s.Debug("initializing logo sprites")

	for _, id := range logoSprites {
		sprite, _ := s.Components.Sprite.Get(id)
		if sprite.Sprite == nil {
			continue
		}

		sprite.PlayForward()
	}

	s.logoInit = true
}
