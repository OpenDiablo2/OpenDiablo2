package d2coreinterface

import "github.com/OpenDiablo2/OpenDiablo2/d2render/d2surface"

// Scene defines the function necessary for scene management
type Scene interface {
	Load() []func()
	Unload()
	Render(target *d2surface.Surface)
	Update(tickTime float64)
}
