package d2input

import "github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"

type KeyCharsHandler struct{}

func (KeyCharsHandler) OnKeyChars(event d2interface.KeyCharsEvent) bool {
	panic("implement me")
}
