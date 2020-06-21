package d2map

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2term"
)

type MapRenderer struct {
	mapEngine     *MapEngine
	viewport      *Viewport
	camera        Camera
	debugVisLevel int
}

func CreateMapRenderer(mapEngine *MapEngine) *MapRenderer {
	result := &MapRenderer{
		mapEngine: mapEngine,
		viewport:  NewViewport(0, 0, 800, 600),
	}
	result.viewport.SetCamera(&result.camera)
	d2term.BindAction("mapdebugvis", "set map debug visualization level", func(level int) {
		result.debugVisLevel = level
	})
	return result
}

func (m *MapRenderer) SetMapEngine(mapEngine *MapEngine) {
	m.mapEngine = mapEngine
}

func (m *MapRenderer) Render(target d2render.Surface) {
	for _, region := range m.mapEngine.regions {
		if region.isVisbile(m.viewport) {
			region.renderPass1(m.viewport, target)
			region.renderDebug(m.debugVisLevel, m.viewport, target)
			region.renderPass2(m.mapEngine.entities, m.viewport, target)
			region.renderPass3(m.viewport, target)
		}
	}
}

func (m *MapRenderer) MoveCameraTo(x, y float64) {
	m.camera.MoveTo(x, y)
}

func (m *MapRenderer) MoveCameraBy(x, y float64) {
	m.camera.MoveBy(x, y)
}

func (m *MapRenderer) ScreenToWorld(x, y int) (float64, float64) {
	return m.viewport.ScreenToWorld(x, y)
}

func (m *MapRenderer) ScreenToOrtho(x, y int) (float64, float64) {
	return m.viewport.ScreenToOrtho(x, y)
}

func (m *MapRenderer) WorldToOrtho(x, y float64) (float64, float64) {
	return m.viewport.WorldToOrtho(x, y)
}

func (m *MapRenderer) ViewportToLeft() {
	m.viewport.toLeft()
}

func (m *MapRenderer) ViewportToRight() {
	m.viewport.toRight()
}

func (m *MapRenderer) ViewportDefault() {
	m.viewport.resetAlign()
}
