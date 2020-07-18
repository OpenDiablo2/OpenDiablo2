package d2gamescreen

import (
	"fmt"
	"image/color"
	"math"
	"os"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2gui"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2inventory"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
)

// CharacterSelect represents the character select screen
type CharacterSelect struct {
	background             *d2ui.Sprite
	newCharButton          d2ui.Button
	convertCharButton      d2ui.Button
	deleteCharButton       d2ui.Button
	exitButton             d2ui.Button
	okButton               d2ui.Button
	deleteCharCancelButton d2ui.Button
	deleteCharOkButton     d2ui.Button
	selectionBox           *d2ui.Sprite
	okCancelBox            *d2ui.Sprite
	d2HeroTitle            d2ui.Label
	deleteCharConfirmLabel d2ui.Label
	charScrollbar          d2ui.Scrollbar
	characterNameLabel     [8]d2ui.Label
	characterStatsLabel    [8]d2ui.Label
	characterExpLabel      [8]d2ui.Label
	characterImage         [8]*d2mapentity.Player
	gameStates             []*d2player.PlayerState
	selectedCharacter      int
	showDeleteConfirmation bool
	connectionType         d2clientconnectiontype.ClientConnectionType
	connectionHost         string

	inputManager  d2interface.InputManager
	audioProvider d2interface.AudioProvider
	renderer      d2interface.Renderer
	navigator     Navigator
}

