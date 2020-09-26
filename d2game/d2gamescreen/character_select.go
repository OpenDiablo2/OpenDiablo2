package d2gamescreen

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
)

// CharacterSelect represents the character select screen
type CharacterSelect struct {
	asset *d2asset.AssetManager
	*d2mapentity.MapEntityFactory
	*d2hero.HeroStateFactory
	background             *d2ui.Sprite
	newCharButton          *d2ui.Button
	convertCharButton      *d2ui.Button
	deleteCharButton       *d2ui.Button
	exitButton             *d2ui.Button
	okButton               *d2ui.Button
	deleteCharCancelButton *d2ui.Button
	deleteCharOkButton     *d2ui.Button
	selectionBox           *d2ui.Sprite
	okCancelBox            *d2ui.Sprite
	d2HeroTitle            *d2ui.Label
	deleteCharConfirmLabel *d2ui.Label
	charScrollbar          *d2ui.Scrollbar
	characterNameLabel     [8]*d2ui.Label
	characterStatsLabel    [8]*d2ui.Label
	characterExpLabel      [8]*d2ui.Label
	characterImage         [8]*d2mapentity.Player
	gameStates             []*d2hero.HeroState
	selectedCharacter      int
	tickTimer              float64
	storedTickTimer        float64
	showDeleteConfirmation bool
	connectionType         d2clientconnectiontype.ClientConnectionType
	connectionHost         string

	uiManager     *d2ui.UIManager
	inputManager  d2interface.InputManager
	audioProvider d2interface.AudioProvider
	renderer      d2interface.Renderer
	navigator     Navigator
}

// CreateCharacterSelect creates the character select screen and returns a pointer to it
func CreateCharacterSelect(
	navigator Navigator,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
	ui *d2ui.UIManager,
	connectionType d2clientconnectiontype.ClientConnectionType,
	connectionHost string,
) *CharacterSelect {
	playerStateFactory, _ := d2hero.NewHeroStateFactory(asset) // TODO: handle errors
	entityFactory, _ := d2mapentity.NewMapEntityFactory(asset)

	return &CharacterSelect{
		selectedCharacter: -1,
		asset:             asset,
		MapEntityFactory:  entityFactory,
		renderer:          renderer,
		connectionType:    connectionType,
		connectionHost:    connectionHost,
		inputManager:      inputManager,
		audioProvider:     audioProvider,
		navigator:         navigator,
		uiManager:         ui,
		HeroStateFactory:  playerStateFactory,
	}
}

const (
	tenPercent = 0.1 * iota
	twentyPercent
	thirtyPercent
	fourtyPercent
	fiftyPercent
	sixtyPercent
	seventyPercent
	eightyPercent
	ninetyPercent
)

const (
	rootLabelOffsetX = 115
	rootLabelOffsetY = 100
	labelHeight      = 15
)

const (
	selectionBoxNumColumns   = 2
	selectionBoxNumRows      = 4
	selectionBoxWidth        = 272
	selectionBoxHeight       = 92
	selectionBoxOffsetX      = 37
	selectionBoxOffsetY      = 86
	selectionBoxImageOffsetX = 40
	selectionBoxImageOffsetY = 50
)

const (
	blackHalfOpacity = 0x0000007f
	lightBrown       = 0xbca88cff
	lightGreen       = 0x18ff00ff
)

const (
	screenWidth  = 800
	screenHeight = 600
)

const (
	newCharBtnX, newCharBtnY         = 33, 468
	convertCharBtnX, convertCharBtnY = 233, 468
	deleteCharBtnX, deleteCharBtnY   = 433, 468
	deleteCancelX, deleteCancelY     = 282, 308
	deleteOkX, deleteOkY             = 422, 308
	exitBtnX, exitBtnY               = 33, 537
	okBtnX, okBtnY                   = 625, 537
)

const (
	doubleClickTime = 1.25
)

