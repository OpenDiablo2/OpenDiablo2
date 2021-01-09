package d2gamescreen

import (
	"image/color"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2hero"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2map/d2mapentity"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
)

// CreateCharacterSelect creates the character select screen and returns a pointer to it
func CreateCharacterSelect(
	navigator d2interface.Navigator,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	inputManager d2interface.InputManager,
	audioProvider d2interface.AudioProvider,
	ui *d2ui.UIManager,
	connectionType d2clientconnectiontype.ClientConnectionType,
	l d2util.LogLevel,
	connectionHost string,
) (*CharacterSelect, error) {
	playerStateFactory, err := d2hero.NewHeroStateFactory(asset)
	if err != nil {
		return nil, err
	}

	entityFactory, err := d2mapentity.NewMapEntityFactory(asset)
	if err != nil {
		return nil, err
	}

	characterSelect := &CharacterSelect{
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

	characterSelect.Logger = d2util.NewLogger()
	characterSelect.Logger.SetLevel(l)
	characterSelect.Logger.SetPrefix(logPrefix)

	return characterSelect, nil
}

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
	loaded                 bool
	connectionType         d2clientconnectiontype.ClientConnectionType
	connectionHost         string

	uiManager     *d2ui.UIManager
	inputManager  d2interface.InputManager
	audioProvider d2interface.AudioProvider
	renderer      d2interface.Renderer
	navigator     d2interface.Navigator

	*d2util.Logger
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
	v.audioProvider.PlayBGM(d2resource.BGMTitle)

	err := v.inputManager.BindHandler(v)
	if err != nil {
		v.Error("failed to add Character Select screen as event handler")
	}

	loading.Progress(tenPercent)

	v.loadBackground()
	v.createButtons(loading)
	v.loadHeroTitle()

	loading.Progress(thirtyPercent)

	v.loadDeleteCharConfirm()
	v.loadSelectionBox()
	v.loadOkCancelBox()
	v.loadCharScrollbar()

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
	v.loaded = true
}

func (v *CharacterSelect) loadBackground() {
	var err error

	bgX, bgY := 0, 0

	v.background, err = v.uiManager.NewSprite(d2resource.CharacterSelectionBackground, d2resource.PaletteSky)
	if err != nil {
		v.Error(err.Error())
	}

	v.background.SetPosition(bgX, bgY)
}

func (v *CharacterSelect) loadHeroTitle() {
	heroTitleX, heroTitleY := 320, 23
	v.d2HeroTitle = v.uiManager.NewLabel(d2resource.Font42, d2resource.PaletteUnits)
	v.d2HeroTitle.SetPosition(heroTitleX, heroTitleY)
	v.d2HeroTitle.Alignment = d2ui.HorizontalAlignCenter
}

func (v *CharacterSelect) loadDeleteCharConfirm() {
	v.deleteCharConfirmLabel = v.uiManager.NewLabel(d2resource.Font16, d2resource.PaletteUnits)
	lines := strings.Join(d2util.SplitIntoLinesWithMaxWidth(v.asset.TranslateLabel(d2enum.DelCharConfLabel), 29), "\n")
	v.deleteCharConfirmLabel.SetText(lines)
	v.deleteCharConfirmLabel.Alignment = d2ui.HorizontalAlignCenter
	deleteConfirmX, deleteConfirmY := 400, 185
	v.deleteCharConfirmLabel.SetPosition(deleteConfirmX, deleteConfirmY)
}

func (v *CharacterSelect) loadSelectionBox() {
	var err error

	v.selectionBox, err = v.uiManager.NewSprite(d2resource.CharacterSelectionSelectBox, d2resource.PaletteSky)
	if err != nil {
		v.Error(err.Error())
	}

	selBoxX, selBoxY := 37, 86
	v.selectionBox.SetPosition(selBoxX, selBoxY)
}

func (v *CharacterSelect) loadOkCancelBox() {
	var err error

	v.okCancelBox, err = v.uiManager.NewSprite(d2resource.PopUpOkCancel, d2resource.PaletteFechar)
	if err != nil {
		v.Error(err.Error())
	}

	okCancelX, okCancelY := 270, 175
	v.okCancelBox.SetPosition(okCancelX, okCancelY)
}

