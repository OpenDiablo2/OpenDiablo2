package d2player

import (
	"log"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2geom"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	miniPanelX      = 325
	miniPanelY      = 526
	miniPanelWidth  = 156
	miniPanelHeight = 26
)

const (
	containerOffsetX = -75
	containerOffsetY = -48

	buttonOffsetX = -72
	buttonOffsetY = -51
)

type miniPanel struct {
	asset          *d2asset.AssetManager
	container      *d2ui.Sprite
	button         *d2ui.Sprite
	isOpen         bool
	isSinglePlayer bool
	rectangle      d2geom.Rectangle
}

func newMiniPanel(asset *d2asset.AssetManager, uiManager *d2ui.UIManager, isSinglePlayer bool) *miniPanel {
	miniPanelContainerPath := d2resource.Minipanel
	if isSinglePlayer {
		miniPanelContainerPath = d2resource.MinipanelSmall
	}

	containerSprite, err := uiManager.NewSprite(miniPanelContainerPath, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
		return nil
	}

	buttonSprite, err := uiManager.NewSprite(d2resource.MinipanelButton, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
		return nil
	}

	rectangle := d2geom.Rectangle{
		Left:   miniPanelX,
		Top:    miniPanelY,
		Width:  miniPanelWidth,
		Height: miniPanelHeight,
	}

	if !isSinglePlayer {
		rectangle.Width = 182
	}

	return &miniPanel{
		asset:          asset,
		container:      containerSprite,
		button:         buttonSprite,
		isOpen:         false,
		isSinglePlayer: isSinglePlayer,
		rectangle:      rectangle,
	}
}

func (m *miniPanel) IsOpen() bool {
	return m.isOpen
}

func (m *miniPanel) Toggle() {
	m.isOpen = !m.isOpen
}

func (m *miniPanel) Open() {
	m.isOpen = true
}

func (m *miniPanel) Close() {
	m.isOpen = false
}

func (m *miniPanel) Render(target d2interface.Surface) error {
	if !m.isOpen {
		return nil
	}

	if err := m.container.SetCurrentFrame(0); err != nil {
		return err
	}

	width, height := target.GetSize()
	halfW := width >> 1
	x, y := halfW+containerOffsetX, height+containerOffsetY

	m.container.SetPosition(x, y)

	if err := m.container.Render(target); err != nil {
		return err
	}

	buttonWidth, _ := m.button.GetCurrentFrameSize()
	buttonWidth++

	for i, j := 0, 0; j < 16; i++ {
		if m.isSinglePlayer && j == 6 { // skip Party Screen button if the game is single player
			j += 2
		}

		if err := m.button.SetCurrentFrame(j); err != nil {
			return err
		}

		offsetX := buttonOffsetX + (buttonWidth * i)
		x, y := halfW+offsetX, height+buttonOffsetY

		m.button.SetPosition(x, y)

		if err := m.button.Render(target); err != nil {
			return err
		}

		j += 2
	}

	return nil
}

func (m *miniPanel) isInRect(x, y int) bool {
	return m.rectangle.IsInRect(x, y)
}
