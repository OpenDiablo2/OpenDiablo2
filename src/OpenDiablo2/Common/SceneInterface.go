package Common

import (
	"github.com/hajimehoshi/ebiten"
)

type SceneInterface interface {
	Load()
	Unload()
	Render(screen *ebiten.Image)
	Update()
}
