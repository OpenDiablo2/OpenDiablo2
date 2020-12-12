package d2gamescreen

import (
	"bufio"
	"os"
	"path"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	creditsX, creditsY               = 0, 0
	charSelExitBtnX, charSelExitBtnY = 33, 543
)

const secondsPerCycle float64 = 0.02

type labelItem struct {
	Label     *d2ui.Label
	IsHeading bool
	Available bool
}

// CreateCredits creates an instance of the credits screen
func CreateCredits(navigator d2interface.Navigator,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	l d2util.LogLevel,
	ui *d2ui.UIManager) *Credits {
	credits := &Credits{
		asset:              asset,
		labels:             make([]*labelItem, 0),
		cycleTime:          0,
		doneWithCredits:    false,
		cyclesTillNextLine: 0,
		renderer:           renderer,
		navigator:          navigator,
		uiManager:          ui,
	}

	credits.Logger = d2util.NewLogger()
	credits.Logger.SetLevel(l)
	credits.Logger.SetPrefix(logPrefix)

	return credits
}

// Credits represents the credits screen
type Credits struct {
	creditsBackground  *d2ui.Sprite
	exitButton         *d2ui.Button
	creditsText        []string
	labels             []*labelItem
	cycleTime          float64
	cyclesTillNextLine int
	doneWithCredits    bool

	asset     *d2asset.AssetManager
	renderer  d2interface.Renderer
	navigator d2interface.Navigator
	uiManager *d2ui.UIManager

	*d2util.Logger
}

// LoadContributors loads the contributors data from file
func (v *Credits) LoadContributors() []string {
	file, err := os.Open(path.Join("./", "CONTRIBUTORS"))
	if err != nil || file == nil {
		v.Warning("CONTRIBUTORS file is missing")
		return []string{"MISSING CONTRIBUTOR FILES!"}
	}

	defer func() {
		if err = file.Close(); err != nil {
			v.Errorf("an error occurred while closing file: %s, err: %e", file.Name(), err)
		}
	}()

	scanner := bufio.NewScanner(file)

	var contributors []string
	for scanner.Scan() {
		contributors = append(contributors, strings.Trim(scanner.Text(), " "))
	}

	return contributors
}

// OnLoad is called to load the resources for the credits screen
func (v *Credits) OnLoad(loading d2screen.LoadingState) {
	var err error

	v.creditsBackground, err = v.uiManager.NewSprite(d2resource.CreditsBackground, d2resource.PaletteSky)
	if err != nil {
		v.Error(err.Error())
	}

	v.creditsBackground.SetPosition(creditsX, creditsY)
	loading.Progress(twentyPercent)

	v.exitButton = v.uiManager.NewButton(d2ui.ButtonTypeMedium, v.asset.TranslateLabel(d2enum.ExitLabel))
	v.exitButton.SetPosition(charSelExitBtnX, charSelExitBtnY)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
	loading.Progress(fourtyPercent)

	fileData, err := v.asset.LoadFile(d2resource.CreditsText)
	if err != nil {
		loading.Error(err)
		return
	}

	loading.Progress(sixtyPercent)

	creditData, err := d2util.Utf16BytesToString(fileData[2:])
	if err != nil {
		v.Error(err.Error())
	}

	v.creditsText = strings.Split(creditData, "\r\n")

	for i := range v.creditsText {
		v.creditsText[i] = strings.Trim(v.creditsText[i], " ")
	}

	loading.Progress(eightyPercent)

	v.creditsText = append(v.LoadContributors(), v.creditsText...)
}

// Render renders the credits screen
func (v *Credits) Render(screen d2interface.Surface) {
	v.creditsBackground.RenderSegmented(screen, 4, 3, 0)

	for _, label := range v.labels {
		if label.Available {
			continue
		}

		label.Label.Render(screen)
	}
}

// Advance runs the update logic on the credits screen
func (v *Credits) Advance(tickTime float64) error {
	v.cycleTime += tickTime
	for v.cycleTime >= secondsPerCycle {
		v.cycleTime -= secondsPerCycle
		v.cyclesTillNextLine--

		if !v.doneWithCredits && v.cyclesTillNextLine <= 0 {
			v.addNextItem()
		}

		for _, label := range v.labels {
			if label.Available {
				continue
			}

			_, y := label.Label.GetPosition()

			if y-1 < -15 {
				label.Available = true
				continue
			}

			label.Label.OffsetPosition(0, -1)
		}
	}

	return nil
}

func (v *Credits) onExitButtonClicked() {
	v.navigator.ToMainMenu()
}

func (v *Credits) addNextItem() {
	if len(v.creditsText) == 0 {
		v.doneWithCredits = true
		return
	}

	text := v.creditsText[0]
	v.creditsText = v.creditsText[1:]

	if text == "" {
		if v.creditsText[0][0] == '*' {
			v.cyclesTillNextLine = 38
			return
		}

		v.cyclesTillNextLine = 19

		return
	}

	isHeading := text[0] == '*'
	isNextHeading := len(v.creditsText) > 0 && len(v.creditsText[0]) > 0 && v.creditsText[0][0] == '*'
	isNextSpace := len(v.creditsText) > 0 && v.creditsText[0] == ""

	var label = v.getNewFontLabel(isHeading)

	if isHeading {
		label.SetText(d2ui.ColorTokenize(text[1:], d2ui.ColorTokenRed))
	} else {
		label.SetText(d2ui.ColorTokenize(text, d2ui.ColorTokenGold))
	}

	isDoubled, isNextHeading := v.setItemLabelPosition(label, isHeading, isNextHeading, isNextSpace)

	switch {
	case isHeading && isNextHeading:
		v.cyclesTillNextLine = 38
	case isNextHeading:
		if isDoubled {
			v.cyclesTillNextLine = 38
		} else {
			v.cyclesTillNextLine = 57
		}
	case isHeading:
		v.cyclesTillNextLine = 38
	default:
		v.cyclesTillNextLine = 19
	}
}

const (
	itemLabelY            = 605
	itemLabelX            = 400
	itemLabel2offsetX     = 10
	halfItemLabel2offsetX = itemLabel2offsetX / 2
)

func (v *Credits) setItemLabelPosition(label *d2ui.Label, isHeading, isNextHeading, isNextSpace bool) (isDoubled, nextHeading bool) {
	width, _ := label.GetSize()
	half := 2
	halfWidth := width / half

	if !isHeading && !isNextHeading && !isNextSpace {
		isDoubled = true
		// Gotta go side by side
		label.SetPosition(itemLabelX-width, itemLabelY)

		text2 := v.creditsText[0]
		v.creditsText = v.creditsText[1:]

		nextHeading = len(v.creditsText) > 0 && len(v.creditsText[0]) > 0 && v.creditsText[0][0] == '*'
		label2 := v.getNewFontLabel(isHeading)
		label2.SetText(d2ui.ColorTokenize(text2, d2ui.ColorTokenGold))

		label2.SetPosition(itemLabelX+itemLabel2offsetX, itemLabelY)

		return isDoubled, nextHeading
	}

	label.SetPosition(itemLabelX+halfItemLabel2offsetX-halfWidth, itemLabelY)

	return isDoubled, isNextHeading
}

func (v *Credits) getNewFontLabel(isHeading bool) *d2ui.Label {
	newLabelItem := &labelItem{
		Available: false,
		IsHeading: isHeading,
		Label:     v.uiManager.NewLabel(d2resource.FontFormal10, d2resource.PaletteSky),
	}
	v.labels = append(v.labels, newLabelItem)

	return newLabelItem.Label
}
