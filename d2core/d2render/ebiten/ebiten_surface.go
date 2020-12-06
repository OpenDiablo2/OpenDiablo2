package ebiten

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2enum"
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2interface"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

// static check that we implement our interface
var _ d2interface.Surface = &ebitenSurface{}

const (
	maxAlpha       = 0xff
	cacheLimit     = 512
	transparency25 = 0.25
	transparency50 = 0.50
	transparency75 = 0.75
)

type colorMCacheKey uint32

type colorMCacheEntry struct {
	colorMatrix ebiten.ColorM
	atime       int64
}

type ebitenSurface struct {
	renderer       *Renderer
	stateStack     []surfaceState
	stateCurrent   surfaceState
	image          *ebiten.Image
	colorMCache    map[colorMCacheKey]*colorMCacheEntry
	monotonicClock int64
}

func createEbitenSurface(r *Renderer, img *ebiten.Image, currentState ...surfaceState) *ebitenSurface {
	state := surfaceState{
		effect:     d2enum.DrawEffectNone,
		saturation: defaultSaturation,
		brightness: defaultBrightness,
		skewX:      defaultSkewX,
		skewY:      defaultSkewY,
		scaleX:     defaultScaleX,
		scaleY:     defaultScaleY,
	}
	if len(currentState) > 0 {
		state = currentState[0]
	}

	return &ebitenSurface{
		renderer:     r,
		image:        img,
		stateCurrent: state,
		colorMCache:  make(map[colorMCacheKey]*colorMCacheEntry),
	}
}

// Renderer returns the renderer
func (s *ebitenSurface) Renderer() d2interface.Renderer {
	return s.renderer
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

// PushRotate pushes a rotation to the state stack
func (s *ebitenSurface) PushRotate(theta float64) {
	s.stateStack = append(s.stateStack, s.stateCurrent)
	s.stateCurrent.rotate = theta
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

func (s *ebitenSurface) RenderSprite(sprite *d2ui.Sprite) {
	opts := s.createDrawImageOptions()

	if s.stateCurrent.brightness != 1 || s.stateCurrent.saturation != 1 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	s.handleStateEffect(opts)

	sprite.Render(s)
}

// Render renders the given surface
func (s *ebitenSurface) Render(sfc d2interface.Surface) {
	opts := s.createDrawImageOptions()

	if s.stateCurrent.brightness != 1 || s.stateCurrent.saturation != 1 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	s.handleStateEffect(opts)

	s.image.DrawImage(sfc.(*ebitenSurface).image, opts)
}

// Renders the section of the surface, given the bounds
func (s *ebitenSurface) RenderSection(sfc d2interface.Surface, bound image.Rectangle) {
	opts := s.createDrawImageOptions()

	if s.stateCurrent.brightness != 0 {
		opts.ColorM.ChangeHSV(0, s.stateCurrent.saturation, s.stateCurrent.brightness)
	}

	s.handleStateEffect(opts)

	s.image.DrawImage(sfc.(*ebitenSurface).image.SubImage(bound).(*ebiten.Image), opts)
}

func (s *ebitenSurface) createDrawImageOptions() *ebiten.DrawImageOptions {
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Skew(s.stateCurrent.skewX, s.stateCurrent.skewY)
	opts.GeoM.Scale(s.stateCurrent.scaleX, s.stateCurrent.scaleY)
	opts.GeoM.Rotate(s.stateCurrent.rotate)
	opts.GeoM.Translate(float64(s.stateCurrent.x), float64(s.stateCurrent.y))

	opts.Filter = s.stateCurrent.filter

	if s.stateCurrent.color != nil {
		opts.ColorM = s.colorToColorM(s.stateCurrent.color)
	}

	return opts
}

func (s *ebitenSurface) handleStateEffect(opts *ebiten.DrawImageOptions) {
	switch s.stateCurrent.effect {
	case d2enum.DrawEffectPctTransparency25:
		opts.ColorM.Translate(0, 0, 0, -transparency25)
	case d2enum.DrawEffectPctTransparency50:
		opts.ColorM.Translate(0, 0, 0, -transparency50)
	case d2enum.DrawEffectPctTransparency75:
		opts.ColorM.Translate(0, 0, 0, -transparency75)
	case d2enum.DrawEffectModulate:
		opts.CompositeMode = ebiten.CompositeModeLighter
	// https://github.com/OpenDiablo2/OpenDiablo2/issues/822
	case d2enum.DrawEffectBurn:
	case d2enum.DrawEffectNormal:
	case d2enum.DrawEffectMod2XTrans:
	case d2enum.DrawEffectMod2X:
	case d2enum.DrawEffectNone:
		opts.CompositeMode = ebiten.CompositeModeSourceOver
	}
}

// DrawTextf renders the string to the surface with the given format string and a set of parameters
func (s *ebitenSurface) DrawTextf(format string, params ...interface{}) {
	str := fmt.Sprintf(format, params...)
	s.Renderer().PrintAt(s.image, str, s.stateCurrent.x, s.stateCurrent.y)
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
func (s *ebitenSurface) Clear(fillColor color.Color) {
	s.image.Fill(fillColor)
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
func (s *ebitenSurface) ReplacePixels(pixels []byte) {
	s.image.ReplacePixels(pixels)
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
