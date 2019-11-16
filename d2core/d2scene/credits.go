package d2scene

import (
	"bufio"
	"image/color"
	"log"
	"os"
	"path"
	"strings"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"

	"github.com/OpenDiablo2/OpenDiablo2/d2data/d2datadict"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"

	"github.com/OpenDiablo2/OpenDiablo2/d2render"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

	"github.com/OpenDiablo2/OpenDiablo2/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2common"
	dh "github.com/OpenDiablo2/OpenDiablo2/d2helper"
	"github.com/OpenDiablo2/OpenDiablo2/d2render/d2ui"
	"github.com/hajimehoshi/ebiten"
)

type labelItem struct {
	Label     d2ui.Label
	IsHeading bool
	Available bool
}

// Credits represents the credits scene
type Credits struct {
	uiManager          *d2ui.Manager
	soundManager       *d2audio.Manager
	fileProvider       d2interface.FileProvider
	sceneProvider      d2interface.SceneProvider
	creditsBackground  d2render.Sprite
	exitButton         d2ui.Button
	creditsText        []string
	labels             []*labelItem
	cycleTime          float64
	cyclesTillNextLine int
	doneWithCredits    bool
}

// CreateCredits creates an instance of the credits scene
func CreateCredits(fileProvider d2interface.FileProvider, sceneProvider d2interface.SceneProvider, uiManager *d2ui.Manager, soundManager *d2audio.Manager) *Credits {
	result := &Credits{
		fileProvider:       fileProvider,
		uiManager:          uiManager,
		soundManager:       soundManager,
		sceneProvider:      sceneProvider,
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
	contributors := []string{}
	file, err := os.Open(path.Join("./", "CONTRIBUTORS"))
	if err != nil {
		log.Print("CONTRIBUTORS file is missing")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		contributors = append(contributors, strings.Trim(scanner.Text(), " "))
	}
	return contributors
}

// Load is called to load the resources for the credits scene
func (v *Credits) Load() []func() {
	return []func(){
		func() {
			v.creditsBackground = d2render.CreateSprite(v.fileProvider.LoadFile(d2resource.CreditsBackground), d2datadict.Palettes[d2enum.Sky])
			v.creditsBackground.MoveTo(0, 0)
		},
		func() {
			v.exitButton = d2ui.CreateButton(d2ui.ButtonTypeMedium, v.fileProvider, d2common.TranslateString("#970"))
			v.exitButton.MoveTo(33, 543)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(&v.exitButton)
		},
		func() {
			fileData, _ := dh.Utf16BytesToString(v.fileProvider.LoadFile(d2resource.CreditsText)[2:])
			v.creditsText = strings.Split(fileData, "\r\n")
			for i := range v.creditsText {
				v.creditsText[i] = strings.Trim(v.creditsText[i], " ")
			}
			v.creditsText = append(v.LoadContributors(), v.creditsText...)
		},
	}
}

// Unload unloads the data for the credits scene
func (v *Credits) Unload() {

}

// Render renders the credits scene
func (v *Credits) Render(screen *ebiten.Image) {
	v.creditsBackground.DrawSegments(screen, 4, 3, 0)
	for _, label := range v.labels {
		if label.Available {
			continue
		}
		label.Label.Draw(screen)
	}
}

const secondsPerCycle = float64(0.02)

// Update runs the update logic on the credits scene
func (v *Credits) Update(tickTime float64) {
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
}

func (v *Credits) onExitButtonClicked() {
	mainMenu := CreateMainMenu(v.fileProvider, v.sceneProvider, v.uiManager, v.soundManager)
	mainMenu.ShowTrademarkScreen = false
	v.sceneProvider.SetNextScene(mainMenu)
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
		label.MoveTo(400-int(width), 605)

		text2 := v.creditsText[0]
		v.creditsText = v.creditsText[1:]

		isNextHeading = len(v.creditsText) > 0 && len(v.creditsText[0]) > 0 && v.creditsText[0][0] == '*'
		label2 := v.getNewFontLabel(isHeading)
		label2.SetText(text2)

		label2.MoveTo(410, 605)
	} else {
		label.MoveTo(405-int(width/2), 605)
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
				label.Label.Color = color.RGBA{255, 88, 82, 255}
			} else {
				label.Label.Color = color.RGBA{198, 178, 150, 255}
			}
			return &label.Label
		}
	}

	newLabelItem := &labelItem{
		Available: false,
		IsHeading: isHeading,
		Label:     d2ui.CreateLabel(v.fileProvider, d2resource.FontFormal10, d2enum.Sky),
	}

	if isHeading {
		newLabelItem.Label.Color = color.RGBA{255, 88, 82, 255}
	} else {
		newLabelItem.Label.Color = color.RGBA{198, 178, 150, 255}

	}

	v.labels = append(v.labels, newLabelItem)
	return &newLabelItem.Label
}
