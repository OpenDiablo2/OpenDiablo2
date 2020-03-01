package d2input_ebiten

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/keyboard"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2input/mouse"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func keyToEbiten(k keyboard.Key) ebiten.Key {
	return ebiten.Key(k)
}

func keyFromEbiten(k ebiten.Key) keyboard.Key {
	return keyboard.Key(k)
}

func mouseToEbiten(mb mouse.MouseButton) ebiten.MouseButton {
	return ebiten.MouseButton(mb)
}

func mouseFromEbiten(mb ebiten.MouseButton) mouse.MouseButton {
	return mouse.MouseButton(mb)
}

type Backend struct{}

func (b *Backend) CursorPosition() (int, int) {
	return ebiten.CursorPosition()
}

func (b *Backend) IsKeyPressed(k keyboard.Key) bool {
	ebitenKey := keyToEbiten(k)
	return ebiten.IsKeyPressed(ebitenKey)
}

func (b *Backend) IsKeyJustPressed(k keyboard.Key) bool {
	ebitenKey := keyToEbiten(k)
	return inpututil.IsKeyJustPressed(ebitenKey)
}

func (b *Backend) IsKeyJustReleased(k keyboard.Key) bool {
	ebitenKey := keyToEbiten(k)
	return inpututil.IsKeyJustReleased(ebitenKey)
}

func (b *Backend) IsMouseButtonPressed(mb mouse.MouseButton) bool {
	ebitenMouseButton := mouseToEbiten(mb)
	return ebiten.IsMouseButtonPressed(ebitenMouseButton)
}

func (b *Backend) IsMouseButtonJustPressed(mb mouse.MouseButton) bool {
	ebitenMouseButton := mouseToEbiten(mb)
	return inpututil.IsMouseButtonJustPressed(ebitenMouseButton)
}

func (b *Backend) IsMouseButtonJustReleased(mb mouse.MouseButton) bool {
	ebitenMouseButton := mouseToEbiten(mb)
	return inpututil.IsMouseButtonJustReleased(ebitenMouseButton)
}

func (b *Backend) InputChars() []rune {
	return ebiten.InputChars()
}