func (v *CharacterSelect) loadCharScrollbar() {
	scrollBarX, scrollBarY, scrollBarHeight := 586, 87, 369
	v.charScrollbar = v.uiManager.NewScrollbar(scrollBarX, scrollBarY, scrollBarHeight)
	v.charScrollbar.OnActivated(func() { v.onScrollUpdate() })
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
	v.newCharButton = v.uiManager.NewButton(d2ui.ButtonTypeTall, strings.Join(
		d2util.SplitIntoLinesWithMaxWidth(v.asset.TranslateString("#831"), 13), "\n"))
	v.newCharButton.SetPosition(newCharBtnX, newCharBtnY)
	v.newCharButton.OnActivated(func() { v.onNewCharButtonClicked() })

	v.convertCharButton = v.uiManager.NewButton(d2ui.ButtonTypeTall,
		strings.Join(d2util.SplitIntoLinesWithMaxWidth(v.asset.TranslateString("#825"), 13), "\n"))
	v.convertCharButton.SetPosition(convertCharBtnX, convertCharBtnY)
	v.convertCharButton.SetEnabled(false)

	v.deleteCharButton = v.uiManager.NewButton(d2ui.ButtonTypeTall,
		strings.Join(d2util.SplitIntoLinesWithMaxWidth(v.asset.TranslateString("#832"), 13), "\n"))
	v.deleteCharButton.OnActivated(func() { v.onDeleteCharButtonClicked() })
	v.deleteCharButton.SetPosition(deleteCharBtnX, deleteCharBtnY)

	v.exitButton = v.uiManager.NewButton(d2ui.ButtonTypeMedium, v.asset.TranslateLabel(d2enum.ExitLabel))
	v.exitButton.SetPosition(exitBtnX, exitBtnY)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })

	loading.Progress(twentyPercent)

	v.deleteCharCancelButton = v.uiManager.NewButton(d2ui.ButtonTypeOkCancel, v.asset.TranslateLabel(d2enum.NoLabel))
	v.deleteCharCancelButton.SetPosition(deleteCancelX, deleteCancelY)
	v.deleteCharCancelButton.SetVisible(false)
	v.deleteCharCancelButton.OnActivated(func() { v.onDeleteCharacterCancelClicked() })

	v.deleteCharOkButton = v.uiManager.NewButton(d2ui.ButtonTypeOkCancel, v.asset.TranslateLabel(d2enum.YesLabel))
	v.deleteCharOkButton.SetPosition(deleteOkX, deleteOkY)
	v.deleteCharOkButton.SetVisible(false)
	v.deleteCharOkButton.OnActivated(func() { v.onDeleteCharacterConfirmClicked() })

	v.okButton = v.uiManager.NewButton(d2ui.ButtonTypeMedium, v.asset.TranslateLabel(d2enum.OKLabel))
	v.okButton.SetPosition(okBtnX, okBtnY)
	v.okButton.OnActivated(func() { v.onOkButtonClicked() })
}

func (v *CharacterSelect) onScrollUpdate() {
	v.moveSelectionBox()
	v.updateCharacterBoxes()
}

func (v *CharacterSelect) updateCharacterBoxes() {
	expText := v.asset.TranslateString("#803")

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
		heroInfo := v.asset.TranslateString("level") + " " + strconv.FormatInt(int64(v.gameStates[idx].Stats.Level), 10) +
			" " + v.asset.TranslateString(v.gameStates[idx].HeroType.String())

		v.characterNameLabel[i].SetText(d2ui.ColorTokenize(heroName, d2ui.ColorTokenGold))
		v.characterStatsLabel[i].SetText(d2ui.ColorTokenize(heroInfo, d2ui.ColorTokenWhite))
		v.characterExpLabel[i].SetText(d2ui.ColorTokenize(expText, d2ui.ColorTokenGreen))

		heroType := v.gameStates[idx].HeroType
		equipment := v.DefaultHeroItems[heroType]

		// https://github.com/OpenDiablo2/OpenDiablo2/issues/791
		v.characterImage[i] = v.NewPlayer("", "", 0, 0, 0,
			v.gameStates[idx].HeroType,
			v.gameStates[idx].Stats,
			v.gameStates[idx].Skills,
			&equipment,
			v.gameStates[idx].LeftSkill,
			v.gameStates[idx].RightSkill,
			v.gameStates[idx].Gold,
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
func (v *CharacterSelect) Render(screen d2interface.Surface) {
	v.background.RenderSegmented(screen, 4, 3, 0)
	v.d2HeroTitle.Render(screen)

	actualSelectionIndex := v.selectedCharacter - (v.charScrollbar.GetCurrentOffset() * 2)

	if v.selectedCharacter > -1 && actualSelectionIndex >= 0 && actualSelectionIndex < 8 {
		v.selectionBox.RenderSegmented(screen, 2, 1, 0)
	}

	for i := 0; i < 8; i++ {
		idx := i + (v.charScrollbar.GetCurrentOffset() * 2)
		if idx >= len(v.gameStates) {
			continue
		}

		v.characterNameLabel[i].Render(screen)
		v.characterStatsLabel[i].Render(screen)
		v.characterExpLabel[i].Render(screen)

		x, y := v.characterNameLabel[i].GetPosition()
		charImgX := x - selectionBoxImageOffsetX
		charImgY := y + selectionBoxImageOffsetY
		screen.PushTranslation(charImgX, charImgY)
		v.characterImage[i].Render(screen)
		screen.Pop()
	}

	if v.showDeleteConfirmation {
		screen.DrawRect(screenWidth, screenHeight, rgbaColor(blackHalfOpacity))
		v.okCancelBox.RenderSegmented(screen, 2, 1, 0)
		v.deleteCharConfirmLabel.Render(screen)
	}
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
	if !v.loaded {
		return false
	}

	if v.showDeleteConfirmation {
		return false
	}

	if event.Button() != d2enum.MouseButtonLeft {
		return false
	}

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
		v.Error(err.Error())
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

// OnUnload candles cleanup when this screen is closed
func (v *CharacterSelect) OnUnload() error {
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/792
	if err := v.inputManager.UnbindHandler(v); err != nil {
		return err
	}

	v.loaded = false

	return nil
}
