package Common

import "github.com/essial/OpenDiablo2/Palettes"

// SpriteProvider is an instance that can provide sprites
type SpriteProvider interface {
	LoadSprite(fileName string, palette Palettes.Palette) *Sprite
}
