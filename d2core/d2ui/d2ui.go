package d2ui

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
)

// CursorButton represents a mouse button
type CursorButton uint8

const (
	// CursorButtonLeft represents the left mouse button
	CursorButtonLeft CursorButton = 1
	// CursorButtonRight represents the right mouse button
	CursorButtonRight CursorButton = 2
)

// NewUIManager creates a UIManager instance with the given input and audio provider
func NewUIManager(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	input d2interface.InputManager,
	audio d2interface.AudioProvider,
) *UIManager {
	ui := &UIManager{
		asset:        asset,
		renderer:     renderer,
		inputManager: input,
		audio:        audio,
	}

	return ui
}