// OnLoad loads the resources for the Character Select screen
func (v *CharacterSelect) OnLoad(loading d2screen.LoadingState) {
	var err error
	v.audioProvider.PlayBGM(d2resource.BGMTitle)

	if err := v.inputManager.BindHandler(v); err != nil {
		fmt.Println("failed to add Character Select screen as event handler")
	}

	loading.Progress(tenPercent)

	bgX, bgY := 0, 0
	v.background, err = v.uiManager.NewSprite(d2resource.CharacterSelectionBackground, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}
	v.background.SetPosition(bgX, bgY)

	v.createButtons(loading)

	heroTitleX, heroTitleY := 320, 23
	v.d2HeroTitle = v.uiManager.NewLabel(d2resource.Font42, d2resource.PaletteUnits)
	v.d2HeroTitle.SetPosition(heroTitleX, heroTitleY)
	v.d2HeroTitle.Alignment = d2gui.HorizontalAlignCenter

	loading.Progress(thirtyPercent)

	v.deleteCharConfirmLabel = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	lines := "Are you sure that you want\nto delete this character?\nTake note: this will delete all\nversions of this Character."
	v.deleteCharConfirmLabel.SetText(lines)
	v.deleteCharConfirmLabel.Alignment = d2gui.HorizontalAlignCenter
	deleteConfirmX, deleteConfirmY := 400, 185
	v.deleteCharConfirmLabel.SetPosition(deleteConfirmX, deleteConfirmY)

	v.selectionBox, err = v.uiManager.NewSprite(d2resource.CharacterSelectionSelectBox, d2resource.PaletteSky)
	if err != nil {
		log.Print(err)
	}
	selBoxX, selBoxY := 37, 86
	v.selectionBox.SetPosition(selBoxX, selBoxY)

	v.okCancelBox, err = v.uiManager.NewSprite(d2resource.PopUpOkCancel, d2resource.PaletteFechar)
	if err != nil {
		log.Print(err)
	}
	okCancelX, okCancelY := 270, 175
	v.okCancelBox.SetPosition(okCancelX, okCancelY)

	scrollBarX, scrollBarY, scrollBarHeight := 586, 87, 369
	v.charScrollbar = v.uiManager.NewScrollbar(scrollBarX, scrollBarY, scrollBarHeight)
	v.charScrollbar.OnActivated(func() { v.onScrollUpdate() })

	loading.Progress(fiftyPercent)

	for i := 0; i < 8; i++ {
		offsetX, offsetY := rootLabelOffsetX, rootLabelOffsetY+((i/2)*95)
		if i&1 > 0 {
			offsetX = 385
		}

		v.characterNameLabel[i] = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
		v.characterNameLabel[i].SetPosition(offsetX, offsetY)
		v.characterNameLabel[i].Color[0] = rgbaColor(lightBrown)

		offsetY += labelHeight
		v.characterStatsLabel[i] = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
		v.characterStatsLabel[i].SetPosition(offsetX, offsetY)

		offsetY += labelHeight
		v.characterExpLabel[i] = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteStatic)
		v.characterExpLabel[i].SetPosition(offsetX, offsetY)
		v.characterExpLabel[i].Color[0] = rgbaColor(lightGreen)
	}
	v.refreshGameStates()
}

func rgbaColor(rgba uint32) color.RGBA {
	result := color.RGBA{}
	a, b, g, r := 0, 1, 2, 3
	byteWidth := 8
	byteMask := 0xff

	for idx := 0; idx < 4; idx++ {
		shift := idx * byteWidth
		component := uint8(rgba>>shift) & uint8(byteMask)

		switch idx {
		case a:
			result.A = component
		case b:
			result.B = component
		case g:
			result.G = component
		case r:
			result.R = component
		}
	}

	return result
}

func (v *CharacterSelect) createButtons(loading d2screen.LoadingState) {
	v.newCharButton = v.uiManager.NewButton(d2ui.ButtonTypeTall, "CREATE NEW\nCHARACTER")

	v.newCharButton.SetPosition(newCharBtnX, newCharBtnY)
	v.newCharButton.OnActivated(func() { v.onNewCharButtonClicked() })

	v.convertCharButton = v.uiManager.NewButton(d2ui.ButtonTypeTall, "CONVERT TO\nEXPANSION")

	v.convertCharButton.SetPosition(convertCharBtnX, convertCharBtnY)
	v.convertCharButton.SetEnabled(false)

	v.deleteCharButton = v.uiManager.NewButton(d2ui.ButtonTypeTall, "DELETE\nCHARACTER")
	v.deleteCharButton.OnActivated(func() { v.onDeleteCharButtonClicked() })
	v.deleteCharButton.SetPosition(deleteCharBtnX, deleteCharBtnY)

	v.exitButton = v.uiManager.NewButton(d2ui.ButtonTypeMedium, "EXIT")
	v.exitButton.SetPosition(exitBtnX, exitBtnY)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })

	loading.Progress(twentyPercent)

	v.deleteCharCancelButton = v.uiManager.NewButton(d2ui.ButtonTypeOkCancel, "NO")
	v.deleteCharCancelButton.SetPosition(deleteCancelX, deleteCancelY)
	v.deleteCharCancelButton.SetVisible(false)
	v.deleteCharCancelButton.OnActivated(func() { v.onDeleteCharacterCancelClicked() })

	v.deleteCharOkButton = v.uiManager.NewButton(d2ui.ButtonTypeOkCancel, "YES")
	v.deleteCharOkButton.SetPosition(deleteOkX, deleteOkY)
	v.deleteCharOkButton.SetVisible(false)
	v.deleteCharOkButton.OnActivated(func() { v.onDeleteCharacterConfirmClicked() })

	v.okButton = v.uiManager.NewButton(d2ui.ButtonTypeMedium, "OK")
	v.okButton.SetPosition(okBtnX, okBtnY)
	v.okButton.OnActivated(func() { v.onOkButtonClicked() })
}

