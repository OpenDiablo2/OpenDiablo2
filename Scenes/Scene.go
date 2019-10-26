package Scenes

import (
	"github.com/hajimehoshi/ebiten"
)

// Scene defines the function necessary for scene management
type Scene interface {
	Load() []func()
	Unload()
	Render(screen *ebiten.Image)
	Update()
}
