package Scenes

import (
	"image/color"
	"strings"

	"github.com/essial/OpenDiablo2/Common"
	"github.com/essial/OpenDiablo2/PaletteDefs"
	"github.com/essial/OpenDiablo2/ResourcePaths"
	"github.com/essial/OpenDiablo2/Sound"
	"github.com/essial/OpenDiablo2/UI"
	"github.com/hajimehoshi/ebiten"
)

type labelItem struct {
	Label     *UI.Label
	IsHeading bool
	Available bool
}

// Credits represents the credits scene
type Credits struct {
	uiManager          *UI.Manager
	soundManager       *Sound.Manager
	fileProvider       Common.FileProvider
	sceneProvider      SceneProvider
	creditsBackground  *Common.Sprite
	exitButton         *UI.Button
	creditsText        []string
	labels             []*labelItem
	cycleTime          float64
	cyclesTillNextLine int
	doneWithCredits    bool
}

// CreateCredits creates an instance of the credits scene
func CreateCredits(fileProvider Common.FileProvider, sceneProvider SceneProvider, uiManager *UI.Manager, soundManager *Sound.Manager) *Credits {
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

// Load is called to load the resources for the credits scene
func (v *Credits) Load() []func() {
	return []func(){
		func() {
			v.creditsBackground = v.fileProvider.LoadSprite(ResourcePaths.CreditsBackground, PaletteDefs.Sky)
			v.creditsBackground.MoveTo(0, 0)
		},
		func() {
			v.exitButton = UI.CreateButton(UI.ButtonTypeMedium, v.fileProvider, Common.TranslateString("#970"))
			v.exitButton.MoveTo(30, 550)
			v.exitButton.OnActivated(func() { v.onExitButtonClicked() })
			v.uiManager.AddWidget(v.exitButton)
		},
		func() {
			fileData, _ := Common.Utf16BytesToString(v.fileProvider.LoadFile(ResourcePaths.CreditsText)[2:])
			v.creditsText = strings.Split(fileData, "\r\n")
			for i := range v.creditsText {
				v.creditsText[i] = strings.Trim(v.creditsText[i], " ")
			}
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
	if len(text) == 0 {
		v.cyclesTillNextLine = 18
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
		label.MoveTo(390-int(width), 605)

		text2 := v.creditsText[0]
		v.creditsText = v.creditsText[1:]

		isNextHeading = len(v.creditsText) > 0 && len(v.creditsText[0]) > 0 && v.creditsText[0][0] == '*'
		label2 := v.getNewFontLabel(isHeading)
		label2.SetText(text2)

		label2.MoveTo(410, 605)
	} else {
		label.MoveTo(400-int(width/2), 605)
	}

	if isHeading && isNextHeading {
		v.cyclesTillNextLine = 40
	} else if isNextHeading {
		if isDoubled {
			v.cyclesTillNextLine = 40
		} else {
			v.cyclesTillNextLine = 70
		}
	} else if isHeading {
		v.cyclesTillNextLine = 40
	} else {
		v.cyclesTillNextLine = 18
	}
}

func (v *Credits) getNewFontLabel(isHeading bool) *UI.Label {
	for _, label := range v.labels {
		if label.Available {
			label.Available = false
			if isHeading {
				label.Label.Color = color.RGBA{255, 88, 82, 255}
			} else {
				label.Label.Color = color.RGBA{198, 178, 150, 255}
			}
			return label.Label
		}
	}

	newLabelItem := &labelItem{
		Available: false,
		IsHeading: isHeading,
		Label:     UI.CreateLabel(v.fileProvider, ResourcePaths.FontFormal10, PaletteDefs.Sky),
	}

	if isHeading {
		newLabelItem.Label.Color = color.RGBA{255, 88, 82, 255}
	} else {
		newLabelItem.Label.Color = color.RGBA{198, 178, 150, 255}

	}

	v.labels = append(v.labels, newLabelItem)
	return newLabelItem.Label
}