func (v *CharacterSelect) onScrollUpdate() {
	v.moveSelectionBox()
	v.updateCharacterBoxes()
}

func (v *CharacterSelect) updateCharacterBoxes() {
	expText := "EXPANSION CHARACTER"

	for i := 0; i < 8; i++ {
		idx := i + (v.charScrollbar.GetCurrentOffset() * 2)
		if idx >= len(v.gameStates) {
			v.characterNameLabel[i].SetText("")
			v.characterStatsLabel[i].SetText("")
			v.characterExpLabel[i].SetText("")
			v.characterImage[i] = nil

			continue
		}

		heroName := v.gameStates[idx].HeroName
		heroInfo := "Level 1 " + v.gameStates[idx].HeroType.String()

		v.characterNameLabel[i].SetText(d2ui.ColorTokenize(heroName, d2ui.ColorTokenGold))
		v.characterStatsLabel[i].SetText(d2ui.ColorTokenize(heroInfo, d2ui.ColorTokenWhite))
		v.characterExpLabel[i].SetText(d2ui.ColorTokenize(expText, d2ui.ColorTokenGreen))

		heroType := v.gameStates[idx].HeroType
		equipment := v.DefaultHeroItems[heroType]

		// TODO: Generate or load the object from the actual player data...
		v.characterImage[i] = v.NewPlayer("", "", 0, 0, 0,
			v.gameStates[idx].HeroType,
			v.gameStates[idx].Stats,
			v.gameStates[idx].Skills,
			&equipment,
		)
	}
}

func (v *CharacterSelect) onNewCharButtonClicked() {
	v.navigator.ToSelectHero(v.connectionType, v.connectionHost)
}

func (v *CharacterSelect) onExitButtonClicked() {
	v.navigator.ToMainMenu()
}

// Render renders the Character Select screen
func (v *CharacterSelect) Render(screen d2interface.Surface) error {
	if err := v.background.RenderSegmented(screen, 4, 3, 0); err != nil {
		return err
	}

	v.d2HeroTitle.Render(screen)
	actualSelectionIndex := v.selectedCharacter - (v.charScrollbar.GetCurrentOffset() * 2)

	if v.selectedCharacter > -1 && actualSelectionIndex >= 0 && actualSelectionIndex < 8 {
		if err := v.selectionBox.RenderSegmented(screen, 2, 1, 0); err != nil {
			return err
		}
	}

	for i := 0; i < 8; i++ {
		idx := i + (v.charScrollbar.GetCurrentOffset() * 2)
		if idx >= len(v.gameStates) {
			continue
		}

		v.characterNameLabel[i].Render(screen)
		v.characterStatsLabel[i].Render(screen)
		v.characterExpLabel[i].Render(screen)

		charImgX := v.characterNameLabel[i].X - selectionBoxImageOffsetX
		charImgY := v.characterNameLabel[i].Y + selectionBoxImageOffsetY
		screen.PushTranslation(charImgX, charImgY)
		v.characterImage[i].Render(screen)
		screen.Pop()
	}

	if v.showDeleteConfirmation {
		screen.DrawRect(screenWidth, screenHeight, rgbaColor(blackHalfOpacity))

		if err := v.okCancelBox.RenderSegmented(screen, 2, 1, 0); err != nil {
			return err
		}

		v.deleteCharConfirmLabel.Render(screen)
	}

	return nil
}

func (v *CharacterSelect) moveSelectionBox() {
	if v.selectedCharacter == -1 {
		v.d2HeroTitle.SetText("")
		return
	}

	bw := 272
	bh := 92
	selectedIndex := v.selectedCharacter - (v.charScrollbar.GetCurrentOffset() * 2)

	selBoxX := selectionBoxOffsetX + ((selectedIndex & 1) * bw)
	selBoxY := selectionBoxOffsetY + (bh * (selectedIndex / 2))
	v.selectionBox.SetPosition(selBoxX, selBoxY)
	v.d2HeroTitle.SetText(v.gameStates[v.selectedCharacter].HeroName)
}