// CreateCharacterSelect creates the character select screen and returns a pointer to it
func CreateCharacterSelect(
	navigator Navigator,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
	connectionType d2clientconnectiontype.ClientConnectionType,
	connectionHost string,
) *CharacterSelect {
	return &CharacterSelect{
		selectedCharacter: -1,
		renderer:          renderer,
		connectionType:    connectionType,
		connectionHost:    connectionHost,
		inputManager:      inputManager,
		audioProvider:     audioProvider,
		navigator:         navigator,
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

// OnLoad loads the resources for the Character Select screen
func (v *CharacterSelect) OnLoad(loading d2screen.LoadingState) {
	v.audioProvider.PlayBGM(d2resource.BGMTitle)

	if err := v.inputManager.BindHandler(v); err != nil {
		fmt.Println("failed to add Character Select screen as event handler")
	}

	loading.Progress(tenPercent)

	animation, _ := d2asset.LoadAnimation(d2resource.CharacterSelectionBackground, d2resource.PaletteSky)
	bgX, bgY := 0, 0
	v.background, _ = d2ui.LoadSprite(animation)
	v.background.SetPosition(bgX, bgY)

	v.createButtons(loading)

	heroTitleX, heroTitleY := 320, 23
	v.d2HeroTitle = d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteUnits)
	v.d2HeroTitle.SetPosition(heroTitleX, heroTitleY)
	v.d2HeroTitle.Alignment = d2gui.HorizontalAlignCenter

	loading.Progress(thirtyPercent)

	v.deleteCharConfirmLabel = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
	lines := "Are you sure that you want\nto delete this character?\nTake note: this will delete all\nversions of this Character."
	v.deleteCharConfirmLabel.SetText(lines)
	v.deleteCharConfirmLabel.Alignment = d2gui.HorizontalAlignCenter
	deleteConfirmX, deleteConfirmY := 400, 185
	v.deleteCharConfirmLabel.SetPosition(deleteConfirmX, deleteConfirmY)

	animation, _ = d2asset.LoadAnimation(d2resource.CharacterSelectionSelectBox, d2resource.PaletteSky)
	v.selectionBox, _ = d2ui.LoadSprite(animation)
	selBoxX, selBoxY := 37, 86
	v.selectionBox.SetPosition(selBoxX, selBoxY)

	animation, _ = d2asset.LoadAnimation(d2resource.PopUpOkCancel, d2resource.PaletteFechar)
	v.okCancelBox, _ = d2ui.LoadSprite(animation)
	okCancelX, okCancelY := 270, 175
	v.okCancelBox.SetPosition(okCancelX, okCancelY)

	scrollBarX, scrollBarY, scrollBarHeight := 586, 87, 369
	v.charScrollbar = d2ui.CreateScrollbar(scrollBarX, scrollBarY, scrollBarHeight)
	v.charScrollbar.OnActivated(func() { v.onScrollUpdate() })
	d2ui.AddWidget(&v.charScrollbar)

	loading.Progress(fiftyPercent)

	for i := 0; i < 8; i++ {
		offsetX, offsetY := rootLabelOffsetX, rootLabelOffsetY+((i/2)*95)
		if i&1 > 0 {
			offsetX = 385
		}

		v.characterNameLabel[i] = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
		v.characterNameLabel[i].Color = rgbaColor(lightBrown)
		v.characterNameLabel[i].SetPosition(offsetX, offsetY)

		offsetY += labelHeight
		v.characterStatsLabel[i] = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteUnits)
		v.characterStatsLabel[i].SetPosition(offsetX, offsetY)

		offsetY += labelHeight
		v.characterExpLabel[i] = d2ui.CreateLabel(d2resource.Font16, d2resource.PaletteStatic)
		v.characterExpLabel[i].Color = rgbaColor(lightGreen)
		v.characterExpLabel[i].SetPosition(offsetX, offsetY)
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
	v.newCharButton = d2ui.CreateButton(
		v.renderer,
		d2ui.ButtonTypeTall,
		"CREATE NEW\nCHARACTER",
	)

	v.newCharButton.SetPosition(newCharBtnX, newCharBtnY)
	v.newCharButton.OnActivated(func() { v.onNewCharButtonClicked() })
	d2ui.AddWidget(&v.newCharButton)

	v.convertCharButton = d2ui.CreateButton(
		v.renderer,
		d2ui.ButtonTypeTall,
		"CONVERT TO\nEXPANSION",
	)

	v.convertCharButton.SetPosition(convertCharBtnX, convertCharBtnY)
	v.convertCharButton.SetEnabled(false)
	d2ui.AddWidget(&v.convertCharButton)

	v.deleteCharButton = d2ui.CreateButton(
		v.renderer,
		d2ui.ButtonTypeTall,
		"DELETE\nCHARACTER",
	)
	v.deleteCharButton.OnActivated(func() { v.onDeleteCharButtonClicked() })
	v.deleteCharButton.SetPosition(deleteCharBtnX, deleteCharBtnY)
	d2ui.AddWidget(&v.deleteCharButton)

	v.exitButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeMedium, "EXIT")
	v.exitButton.SetPosition(exitBtnX, exitBtnY)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
	d2ui.AddWidget(&v.exitButton)

	loading.Progress(twentyPercent)

	v.deleteCharCancelButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeOkCancel, "NO")
	v.deleteCharCancelButton.SetPosition(deleteCancelX, deleteCancelY)
	v.deleteCharCancelButton.SetVisible(false)
	v.deleteCharCancelButton.OnActivated(func() { v.onDeleteCharacterCancelClicked() })
	d2ui.AddWidget(&v.deleteCharCancelButton)

	v.deleteCharOkButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeOkCancel, "YES")
	v.deleteCharOkButton.SetPosition(deleteOkX, deleteOkY)
	v.deleteCharOkButton.SetVisible(false)
	v.deleteCharOkButton.OnActivated(func() { v.onDeleteCharacterConfirmClicked() })
	d2ui.AddWidget(&v.deleteCharOkButton)

	v.okButton = d2ui.CreateButton(v.renderer, d2ui.ButtonTypeMedium, "OK")
	v.okButton.SetPosition(okBtnX, okBtnY)
	v.okButton.OnActivated(func() { v.onOkButtonClicked() })
	d2ui.AddWidget(&v.okButton)
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

		v.characterNameLabel[i].SetText(v.gameStates[idx].HeroName)
		v.characterStatsLabel[i].SetText("Level 1 " + v.gameStates[idx].HeroType.String())
		v.characterExpLabel[i].SetText(expText)

		heroType := v.gameStates[idx].HeroType
		equipment := d2inventory.HeroObjects[heroType]

		// TODO: Generate or load the object from the actual player data...
		v.characterImage[i] = d2mapentity.CreatePlayer("", "", 0, 0, 0,
			v.gameStates[idx].HeroType,
			v.gameStates[idx].Stats,
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

			if localMouseX > 0 && localMouseX < bw*2 && localMouseY >= 0 && localMouseY < bh*4 {
				adjustY := localMouseY / bh
				selectedIndex := adjustY * selectionBoxNumColumns

				if localMouseX > bw {
					selectedIndex++
				}

				if (v.charScrollbar.GetCurrentOffset()*2)+selectedIndex < len(v.gameStates) {
					v.selectedCharacter = (v.charScrollbar.GetCurrentOffset() * 2) + selectedIndex
					v.moveSelectionBox()
				}
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
			hero.Advance(tickTime)
		}
	}

	return nil
}

func (v *CharacterSelect) onDeleteCharButtonClicked() {
	v.toggleDeleteCharacterDialog(true)
}

func (v *CharacterSelect) onDeleteCharacterConfirmClicked() {
	_ = os.Remove(v.gameStates[v.selectedCharacter].FilePath)
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
	v.gameStates = d2player.GetAllPlayerStates()
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
