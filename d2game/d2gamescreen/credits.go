package d2gamescreen

import (
	"bufio"
	"image/color"
	"log"
	"os"
	"path"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2screen"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type labelItem struct {
	Label     d2ui.Label
	IsHeading bool
	Available bool
}

// Credits represents the credits screen
type Credits struct {
	creditsBackground  *d2ui.Sprite
	exitButton         d2ui.Button
	creditsText        []string
	labels             []*labelItem
	cycleTime          float64
	cyclesTillNextLine int
	doneWithCredits    bool
}

// CreateCredits creates an instance of the credits screen
func CreateCredits() *Credits {
	result := &Credits{
		labels:             make([]*labelItem, 0),
		cycleTime:          0,
		doneWithCredits:    false,
		cyclesTillNextLine: 0,
	}
	return result
}

// Load is called to load the contributors data from file
// TODO: use markdown for file and convert it to the suitable format
func (v *Credits) LoadContributors() []string {
	var contributors []string
	file, err := os.Open(path.Join("./", "CONTRIBUTORS"))
	if err != nil || file == nil {
		log.Print("CONTRIBUTORS file is missing")
		return []string{"MISSING CONTRIBUTOR FILES!"}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contributors = append(contributors, strings.Trim(scanner.Text(), " "))
	}
	return contributors
}

// OnLoad is called to load the resources for the credits screen
func (v *Credits) OnLoad(loading d2screen.LoadingState) {
	animation, _ := d2asset.LoadAnimation(d2resource.CreditsBackground, d2resource.PaletteSky)
	v.creditsBackground, _ = d2ui.LoadSprite(animation)
	v.creditsBackground.SetPosition(0, 0)
	loading.Progress(0.2)

	v.exitButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, "EXIT")
	v.exitButton.SetPosition(33, 543)
	v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
	d2ui.AddWidget(&v.exitButton)
	loading.Progress(0.4)

	fileData, err := d2asset.LoadFile(d2resource.CreditsText)
	if err != nil {
		loading.Error(err)
		return
	}
	loading.Progress(0.6)

	creditData, _ := d2common.Utf16BytesToString(fileData[2:])
	v.creditsText = strings.Split(creditData, "\r\n")
	for i := range v.creditsText {
		v.creditsText[i] = strings.Trim(v.creditsText[i], " ")
	}
	loading.Progress(0.8)

	v.creditsText = append(v.LoadContributors(), v.creditsText...)
}

// Render renders the credits screen
func (v *Credits) Render(screen d2render.Surface) error {
	v.creditsBackground.RenderSegmented(screen, 4, 3, 0)
	for _, label := range v.labels {
		if label.Available {
			continue
		}
		label.Label.Render(screen)
	}

	return nil
}

const secondsPerCycle = float64(0.02)

// Update runs the update logic on the credits screen
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
	mainMenu := CreateMainMenu()
	mainMenu.SetScreenMode(ScreenModeMainMenu)
	d2screen.SetNextScreen(mainMenu)
}

func (v *Credits) addNextItem() {
	if len(v.creditsText) == 0 {
		v.doneWithCredits = true
		return
	}

	text := v.creditsText[0]
	v.creditsText = v.creditsText[1:]
	if len(text) == 0 && v.creditsText[0][0] != '*' {
		v.cyclesTillNextLine = 19
		return
	} else if len(text) == 0 && v.creditsText[0][0] == '*' {
		v.cyclesTillNextLine = 38
		return
	}
	isHeading := text[0] == '*'
	isNextHeading := len(v.creditsText) > 0 && len(v.creditsText[0]) > 0 && v.creditsText[0][0] == '*'
	isNextSpace := len(v.creditsText) > 0 && len(v.creditsText[0]) == 0
	var label = v.getNewFontLabel(isHeading)
	if isHeading {
		label.SetText(text[1:])
	} else {
		label.SetText(text)
	}
	width, _ := label.GetSize()
	isDoubled := false
	if !isHeading && !isNextHeading && !isNextSpace {
		isDoubled = true

		// Gotta go side by side
		label.SetPosition(400-width, 605)

		text2 := v.creditsText[0]
		v.creditsText = v.creditsText[1:]

		isNextHeading = len(v.creditsText) > 0 && len(v.creditsText[0]) > 0 && v.creditsText[0][0] == '*'
		label2 := v.getNewFontLabel(isHeading)
		label2.SetText(text2)

		label2.SetPosition(410, 605)
	} else {
		label.SetPosition(405-width/2, 605)
	}

	if isHeading && isNextHeading {
		v.cyclesTillNextLine = 38
	} else if isNextHeading {
		if isDoubled {
			v.cyclesTillNextLine = 38
		} else {
			v.cyclesTillNextLine = 57
		}
	} else if isHeading {
		v.cyclesTillNextLine = 38
	} else {
		v.cyclesTillNextLine = 19
	}
}

func (v *Credits) getNewFontLabel(isHeading bool) *d2ui.Label {
	for _, label := range v.labels {
		if label.Available {
			label.Available = false
			if isHeading {
				label.Label.Color = color.RGBA{R: 255, G: 88, B: 82, A: 255}
			} else {
				label.Label.Color = color.RGBA{R: 198, G: 178, B: 150, A: 255}
			}
			return &label.Label
		}
	}

	newLabelItem := &labelItem{
		Available: false,
		IsHeading: isHeading,
		Label:     d2ui.CreateLabel(d2resource.FontFormal10, d2resource.PaletteSky),
	}

	if isHeading {
		newLabelItem.Label.Color = color.RGBA{R: 255, G: 88, B: 82, A: 255}
	} else {
		newLabelItem.Label.Color = color.RGBA{R: 198, G: 178, B: 150, A: 255}

	}

	v.labels = append(v.labels, newLabelItem)
	return &newLabelItem.Label
}