// OnMouseButtonDown is called when a mouse button is clicked
func (v *CharacterSelect) OnMouseButtonDown(event d2interface.MouseEvent) bool {
	if !v.showDeleteConfirmation {
		if event.Button() == d2enum.MouseButtonLeft {
			mx, my := event.X(), event.Y()

			bw := selectionBoxWidth
			bh := selectionBoxHeight
			localMouseX := mx - selectionBoxOffsetX
			localMouseY := my - selectionBoxOffsetY

			// if Mouse is within character selection bounds.
			if localMouseX > 0 && localMouseX < bw*2 && localMouseY >= 0 && localMouseY < bh*4 {
				adjustY := localMouseY / bh
				// sets current verticle index for selected character in left column.
				selectedIndex := adjustY * selectionBoxNumColumns

				// if selected character in left column should be in right column, add 1.
				if localMouseX > bw {
					selectedIndex++
				}

				// Make sure selection takes the scrollbar into account to make proper selection.
				if (v.charScrollbar.GetCurrentOffset()*2)+selectedIndex < len(v.gameStates) {
					selectedIndex = (v.charScrollbar.GetCurrentOffset() * 2) + selectedIndex
				}

				// if the selection box didn't move, check if it was a double click, otherwise set selectedCharacter to
				// selectedIndex and move selection box over both.
				if v.selectedCharacter == selectedIndex {
					// We clicked twice within character selection box within  v.doubleClickTime seconds.
					if (v.tickTimer - v.storedTickTimer) < doubleClickTime {
						v.onOkButtonClicked()
					}
				} else if selectedIndex < len(v.gameStates) {
					v.selectedCharacter = selectedIndex
					v.moveSelectionBox()
				}
				// Keep track of when we last clicked so we can determine if we double clicked a character.
				v.storedTickTimer = v.tickTimer
			}

			return true
		}
	}

	return false
}

// Advance runs the update logic on the Character Select screen
func (v *CharacterSelect) Advance(tickTime float64) error {
	for _, hero := range v.characterImage {
		if hero != nil {
			v.tickTimer += tickTime
			hero.Advance(tickTime)
		}
	}

	return nil
}

func (v *CharacterSelect) onDeleteCharButtonClicked() {
	v.toggleDeleteCharacterDialog(true)
}

func (v *CharacterSelect) onDeleteCharacterConfirmClicked() {
	err := os.Remove(v.gameStates[v.selectedCharacter].FilePath)
	if err != nil {
		log.Print(err)
	}
	v.charScrollbar.SetCurrentOffset(0)
	v.refreshGameStates()
	v.toggleDeleteCharacterDialog(false)
	v.deleteCharButton.SetEnabled(len(v.gameStates) > 0)
	v.okButton.SetEnabled(len(v.gameStates) > 0)
}

func (v *CharacterSelect) onDeleteCharacterCancelClicked() {
	v.toggleDeleteCharacterDialog(false)
}

func (v *CharacterSelect) toggleDeleteCharacterDialog(showDialog bool) {
	v.showDeleteConfirmation = showDialog
	v.okButton.SetEnabled(!showDialog)
	v.deleteCharButton.SetEnabled(!showDialog)
	v.exitButton.SetEnabled(!showDialog)
	v.newCharButton.SetEnabled(!showDialog)
	v.deleteCharOkButton.SetVisible(showDialog)
	v.deleteCharCancelButton.SetVisible(showDialog)
}

func (v *CharacterSelect) refreshGameStates() {
	gameStates, err := v.HeroStateFactory.GetAllHeroStates()
	if err == nil {
		v.gameStates = gameStates
	}

	v.updateCharacterBoxes()

	if len(v.gameStates) > 0 {
		v.selectedCharacter = 0
		numStates := selectionBoxNumColumns * selectionBoxNumRows
		byHalf := 2.0

		v.d2HeroTitle.SetText(v.gameStates[0].HeroName)
		v.charScrollbar.SetMaxOffset(int(math.Ceil(float64(len(v.gameStates)-numStates) / byHalf)))
	} else {
		v.selectedCharacter = -1
		v.charScrollbar.SetMaxOffset(0)
	}

	v.moveSelectionBox()
}

func (v *CharacterSelect) onOkButtonClicked() {
	v.navigator.ToCreateGame(v.gameStates[v.selectedCharacter].FilePath, v.connectionType, v.connectionHost)
}

func (v *CharacterSelect) OnUnload() error {
	if err := v.inputManager.UnbindHandler(v); err != nil { // TODO: hack
		return err
	}

	return nil
}
