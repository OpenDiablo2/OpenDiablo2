package ebiten

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2debugutil"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
)

const (
	maxAlpha   = 0xff
	cacheLimit = 512
)

type colorMCacheKey uint32

type colorMCacheEntry struct {
	colorMatrix ebiten.ColorM
	atime       int64
}

type ebitenSurface struct {
	stateStack     []surfaceState
	stateCurrent   surfaceState
	image          *ebiten.Image
	colorMCache    map[colorMCacheKey]*colorMCacheEntry
	monotonicClock int64
}

func createEbitenSurface(img *ebiten.Image, currentState ...surfaceState) *ebitenSurface {
	state := surfaceState{
		effect:     d2enum.DrawEffectNone,
		saturation: 1.0,
		brightness: 1.0,
		skewX:      0.0,
		skewY:      0.0,
		scaleX:     1.0,
		scaleY:     1.0,
	}
	if len(currentState) > 0 {
		state = currentState[0]
	}

	return &ebitenSurface{
		image:        img,
		stateCurrent: state,
		colorMCache:  make(map[colorMCacheKey]*colorMCacheEntry),
	}
}

// PushTranslation pushes an x,y translation to the state stack
func (s *ebitenSurface) PushTranslation(x, y int) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.x += x
	s.stateCurrent.y += y
}

// PushSkew pushes a skew to the state stack
func (s *ebitenSurface) PushSkew(skewX, skewY float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.skewX = skewX
	s.stateCurrent.skewY = skewY
}

// PushScale pushes a scale to the state stack
func (s *ebitenSurface) PushScale(scaleX, scaleY float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.scaleX = scaleX
	s.stateCurrent.scaleY = scaleY
}

// PushEffect pushes an effect to the state stack
func (s *ebitenSurface) PushEffect(effect d2enum.DrawEffect) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.effect = effect
}

// PushFilter pushes a filter to the state stack
func (s *ebitenSurface) PushFilter(filter d2enum.Filter) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.filter = d2ToEbitenFilter(filter)
}

// PushColor pushes a color to the stat stack
func (s *ebitenSurface) PushColor(c color.Color) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.color = c
}

// PushBrightness pushes a brightness value to the state stack
func (s *ebitenSurface) PushBrightness(brightness float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.brightness = brightness
}

// PushSaturation pushes a saturation value to the state stack
func (s *ebitenSurface) PushSaturation(saturation float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.saturation = saturation
}

// Pop pops a state off of the state stack
func (s *ebitenSurface) Pop() {
	count := len(s.stateStack)
	if count == 0 {
		panic("empty stack")
	}

	s.stateCurrent = s.stateStack[count-1]
	s.stateStack = s.stateStack[:count-1]
}

// PopN pops n states off the the state stack
func (s *ebitenSurface) PopN(n int) {
	for i := 0; i < n; i++ {
		s.Pop()
	}
}

// Render renders the given surface
func (s *ebitenSurface) Render(sfc d2interface.Surface) error {
	opts := &ebiten.DrawImageOptions{}

	if s.stateCurrent.skewX != 0 || s.stateCurrent.skewY != 0 {
		opts.GeoM.Skew(s.stateCurrent.skewX, s.stateCurrent.skewY)
	}

	if s.stateCurrent.scaleX != 1.0 || s.stateCurrent.scaleY != 1.0 {
		opts.GeoM.Scale(s.stateCurrent.scaleX, s.stateCurrent.scaleY)
	}

	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))

	opts.Filter = s.stateCurrent.filter

	if s.stateCurrent.color != nil {
		opts.ColorM = s.colorToColorM(s.stateCurrent.color)
	}

	if s.stateCurrent.brightness != 1 || s.stateCurrent.saturation != 1 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	// Are these correct? who even knows
	switch s.stateCurrent.effect {
	case d2enum.DrawEffectPctTransparency25:
		opts.ColorM.Translate(0, 0, 0, -0.25)
	case d2enum.DrawEffectPctTransparency50:
		opts.ColorM.Translate(0, 0, 0, -0.50)
	case d2enum.DrawEffectPctTransparency75:
		opts.ColorM.Translate(0, 0, 0, -0.75)
	case d2enum.DrawEffectModulate:
		opts.CompositeMode = ebiten.CompositeModeLighter
	// TODO: idk what to do when ebiten doesn't exactly match, pick closest?
	case d2enum.DrawEffectBurn:
	case d2enum.DrawEffectNormal:
	case d2enum.DrawEffectMod2XTrans:
	case d2enum.DrawEffectMod2X:
	case d2enum.DrawEffectNone:
		opts.CompositeMode = ebiten.CompositeModeSourceOver
	}

	var img = sfc.(*ebitenSurface).image

	return s.image.DrawImage(img, opts)
}

