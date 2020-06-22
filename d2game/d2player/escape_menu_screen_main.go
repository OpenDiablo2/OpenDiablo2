package d2player

//
//import (
//	"log"
//
//	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
//	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
//	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2audio"
//	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input"
//	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
//	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
//)
//
//func newMainScreen() *Screen {
//	return &Screen{}
//}
//
//// ScreenLoadHandler
//func (s *Screen) OnLoad() error {
//	s.labels = []d2ui.Label{
//		d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
//		d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
//		d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteSky),
//	}
//
//	s.labels[EscapeOptions].SetText("OPTIONS")
//	s.labels[EscapeSaveExit].SetText("SAVE AND EXIT GAME")
//	s.labels[EscapeReturn].SetText("RETURN TO GAME")
//
//	for i := range s.labels {
//		s.labels[i].Alignment = d2ui.LabelAlignCenter
//	}
//
//	animation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteUnits)
//	s.pentLeft, _ = d2ui.LoadSprite(animation)
//	s.pentLeft.SetBlend(false)
//	s.pentLeft.PlayBackward()
//
//	s.pentRight, _ = d2ui.LoadSprite(animation)
//	s.pentRight.SetBlend(false)
//	s.pentRight.PlayForward()
//
//	s.pentWidth, s.pentHeight = s.pentLeft.GetFrameBounds()
//	_, s.textHeight = s.labels[EscapeOptions].GetSize()
//
//	s.selectSound, _ = d2audio.LoadSoundEffect(d2resource.SFXCursorSelect)
//
//	return nil
//}
//
//// ScreenRenderHandler
//func (s *Screen) Render(target d2render.Surface) error {
//	if !s.isActive {
//		return nil
//	}
//
//	tw, _ := target.GetSize()
//	// X Position of the mid-render target.
//	midX := tw / 2
//
//	// Y Coordinates for the center of the first option
//	choiceStart := 210
//	// Y Delta, in pixels, between center of choices
//	choiceDx := 50
//	// X Delta, in pixels, between center of pentagrams
//	betwPentDist := 275
//
//	for i := range s.labels {
//		s.labels[i].SetPosition(midX, choiceStart+i*choiceDx-s.textHeight/2)
//		s.labels[i].Render(target)
//	}
//
//	s.pentLeft.SetPosition(midX-(betwPentDist+s.pentWidth/2), choiceStart+int(s.current)*choiceDx+s.pentHeight/2)
//	s.pentRight.SetPosition(midX+(betwPentDist-s.pentWidth/2), choiceStart+int(s.current)*choiceDx+s.pentHeight/2)
//
//	s.pentLeft.Render(target)
//	s.pentRight.Render(target)
//
//	return nil
//}
//
//// ScreenAdvanceHandler
//func (s *Screen) Advance(elapsed float64) error {
//	if !s.isActive {
//		return nil
//	}
//
//	s.pentLeft.Advance(elapsed)
//	s.pentRight.Advance(elapsed)
//	return nil
//}
//
//func (s *Screen) IsActive() bool {
//	return s.isActive
//}
//
//func (s *Screen) reset() {
//	s.current = EscapeOptions
//}
//
//func (s *Screen) OnUpKey() {
//	switch s.current {
//	case EscapeSaveExit:
//		s.current = EscapeOptions
//	case EscapeReturn:
//		s.current = EscapeSaveExit
//	}
//}
//
//func (s *Screen) OnDownKey() {
//	switch s.current {
//	case EscapeOptions:
//		s.current = EscapeSaveExit
//	case EscapeSaveExit:
//		s.current = EscapeReturn
//	}
//}
//
//func (s *Screen) OnEnterKey() {
//	s.selectCurrent()
//}
//
//// Moves current selection marker to closes option to mouse.
//func (s *Screen) OnMouseMove(event d2input.MouseMoveEvent) bool {
//	if !s.isActive {
//		return false
//	}
//	lbl := &s.labels[EscapeSaveExit]
//	reg := s.toMouseRegion(event.HandlerEvent, lbl)
//
//	switch reg {
//	case regAbove:
//		s.current = EscapeOptions
//	case regIn:
//		s.current = EscapeSaveExit
//	case regBelow:
//		s.current = EscapeReturn
//	}
//
//	return false
//}
//
//// Allows user to click on menu options in Y coord. of mouse is over label.
//func (s *Screen) OnMouseButtonDown(event d2input.MouseEvent) bool {
//	if !s.isActive {
//		return false
//	}
//
//	lbl := &s.labels[EscapeOptions]
//	if s.toMouseRegion(event.HandlerEvent, lbl) == regIn {
//		s.current = EscapeOptions
//		s.selectCurrent()
//		return false
//	}
//
//	lbl = &s.labels[EscapeSaveExit]
//	if s.toMouseRegion(event.HandlerEvent, lbl) == regIn {
//		s.current = EscapeSaveExit
//		s.selectCurrent()
//		return false
//	}
//
//	lbl = &s.labels[EscapeReturn]
//	if s.toMouseRegion(event.HandlerEvent, lbl) == regIn {
//		s.current = EscapeReturn
//		s.selectCurrent()
//		return false
//	}
//
//	return false
//}
//
//func (s *Screen) selectCurrent() {
//	switch s.current {
//	case EscapeOptions:
//		s.onOptions()
//		s.selectSound.Play()
//	case EscapeSaveExit:
//		s.onSaveAndExit()
//		s.selectSound.Play()
//	case EscapeReturn:
//		s.onReturnToGame()
//		s.selectSound.Play()
//	}
//}
//
//// User clicked on "OPTIONS"
//func (s *Screen) onOptions() error {
//	log.Println("OPTIONS Clicked from Escape Menu")
//	return nil
//}
//
//// User clicked on "SAVE AND EXIT"
//func (s *Screen) onSaveAndExit() error {
//	log.Println("SAVE AND EXIT GAME Clicked from Escape Menu")
//	return nil
//}
//
//// User clicked on "RETURN TO GAME"
//func (s *Screen) onReturnToGame() error {
//	return nil
//}
//
//// Where is the Y coordinate of the mouse compared to this label.
//func (s *Screen) toMouseRegion(event d2input.HandlerEvent, lbl *d2ui.Label) mouseRegion {
//	_, h := lbl.GetSize()
//	y := lbl.Y
//	my := event.Y
//
//	if my < y {
//		return regAbove
//	}
//	if my > (y + h) {
//		return regBelow
//	}
//	return regIn
//}
