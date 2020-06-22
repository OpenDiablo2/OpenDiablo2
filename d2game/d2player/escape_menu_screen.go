package d2player

import (
	"fmt"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type itemID int

const (
	itemOptions      itemID = 0
	itemSaveAndExit         = 1
	itemReturnToGame        = 2

	itemSoundOptions      itemID = 0
	itemVideoOptions             = 1
	itemAutomapOptions           = 2
	itemConfigureControls        = 3
	itemPreviousMenu             = 4
)

type Screen interface {
	Load(pentLeft, pentRight *d2ui.Sprite, selectSound d2audio.SoundEffect)
	Render(target d2render.Surface)
	Advance(elapsed float64) error
	OnUpKey()
	OnDownKey()
	OnEnterKey()
	OnMouseMove(event d2input.MouseMoveEvent) bool
	OnMouseButtonDown(event d2input.MouseEvent) bool
	Reset()
}

type baseScreen struct {
	totalHeight    int
	labels         []d2ui.Label
	current        itemID
	defaultItem    itemID
	pentLeft       *d2ui.Sprite
	pentRight      *d2ui.Sprite
	selectSound    d2audio.SoundEffect
	switchScreenFn func(id screenID)
}

func (s *baseScreen) Render(target d2render.Surface) {
	tw, th := target.GetSize()
	midX := tw / 2
	startY := (th - s.totalHeight) / 2

	for i, label := range s.labels {
		_, lh := label.GetSize()
		ly := startY + i*lh

		label.SetPosition(midX, ly)
		label.Render(target)

		if i == int(s.current) {
			_, sh := s.pentLeft.GetCurrentFrameSize()
			s.pentLeft.SetPosition(100, label.Y+sh)
			s.pentRight.SetPosition(tw-100, label.Y+sh)

			s.pentLeft.Render(target)
			s.pentRight.Render(target)
		}
	}
}

func (s *baseScreen) Load(pentLeft, pentRight *d2ui.Sprite, selectSound d2audio.SoundEffect) {
	s.pentLeft = pentLeft
	s.pentLeft.PlayForward()
	s.pentRight = pentRight
	s.pentRight.PlayBackward()
	s.selectSound = selectSound

	totalHeight := 0
	for _, item := range s.labels {
		_, ih := item.GetSize()
		totalHeight += ih
	}
	s.totalHeight = totalHeight
}

func (s *baseScreen) Advance(elapsed float64) error {
	s.pentLeft.Advance(elapsed)
	s.pentRight.Advance(elapsed)
	return nil
}

func (s *baseScreen) OnUpKey() {
	if s.current == 0 {
		return
	}
	s.current--
}

func (s *baseScreen) OnDownKey() {
	if int(s.current) == len(s.labels)-1 {
		return
	}
	s.current++
}

func (s *baseScreen) OnMouseMove(event d2input.MouseMoveEvent) bool {
	return false
}

func (s *baseScreen) OnMouseButtonDown(event d2input.MouseEvent) bool {
	return false
}

func (s *baseScreen) Reset() {
	s.current = s.defaultItem
}

type optionsScreen struct {
	*baseScreen
}

func newOptionsScreen(switchScreenFn func(screenID)) *optionsScreen {
	labels := []d2ui.Label{
		itemSoundOptions:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		itemVideoOptions:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		itemAutomapOptions:    d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		itemConfigureControls: d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		itemPreviousMenu:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
	}
	labels[itemSoundOptions].SetText("sound options")
	labels[itemSoundOptions].Alignment = d2ui.LabelAlignCenter
	labels[itemVideoOptions].SetText("video options")
	labels[itemVideoOptions].Alignment = d2ui.LabelAlignCenter
	labels[itemAutomapOptions].SetText("automap options")
	labels[itemAutomapOptions].Alignment = d2ui.LabelAlignCenter
	labels[itemConfigureControls].SetText("configure controls")
	labels[itemConfigureControls].Alignment = d2ui.LabelAlignCenter
	labels[itemPreviousMenu].SetText("previous menu")
	labels[itemPreviousMenu].Alignment = d2ui.LabelAlignCenter

	return &optionsScreen{
		baseScreen: &baseScreen{
			defaultItem:    itemSoundOptions,
			labels:         labels,
			switchScreenFn: switchScreenFn,
		},
	}
}

func (s *optionsScreen) OnEnterKey() {
	s.selectSound.Play()
	switch s.current {
	case itemSoundOptions:
		fmt.Println("TODO: sound options menu")
	case itemVideoOptions:
		fmt.Println("TODO: video options menu")
	case itemAutomapOptions:
		fmt.Println("TODO: automap options menu")
	case itemConfigureControls:
		fmt.Println("TODO: configure controls menu")
	case itemPreviousMenu:
		s.switchScreenFn(mainScreenID)
	}
}

type mainScreen struct {
	*baseScreen
}

func newMainScreen(switchScreenFn func(screenID)) *mainScreen {
	labels := []d2ui.Label{
		itemOptions:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		itemSaveAndExit:  d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
		itemReturnToGame: d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
	}
	labels[itemOptions].SetText("options")
	labels[itemOptions].Alignment = d2ui.LabelAlignCenter
	labels[itemSaveAndExit].SetText("save and exit game")
	labels[itemSaveAndExit].Alignment = d2ui.LabelAlignCenter
	labels[itemReturnToGame].SetText("return to game")
	labels[itemReturnToGame].Alignment = d2ui.LabelAlignCenter

	return &mainScreen{
		baseScreen: &baseScreen{
			labels:         labels,
			switchScreenFn: switchScreenFn,
			defaultItem:    itemOptions,
		},
	}
}

func (s *mainScreen) OnEnterKey() {
	s.selectSound.Play()
	switch s.current {
	case itemOptions:
		s.switchScreenFn(optionsScreenID)
	case itemSaveAndExit:
		fmt.Println("TODO: save and exit menu")
	case itemReturnToGame:
		s.switchScreenFn(exitScreenID)
	}
}