// Renders the section of the surface, given the bounds
func (s *ebitenSurface) RenderSection(sfc d2interface.Surface, bound image.Rectangle) error {
	opts := &ebiten.DrawImageOptions{}

	if s.stateCurrent.skewX != 0 || s.stateCurrent.skewY != 0 {
		opts.GeoM.Skew(s.stateCurrent.skewX, s.stateCurrent.skewY)
	}

	if s.stateCurrent.scaleX != 1.0 || s.stateCurrent.scaleY != 1.0 {
		opts.GeoM.Scale(s.stateCurrent.scaleX, s.stateCurrent.scaleY)
	}

	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))

	opts.Filter = s.stateCurrent.filter

	if s.stateCurrent.color != nil {
		opts.ColorM = s.colorToColorM(s.stateCurrent.color)
	}

	if s.stateCurrent.brightness != 0 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	// Are these correct? who even knows
	switch s.stateCurrent.effect {
	case d2enum.DrawEffectPctTransparency25:
		opts.ColorM.Translate(0, 0, 0, -0.25)
	case d2enum.DrawEffectPctTransparency50:
		opts.ColorM.Translate(0, 0, 0, -0.50)
	case d2enum.DrawEffectPctTransparency75:
		opts.ColorM.Translate(0, 0, 0, -0.75)
	case d2enum.DrawEffectModulate:
		opts.CompositeMode = ebiten.CompositeModeLighter
	// TODO: idk what to do when ebiten doesn't exactly match, pick closest?
	case d2enum.DrawEffectBurn:
	case d2enum.DrawEffectNormal:
	case d2enum.DrawEffectMod2XTrans:
	case d2enum.DrawEffectMod2X:
	case d2enum.DrawEffectNone:
		opts.CompositeMode = ebiten.CompositeModeSourceOver
	}

	var img = sfc.(*ebitenSurface).image

	return s.image.DrawImage(img.SubImage(bound).(*ebiten.Image), opts)
}

// DrawTextf renders the string to the surface with the given format string and a set of parameters
func (s *ebitenSurface) DrawTextf(format string, params ...interface{}) {
	d2debugutil.D2DebugPrintAt(s.image, fmt.Sprintf(format, params...), s.stateCurrent.x, s.stateCurrent.y)
}

// DrawLine draws a line
func (s *ebitenSurface) DrawLine(x, y int, fillColor color.Color) {
	ebitenutil.DrawLine(
		s.image,
		float64(s.stateCurrent.x),
		float64(s.stateCurrent.y),
		float64(s.stateCurrent.x+x),
		float64(s.stateCurrent.y+y),
		fillColor,
	)
}

// DrawRect draws a rectangle
func (s *ebitenSurface) DrawRect(width, height int, fillColor color.Color) {
	ebitenutil.DrawRect(
		s.image,
		float64(s.stateCurrent.x),
		float64(s.stateCurrent.y),
		float64(width),
		float64(height),
		fillColor,
	)
}

// Clear clears the entire surface, filling with the given color
func (s *ebitenSurface) Clear(fillColor color.Color) error {
	return s.image.Fill(fillColor)
}

// GetSize gets the size of the surface
func (s *ebitenSurface) GetSize() (x, y int) {
	return s.image.Size()
}

// GetDepth returns the depth of this surface in the stack
func (s *ebitenSurface) GetDepth() int {
	return len(s.stateStack)
}

// ReplacePixels replaces pixels in the surface with the given pixels
func (s *ebitenSurface) ReplacePixels(pixels []byte) error {
	return s.image.ReplacePixels(pixels)
}

// Screenshot returns an *image.RGBA of the surface
func (s *ebitenSurface) Screenshot() *image.RGBA {
	width, height := s.GetSize()
	bounds := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: width, Y: height}}
	rgba := image.NewRGBA(bounds)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rgba.Set(x, y, s.image.At(x, y))
		}
	}

	return rgba
}

func (s *ebitenSurface) now() int64 {
	s.monotonicClock++
	return s.monotonicClock
}

// colorToColorM converts a normal color to a color matrix
func (s *ebitenSurface) colorToColorM(clr color.Color) ebiten.ColorM {
	// RGBA() is in [0 - 0xffff]. Adjust them in [0 - 0xff].
	cr, cg, cb, ca := clr.RGBA()
	cr >>= 8
	cg >>= 8
	cb >>= 8
	ca >>= 8

	if ca == 0 {
		emptyColorM := ebiten.ColorM{}
		emptyColorM.Scale(0, 0, 0, 0)

		return emptyColorM
	}

	key := colorMCacheKey(cr | (cg << 8) | (cb << 16) | (ca << 24))
	e, ok := s.colorMCache[key]

	if ok {
		e.atime = s.now()
		return e.colorMatrix
	}

	if len(s.colorMCache) > cacheLimit {
		oldest := int64(math.MaxInt64)
		oldestKey := colorMCacheKey(0)

		for key, c := range s.colorMCache {
			if c.atime < oldest {
				oldestKey = key
				oldest = c.atime
			}
		}

		delete(s.colorMCache, oldestKey)
	}

	cm := ebiten.ColorM{}
	rf := float64(cr) / float64(ca)
	gf := float64(cg) / float64(ca)
	bf := float64(cb) / float64(ca)
	af := float64(ca) / maxAlpha
	cm.Scale(rf, gf, bf, af)

	e = &colorMCacheEntry{
		colorMatrix: cm,
		atime:       s.now(),
	}

	s.colorMCache[key] = e

	return e.colorMatrix
}
