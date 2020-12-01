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

	s.Info("initializing ...")
}

func (s *MainMenuScene) boot() {
	if !s.BaseScene.booted {
		s.BaseScene.boot()
		return
	}

	s.createBackground()
	s.createButtons()
	s.createTrademarkScreen()
	s.createLogo()

	s.booted = true
}

func (s *MainMenuScene) createBackground() {
	s.Info("creating background")

	imgPath := d2resource.GameSelectScreen
	palPath := d2resource.PaletteSky

	s.sprites.mainBackground = s.Add.SegmentedSprite(0, 0, imgPath, palPath, 4, 3, 0)
}

func (s *MainMenuScene) createLogo() {
	s.Info("creating logo")

	const (
		logoX, logoY = 400, 120
	)

	s.sprites.logoBlackLeft = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoBlackLeft, d2resource.PaletteUnits)
	s.sprites.logoBlackRight = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoBlackRight, d2resource.PaletteUnits)

	s.sprites.logoFireLeft = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoFireLeft, d2resource.PaletteUnits)
	s.sprites.logoFireRight = s.Add.Sprite(logoX, logoY, d2resource.Diablo2LogoFireRight, d2resource.PaletteUnits)

	s.AddDrawEffect(s.sprites.logoFireLeft).DrawEffect = d2enum.DrawEffectModulate
	s.AddDrawEffect(s.sprites.logoFireRight).DrawEffect = d2enum.DrawEffectModulate
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

		alpha := s.AddAlpha(s.sprites.trademark)

		go func() {
			a := 1.0
			for a > 0 {
				a -= 0.125
				alpha.Alpha = a
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
		s.AddPriority(id).Priority = scenePriorityMainMenu
	}

	if s.Paused() {
		return
	}

	if !s.booted {
		s.boot()
	}

	if !s.logoInit {
		s.Info("attempting logo sprite init")
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
		sprite, found := s.GetSprite(id)
		if !found {
			return
		}

		if sprite.Sprite == nil {
			return
		}

		texture, found := s.GetTexture(id)
		if !found {
			return
		}

		if texture.Texture == nil {
			return
		}
	}

	s.Info("initializing logo sprites")

	for _, id := range logoSprites {
		sprite, _ := s.GetSprite(id)
		if sprite.Sprite == nil {
			continue
		}

		sprite.SetEffect(d2enum.DrawEffectModulate)
		sprite.PlayForward()
	}

	s.logoInit = true
}
