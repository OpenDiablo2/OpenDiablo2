package Common

import (
	"github.com/hajimehoshi/ebiten"
)

// SceneInterface defines the function necessary for scene management
type SceneInterface interface {
	Load()
	Unload()
	Render(screen *ebiten.Image)
	Update()
}
