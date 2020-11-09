package d2gamescreen

import (
	//"bufio"
	"fmt"
	"log"
	//"os"
	//"path"
	//"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	//"github.com/OpenDiablo2/OpenDiablo2/d2common/d2util"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

const (
	cinematicsX, cinematicsY               = 100, 100
	a1BtnX, a1BtnY                         = 0, 0
	a2BtnX, a2BtnY                         = 0, 0
	a3BtnX, a3BtnY                         = 0, 0
	a4BtnX, a4BtnY                         = 0, 0
	a5BtnX, a5BtnY                         = 0, 0
	endCreditBtnX, endCreditBtnY           = 0, 0
	cinematicsExitBtnX, cinematicsExitBtnY = 33, 543
)

// Cinematics represents the cinematics screen
type Cinematics struct {
	cinematicsBackground *d2ui.Sprite
	a1Btn                *d2ui.Sprite
	a2Btn                *d2ui.Sprite
	a3Btn                *d2ui.Sprite
	a4Btn                *d2ui.Sprite
	a5Btn                *d2ui.Sprite
	endCreditBtn         *d2ui.Sprite
	cinematicsExitBtn    *d2ui.Button

	asset     *d2asset.AssetManager
	renderer  d2interface.Renderer
	navigator d2interface.Navigator
	uiManager *d2ui.UIManager
}

// CreateCinematics creates an instance of the credits screen
func CreateCinematics(
	navigator d2interface.Navigator,
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager) *Cinematics {
	result := &Cinematics{
		asset:     asset,
		renderer:  renderer,
		navigator: navigator,
		uiManager: ui,
	}

	return result
}

// OnLoad is called to load the resources for the credits screen
func (v *Cinematics) OnLoad(loading d2screen.LoadingState) {
	var err error

	fmt.Println("\n\ninitializing background\n\n")
	v.cinematicsBackground, err = v.uiManager.NewSprite(d2resource.CinematicsBackground, d2resource.PaletteSky)
	if err != nil {
		fmt.Println("\n\nerror occured\n\n")
		log.Print(err)
	}
	fmt.Println("\n\nsetting position\n\n")
	v.cinematicsBackground.SetPosition(0, 0)

	loading.Progress(twentyPercent)

	v.cinematicsExitBtn = v.uiManager.NewButton(d2ui.ButtonTypeMedium, "EXIT")
	v.cinematicsExitBtn.SetPosition(charSelExitBtnX, charSelExitBtnY)
	//	v.cinematicsExitBtn.OnActivated(func() { })
	/*loading.Progress(fourtyPercent)

	fileData, err := v.asset.LoadFile(d2resource.CreditsText)
	if err != nil {
		loading.Error(err)
		return
	}

	loading.Progress(sixtyPercent)

	creditData, err := d2util.Utf16BytesToString(fileData[2:])
	if err != nil {
		log.Print(err)
	}

	v.creditsText = strings.Split(creditData, "\r\n")

	for i := range v.creditsText {
		v.creditsText[i] = strings.Trim(v.creditsText[i], " ")
	}

	loading.Progress(eightyPercent)

	v.creditsText = append(v.LoadContributors(), v.creditsText...)*/
}

/*
// Render renders the credits screen
func (v *Credits) Render(screen d2interface.Surface) {
	err := v.creditsBackground.RenderSegmented(screen, 4, 3, 0)
	if err != nil {
		return
	}

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

			if label.Label.Y-1 < -15 {
				label.Available = true
				continue
			}
			label.Label.Y--
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
		label.SetText(text[1:])
	} else {
		label.SetText(text)
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
		label2.SetText(text2)

		label2.SetPosition(itemLabelX+itemLabel2offsetX, itemLabelY)

		return isDoubled, nextHeading
	}

	label.SetPosition(itemLabelX+halfItemLabel2offsetX-halfWidth, itemLabelY)

	return isDoubled, isNextHeading
}

const (
	lightRed = 0xff5852ff
	beige    = 0xc6b296ff
)

func (v *Credits) getNewFontLabel(isHeading bool) *d2ui.Label {
	for _, label := range v.labels {
		if label.Available {
			label.Available = false
			if isHeading {
				label.Label.Color[0] = rgbaColor(lightRed)
			} else {
				label.Label.Color[0] = rgbaColor(beige)
			}

			return label.Label
		}
	}

	newLabelItem := &labelItem{
		Available: false,
		IsHeading: isHeading,
		Label:     v.uiManager.NewLabel(d2resource.FontFormal10, d2resource.PaletteSky),
	}

	if isHeading {
		newLabelItem.Label.Color[0] = rgbaColor(lightRed)
	} else {
		newLabelItem.Label.Color[0] = rgbaColor(beige)
	}

	v.labels = append(v.labels, newLabelItem)

	return newLabelItem.Label
}*/
