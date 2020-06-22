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

type mouseRegion int

const (
	regAbove mouseRegion = 0
	regIn                = 1
	regBelow             = 2
)

type Screen interface {
	Load(pentLeft, pentRight *d2ui.Sprite, selectSound d2audio.SoundEffect)
	Render(target d2render.Surface)
	Advance(elapsed float64) error
	OnUpKey()
	OnDownKey()
	OnEnterKey()
	OnMouseMove(event d2input.MouseMoveEvent) bool
	OnLeftClick(event d2input.MouseEvent) bool
	Reset()
}

type baseScreen struct {
	totalHeight    int
	labels         []*d2ui.Label
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
	s.pentRight = pentRight
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
	for i, label := range s.labels {
		region := s.toMouseRegion(event.HandlerEvent, label)
		if region == regIn {
			s.current = itemID(i)
			return true
		}
		if i == 0 && region == regAbove {
			s.current = 0
			return true
		}
		if i == len(s.labels)-1 && region == regBelow {
			s.current = itemID(len(s.labels) - 1)
			return true
		}
	}
	return false
}

func (s *baseScreen) toMouseRegion(event d2input.HandlerEvent, lbl *d2ui.Label) mouseRegion {
	_, h := lbl.GetSize()
	y := lbl.Y
	my := event.Y

	if my < y {
		return regAbove
	}
	if my > (y + h) {
		return regBelow
	}
	return regIn
}

func (s *baseScreen) Reset() {
	s.current = s.defaultItem
}

type optionsScreen struct {
	*baseScreen
}

func newOptionsScreen(switchScreenFn func(screenID)) *optionsScreen {
	labels := make([]*d2ui.Label, 5)

	labelSound := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelSound.SetText("sound options")
	labelSound.Alignment = d2ui.LabelAlignCenter
	labels[itemSoundOptions] = &labelSound

	labelVideo := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelVideo.SetText("video options")
	labelVideo.Alignment = d2ui.LabelAlignCenter
	labels[itemVideoOptions] = &labelVideo

	labelAutomap := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelAutomap.SetText("automap options")
	labelAutomap.Alignment = d2ui.LabelAlignCenter
	labels[itemAutomapOptions] = &labelAutomap

	labelConfigureOpts := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelConfigureOpts.SetText("configure controls")
	labelConfigureOpts.Alignment = d2ui.LabelAlignCenter
	labels[itemConfigureControls] = &labelConfigureOpts

	labelPrevMenu := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelPrevMenu.SetText("previous menu")
	labelPrevMenu.Alignment = d2ui.LabelAlignCenter
	labels[itemPreviousMenu] = &labelPrevMenu

	return &optionsScreen{
		baseScreen: &baseScreen{
			defaultItem:    itemSoundOptions,
			labels:         labels,
			switchScreenFn: switchScreenFn,
		},
	}
}

func (s *optionsScreen) OnEnterKey() {
	s.selectItem(s.current)
}

func (s *optionsScreen) OnLeftClick(event d2input.MouseEvent) bool {
	for i, label := range s.labels {
		reg := s.toMouseRegion(event.HandlerEvent, label)
		if reg != regIn {
			continue
		}
		s.selectItem(itemID(i))
	}
	return true
}

func (s *optionsScreen) selectItem(opt itemID) {
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
	labels := make([]*d2ui.Label, 3)

	labelOptions := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelOptions.SetText("options")
	labelOptions.Alignment = d2ui.LabelAlignCenter
	labels[itemOptions] = &labelOptions

	labelSaveExit := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelSaveExit.SetText("save and exit game")
	labelSaveExit.Alignment = d2ui.LabelAlignCenter
	labels[itemSaveAndExit] = &labelSaveExit

	labelReturn := d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky)
	labelReturn.SetText("return to game")
	labelReturn.Alignment = d2ui.LabelAlignCenter
	labels[itemReturnToGame] = &labelReturn

	return &mainScreen{
		baseScreen: &baseScreen{
			labels:         labels,
			switchScreenFn: switchScreenFn,
			defaultItem:    itemOptions,
		},
	}
}

func (s *mainScreen) OnEnterKey() {
	s.selectItem(s.current)
}

func (s *mainScreen) OnLeftClick(event d2input.MouseEvent) bool {
	for i, label := range s.labels {
		reg := s.toMouseRegion(event.HandlerEvent, label)
		if reg != regIn {
			continue
		}
		s.selectItem(itemID(i))
	}
	return true
}

func (s *mainScreen) selectItem(opt itemID) {
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
