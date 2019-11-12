package d2mapengine

import "github.com/hajimehoshi/ebiten"

type TileCacheRecord struct {
	Image   *ebiten.Image
	XOffset int
	YOffset int
}
