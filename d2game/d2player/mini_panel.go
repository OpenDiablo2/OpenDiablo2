package d2player

import (
	"log"

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
	containerOffsetY = -49

	buttonOffsetX = -72
	buttonOffsetY = -52
)

type miniPanelActions struct {
	characterToggle func()
	inventoryToggle func()
	skilltreeToggle func()
	partyToggle     func()
	automapToggle   func()
	messageToggle   func()
	questToggle     func()
	menuToggle      func()
}

type miniPanel struct {
	ui             *d2ui.UIManager
	asset          *d2asset.AssetManager
	container      *d2ui.Sprite
	sprite         *d2ui.Sprite
	isOpen         bool
	isSinglePlayer bool
	panelGroup     *d2ui.WidgetGroup
}

func newMiniPanel(asset *d2asset.AssetManager, uiManager *d2ui.UIManager, isSinglePlayer bool) *miniPanel {
	return &miniPanel{
		ui: uiManager,
		asset:          asset,
		isOpen:         false,
		isSinglePlayer: isSinglePlayer,
	}
}

func (m *miniPanel) load(actions *miniPanelActions) {
	var err error

	m.sprite, err = m.ui.NewSprite(d2resource.MinipanelButton, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
		return
	}

	m.createWidgets(actions)
}

func (m *miniPanel) createWidgets(actions *miniPanelActions) {
	var err error

	m.panelGroup = m.ui.NewWidgetGroup(d2ui.RenderPriorityMinipanel)
	m.panelGroup.SetPosition(miniPanelX, miniPanelY)

	miniPanelContainerPath := d2resource.Minipanel
	if m.isSinglePlayer {
		miniPanelContainerPath = d2resource.MinipanelSmall
	}
	m.container, err = m.ui.NewSprite(miniPanelContainerPath, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
		return
	}
	if err:=m.container.SetCurrentFrame(0); err != nil {
		log.Print(err)
		return
	}
	x, y := screenWidth/2+containerOffsetX, screenHeight+containerOffsetY
	m.container.SetPosition(x, y)
	m.panelGroup.AddWidget(m.container)


	buttonWidth, buttonHeight, err := m.sprite.GetFrameSize(0)
	if err != nil {
		log.Print(err)
		return
	}

	buttonWidth++

	x, y = screenWidth/2 + buttonOffsetX, screenHeight + buttonOffsetY - buttonHeight
	buttonsFirst := []struct{t d2ui.ButtonType; f func()} {
		{d2ui.ButtonTypeMinipanelCharacter, actions.characterToggle},
		{d2ui.ButtonTypeMinipanelInventory, actions.inventoryToggle},
		{d2ui.ButtonTypeMinipanelSkill, actions.skilltreeToggle},
	}

	for i := range buttonsFirst {
		btn := m.ui.NewButton(buttonsFirst[i].t, "")
		btn.SetPosition(x + (i * buttonWidth), y)
		btn.OnActivated(buttonsFirst[i].f)
		m.panelGroup.AddWidget(btn)
	}
	idxOffset := len(buttonsFirst)

	if !m.isSinglePlayer {
		partyButton := m.ui.NewButton(d2ui.ButtonTypeMinipanelParty, "")
		partyButton.SetPosition(x + (3 * buttonWidth), y)
		partyButton.OnActivated(actions.partyToggle)
		m.panelGroup.AddWidget(partyButton)
		idxOffset += 1
	}

	buttonsLast := []struct{t d2ui.ButtonType; f func()} {
		{d2ui.ButtonTypeMinipanelAutomap, actions.automapToggle},
		{d2ui.ButtonTypeMinipanelMessage, actions.messageToggle },
		{d2ui.ButtonTypeMinipanelQuest, actions.questToggle },
		{d2ui.ButtonTypeMinipanelMen, actions.menuToggle},
	}

	for i := range buttonsLast {
		idx := i + idxOffset
		btn := m.ui.NewButton(buttonsLast[i].t, "")
		btn.SetPosition(x + (idx * buttonWidth), y)
		btn.OnActivated(buttonsLast[i].f)
		m.panelGroup.AddWidget(btn)
	}

	m.panelGroup.SetVisible(false)
}

func (m *miniPanel) IsOpen() bool {
	return m.isOpen
}

func (m *miniPanel) Toggle() {
	if m.isOpen {
		m.Close()
	} else {
		m.Open()
	}
}

func (m *miniPanel) Open() {
	m.panelGroup.SetVisible(true)
	m.isOpen = true
}

func (m *miniPanel) Close() {
	m.panelGroup.SetVisible(false)
	m.isOpen = false
}

func (m *miniPanel) IsInRect(px, py int) bool {
	return m.panelGroup.Contains(px, py)
}
