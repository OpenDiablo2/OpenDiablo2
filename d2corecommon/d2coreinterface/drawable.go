package d2coreinterface

import "github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"

// Drawable represents an instance that can be drawn
type Drawable interface {
	Render(target *d2surface.Surface)
	GetSize() (width, height int)
	SetPosition(x, y int)
	GetPosition() (x, y int)
	GetVisible() bool
	SetVisible(visible bool)
}
