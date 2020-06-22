package d2player

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Screen interface {
	OnLoad()
	Render(target d2render.Surface)
	Advance(elapsed float64) error
	OnUpKey()
	OnDownKey()
	OnEnterKey()
	OnMouseMove(event d2input.MouseMoveEvent) bool
	OnMouseButtonDown(event d2input.MouseEvent) bool
	PrevScreen() Screen
	Reset()
}

type baseScreen struct {
	items       []*Item
	currentItem int
	prevScreen  Screen
}

func (s *baseScreen) Render(target d2render.Surface) {
	//tw, _ := target.GetSize()
	//// X Position of the mid-render target.
	//midX := tw / 2
	//
	//// Y Coordinates for the center of the first option
	//choiceStart := 210
	//// Y Delta, in pixels, between center of choices
	//choiceDx := 50

	for _, item := range s.items {
		item.Render(target)
	}

	//_, th := s.labels[EscapeOptions].GetSize()
	//for i := range s.labels {
	//	s.labels[i].SetPosition(midX, choiceStart+i*choiceDx-th/2)
	//	s.labels[i].Render(target)
	//}

	//s.pentLeft.SetPosition(midX-(betwPentDist+s.pentWidth/2), choiceStart+int(s.current)*choiceDx+s.pentHeight/2)
	//s.pentRight.SetPosition(midX+(betwPentDist-s.pentWidth/2), choiceStart+int(s.current)*choiceDx+s.pentHeight/2)
	//
	//s.pentLeft.Render(target)
	//s.pentRight.Render(target)
}

func (s *baseScreen) Advance(elapsed float64) error {
	return nil
}

func (s *baseScreen) OnUpKey() {

}

func (s *baseScreen) OnDownKey() {

}

func (s *baseScreen) OnEnterKey() {

}

func (s *baseScreen) OnMouseMove(event d2input.MouseMoveEvent) bool {
	return false
}

func (s *baseScreen) OnMouseButtonDown(event d2input.MouseEvent) bool {
	return false
}

func (s *baseScreen) PrevScreen() Screen {
	return s.prevScreen
}

func (s *baseScreen) Reset() {
	s.currentItem = 0
}

type mainScreen struct {
	*baseScreen
}

func newMainScreen() *mainScreen {
	return &mainScreen{
		baseScreen: &baseScreen{
			items: []*Item{
				newItem("options"),
				newItem("save and exit game"),
				newItem("return to game"),
			},
		},
	}
}

func (s *mainScreen) OnLoad() {
	for _, item := range s.items {
		item.Load()
	}
}
